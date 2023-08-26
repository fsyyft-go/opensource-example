package testing

import (
	"fmt"
)

const (
	logHeader = "=-=       "
)

func Println(a ...interface{}) {
	fmt.Print(logHeader)
	// fmt.Print(a...) 所有参数连在一起。
	// fmt.Println(a...) 参数之间空格分割。
	fmt.Println(a...)
}

func Printf(format string, a ...interface{}) {
	fmt.Print(logHeader)
	fmt.Printf(format, a...)
}
