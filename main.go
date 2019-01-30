package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
  "time"
)

const restartTimeout = 200

func main() {
  // Get parameters out of program arguments 
  timeoutPtr := flag.Int("timeout", 180, "Seconds before all unfinished downloads are killed")
  var filename string
  var dest string
  flag.StringVar(&filename, "filename", "", "The path of the file with all of the urls")
  flag.StringVar(&dest, "destination", "", "The path of where to save the images")
  flag.Parse()

  // Print error and end program if improper arguments given
  if len(filename) == 0 || len(dest) == 0 {
    fmt.Println("Requred parameter not given, both filename and destination are required")
    return
  }

  // Prepare the destination by creating the directory and removing contents from it
  PrepareDestination(dest)

  // Create the timer process with the given timeout time
  timeoutDelay := time.Duration(*timeoutPtr) * time.Second
  timer := time.NewTimer(timeoutDelay).C

  file, err := os.Open(filename)

  if err != nil {
    fmt.Println("Failed to open file")
    return
  }

  // Close file after this function is over
  defer file.Close()

  // Create scanner to read urls file
  scanner := bufio.NewScanner(file)

  // Create tracker channel for the download processes to communicate
  // with the main process on
  tracker := make(chan int)

  // Loop over all lines in the urls file and start a download process for each
  // using i to keep track of the number of files remaining
  i := 0
  for scanner.Scan() {
    go downloadProcess(scanner.Text(), dest, i, tracker)
    i++
  }

  // While there are still files to download continue watching for messages
  for i > 0 {
    select {
    // If recieved message on timer channel then log and end program
    case <-timer:
      fmt.Println("Process killed by timer")
      os.Exit(0)
    // If recieved message from tracker channel then log and subtract one
    // from pending files count
    case <-tracker:
      fmt.Printf("Files remaining: %d\n", i)
      i--
    }
  }

  // If scanner bombs out then print errors
  if err := scanner.Err(); err != nil {
    fmt.Println(err)
    return
  }
}

func downloadProcess(url string, dest string, i int, c chan int) {
  // Download the image
  _, err := DownloadImage(url, dest, i)

  if err != nil {
    // If there is an error then sleep for a little and try again
    time.Sleep(restartTimeout * time.Millisecond)
    go downloadProcess(url, dest, i, c)
  } else {
    // If success then send a message on the tracker channel back to the main process
    c <- i
  }
}
