package common

import (
	"runtime"
	"strings"
)

func GetCurrentFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}

	fullFuncName := fn.Name()
	splitName := strings.Split(fullFuncName, ".")
	return splitName[len(splitName)-1]
}
