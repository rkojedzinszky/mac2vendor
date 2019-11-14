package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type formatterFn func(*OUIDescr, http.ResponseWriter)

func plainFormatter(oui *OUIDescr, w http.ResponseWriter) {
	fmt.Fprint(w, oui.Vendor)
}

func jsonFormatter(oui *OUIDescr, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

	txt, _ := json.Marshal(oui)
	fmt.Fprint(w, string(txt))
}

var formatters = map[string]formatterFn{
	"plain": plainFormatter,
	"json":  jsonFormatter,
}
