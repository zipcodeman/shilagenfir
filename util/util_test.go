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

func TestGetResponseGreaterThan(t *testing.T) {
  ze := []byte{0}
  fi := []byte{50}

  zero_fifty := GetResponse(ze, fi, GREATER_THAN)
  if !boolArrayEquals(zero_fifty, []bool{false}) {
    t.Error("0 is indeed greater than 50, sir", zero_fifty)
  }

  fifty_zero := GetResponse(fi, ze, GREATER_THAN)
  if !boolArrayEquals(fifty_zero, []bool{true}) {
    t.Error("50 is indeed greater than 0, sir", fifty_zero)
  }

  zero_zero := GetResponse(ze, ze, GREATER_THAN)
  if !boolArrayEquals(zero_zero, []bool{false}) {
    t.Error("0 is actually equal to, sir", zero_zero);
  }
}

func TestGetResponseLessThan(t *testing.T) {
  ze := []byte{0}
  fi := []byte{50}

  zero_fifty := GetResponse(ze, fi, LESS_THAN)
  if !boolArrayEquals(zero_fifty, []bool{true}) {
    t.Error("0 is indeed less than 50, sir", zero_fifty)
  }

  fifty_zero := GetResponse(fi, ze, LESS_THAN)
  if !boolArrayEquals(fifty_zero, []bool{false}) {
    t.Error("50 is indeed less than 0, sir", fifty_zero)
  }

  zero_zero := GetResponse(ze, ze, LESS_THAN)
  if !boolArrayEquals(zero_zero, []bool{false}) {
    t.Error("0 is actually equal to, sir", zero_zero);
  }
}

func TestUpdateAt(t *testing.T) {
  ff := NewFuzzyFile(1)

  ff.updateAt(0, true)

  if ff.Min[0] != 0 && ff.Max[0] != 126 {
    t.Error("Min and max did not match", 0, 126, ff.Min, ff.Max)
  }
}
