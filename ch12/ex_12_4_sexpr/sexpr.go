package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

// Marshal encodePretty Go value to S-expression.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encodePretty(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encodePretty(buf *bytes.Buffer, v reflect.Value, level int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("t")
		} else {
			buf.WriteString("nil")
		}
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		buf.WriteString(strconv.FormatFloat(v.Float(), 'f', -1, 64))
	case reflect.Complex64, reflect.Complex128:
		buf.WriteString("#C(")
		if err := encodePretty(buf, reflect.ValueOf(real(v.Complex())), level); err != nil {
			return err
		}
		buf.WriteByte(' ')
		if err := encodePretty(buf, reflect.ValueOf(imag(v.Complex())), level); err != nil {
			return err
		}
		buf.WriteByte(')')
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return encodePretty(buf, v.Elem(), level)
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encodePretty(buf, v.Index(i), level); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				fmt.Fprintf(buf, "\n%*s", level+1, "")
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			offset := len(v.Type().Field(i).Name) + 3
			if err := encodePretty(buf, v.Field(i), level+offset); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Interface:
		buf.WriteString(`("`)
		buf.WriteString(v.Elem().Type().String())
		buf.WriteString(`" `)
		if err := encodePretty(buf, v.Elem(), level); err != nil {
			return err
		}
		buf.WriteByte(')')
	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i != 0 {
				fmt.Fprintf(buf, "\n%*s", level+1, "")
			}
			buf.WriteByte('(')
			if err := encodePretty(buf, key, level); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encodePretty(buf, v.MapIndex(key), level); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
