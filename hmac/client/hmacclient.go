package main

 import (
         "crypto/hmac"
         "crypto/md5"
         "encoding/base64"
         "fmt"
         "net"
 )

 // both client and server MUST have the same secret key
 // to authenticate

 var secret = "GolangIsAwesome!"
 // change the secret to something else and the authentication will fail
 //var secret = "GolangIsTerrible!"

 func clientSideAuthenticate(serverConn net.Conn, secretKey string, message string) {

         // prepare client side hmac digest
         // with secret key and message received from server
         // if the secret key in client and server is the same
         // the digest should be the same.

         hasher := hmac.New(md5.New, []byte(secretKey))
         hasher.Write([]byte(message))
         clientHMACdigest := hasher.Sum(nil)
         fmt.Println("Digest send to server : ", base64.StdEncoding.EncodeToString(clientHMACdigest))

         // send hmacDigest back to server to authenticate

         n, err := serverConn.Write(clientHMACdigest)
         if err != nil || n == 0 {
                 serverConn.Close()
                 return
         }

 }

 func handleConnection(c net.Conn) {

         buffer := make([]byte, 4096)

         for {
                 n, err := c.Read(buffer)
                 if err != nil || n == 0 {
                         c.Close()
                         break
                 }
                 // don't over read n length
                 msg := string(buffer[:n])

                 fmt.Println("\nData received from server : ", msg)
                 clientSideAuthenticate(c, secret, msg)
         }
         fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
 }

 func main() {
         hostName := "localhost" // change this to your server domain name 
         portNum := "6000"

         for {
                 dialConn, err := net.Dial("tcp", hostName+":"+portNum)

                 if err != nil {
                         fmt.Println(err)
                         continue
                 }

                 fmt.Printf("\nConnection established between %s and localhost.\n", hostName)
                 fmt.Printf("Remote Address : %s \n", dialConn.RemoteAddr().String())
                 fmt.Printf("Local Address : %s \n", dialConn.LocalAddr().String())

                 go handleConnection(dialConn)
         }

 }