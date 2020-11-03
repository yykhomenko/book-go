package eval

import "fmt"

func Filter(e Expr, f func(Expr) bool) (out []Expr) {
	if f != nil && f(e) {
		out = append(out, e)
	}
	switch e.(type) {
	case Var, literal:
	case Unary:
		exp := e.(Unary)
		out = append(out, Filter(exp.x, f)...)
	case Binary:
		exp := e.(Binary)
		out = append(out, Filter(exp.x, f)...)
		out = append(out, Filter(exp.y, f)...)
	case Call:
		exp := e.(Call)
		for _, arg := range exp.args {
			out = append(out, Filter(arg, f)...)
		}
	default:
		panic(fmt.Errorf("unknown expr type: %v", e))
	}
	return
}

var FilterVars = func(expr Expr) bool {
	switch expr.(type) {
	case Var:
		return true
	default:
		return false
	}
}
