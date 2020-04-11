package runtime

import (
	"fmt"
	"runtime"
)

func Display() {
	fmt.Println("Go OS", runtime.GOOS)
	fmt.Println("Go OS", runtime.GOARCH)
}
