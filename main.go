package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/rkojedzinszky/mac2vendor/radix"
)

// Mac2VendorServer serves requests
type Mac2VendorServer struct {
	tree    radix.Readonly
	macre   *regexp.Regexp
	stripre *regexp.Regexp
}

func (s *Mac2VendorServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mac := strings.ToLower(s.macre.FindString(s.stripre.ReplaceAllString(r.URL.Path, "")))

	ret := s.tree.Get(mac)
	if ret != nil {
		q := r.URL.Query()
		var format string

		if format = q.Get("format"); format == "" {
			// Default format is plain
			format = "plain"
		}

		if f, ok := formatters[format]; ok {
			f(ret.(*OUIDescr), w)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	t, err := readOuiFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	mac2vendor := &Mac2VendorServer{
		tree:    t,
		macre:   regexp.MustCompile("[0-9a-fA-F]+"),
		stripre: regexp.MustCompile("[:.-]"),
	}

	http.ListenAndServe(":3000", mac2vendor)
}
