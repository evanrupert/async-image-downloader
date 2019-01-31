package main

import (
  "fmt"
  "time"
)

// ===========EXAMPLE 1==============
// func main() {
//   go func() {
//     time.Sleep(1000 * time.Millisecond)
//     fmt.Println("Hello, Process!")
//   }()

//   fmt.Println("Hello, World!")

//   var input string
//   fmt.Scanln(&input)
// }

// ==============EXAMPLE 2==============
// func main() {
//   channel := make(chan string)

//   go func() {
//     msg := <-channel
//     fmt.Printf("Received message: '%s'", msg)
//   }()

//   fmt.Println("Hello, World!")

//   time.Sleep(1000 * time.Millisecond)
//   channel <- "Hello, Process!"

//   var input string
//   fmt.Scanln(&input)
// }

// =============EXAMPLE 3===========
// func main() {
//   channel := make(chan string)

//   go listenForMessages(channel)

//   fmt.Println("Hello, World!")

//   for i := 0; i < 10; i++ {
//     time.Sleep(1000 * time.Millisecond)
//     channel <- fmt.Sprintf("Hello for the %dth time", i)
//   }

//   var input string
//   fmt.Scanln(&input)
// }

// func listenForMessages(channel chan string) {
//   for {
//     msg := <- channel
//     fmt.Printf("Received message: '%s' on channel: '%v'\n", msg, channel)
//   }
// }

// ============EXAMPLE 4===========
// func main() {
//   channel := make(chan string)

//   go relayMessages(channel)

//   fmt.Println("Hello, World!")

//   time.Sleep(1000 * time.Millisecond)
//   channel <- "Hello, Process!"

//   msg := <- channel

//   fmt.Printf("Received message: '%s' on channel: '%v'", msg, channel)

//   var input string
//   fmt.Scanln(&input)
// }

// func relayMessages(channel chan string) {
//   for {
//     msg := <- channel
//     channel <- fmt.Sprintf("Received message: '%s' on channel: '%v'", msg, channel)
//   }
// }