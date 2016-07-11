package main

import (
    "fmt"
    "net"
    "time"
)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

var z1 string =
`...                                  ...
.                                      .
.                                      .
.                                      .
. u    u  uuu  uuuuu uuuu  uuuuu u   u .
. mm  mm m   m   m   m   m   m    m m  .
. m mm m mmmmm   m   mmmm    m     m   .
. m mm m mmmmm   m   mmmm    m     m   .
. m    m m   m   m   m  m    m    m m  .
. m    m m   m   m   m  m    m    m m  .
. m    m m   m   m   m   m mmmmm m   m .
. u    u u   u   u   u   u uuuuu u   u .
.                                      .
.                                      .
.                                      .
...                                  ...
`
var z2 string =
`...                                  ...
.                                      .
.                                      .
.                                      .
. u       uuu     uuu uuu   uuuuu uuu  .
. u      m   m    u uuu u  m      m  uu.
. u     mm   m    m  m  u  m      m   u.
. m    m mmmmm    m  m  m  mu     mm uu.
. m    m m   m    m     m    mu   mu m .
. m    m m   m    m     m  mu     m  m .
. m    m m   m    m     m  mm     m   u.
. uuuu u u   u    u     u   uuuuu u   u.
.                                      .
.                                      .
.                                      .
...                                  ...
`
var z3 string =
`........................................
........................................
........................................
........................................
............b...........................
..............b.........................
........................................
..............bbbb.b...r..r..bbbbb......
..............b..b.b...b..b..b..........
..............b.b..b...b..b..b..........
..............br...b...b..b..bbbb.......
..............b.b..b...b..b..b..........
..............b..b.b...b..b..b..........
..............bbbb.bbb.bbbb..bbbbb......
........................................
........................................
`

func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:1337")
    CheckError(err)

    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)

    defer Conn.Close()
    for i:=0;i<9999999; i++ {

        img := Dots()
        if i%3==0 { img=Logo(z1) }
        if i%4==0 { img=Logo(z3) }
        if i%7==0 { img=Logo(z2) }

        n,err := Conn.Write(img.e)
        if err != nil {
            fmt.Println(n, err)
            break
        }
        time.Sleep(time.Second * 27)
    }
}
