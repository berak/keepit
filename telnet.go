package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

type Telnet struct {
	con     *net.TCPConn
}

func (telnet *Telnet) Connect(host string) error {
	addr, e := net.ResolveTCPAddr("tcp", host)
	if addr == nil || e != nil {
		fmt.Printf("!addr!%s.\n", e)
		return e
	}

	c, e := net.DialTCP("tcp", nil, addr)
	telnet.con = c
	if telnet.con == nil || e != nil {
		fmt.Printf("!con!%s.\n", e)
		return e
	}
	fmt.Printf("connected to %s\r\n", host)
	return nil
}

func (telnet *Telnet) Read() string {
	buf := make([]byte, 8021)
	r, e := telnet.con.Read(buf)
	if r < 1 || e != nil {
		//fmt.Printf("!read!%s.\n", e)
		return ""
	}
	return string(buf[0:r])
}

func (telnet *Telnet) Write(s string) {
	r, e := telnet.con.Write([]byte(s))
	if r < 1 || e != nil {
		//fmt.Printf("!read!%s.\n", e)
	}
}

func (telnet *Telnet) Listen() {
	for {
		s := telnet.Read()
		if s == "" {
			return
		}
		fmt.Printf("%s\r\n", s)
	}
}

func(telnet *Telnet) Console() {
	con := bufio.NewReader(os.Stdin)
	for {
		data, err := con.ReadBytes('\n')
		if err != nil {
			break
		}
		line := string(data)
		//if len(line) == 0 {
		//	continue
		//}
		if line == "." {
			os.Exit(0)
		}
		telnet.Write(line + "\r\n")
	}
}

func main() {
	var telnet Telnet
	// flags have to go BEFORE other args !
	//var raw *bool = flag.Bool("raw", false, "do not parse input")
	flag.Parse()
	var host="127.0.0.1"
	if flag.NArg() > 0 {
		host = flag.Arg(0)
	}
	var port="23"
	if flag.NArg() > 1 {
		port = flag.Arg(1)
	}
	var initialCmd=""
    if flag.NArg() > 2 {
		initialCmd = flag.Arg(2)
	}

	e := telnet.Connect(host + ":" + port)
	if e != nil {
		return
	}

	if (len(initialCmd) > 0) {
		telnet.Write(initialCmd + "\r\n")
	}

	go telnet.Console()
	telnet.Listen()
}
