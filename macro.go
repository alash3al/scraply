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
func (m *Macro) Exec(params map[string]interface{}) (interface{}, error) {
	if m.URL != "" {
		return m.execURL(params)
	}

	return m.execCodeOnly(params)
}

// execCodeOnly ...
func (m *Macro) execCodeOnly(params map[string]interface{}) (interface{}, error) {
	return m.execJS(map[string]interface{}{
		"scraply": map[string]interface{}{
			"params": params,
		},
	})
}

// Exec ...
func (m *Macro) execURL(params map[string]interface{}) (interface{}, error) {
	doc, err := goquery.NewDocument(m.URL)
	if err != nil {
		return nil, err
	}

	return m.execJS(map[string]interface{}{
		"document": doc,
		"scraply": map[string]interface{}{
			"params": params,
		},
	})
}

func (m *Macro) execJS(ctx map[string]interface{}) (interface{}, error) {
	vm := goja.New()

	for k, v := range ctx {
		vm.Set(k, v)
	}

	vm.Set("println", fmt.Println)
	vm.Set("fetch", goquery.NewDocument)
	vm.Set("time", func() int64 {
		return time.Now().Unix()
	})

	vm.Set("macro", func(macroName string, params map[string]interface{}) (interface{}, error) {
		m := m.config.Macros[macroName]

		if m == nil {
			return nil, errors.New(macroName + " : macro not found")
		}

		return m.Exec(params)
	})

	vm.Set("sleep", func(ms int) {
		time.Sleep(time.Duration(ms) * time.Millisecond)
	})

	vm.RunString(`
		var console = {log: println};
		var exports = {};
		var $ = function(s){
			if ( ! document ) {
				throw "none document context";
			}
			return document.Find(s);
		};

		scraply.macro = macro;
	`)

	if _, err := vm.RunString(m.Code); err != nil {
		return nil, err
	}

	return vm.Get("exports").ToObject(vm).Export(), nil
}
