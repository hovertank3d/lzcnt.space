package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"embed"
)

/*
#include <stdint.h>
extern unsigned long int log2lzcnt(unsigned long int);
*/
import "C"

//go:embed templates
var templates embed.FS

func main() {
	counter := atomic.Uint64{}
	host := ":8080"
	if len(os.Args) == 2 {
		host = os.Args[1]
	}

	tmpl, err := template.ParseFS(templates, "templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		counter.Store(0)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := counter.Add(1)
		payload := struct {
			Requests  uint64
			Log2lzcnt uint64
		}{
			Requests:  n,
			Log2lzcnt: uint64(C.log2lzcnt(C.uint64_t(n))),
		}

		tmpl.Execute(w, payload)
	})

	http.ListenAndServe(host, nil)
}
