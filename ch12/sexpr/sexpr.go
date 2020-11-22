package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Printf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Printf("%d", v.Uint())
	case reflect.String:
		fmt.Printf("%q", v.String())
	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				buf.WriteByte(' ')
			}
			buf.WriteString(FormatAtom(v.Field(i)))
		}
		buf.WriteByte(')')

	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Chan, reflect.Func,
		reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
		// ex12.1

		// ex12.1

	default:
		return v.Type().String() + " value"
	}
}
