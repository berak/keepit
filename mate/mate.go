package main

import (
    "bufio"
    "fmt"
    "image"
    "image/color"
    "image/jpeg"
    "image/png"
    "net"
    "os"
    "strings"
)


type Matrix struct {
    b []uint8
}

func createMatrix() *Matrix {
    var m Matrix
    m.b = make([]uint8, 1923)
    return &m
}

func decode(filename string) (image.Image, string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, "no image", err
    }
    defer f.Close()
    rd := bufio.NewReader(f)
    return image.Decode(rd)
}

func (ih Matrix) Bytes() int {
    return 1920;
}

func (ih Matrix) ColorModel() color.Model {
    return color.NRGBAModel
}
func (ih Matrix) Bounds() image.Rectangle {
    return image.Rect(0, 0, 40, 16)
}
func (ih Matrix) At(x, y int) color.Color {
    p := 3 * (y * 40 + x)
    r := ih.b[p]
    g := ih.b[p+1]
    b := ih.b[p+2]
    return color.NRGBA{r,g,b,0xff}
}

func(ih Matrix) WriteImg(fname string) (*Matrix) {
    f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Printf("Couldn't open %s (%s)\n", fname, err)
        return &ih
    }
    defer f.Close()
    if strings.Contains(fname,"png") {
        err = png.Encode(f, ih)
    } else { //if strings.Contains(fname,"jpg")
        err = jpeg.Encode(f, ih, &jpeg.Options{100})
    }
    if err != nil {
        fmt.Printf("Couldn't encode png (%s)\n", err)
    }
    return &ih
}

func Blank() (*Matrix) {
    return createMatrix()
}
func Dots() (*Matrix) {
    var im = createMatrix()
    for r:=0; r<1920; r++ {
        im.b[r] = 10
        if r % 16 == 0  {
            im.b[r] = 255
        }
    }
    return im
}

func Logo(z string) (*Matrix) {
    z = strings.Replace(z,"\r","",-1)
    z = strings.Replace(z,"\n","",-1)
    var im = createMatrix()
    if 3*(len(z)) != im.Bytes() {
        fmt.Println(3*len(z), " != ", im.Bytes())
        return im
    }
    //fmt.Printf("%d %d\n", im.bytes()/3, len(z) )
    for r:=0; r<1920; r+=3 {
        im.b[r] = 10; im.b[r+1] = 10;  im.b[r+2] = 10
        if z[r/3] == '#' { im.b[r] = 255; im.b[r+1] = 125;  im.b[r+2] = 125 }
        if z[r/3] == 'm' { im.b[r] = 255; im.b[r+1] = 225;  im.b[r+2] =  25 }
        if z[r/3] == 'u' { im.b[r] =  25; im.b[r+1] = 225;  im.b[r+2] = 125 }
        if z[r/3] == '.' { im.b[r] =  30; im.b[r+1] =  30;  im.b[r+2] =  30 }
        if z[r/3] == 'r' { im.b[r] = 255; im.b[r+1] =  10;  im.b[r+2] =  10 }
        if z[r/3] == 'g' { im.b[r] =  10; im.b[r+1] = 255;  im.b[r+2] =  10 }
        if z[r/3] == 'b' { im.b[r] =  10; im.b[r+1] =  10;  im.b[r+2] = 255 }
    }
    return im
}


func CrapSend(host string, bytes []byte) string {
    ServerAddr,err := net.ResolveUDPAddr("udp",host)
    if err != nil { return fmt.Sprintf("no route to %s %s", host, err) }

    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    if err != nil { return "no local socket" }

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    if err != nil { return "dialup fail" }
    defer Conn.Close()

    n,err := Conn.Write(bytes)
    if err != nil { return fmt.Sprintf("write fail (%d bytes)",n) }

    return "ok"
}


func CrapServer(cb func ([]byte) bool) {
    ServerAddr,err := net.ResolveUDPAddr("udp",":1337")
    if err != nil {
        fmt.Println("resolve :1337 failed",err)
        return
    }

    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    if err != nil {
        fmt.Println("can't listen to", ServerAddr);
        return
    }

    defer ServerConn.Close()

    for {
        buf := make([]byte, 1923)
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ", n, " from ", addr)
        if err != nil || n != 1923 {
            fmt.Println("Error: ",err, " " , n, " != 1923")
            continue
        }
        if cb(buf) == false {
            break
        }
    }
}
