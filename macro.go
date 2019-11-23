package main

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/dop251/goja"
)

// Macro ...
type Macro struct {
	URL      string `hcl:"url"`
	Code     string `hcl:"exec"`
	TTL      int64  `hcl:"ttl"`
	Schedule string `hcl:"schedule"`
	Webhook  string `hcl:"webhook"`
	Private  bool   `hcl:"private"`
}

// Exec ...
func (m *Macro) Exec() (interface{}, error) {
	doc, err := goquery.NewDocument(m.URL)
	if err != nil {
		return nil, err
	}

	vm := goja.New()

	vm.Set("println", fmt.Println)
	vm.Set("fetch", goquery.NewDocument)
	vm.Set("document", doc)
	vm.Set("time", func() int64 {
		return time.Now().Unix()
	})
	vm.Set("sleep", func(ms int) {
		time.Sleep(time.Duration(ms) * time.Millisecond)
	})

	vm.RunString(`
		var console = {log: println};
		var exports = {};
		var $ = function(s){
			return document.Find(s);
		};
	`)

	if _, err := vm.RunString(m.Code); err != nil {
		return nil, err
	}

	return vm.Get("exports").ToObject(vm).Export(), nil
}
