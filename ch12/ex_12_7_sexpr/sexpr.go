package sexpr

import "io"

type Decoder struct {
	r io.Reader
}

func (d Decoder) Decode(v interface{}) interface{} {

}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}
