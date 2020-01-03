package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func execMacro(macroKey string, params map[string]interface{}) (interface{}, bool, error) {
	macro := configs.Macros[macroKey]
	if nil == macro {
		return nil, false, fmt.Errorf("macro %s cannot be found", macroKey)
	}

	jsnBin, _ := json.Marshal(params)
	jsnStr := string(jsnBin)

	if macro.TTL > 0 {
		cachedResult, found := cacher.Get(macroKey + jsnStr)
		if found && cachedResult != nil {
			return cachedResult, true, nil
		}
	}

	resp, err := macro.Exec(params)
	if nil != err {
		return nil, false, err
	}

	if macro.TTL > 0 {
		cacher.Set(macroKey+jsnStr, resp, time.Duration(macro.TTL)*time.Second)
	}

	return resp, false, nil
}
