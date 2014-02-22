package util

import "testing"

func TestMaxByte(t *testing.T) {
  if maxByte() != 255 {
    t.Error("Max byte should be 255. It was instead ", maxByte())
  }
}
