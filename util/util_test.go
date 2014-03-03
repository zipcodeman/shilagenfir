package util

import "testing"

func TestMaxByte(t *testing.T) {
  if maxByte() != 255 {
    t.Error("Max byte should be 255. It was instead ", maxByte())
  }
}

func TestGetMid(t *testing.T) {
  if getMid(0, 0) != 0 {
    t.Error("getMid(0,0) should be 0. It was instead ", getMid(0,0))
  }

  if getMid(0, 10) != 5 {
    t.Error("getMid(0,10) should be 5. It was instead ", getMid(0,10))
  }

  if getMid(0, 11) != 5 {
    t.Error("getMid(0,11) should be 5. It was instead ", getMid(0,11))
  }

  if getMid(5, 10) != 7 {
    t.Error("getMid(5,10) should be 7. It was instead ", getMid(5,10))
  }
}

func boolArrayEquals(a, b []bool) bool {
  if len(a) != len(b) {
    return false
  }
  for i, v := range a {
    if v != b[i] {
      return false
    }
  }
  return true
}

func TestGetResponse(t *testing.T) {
  if !boolArrayEquals(getResponse([]byte{0}, []byte{50}), []bool{true}) {
    t.Error("0 is indeed less than 50, sir", getResponse([]byte{0}, []byte{50}))
  }

  if !boolArrayEquals(getResponse([]byte{50}, []byte{0}), []bool{false}) {
    t.Error("50 is indeed less than 0, sir", getResponse([]byte{50}, []byte{0}))
  }

  if !boolArrayEquals(getResponse([]byte{0}, []byte{0}), []bool{false}) {
    t.Error("0 is actually equal to, sir", getResponse([]byte{0}, []byte{0}))
  }
}

func TestUpdateMinMax(t *testing.T) {
  min, max := updateMinMax(0, 10, true)
  if min != 0 && max != 4 {
    t.Error("Min and max did not match", 0, 4, min, max)
  }

  min, max = updateMinMax(0, 10, false)
  if min != 5 && max != 10 {
    t.Error("Min and max did not match", 5, 10, min, max)
  }
}

func TestFindRangeWithMidpoint(t *testing.T) {
  min, max := findRangeWithMidpoint(5, 0, 10)
  if min != 0 && max != 10 {
    t.Error("Min and max did not match", 0, 10, min, max)
  }

  min, max = findRangeWithMidpoint(2, 0, 10)
  if min != 0 && max != 4 {
    t.Error("Min and max did not match", 0, 4, min, max)
  }

  min, max = findRangeWithMidpoint(7, 0, 10)
  if min != 5 && max != 10 {
    t.Error("Min and max did not match", 5, 10, min, max)
  }
}

func BenchmarkFindRangeWithMidpoint(b *testing.B) {
  for i := 0; i < b.N; i++ {
    min, max := findRangeWithMidpoint(7, 0, 10)
    if min != 5 && max != 10 {
      b.Error("Min and max did not match", 5, 10, min, max)
    }
  }
}
