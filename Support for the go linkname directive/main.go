//go:build unix

package main

import (
	"fmt"
	_ "time"
	_ "unsafe"
)

func main() {
	fmt.Println(now())
}

//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)
