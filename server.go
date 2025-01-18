package lzcnt

import (
	"embed"
	"html/template"
	"net/http"
	"sync/atomic"
)

/*
#include <stdint.h>
extern unsigned long int log2lzcnt(unsigned long int);
*/
import "C"

var (
	//go:embed templates
	templates embed.FS

	tmpl = template.New("")
)

func init() {
	tmpl = tmpl.Funcs(template.FuncMap{
		"arr": func(els ...any) []any {
			return els
		},
	})

	tmpl = template.Must(tmpl.ParseFS(templates, "templates/*.html"))
}

type Server struct {
	mux            *http.ServeMux
	requestCounter atomic.Uint64
}

func New() *Server {
	srv := &Server{
		mux: http.NewServeMux(),
	}

	srv.mux.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		srv.requestCounter.Store(0)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	srv.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := srv.requestCounter.Add(1)
		payload := map[string]any{
			"Page":      "log2.s.html",
			"Title":     "Count the Number of Leading Zero Bits",
			"Requests":  n,
			"Log2lzcnt": uint64(C.log2lzcnt(C.uint64_t(n))),
		}

		tmpl.ExecuteTemplate(w, "index.html", payload)
	})

	srv.mux.HandleFunc("/about", srv.aboutPage)

	return srv
}

func (s *Server) ListenAndServe(host string) error {
	return http.ListenAndServe(host, s.mux)
}
