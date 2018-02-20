package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/intelfike/comet"
	"github.com/intelfike/shufflechat/io/output"
)

var port = flag.String("http", ":8888", "HTTP port number.")

// cookieのキーを定義
var cmt = comet.NewComet("realtimesession")

func init() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			return
		}
		// セッションを開始
		cmt.Start(w, r)
		f, err := os.Open("./index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		io.Copy(w, f)
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		// Bodyを読み取る
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// 全てのチャンネルに通知
		cmt.DoneOther(r, string(b))
	})
	http.HandleFunc("/comet", func(w http.ResponseWriter, r *http.Request) {
		// 待機
		i, err := cmt.Wait(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		switch t := i.(type) {
		case string:
			chatHTML := t
			output.WriteString(w, chatHTML)
		default:
		}
	})
	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		cmt.End(r)
	})
}
func main() {
	fmt.Printf("Start HTTP Server localhost%s\n", *port)
	fmt.Println(http.ListenAndServe(*port, nil))
}
