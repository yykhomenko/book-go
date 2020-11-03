package eval

import "fmt"

func Filter(e Expr, f func(Expr) bool) (out []Expr) {
	if f != nil && f(e) {
		out = append(out, e)
	}
	switch e := e.(type) {
	case Var, literal:
	case Unary:
		out = append(out, Filter(e.x, f)...)
	case Binary:
		out = append(out, Filter(e.x, f)...)
		out = append(out, Filter(e.y, f)...)
	case Call:
		for _, arg := range e.args {
			out = append(out, Filter(arg, f)...)
		}
	default:
		panic(fmt.Errorf("unknown expr type: %v", e))
	}
	return
}

var FilterVars = func(expr Expr) bool {
	_, ok := expr.(Var)
	return ok
}
