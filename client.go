package main

import (
  "fmt"
  "net"
  "encoding/binary"
  "github.com/zipcodeman/shilagenfir/util"
  "bufio"
)

func ExtractResponse(resp []byte) []bool {
  retval := make([]bool, len(resp) * 8)

  for i := 0; i < len(resp); i++ {
    for j := 0; j < 8; j++ {
      retval[8*i + j] = (resp[i] & 0x80) != 0
      resp[i] <<= 1
    }
  }

  return retval
}

func runClient() {
  conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))

  if err != nil {
    logger.Println(err)
    return
  }

  fil := util.NewFuzzyFile(size)

  for len(fil.GetUnconvergedRanges()) > 0 {
    fmt.Println(string(fil.Mid))

    conn.Write([]byte{byte(fil.Round)})

    b := make([]byte, 4)
    binary.LittleEndian.PutUint32(b, uint32(size))
    conn.Write(b)
    conn.Write(fil.Mid)

    br := bufio.NewReader(conn)

    b, err = ReadNBytesFromBuffer(4, br)

    responseLength := binary.LittleEndian.Uint32(b)

    response, _ := ReadNBytesFromBuffer(int(responseLength), br)

    real_response := ExtractResponse(response)
    fmt.Println(real_response)

    fil.Update(real_response)
  }

  fmt.Println(string(fil.Mid))
  conn.Close()
}
