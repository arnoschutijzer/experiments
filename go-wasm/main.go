//go:build js && wasm

package main

import (
	"encoding/json"
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

	js.Global().Set("add", js.FuncOf(addFunction))
	js.Global().Set("getJohn", js.FuncOf(getJohn))

	// Prevent from exiting
	select {}
}

func addFunction(this js.Value, p []js.Value) interface{} {
	sum := p[0].Int() + p[1].Int()
	return js.ValueOf(sum)
}

type Person struct {
	Name string `json:"name"`
}

func getJohn(this js.Value, p []js.Value) interface{} {
	objJSON, err := json.Marshal(&Person{Name: "John"})
	if err != nil {
		return err.Error()
	}

	return js.ValueOf(string(objJSON))
}
