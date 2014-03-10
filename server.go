package main

import (
  "github.com/zipcodeman/shilagenfir/util"
  "net"
  "fmt"
  "bufio"
  "encoding/binary"
  "io"
  "math"
)

var data []byte

func init() {
  data = make([]byte, 81920)
}

func runServer() {
  ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

  if err != nil {
    fmt.Println(err)
  }

  for {
    logger.Println("Waiting for connection")
    conn, err := ln.Accept()

    if err != nil {
      fmt.Println(err)
      continue
    }

    go handleConnection(conn)
  }
}

func ReadNBytesFromBuffer(n int, br *bufio.Reader) ([]byte, error) {
  b, err := br.Peek(n)
  for ;err != nil; b, err = br.Peek(n) {
    if err == io.EOF {
      return nil, err
    }
  }

  for i := 0; i < n; i++ {
    br.ReadByte()
  }

  return b, nil
}

func CompressResponse(resp []bool) []byte {
  retval := make([]byte, int(math.Ceil(float64(len(resp)) / 8.0)))


  for i := 0; i < len(resp); i++ {
    byteNum := i / 8
    byteIx := i % 8
    if resp[i] {
      fmt.Println("Oring with: ", 1 << uint(7 - byteIx))
      retval[byteNum] |= 1 << uint(7 - byteIx)
      fmt.Printf("%b\n", retval[byteNum])
    }
  }


  return retval
}

func handleConnection(conn net.Conn) {
  logger.Println("New connection made", conn)
  br := bufio.NewReader(conn)
  var round byte
  var messageLength uint32
  for {
    b, err := ReadNBytesFromBuffer(1, br)
    if err != nil {
      break
    }
    round = b[0]

    b, err = ReadNBytesFromBuffer(4, br)

    if err != nil {
      break
    }

    messageLength = binary.LittleEndian.Uint32(b)
    logger.Println("Round:", round)
    logger.Println("Size:", messageLength)

    guess, err := ReadNBytesFromBuffer(int(messageLength), br)

    response := util.GetResponse(guess, targ, int(round))

    logger.Println("Responding with:", response)
    comp_resp := CompressResponse(response)
    b = make([]byte, 4)
    binary.LittleEndian.PutUint32(b, uint32(len(comp_resp)))

    conn.Write(b)
    conn.Write(comp_resp)
  }
  logger.Println("Closing connection. Nothing to read")
  conn.Close()
}
