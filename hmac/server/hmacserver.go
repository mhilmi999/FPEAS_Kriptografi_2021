package main

 import (
         "crypto/hmac"
         "crypto/md5"
         "crypto/rand"
         "encoding/base64"
         "fmt"
         "log"
         "net"
 )

 // both client and server MUST have the same secret key
 // to authenticate

 var secret = "GolangIsAwesome!"

 func randStr(strSize int, randType string) string {

         var dictionary string

         if randType == "alphanum" {
                 dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
         }

         if randType == "alpha" {
                 dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
         }

         if randType == "number" {
                 dictionary = "0123456789"
         }

         var bytes = make([]byte, strSize)
         rand.Read(bytes)
         for k, v := range bytes {
                 bytes[k] = dictionary[v%byte(len(dictionary))]
         }
         return string(bytes)
 }

 func serverSideAuthenticate(clientConn net.Conn, secretKey string) {
         // request client authentication
         // send client a random string as message

         message := randStr(16, "alphanum")

         _, err := clientConn.Write([]byte(message))

         if err != nil {
                 clientConn.Close()
         }

         fmt.Println("Data send to client : ", message)

         // prepare server side hmac digest
         // with secret key and msg

         hasher := hmac.New(md5.New, []byte(secretKey))
         hasher.Write([]byte(message))
         serverHMACdigest := hasher.Sum(nil)
         fmt.Println("Server : ", base64.StdEncoding.EncodeToString(serverHMACdigest))

         // receive hmacDigest from client

         buffer := make([]byte, 4096)
         n, err := clientConn.Read(buffer)
         if err != nil || n == 0 {
                 clientConn.Close()
                 return
         }


         // don't over read n length
         clientHMACdigest := buffer[:n]

         fmt.Println("Client : ", base64.StdEncoding.EncodeToString(clientHMACdigest))

         // compare if the server and client HMAC digests are the same or not
         // the HANDSHAKING part!
         fmt.Println("Connection authenticated : ", hmac.Equal(serverHMACdigest, clientHMACdigest))

         // this is where you want to do stuff like disconnect client if the authentication failed
         // or proceed
 }

 func handleConnection(c net.Conn) {

         log.Printf("Client %v connected.", c.RemoteAddr())

         serverSideAuthenticate(c, secret)

         log.Printf("Connection from %v closed.", c.RemoteAddr())
 }

 func main() {
         ln, err := net.Listen("tcp", ":6000")
         if err != nil {
                 log.Fatal(err)
         }

         fmt.Println("Server up and listening on port 6000")

         for {
                 conn, err := ln.Accept()
                 if err != nil {
                         log.Println(err)
                         continue
                 }
                 go handleConnection(conn)
         }
 }