package main

// this is a helper executable to hit our
// healthcheck endpoint in a scratch container.

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/health", os.Getenv("HTTPPORT")))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// os.Exit(0)
}
