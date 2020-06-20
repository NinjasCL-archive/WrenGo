// +build ignore

// Meant to be run from "go generate"
//
// Creates Go functions that are visible from C
// So that WrenGo can bind Go functions to Wren
package main

import (
	"os"
	// "io/ioutil"
	// "strconv"
	"flag"
	"text/template"
)

var fileTemplate = template.Must(template.New("").Parse(
`// Code generated by go generate; DO NOT EDIT.

package wren

/*
#cgo CFLAGS:
#cgo LDFLAGS: -lm
#include "wren.h"

{{range .}}extern void f{{.}}(WrenVM* vm);
{{end}}
static inline WrenForeignMethodFn get_f(int i) {
	switch (i) {
		{{range .}}case {{.}}: return f{{.}};
		{{end}}default: return (void*)(0);
	}
}
*/
import "C"
import (
	"fmt"
)

const MAX_REGISTRATIONS = {{len .}}

{{ range . -}}
//export f{{.}}
func f{{.}}(v *C.WrenVM) {
	if vm, ok := vmMap[v]; ok {
		if len(vm.bindMap) > {{.}} {
			vm.bindMap[{{.}}](vm)
		}
	}
}

{{end -}}
type MaxBindingsReached struct {
	VM *VM
}

func (err *MaxBindingsReached)Error() string {
	return fmt.Sprintf("Cannot bind more that %v functions or classes", MAX_REGISTRATIONS)
}

func (vm *VM) registerFunc(fn ForeignMethodFn) error {
	if len(vm.bindMap) >= MAX_REGISTRATIONS {
		return &MaxBindingsReached{VM: vm}
	}

	vm.bindMap = append(vm.bindMap, fn)
	return nil
}
`))

func main() {
	var count int
	flag.IntVar(&count, "bindings", 0, "Amount of bindings to create for wrenGo (bindings should be greater than 0)")
	flag.Parse()
	if count <= 0 {
		flag.Usage()
		return
	}

	f, err := os.Create("bindings.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = i
	}
	if err := fileTemplate.Execute(f, data); err != nil {
		panic(err)
	}
}
