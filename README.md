# workout

This script executes the "segments.txt" file (by default) which contains workout segments. Each segment has a warmup period in seconds, a duration in seconds, and a title, each separated by a tab character. Each segment ends with a newline character. See the included "segments.txt" file for an example.

Example segments file:

```
10	30	Jumping Jacks
10	30	Wall Sit
10	30	Push Ups
10	30	Crunches
10	30	Chair Steps
10	30	Squats
10	30	Chair Dips
10	30	Plank
10	30	Knee Run
10	30	Lunges
10	30	Rotating Push Ups
10	30	Left Side Plank
10	30	Right Side Plank
```
 
## Requirements
 
- Go 1.6+
- Mac OS X (or a pre-installed command line utility that behaves like `say` from Mac OS)
- A desire to exercise :)
 
## Usage
 
 
How to print the usage message:
 
```
$ go build
$ ./workout -help
Usage of ./workout:
  -file string
    	The tab-separated file from which to load workout segments. See README.md for details. (default "segments.txt")
```
 
How to run the program:
 
 ```
 go build
 ./workout -file=segments.txt
 
 ```
