//  i made this site to serve some of my future demos and fuuny
//  projects. currently there are only couple of interesting things i wrote:

// * {{ template "link.html" (arr "46load" "https://github.com/hovertank3d/46load")}},
// * {{ template "link.html" (arr "fire" "https://github.com/hovertank3d/fire")}},
// * {{ template "link.html" (arr "bfc" "https://github.com/hovertank3d/bfc")}},
// * {{ template "link.html" (arr "enigma16" "https://github.com/hovertank3d/enigma16")}},

// so, come back when i'll write something

// also, this site uses {{ template "link.html" (arr "genhl" "https://github.com/hovertank3d/genhl")}} to go:generate syntax highlight

package lzcnt

import "net/http"

func (s *Server) aboutPage(w http.ResponseWriter, r *http.Request) {
	payload := map[string]any{
		"Page":  "about.go.html",
		"Title": "About",
	}

	tmpl.ExecuteTemplate(w, "index.html", payload)
}
