package json

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Struct:
		buf.WriteByte('{')

		buf.WriteByte('}')
	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')

		buf.WriteByte(']')
	default:
		return fmt.Errorf("unsupported type: %q", v.Kind())
	}
	return nil
}
