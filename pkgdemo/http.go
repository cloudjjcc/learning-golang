package pkgdemo

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func runHttpServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		content, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(content)
	})
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}
	if err := http.Serve(listener, mux); err != nil {
		panic(err)
	}
}

func clientCall() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:7777", strings.NewReader("hello"))
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
