package main

import (
	"fmt"
	"time"
)

func execMacro(macroKey string) (interface{}, bool, error) {
	macro := configs.Macros[macroKey]
	if nil == macro {
		return nil, false, fmt.Errorf("macro %s cannot be found", macroKey)
	}

	if macro.TTL > 0 {
		cachedResult, found := cacher.Get(macroKey)
		if found && cachedResult != nil {
			return cachedResult, true, nil
		}
	}

	resp, err := macro.Exec()
	if nil != err {
		return nil, false, err
	}

	if macro.TTL > 0 {
		cacher.Set(macroKey, resp, time.Duration(macro.TTL)*time.Second)
	}

	return resp, false, nil
}
