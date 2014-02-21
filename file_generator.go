package main

import (
  "crypto/rand"
  "fmt"
  "os"
  "flag"
  "io/ioutil"
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

func readFile(filename string) []byte {
  f, err := ioutil.ReadFile(filename)

  if err != nil {
    return make([]byte, 0)
  }

  return f
}

var fname string
var size int
var target string

func init() {
  flag.StringVar(&fname, "f", "generated_file", "The file name")
  flag.IntVar(&size, "s", -1,  "Size of the generated file (in bytes)")
  flag.StringVar(&target, "t", "target",  "the target file")
  flag.Parse()
}

func main() {
  targ := readFile(target)

  if (size < 0) {
    size = len(targ)
  }

  fmt.Println("Read: ", targ)

  retval := make([]byte, size)

  for i := 0; i < size; i++ {
    for retval[i] != targ[i] {
      fmt.Println(string(retval))
      retval[i] = nextByte()
    }
  }
  fmt.Println(string(retval))
  writeFile(retval, fname)
}
