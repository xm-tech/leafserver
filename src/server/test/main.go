package main

import "encoding/binary"
import "fmt"
import "net"

func main() {
	fmt.Println("client start")

	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	data := []byte(`{
		"Hello": {
			"Name": "leaf"
		}
	}`)

	m := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)))
	copy(m[2:], data)

	conn.Write(m)

	// not stop
	signal := make(chan int, 1)
	<-signal
}
