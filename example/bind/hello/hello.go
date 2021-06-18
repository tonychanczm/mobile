// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hello is a trivial package for gomobile bind example.
package hello

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

var o sync.Once
var ln net.Listener

func startListening() {
	listen := "0.0.0.0:8081"
	var err error
	ln, err = net.Listen("tcp", listen)
	if err != nil {
		panic(err)
	}
	log.Print("Now listening", listen, "...")
}

func Run() {
	o.Do(func() {
		startListening()
	})
}

func DoAccept() {
	for {
		if ln == nil {
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("Waiting Accept...")
		conn, err := ln.Accept()
		log.Printf("Accept! from %v", conn.RemoteAddr())
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			time.Sleep(1*time.Second)
			io.WriteString(conn, `HTTP/1.1 200 OK
Date: Mon, 27 Jul 2009 12:28:53 GMT
Server: Apache
Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT
ETag: "34aa387-d-1568eb00"
Accept-Ranges: bytes
Content-Length: 51
Vary: Accept-Encoding
Content-Type: text/plain

i am worked`)
			time.Sleep(10*time.Second)
		}()
	}
}