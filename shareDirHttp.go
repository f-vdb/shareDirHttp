package main

import (
	"fmt"
	"net/http"
)

func main() {

	var (
		dir  string
		port string
	)

	fmt.Printf("Please input which directory you want to share,\n start with \"/\":\n")
	fmt.Scanf("%s", &dir)
	handle := http.FileServer(http.Dir(dir))
	fmt.Printf("Please input port Number: \n")
	fmt.Scanf("%s", &port)
	http.ListenAndServe(":"+port, handle)
}
