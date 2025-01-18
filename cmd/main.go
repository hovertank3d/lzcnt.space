package main

import (
	"os"

	"github.com/hovertank3d/lzcnt.space"
)

import "C"

func main() {
	host := ":8080"
	if len(os.Args) == 2 {
		host = os.Args[1]
	}

	srv := lzcnt.New()
	srv.ListenAndServe(host)
}
