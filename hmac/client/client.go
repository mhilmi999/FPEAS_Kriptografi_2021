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
	CONNECT := "127.0.0.1:567"
	c, err := net.Dial("tcp", CONNECT)
	if err != nil { //Jika terdapat eror
		fmt.Println(err)
		return
	}

	for {
		secret := bufio.NewReader(os.Stdin)
		fmt.Print("Masukan secret key nya >> ")
		secretkey, _ := secret.ReadString('\n')
		fmt.Fprintf(c, secretkey+"\n")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukan pesan nya >> ")
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
