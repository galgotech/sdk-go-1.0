package util

import (
	"runtime"
)

func WebAssembly() bool {
	return runtime.GOOS == "js" && runtime.GOARCH == "wasm"
}
