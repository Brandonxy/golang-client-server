package main

import (
    "net"
    "fmt"
    "log"
    "os"
    "bufio"
    "bytes"
    "strings"
)
const address = "127.0.0.1:3000"
const bufferSize = 1024

func main() {
	
	listener, err := net.Listen("tcp", address)

	fmt.Printf("Escuchando en: " + address + "\n")

	if err != nil {
		log.Fatal("Problema al escuchar en : " + address)
		os.Exit(1)
	}

	var conn net.Conn
	var _ net.Conn

	conn, _ = listener.Accept()

	fmt.Println("Session abierta desde: " + conn.RemoteAddr().String())

    for {

		var in *bufio.Reader = bufio.NewReader(os.Stdin)

		var command string
		var _ error


        fmt.Print("Escriba un comando> ")

        command,_ = in.ReadString('\n')

        command = strings.Replace(command,"\n","", -1)

        if len(command) > 0 {

            /**
             * Send the command to the client
             */
            executeCommand(command, conn)

            HandleConnection(conn)
        }
    }

    defer conn.Close()
}

func executeCommand(command string, client net.Conn) {

	client.Write([]byte(command))
}

func HandleConnection(conn net.Conn) {

	var data []byte

	buffer := make([]byte, bufferSize)

    /*
     * Read the recv data into the buffer
     */
    n, err := conn.Read(buffer)

    if err != nil {
        os.Exit(1)
    }

    /**
    * Remove non-used bytes
    */
    buffer = bytes.Trim(buffer[:n], "\x00")

    data = append(data, buffer...)

    fmt.Println(string(data))

    data = make([]byte, 0)
}