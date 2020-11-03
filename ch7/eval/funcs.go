package eval

import "fmt"

func Filter(e Expr, f func(Expr) bool) (out []Expr) {
	if f != nil && f(e) {
		out = append(out, e)
	}
	switch e.(type) {
	case Var, literal:
	case Unary:
		out = append(out, Filter(e.(Unary).x, f)...)
	case Binary:
		out = append(out, Filter(e.(Binary).x, f)...)
		out = append(out, Filter(e.(Binary).y, f)...)
	case Call:
		for _, arg := range e.(Call).args {
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
