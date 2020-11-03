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

// Unary represents expression with a Unary operator.
type Unary struct {
	op rune // '+' or '-'
	x  Expr
}

// Binary represents expression with a Binary operator.
type Binary struct {
	op   rune // '+', '-', '*' or '/'
	x, y Expr
}

// Call represents expression function invocation, for example six(x).
type Call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}
