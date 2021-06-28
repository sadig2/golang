package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(time.Second * 100))
	if err != nil {
		log.Println("connection timeout")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "i heard you say %s \n", ln)

	}
	defer conn.Close()

	fmt.Println("code got here")
}
