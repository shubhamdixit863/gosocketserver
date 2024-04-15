package main

import (
	"fmt"
	"log"
	"net"
)

// connection store for holding the connections

type ConnectionStore struct {
	connection map[net.Conn]bool
	// mutex lock we will be seeing
}

func (store *ConnectionStore) Add(conn net.Conn) {
	// we will add the connection object to the map
	store.connection[conn] = true // this should be into the mutex lock
}

func (store *ConnectionStore) Broadcast(sender net.Conn, message string) {
	// We will iterate over all the connections

	for k, _ := range store.connection {
		// anc we will be writing the message
		_, err := k.Write([]byte(message))
		if err != nil {
			log.Println("error in writing the message")
			continue
		}

	}

}

func handleConnection(conn net.Conn, store *ConnectionStore) {
	fmt.Println("Client connected")
	for {
		b := make([]byte, 1024) // 1 kb
		data, err := conn.Read(b)
		if err != nil {
			log.Println("error reading the data ", err)
		}
		data2 := fmt.Sprintf("Data : %s \n\n", b[:data])
		store.Broadcast(conn, data2)
		log.Println(data) // number of bytes received

	}

}

func main() {

	// create a broadcast channel

	//broadCastChannel := make(chan string, 10)
	//sendChannel := make(chan string, 10)

	storeObj := &ConnectionStore{connection: make(map[net.Conn]bool)}

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		storeObj.Add(conn) // this will store all the connection objects in the memory
		if err != nil {
			// handle error
		}
		go handleConnection(conn, storeObj)
	}

}
