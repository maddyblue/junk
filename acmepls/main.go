package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

//go:generate yarn --cwd frontend build
//go:generate esc -o static.go -prefix frontend/build -ignore \.map frontend/build

var (
	flagAddr = flag.String("addr", ":8041", "HTTP address")
	//	flagLSP  = flag.String("lsp", "go:gopls,js:javascript-typescript-langserver", "LSP server specs")
	flagLSP = flag.String("lsp", "go:gopls", "LSP server specs")
	flagDev = flag.Bool("dev", false, "dev mode")
)

func main() {
	flag.Parse()

	p := NewPLS()
	for _, spec := range strings.Split(*flagLSP, ",") {
		sp := strings.Split(spec, ":")
		if len(sp) != 2 {
			log.Fatal("bad spec")
		}
		if err := p.AddLSP(sp[0], sp[1]); err != nil {
			panic(err)
		}
	}

	http.Handle("/", http.FileServer(FS(*flagDev)))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		u := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := u.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		p.WS(conn)
	})
	http.HandleFunc("/api/command", wrap(p.Command))
	http.HandleFunc("/api/open", func(w http.ResponseWriter, r *http.Request) {
		if n := r.FormValue("name"); n == "" || n == "undefined" {
			log.Println("OPEN PROB", r.Form)
			return
		}
		addr := fmt.Sprintf("%s:%s:%s", r.FormValue("name"), r.FormValue("line"), r.FormValue("char"))
		open([]byte(addr))
	})
	fmt.Println("listen on", *flagAddr)
	log.Fatal(http.ListenAndServe(*flagAddr, nil))
}

func wrap(f func(r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := f(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if res == nil {
			return
		}
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
