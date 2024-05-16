package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	for _, segment := range loadSegments(resolveFilename()) {
		segment.Execute()
	}
	say("Go take a shower!")
}

func resolveFilename() string {
	_, file, _, _ := runtime.Caller(0)
	defaultPath := filepath.Join(filepath.Dir(file), "segments.txt")
	filename := flag.String("file", defaultPath,
		"The tab-separated file from which to load workout segments. See README.md for details.")
	flag.Parse()
	return *filename
}

func loadSegments(filename string) []Segment {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()
	return parseSegments(file)
}

func parseSegments(file io.Reader) (segments []Segment) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		segments = append(segments, parseSegment(scanner.Text()))
	}
	handleScannerError(scanner)
	return segments
}

func parseSegment(line string) Segment {
	parser := NewSegmentLineParser(line)
	if err := parser.Parse(); err != nil {
		log.Fatal(err)
	}
	return parser.Segment()
}

func handleScannerError(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func NewSegmentLineParser(line string) *SegmentLineParser {
	return &SegmentLineParser{line: strings.TrimSpace(line)}
}

type SegmentLineParser struct {
	line string
	err  error

	fields   []string
	warmUp   int
	duration int
	title    string
}

func (this *SegmentLineParser) Parse() error {
	this.fields = strings.Split(this.line, "\t")
	if len(this.fields) != 3 {
		return errors.New("each line must have 3 distinct fields (see README.md for details)")
	}
	this.warmUp, this.err = strconv.Atoi(this.fields[0])
	if this.err != nil {
		return this.err
	}
	this.duration, this.err = strconv.Atoi(this.fields[1])
	if this.err != nil {
		return this.err
	}
	this.title = strings.TrimSpace(this.fields[2])
	if len(this.title) == 0 {
		return errors.New("please provide a non-blank title")
	}
	return nil
}

func (this *SegmentLineParser) Segment() Segment {
	return NewSegment(this.warmUp, this.duration, this.title)
}

func NewSegment(warmUp, duration int, title string) Segment {
	return Segment{
		Title:    title,
		WarmUp:   time.Second * time.Duration(warmUp),
		Duration: time.Second * time.Duration(duration),
	}
}

type Segment struct {
	Title    string
	WarmUp   time.Duration
	Duration time.Duration
}

func (this Segment) Execute() {
	say("Warm up for " + this.Title)
	countdown(seconds(this.WarmUp))
	say(this.Title)
	countdown(seconds(this.Duration))
}

func seconds(duration time.Duration) int {
	return int(duration / time.Second)
}

func countdown(seconds int) {
	for x := seconds; x > 0; x-- {
		if x <= 5 {
			say(strconv.Itoa(x))
		}
		time.Sleep(time.Second)
	}
}

func say(message string) time.Duration {
	started := time.Now()
	fmt.Println(message)
	command := exec.Command("say", message)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	_ = command.Start()
	return time.Since(started)
}
