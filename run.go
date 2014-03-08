package main

import (
  "github.com/zipcodeman/shilagenfir/util"
  "fmt"
  "flag"
  "io/ioutil"
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
var target string
var targ []byte

func init() {
  flag.StringVar(&fname, "f", "generated_file", "The file name")
  flag.IntVar(&size, "s", -1,  "Size of the generated file (in bytes)")
  flag.StringVar(&target, "t", "target",  "the target file")
  flag.Parse()

  targ = readFile(target)

  if (size < 0) {
    size = len(targ)
  }
}

func main() {
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
