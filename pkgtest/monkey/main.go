package main

import (
	"bou.ke/monkey"
	"fmt"
	"time"
	_ "unsafe"
)

//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)

func Now() time.Time {
	sec, nsec, _ := now()
	return time.Unix(sec+60, int64(nsec))
}

func main() {
	patch()
}

func patch() {
	fmt.Println(time.Now().String())
	monkey.Patch(time.Now, Now)
	fmt.Println(time.Now().String())
}
