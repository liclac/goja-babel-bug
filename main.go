package main

import (
	"fmt"
	"github.com/dop251/goja"
	"io/ioutil"
	"log"
)

func main() {
	babelsrc, err := ioutil.ReadFile("babel.js")
	if err != nil {
		log.Fatalf("Couldn't read babel: %s", err)
	}

	src, err := ioutil.ReadFile("script.js")
	if err != nil {
		log.Fatalf("Couldn't load script.js: %s", err)
	}

	vm := goja.New()
	if _, err := vm.RunScript("babel.js", string(babelsrc)); err != nil {
		log.Fatalf("Couldn't run babel: %s", err)
	}

	this := vm.Get("Babel")
	obj := this.ToObject(vm)

	var transform goja.Callable
	if err := vm.ExportTo(obj.Get("transform"), &transform); err != nil {
		log.Fatalf("Couldn't export transform: %s", err)
	}

	opts := map[string]interface{}{
		"presets": []string{"latest"},
	}
	v, err := transform(this, vm.ToValue(string(src)), vm.ToValue(opts))
	if err != nil {
		log.Fatalf("Couldn't transform: %s", err)
	}
	fmt.Println(v.ToObject(vm).Get("code").Export())
}
