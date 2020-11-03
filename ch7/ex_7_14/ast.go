package eval

type Expr interface {
	Eval(env Env) float64          // Eval returns value of Expr in environment env.
	Check(vars map[Var]bool) error // Check report about errors in Expr and add self Vars.
	String() string
}

// Env contains values of Var
type Env map[Var]float64

// Var is variable, for example x.
type Var string

// literal represents numbering constant, for example 3.14.
type literal float64

// unary represents expression with a unary operator.
type unary struct {
	op rune // '+' or '-'
	x  Expr
}

// binary represents expression with a binary operator.
type binary struct {
	op   rune // '+', '-', '*' or '/'
	x, y Expr
}

// call represents expression function invocation, for example six(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

// min returns minimal one of sequence <a1, a2, ...aN>, for example <1, 2, 3>.
type min struct {
	args []Expr
}
