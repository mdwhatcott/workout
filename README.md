# workout

This script executes the "segments.txt" file (by default) which contains workout segments. Each segment has a warmup period in seconds, a duration in seconds, and a title, each separated by a tab character. Each segment ends with a newline character. See the included "segments.txt" file for an example.
 
 ## Requirements:
 
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