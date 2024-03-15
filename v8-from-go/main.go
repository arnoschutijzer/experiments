package main

import (
	"fmt"

	v8 "rogchap.com/v8go"
)

func main() {
	ctx := v8.NewContext()
	defer ctx.Isolate().Dispose()
	defer ctx.Close()

	origin := "main.js"

	ctx.RunScript("const x = () => 123", origin)
	val, err := ctx.RunScript("x()", origin)
	fmt.Println(err, val)
}
