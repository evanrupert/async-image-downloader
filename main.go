package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "time"
)

func main() {
  timeoutPtr := flag.Int("timeout", 180, "Seconds before all unfinished downloads are killed")
  var filename string
  var dest string
  flag.StringVar(&filename, "filename", "", "The path of the file with all of the urls")
  flag.StringVar(&dest, "destination", "", "The path of where to save the images")

  flag.Parse()

  if len(filename) == 0 || len(dest) == 0 {
    fmt.Println("Requred parameter not given, both filename and destination are required")
    return
  }

  createDestination(dest)
  removeContents(dest)

  timeoutDelay := time.Duration(*timeoutPtr) * time.Second
  timer := time.NewTimer(timeoutDelay).C

  tracker := make(chan int)

  file, err := os.Open(filename)

  if err != nil {
    fmt.Println("Failed to open file")
    return
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  i := 0

  for scanner.Scan() {
    go downloadProcess(scanner.Text(), dest, i, tracker)
    i++
  }

  for i > 0 {
    select {
    case <-timer:
      fmt.Println("Process killed by timer")
      os.Exit(0)
    case <-tracker:
      fmt.Printf("Files remaining: %d\n", i)
      i--
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
    return
  }
}

func downloadProcess(url string, dest string, i int, c chan int) {
  _, err := DownloadImage(url, dest, i)
  if err != nil {
    time.Sleep(200 * time.Millisecond)
    go downloadProcess(url, dest, i, c)
  } else {
    c <- i
  }
}

func removeContents(dir string) error {
  d, err := os.Open(dir)
  if err != nil {
    return err
  }
  defer d.Close()
  names, err := d.Readdirnames(-1)
  if err != nil {
    return err
  }
  for _, name := range names {
    err = os.RemoveAll(filepath.Join(dir, name))
    if err != nil {
      return err
    }
  }
  return nil
}

func createDestination(dir string) error {
  cmd := exec.Command("mkdir", "-p", dir)
  return cmd.Run()
}
