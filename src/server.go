package main

import (
"net"
"fmt"
"os"
)

func handleConnect(c net.Conn){


}

func main() {
   ln, err := net.ListenUnix("unix", &net.UnixAddr{"/tmp/funix", "unix"})
   if err != nil {
     panic(err)
   }
   defer os.Remove("/tmp/funix")

   for {
       conn, err := ln.AcceptUnix()
       if err != nil {
         panic(err)
       }
      var buf [1024]byte
      n, err := conn.Read(buf[:])
      if err != nil {
        panic(err)
      }
      fmt.Printf("%s\n", string(buf[:n]));
      conn.Close()
   }

}
