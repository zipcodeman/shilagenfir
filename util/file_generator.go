package util

import (
  "crypto/rand"
  "fmt"
  "os"
  "flag"
  "io/ioutil"
  "math"
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

func maxByte() int {
  return int(math.Pow(2, 8)) - 1
}

func getMid(i, j int) int {
  return (j - i) / 2;
}

func findRangeWithMidpoint(mid, min, max int) (int, int) {

  val := getMid(min, max)

  for val != mid {
    if mid < val {
      min = val
    } else {
      max = val - 1
    }
    val = getMid(min, max)
  }

  return min, max
}

func getResponse(bytes, targ []byte) []bool {
  response := make([]bool, len(bytes))

  for i := 0; i < size; i++ {
    response[i] = bytes[i] < targ[i]
  }

  return response
}

func getBytes(i int) []byte {
  b := make([]byte, i)
  siz := getMid(0, maxByte())

  for i := 0; i < len(b); i++ {
    b[i] = byte(siz)
  }

  return b
}

func updateBytes(bytes []byte, response []bool) {
  for i := 0; i < len(bytes); i++ {
    mid := int(bytes[i])
    min, max := findRangeWithMidpoint(mid, 0, maxByte())
    if response[i] {
      // it was >= the target
      max = mid
    } else {
      // it was < the target
      min = mid - 1
    }
    bytes[i] = byte(getMid(min, max))
  }
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
  bytes := getBytes(size)

  fmt.Println(string(bytes))
  response := getResponse(bytes, targ)
  fmt.Println(response)
  updateBytes(bytes, response)

  fmt.Println(string(bytes))
  response = getResponse(bytes, targ)
  fmt.Println(response)
  updateBytes(bytes, response)

  fmt.Println(string(bytes))
  response = getResponse(bytes, targ)
  fmt.Println(response)
  updateBytes(bytes, response)

  fmt.Println(string(bytes))
  response = getResponse(bytes, targ)
  fmt.Println(response)
  updateBytes(bytes, response)

  fmt.Println(string(bytes))
  response = getResponse(bytes, targ)
  fmt.Println(response)
  updateBytes(bytes, response)

  fmt.Println(string(bytes))
  response = getResponse(bytes, targ)
  fmt.Println(response)
  updateBytes(bytes, response)

  fmt.Println(string(bytes))
  return

  retval := make([]byte, size)

  for i := 0; i < size; i++ {
    for retval[i] != targ[i] {
      retval[i] = nextByte()
    }
    if i % (size/100) == 0 {
      perc := i/(size/100)
      fmt.Println("Generated ", i, " of ", size, " bytes (", perc, "percent complete)")
    }
  }
  fmt.Println(string(retval))

  writeFile(retval, fname)
}
