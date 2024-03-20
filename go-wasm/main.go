//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("button clicked")
		// cb.Release() // release the function if the button will not be clicked again
		// os.Exit(0) // cleans everything up?
		return nil
	})
	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)

	// Prevent from exiting
	select {}
}
