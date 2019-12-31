package main

import (
	"errors"
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
	config   *Config
}

// Exec ...
func (m *Macro) Exec() (interface{}, error) {
	if m.URL != "" {
		return m.execURL()
	}

	return m.execCodeOnly()
}

// execCodeOnly ...
func (m *Macro) execCodeOnly() (interface{}, error) {
	return m.execJS(nil)
}

// Exec ...
func (m *Macro) execURL() (interface{}, error) {
	doc, err := goquery.NewDocument(m.URL)
	if err != nil {
		return nil, err
	}

	return m.execJS(map[string]interface{}{"document": doc})
}

func (m *Macro) execJS(ctx map[string]interface{}) (interface{}, error) {
	vm := goja.New()

	vm.Set("println", fmt.Println)
	vm.Set("fetch", goquery.NewDocument)
	vm.Set("time", func() int64 {
		return time.Now().Unix()
	})

	vm.Set("macro", func(macroName string) (interface{}, error) {
		m := m.config.Macros[macroName]

		if m == nil {
			return nil, errors.New(macroName + " : macro not found")
		}

		return m.Exec()
	})

	vm.Set("sleep", func(ms int) {
		time.Sleep(time.Duration(ms) * time.Millisecond)
	})

	for k, v := range ctx {
		vm.Set(k, v)
	}

	vm.RunString(`
		var console = {log: println};
		var exports = {};
		var $ = function(s){
			if ( ! document ) {
				throw "none document context";
			}
			return document.Find(s);
		};
	`)

	if _, err := vm.RunString(m.Code); err != nil {
		return nil, err
	}

	return vm.Get("exports").ToObject(vm).Export(), nil
}
