package main

import (
	"fmt"
	"go.starlark.net/starlark"
)

func main() {
	// Execute Starlark program in a file.
	thread := &starlark.Thread{Name: "my thread"}
	globals, err := starlark.ExecFile(thread, "test/test.star", nil, nil)
	if err != nil {  }

	// Retrieve a module global.
	hello := globals["hello"]

	// Call Starlark function from Go.
	v, err := starlark.Call(thread, hello, nil, nil)
	if err != nil {  }
	fmt.Printf("hello() = %v\n", v)
}
