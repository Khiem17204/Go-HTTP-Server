package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	// handle concurrent connections
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handle_connection(conn)
	}
}

func handle_connection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		return
	}

	req_line := strings.Split(string(buffer), "\r\n")[0]
	user_agent := strings.Split(string(buffer), "\r\n")[2]
	path := strings.Split(req_line, " ")[1]
	method := strings.Split(req_line, " ")[0]
	if path == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.Split(path, "/")[1] == "echo" {
		message := strings.Split(path, "/")[2]
		conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)))
	} else if strings.Split(path, "/")[1] == "user-agent" {
		user_agent_val := strings.TrimPrefix(user_agent, "User-Agent: ")
		conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(user_agent_val), user_agent_val)))
	} else if strings.Split(path, "/")[1] == "files" && method == "GET" {
		dir := os.Args[2]
		file_name := strings.TrimPrefix(path, "/files/")
		data, err := os.ReadFile(dir + file_name)
		if err != nil {
			conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		} else {
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", len(data), data)))
		}
	} else if strings.Split(path, "/")[1] == "files" && method == "POST" {
		dir := os.Args[2]
		data := strings.Trim(string(buffer[len(buffer)-1]), "\x00")
		file_name := strings.TrimPrefix(path, "/files/")
		_ = os.WriteFile(dir+file_name, []byte(data), 0644)
		conn.Write([]byte("HTTP/1.1 201 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
