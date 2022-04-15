package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type Message struct {
	ID   string
	Data string
}
func send(conn net.Conn) {
	msg := Message{ID: "Yo", Data: "Hello"}
	bin_buf := new(bytes.Buffer)
	gobobj := gob.NewEncoder(bin_buf)
	gobobj.Encode(msg)
	conn.Write(bin_buf.Bytes())
}
func recv(conn net.Conn) {
	tmp := make([]byte, 500)
	conn.Read(tmp)
	tmpbuff := bytes.NewBuffer(tmp)
	tmpstruct := new(Message)
	gobobjdec := gob.NewDecoder(tmpbuff)
	gobobjdec.Decode(tmpstruct)
	fmt.Println(tmpstruct)
}
func main() {
	conn, _ := net.Dial("tcp", ":8081")
	send(conn)
	recv(conn)
}