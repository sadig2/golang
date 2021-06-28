package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
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
	defer conn.Close()
	//read request
	request(conn)
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//request line
			m := strings.Fields(ln)[0]
			fmt.Println("method", m)
		}
		if ln == "" {
			break
		}
		i++

	}

}

func respond(conn net.Conn) {
	body := `<html>

	<head>
		<title></title>
	</head>
	
	<body>
		<strong>
			hello world
		</strong>
	</body>
	
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content length %d\n", len(body))
	fmt.Fprint(conn, "Content-type: text/html\r\n")
	fmt.Fprint(conn, body)

}
