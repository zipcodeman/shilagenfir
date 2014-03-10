package main

import (
  "github.com/zipcodeman/shilagenfir/util"
  "fmt"
  "flag"
  "io/ioutil"
  "log"
  "os"
)

func readFile(filename string) []byte {
  f, err := ioutil.ReadFile(filename)

  if err != nil {
    return make([]byte, 0)
  }

  return f
}

var fname string
var size int
var port int
var target string
var targ []byte
var server bool
var logger *log.Logger

func init() {
  flag.StringVar(&fname, "f", "generated_file", "The file name")
  flag.StringVar(&target, "t", "target",  "the target file")
  flag.IntVar(&port, "p", 15231, "the port to listen/connect to")
  flag.BoolVar(&server, "s", false, "start in server mode")
  flag.Parse()

  targ = readFile(target)

  size = len(targ)

  logger = log.New(os.Stdout, "", log.Ldate | log.Ltime)
}

func main() {
  if server {
    logger.Println("Starting server")
    runServer()
  } else {
    logger.Println("Starting client")
    runClient()
    return
  }

  fil := util.NewFuzzyFile(size)

  i := 0
  for len(fil.GetUnconvergedRanges()) > 0 {
    fmt.Println(string(fil.ConvergedBytes()))
    response := util.GetResponse(fil.Mid, targ, fil.Round)
    fil.Update(response)
    fmt.Println(fil.GetUnconvergedRanges())
    fmt.Println()
    fmt.Println()
    fmt.Println()
    fmt.Println()
    fmt.Println()
    i++
  }

  fmt.Println(string(fil.Mid))
  fmt.Println("Converged in", i, "iterations")
}
