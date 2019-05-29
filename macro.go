package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/dop251/goja"
)

// Macro ...
type Macro struct {
	URL  string `hcl:"url"`
	Code string `hcl:"exec"`
	TTL  int64  `hcl:"ttl"`
}

// Exec ...
func (m *Macro) Exec() (interface{}, error) {
	doc, err := goquery.NewDocument(m.URL)
	if err != nil {
		return nil, err
	}

	vm := goja.New()
	vm.Set("document", doc)
	vm.RunString(`
		var exports = {};
		var $ = function(s){
			return document.Find(s)
		};
	`)

	if _, err := vm.RunString(m.Code); err != nil {
		return nil, err
	}

	return vm.Get("exports").ToObject(vm), nil
}
