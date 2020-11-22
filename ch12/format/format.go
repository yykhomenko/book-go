package format

import (
	"bytes"
	"reflect"
	"strconv"
)

// Any formats any value to string.
func Any(v interface{}) string {
	return FormatAtom(reflect.ValueOf(v))
}

func FormatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ex12.1
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
		// ex12.1
	case reflect.Struct:
		var buf bytes.Buffer
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				buf.WriteByte(' ')
			}
			buf.WriteString(FormatAtom(v.Field(i)))
		}
		buf.WriteByte('}')
		return buf.String()
		// ex12.1
	case reflect.Array:
		var buf bytes.Buffer
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				buf.WriteByte(' ')
			}
			buf.WriteString(FormatAtom(v.Index(i)))
		}
		buf.WriteByte(']')
		return buf.String()
	default:
		return v.Type().String() + " value"
	}
}
