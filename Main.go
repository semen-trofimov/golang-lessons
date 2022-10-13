package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)
	e := http.ListenAndServe(":9090", nil)
	if e != nil {
		fmt.Println(e.Error())
	}
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	str := "hello"
	n, e := io.WriteString(w, str)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(n, "", len(str))
	}
}
