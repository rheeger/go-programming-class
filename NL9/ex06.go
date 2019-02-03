package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Arch used:", runtime.GOARCH)
	fmt.Println("OS used:", runtime.GOOS)

}
