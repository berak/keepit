package main

import (
	"fmt"
	"net/http"
	"os"
)


func main() {
	go CrapServer(func (buf []byte) bool {
	    var img Matrix
	    img.b = buf
		img.WriteImg(img_name)
		return true
	})

	port := os.Getenv("OPENSHIFT_GO_PORT")
	if port == "" { port="8000"; }
	host := os.Getenv("OPENSHIFT_GO_IP")
	if host == "" { host="0.0.0.0"; }

	http.HandleFunc("/", hello)
	http.HandleFunc("/omg", omg)
	http.HandleFunc("/txt", txt_frm)
	http.HandleFunc("/txt/up", txt_up)

	bind := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}


func hello(res http.ResponseWriter, req *http.Request) {
    var front string = head + bar + "\r\n"
    front += "<img src=\"/omg\" width=640 height=360>"
	fmt.Fprintf(res, front)
}


func edit(pat string) string {
	return head + bar + `
		<form action="/txt/up" method="POST">
		<div><textarea name="body" rows="16" cols="40">`+pat+`</textarea></div>
		<div><input type="submit" value="Save"></div>
		</form>`
}


func txt_frm(res http.ResponseWriter, req *http.Request) {
	t := def_txt
	fmt.Fprintln(res, edit(t))
}


func txt_up(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	body := req.FormValue("body")
	img := Logo(body)
	img.WriteImg(img_name)
    http.Redirect(res, req, "/", http.StatusFound)
}

func omg(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open(img_name)
    if err != nil {
	    fmt.Fprintf(res, "404, sorry, %s\r\n", err)
        return
    }
    defer f.Close()
	buf := make([]byte, 8021)
    n, err := f.Read(buf)
	if err == nil {
		body := buf[0:n]
		res.Write(body)
	} else {
		fmt.Fprintf(res, "sorry, %s\r\n", err)
	}
}


var img_name = "omg.png"
var head string = `
<head><style>
 body,textarea,div,a,input,button{ color:#6c6; background-color:#333; border-style:solid; border:0; }
 body,div,a{ margin:15 15 5 5; text-decoration:none;}
 img{ margin:15 15 10 10; text-decoration:none;}
 div{ border:1; border-color:#393; }
</style></head>`
var bar string = `<div><a href="/">omg</a>&nbsp;<a href="/txt">new</a>&nbsp;<a href="javascript:history.back();">back</a></div>`
var def_txt string = `
........................................
........................................
........................................
........................................
........................................
........................................
........................................
........................................
........................................
........................................
........................................
........................................
........................................
........................................
........................................
........................................
`
