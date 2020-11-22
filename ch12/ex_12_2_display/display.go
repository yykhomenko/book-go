package display

import (
	"fmt"
	"reflect"

	"github.com/yykhomenko/book-gopl/ch12/format"
)

const maxDepth = 10

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), maxDepth)
}

func display(path string, v reflect.Value, depth int) {
	if depth == 0 {
		fmt.Printf("%s = %s\n", path, format.FormatAtom(v))
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), depth-1)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), depth-1)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, format.FormatAtom(key)), v.MapIndex(key), depth-1)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), depth-1)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), depth-1)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, format.FormatAtom(v))
	}
}
