package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8972")
	if err != nil {
		panic(err)
	}

	go func() {
		if e := http.ListenAndServe("127.0.0.1:6060", nil); e != nil {
			log.Fatalf("pprof failed: %v", e)
		}
	}()
	var connections []net.Conn

	defer func() {
		for _, c := range connections {
			c.Close()
		}
	}()

	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}
			go handleConn(conn)
			connections = append(connections, conn)
			if len(connections)%100 == 0 {
				log.Printf("total number of connections: %v", len(connections))
			}
		}
	}
}

func handleConn(conn net.Conn) {
	io.Copy(ioutil.Discard, conn)
}
