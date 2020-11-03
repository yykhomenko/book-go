package xmltree

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

type Node interface{}
type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	var b bytes.Buffer
	visit(e, &b, 0)
	return b.String()[1:]
}

func visit(n Node, w io.Writer, depth int) {
	switch n := n.(type) {
	case CharData:
		if n != "\n" {
			fmt.Fprintf(w, ": %q", n)
		}
	case *Element:
		var b bytes.Buffer
		if len(n.Attr) != 0 {
			b.WriteRune('(')
			for i, attr := range n.Attr {
				if i != 0 {
					b.WriteRune(' ')
				}
				b.WriteString(attr.Name.Local)
				b.WriteString(`="`)
				b.WriteString(attr.Value)
				b.WriteRune('"')
			}
			b.WriteRune(')')
		}

		fmt.Fprintf(w, "\n%*s%s%s", depth*2, "", n.Type.Local, b.String())
		for _, c := range n.Children {
			visit(c, w, depth+1)
		}
	}
}

func XMLTree(r io.Reader) (root Node, _ error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			el := &Element{tok.Name, tok.Attr, nil}
			if len(stack) == 0 {
				root = el
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, el)
			}
			stack = append(stack, el)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}
	return
}
