package main

import (
  "crypto/rand"
  "fmt"
  "os"
)

func nextByte() byte {
  return nextBytes(1)[0]
}

func nextBytes(i int) []byte {
  b := make([]byte, i)
  _, err := rand.Read(b)
  if err != nil {
    fmt.Println("error: ", err)
   }
  return b
}

func writeFile(data []byte, filename string) {
  f, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0666)
  if err != nil {
    return
  }

  defer f.Close()

  _, err = f.Write(data)
  if err != nil {
    return
  }

  return
}

func main() {
  writeFile(nextBytes(100), os.Args[1])
}
