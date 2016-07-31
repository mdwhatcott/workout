package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"flag"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	for _, segment := range loadSegments(resolveFilename()) {
		segment.Execute()
	}
	say("Go take a shower!")
}

func resolveFilename() string {
	filename := flag.String("file", "segments.txt", "The tab-separated file from which to load workout segments. See README.md for details.")
	flag.Parse()
	return *filename
}

func loadSegments(filename string) []Segment {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return parseSegments(scanner)
}

func parseSegments(scanner *bufio.Scanner) (segments []Segment) {
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fields := strings.Split(line, "\t")
		warmUp, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		duration, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		title := fields[2]
		segments = append(segments, NewSegment(warmUp, duration, title))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return segments
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
	command.Start()
	return time.Since(started)
}
