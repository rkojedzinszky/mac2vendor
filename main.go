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
	mac := strings.ToLower(s.stripre.ReplaceAllString(s.macre.FindString(r.RequestURI), ""))

	ret := s.tree.Get(mac)
	if ret != nil {
		fmt.Fprint(w, ret.(*OUIDescr).vendor)
	} else {
		w.WriteHeader(404)
	}
}

func main() {
	t, err := readOuiFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	mac2vendor := &Mac2VendorServer{
		tree:    t,
		macre:   regexp.MustCompile("[0-9a-fA-F:-]+"),
		stripre: regexp.MustCompile("[:-]"),
	}

	http.ListenAndServe(":3000", mac2vendor)
}
