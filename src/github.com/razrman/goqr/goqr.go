package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address") // Q=17, R=18

var templQr = template.Must(template.New("qr").Parse(templateStrQr))
var templIm = template.Must(template.New("im").Parse(templateStrIm))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		    http.ServeFile(w, r, r.URL.Path[1:])
	    })
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templQr.Execute(w, req.FormValue("s"))
}

func Image(w http.ResponseWriter, req *http.Request) {
	templIm.Execute(w, req.FormValue("s"))
}


const templateStrQr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`

const templateStrIm = `
<html>
<head>
<title>Go draw example output</title>
</head>
<body>
<br>
<img src="static/out.png" />
<br>
</body>
</html>
`
