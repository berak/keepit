package main

import (
    "fmt"
    "time"
)

func main() {
    server := "127.0.0.1:1337"
    z := [3]string{z1,z2,z3}

    for i:=0;i<9999999; i++ {
        txt := z[i%3]
        fmt.Println(i)
        fmt.Println(txt)
        img := Logo(txt)
        ok := CrapSend(server, img.b)
        if ok != "ok" {
            fmt.Println(i, ok)
            break
        }
        time.Sleep(time.Second * 27)
    }
}


var z1 =
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
...                                  ...`
var z2 =
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
...                                  ...`
var z3 =
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
