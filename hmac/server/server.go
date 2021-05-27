package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"crypto/hmac"
	"crypto/sha256"
)

func main() {
	//Konfigurasi awal Server
	PORT := ":567"
	l, err := net.Listen("tcp", PORT) //Listen via TCP
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil { //Jika terdapat eror
		fmt.Println(err)
		return
	}

	for {
		// secret := "r4h4514"

		secret, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Println(secret)

		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		//Jika user mengetikan "KELUAR" maka tutup server
		if strings.TrimSpace(string(netData)) == "KELUAR" {
			fmt.Println("Menutup server!")
			return
		}

		//Lakukan hashing hmac sha256
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(string(netData)))
		hasil := h.Sum(nil)

		//Cetak hasil hashing hmac sha256
		fmt.Printf("-> %x\n", hasil)

		//Waktu pengiriman
		t := time.Now() //Dapatkan waktu sekarang
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime)) //Cetak waktu
	}
}
