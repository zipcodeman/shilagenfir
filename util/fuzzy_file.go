package util

import (
  "math/rand"
  "os"
  "io/ioutil"
  "math"
)

const (
  GREATER_THAN = iota
  LESS_THAN = iota
  NUMBER_OF_ROUNDS = iota
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

  mid := getNewMid(min, max)

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
  ff.Round = (ff.Round + 1) % NUMBER_OF_ROUNDS
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
  }

  ff.Mid[i] = getNewMid(ff.Min[i], ff.Max[i])
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
    retval = append(retval, start, len(ff.Mid))
  }

  return retval
}

func (ff *FuzzyFile) ConvergedBytes() []byte {
  ranges := ff.GetUnconvergedRanges()
  retval := make([]byte, len(ff.Min))
  copy(retval, ff.Mid)

  for i := 0; i < len(ranges); i += 2 {
    for j := ranges[i]; j < ranges[i + 1]; j++ {
      retval[j] = '_'
    }
  }

  return retval
}

func maxByte() byte {
  return byte(math.Pow(2, 8)) - 1
}

func getNewMid(min, max byte) byte {
  if max - min > 0 {
    return min + byte(rand.Intn(int(max - min)))
  } else {
    return min
  }
}

func GetResponse(bytes, targ []byte, round int) []bool {
  response := make([]bool, len(bytes))

  for i := 0; i < len(bytes); i++ {
    if round == GREATER_THAN {
      response[i] = bytes[i] > targ[i]
    } else if round == LESS_THAN {
      response[i] = bytes[i] < targ[i]
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
