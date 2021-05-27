package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//Konfigurasi awal koneksi ke server
	CONNECT := "127.0.0.1:123"
	c, err := net.Dial("tcp", CONNECT)
	if err != nil { //Jika terdapat eror
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)

		//Jika user mengetikan "KELUAR" maka tutup koneksi ke server
		if strings.TrimSpace(string(text)) == "KELUAR" {
			fmt.Println("Menutup koneksi client ke server...")
			return
		}
	}
}
