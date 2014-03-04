package util

import (
  "crypto/rand"
  "fmt"
  "os"
  "io/ioutil"
  "math"
)

const (
  GREATER_THAN = iota
  LESS_THAN = iota
  EQUAL_TO = iota
)

type FuzzyFile struct {
  Max, Mid, Min []byte
  Round int
}

func NewFuzzyFile(length int) *FuzzyFile {
  if (length < 0) {
    return nil
  }

  ff := new(FuzzyFile)
  ff.Max = make([]byte, length)
  ff.Mid = make([]byte, length)
  ff.Min = make([]byte, length)
  ff.Round = GREATER_THAN

  max := maxByte()
  min := byte(0)

  mid := getMid(min, max)

  for i := 0; i < len(ff.Mid); i++ {
    ff.Mid[i] = mid
    ff.Max[i] = max
    ff.Min[i] = min
  }

  return ff
}

func (ff *FuzzyFile) Update(response []bool) {
  for i := 0; i < len(ff.Mid); i++ {
    ff.updateAt(i, response[i])
  }
  ff.Round = (ff.Round + 1) % 3
}

func (ff *FuzzyFile) updateAt(i int, response bool) {
  if ff.Round == GREATER_THAN {
    if response {
      ff.Max[i] = ff.Mid[i] - 1
    } else {
      ff.Min[i] = ff.Mid[i]
    }
  } else if ff.Round == LESS_THAN {
    if response {
      ff.Min[i] = ff.Mid[i] + 1
    } else {
      ff.Max[i] = ff.Mid[i]
    }
  } else {
    if response {
      ff.Min[i] = ff.Mid[i]
      ff.Max[i] = ff.Mid[i]
    }
  }

  ff.Mid[i] = getMid(ff.Min[i], ff.Max[i])
}

func (ff *FuzzyFile) convergedAt(i int) bool {
  return ff.Min[i] == ff.Max[i]
}

func (ff *FuzzyFile) GetUnconvergedRanges() []int {
  retval := make([]int, 0, len(ff.Min) / 2)

  start := -1

  for i := 0; i < len(ff.Mid); i++ {
    if ff.convergedAt(i) {
      if (start >= 0) {
        retval = append(retval, start, i)
        start = -1
      }
    } else {
      if (start < 0) {
        start = i
      }
    }
  }

  if (start >= 0) {
    retval = append(retval, start, len(ff.Mid) - 1)
  }

  return retval
}

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

func maxByte() byte {
  return byte(math.Pow(2, 8)) - 1
}

func getMid(i, j byte) byte {
  return i + ((j - i) / 2);
}

func GetResponse(bytes, targ []byte, round int) []bool {
  response := make([]bool, len(bytes))

  for i := 0; i < len(bytes); i++ {
    if round == GREATER_THAN {
      response[i] = bytes[i] > targ[i]
    } else if round == LESS_THAN {
      response[i] = bytes[i] < targ[i]
    } else {
      response[i] = bytes[i] == targ[i]
    }
  }

  return response
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
