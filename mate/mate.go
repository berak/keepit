package main

import (
    "bufio"
    "fmt"
    "image"
    "image/color"
    "image/jpeg"
    "image/png"
    "os"
    "strings"
)


type Matrix struct {
    e    []uint8
    w, h, c int
}

func createMatrix(w, h, c int) *Matrix {
    var m Matrix
    m.c = c
    m.w = w
    m.h = h
    m.e = make([]uint8, c*(w*h+1))
    return &m
}

func Decode(filename string) (image.Image, string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, "no image", err
    }
    defer f.Close()
    rd := bufio.NewReader(f)
    return image.Decode(rd)
}

func (ih Matrix) Bytes() int {
    return ih.c * ih.w * ih.h
}

func (ih Matrix) ColorModel() color.Model {
    return color.NRGBAModel
}
func (ih Matrix) Bounds() image.Rectangle {
    return image.Rect(0, 0, ih.w, ih.h)
}
func (ih Matrix) At(x, y int) color.Color {
    p := ih.c*(y * ih.w + x)
    r := ih.e[p]
    g,b := r,r;
    if ih.c>1 {
        g = ih.e[p+1]
    }
    if ih.c>2 {
        b = ih.e[p+2]
    }
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
    return createMatrix(40,16,3)
}
func Dots() (*Matrix) {
    var im = createMatrix(40,16,3)
    for r:=0; r<1920; r++ {
        im.e[r] = 10
        if r % 16 == 0  {
            im.e[r] = 255
        }
    }
    return im
}

func Logo(z string) (*Matrix) {
    z = strings.Replace(z,"\r","",-1)
    z = strings.Replace(z,"\n","",-1)
    var im = createMatrix(40,16,3)
    if 3*(len(z)) != im.Bytes() {
        fmt.Println(3*len(z), " != ", im.Bytes())
        return im
    }
    //fmt.Printf("%d %d\n", im.bytes()/3, len(z) )
    for r:=0; r<1920; r+=3 {
        im.e[r] = 10; im.e[r+1] = 10;  im.e[r+2] = 10
        if z[r/3] == '#' { im.e[r] = 255; im.e[r+1] = 125;  im.e[r+2] = 125 }
        if z[r/3] == 'm' { im.e[r] = 255; im.e[r+1] = 225;  im.e[r+2] =  25 }
        if z[r/3] == 'u' { im.e[r] =  25; im.e[r+1] = 225;  im.e[r+2] = 125 }
        if z[r/3] == '.' { im.e[r] =  30; im.e[r+1] =  30;  im.e[r+2] =  30 }
        if z[r/3] == 'r' { im.e[r] = 255; im.e[r+1] =  10;  im.e[r+2] =  10 }
        if z[r/3] == 'g' { im.e[r] =  10; im.e[r+1] = 255;  im.e[r+2] =  10 }
        if z[r/3] == 'b' { im.e[r] =  10; im.e[r+1] =  10;  im.e[r+2] = 255 }
    }
    return im
}
