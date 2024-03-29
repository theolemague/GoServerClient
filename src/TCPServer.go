package main

import (
  "bytes"
  "encoding/gob"
  "fmt"
  "io"
  "log"
  "net"
  "time"
)
// Create your custom data struct
type Message struct {
  ID   string
  Data string
}

func logerr(err error) bool {
  if err != nil {
    if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
      log.Println("read timeout:", err)
          } else if err == io.EOF {
          } else {
      log.Println("read error:", err)
        }
  return true
    }
  return false
}

func read(conn net.Conn) {
  tmp := make([]byte, 500)
  for {
    _, err := conn.Read(tmp)
    if logerr(err) {break}
    tmpbuff := bytes.NewBuffer(tmp)
    tmpstruct := new(Message)
    gobobj := gob.NewDecoder(tmpbuff)
    gobobj.Decode(tmpstruct)
    fmt.Println(tmpstruct)
    return
  }
}
func resp(conn net.Conn) {
  msg := Message{ID: "Yo", Data: "Hello back"}
  bin_buf := new(bytes.Buffer)
  gobobje := gob.NewEncoder(bin_buf)
  gobobje.Encode(msg)
  conn.Write(bin_buf.Bytes())
  conn.Close()
}

func handle(conn net.Conn) {
  timeoutDuration := 2 * time.Second
  fmt.Println("Launching server...")
  conn.SetReadDeadline(time.Now().Add(timeoutDuration))
  remoteAddr := conn.RemoteAddr().String()
  fmt.Println("Client connected from " + remoteAddr)
  read(conn)
  resp(conn)
}

func main() {
  server, _ := net.Listen("tcp", ":8081")
  for {
    conn, err := server.Accept()
      if err != nil {
        log.Println("Connection error: ", err)
        return
      }
    go handle(conn)
  	}
}