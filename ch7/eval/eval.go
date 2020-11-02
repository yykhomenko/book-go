package main

import (
	"fmt"
	"math"
)

type Expr interface {
	// Eval returns value of Expr in environment env.
	Eval(env Env) float64
}

// Env contains values of Var
type Env map[Var]float64

// Var is variable, for example x.
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

// literal represents numbering constant, for example 3.14.
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// unary represents expression with a unary operator.
type unary struct {
	op rune // '+' or '-'
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	default:
		panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
	}
}

// binary represents expression with a binary operator.
type binary struct {
	op   rune // '+', '-', '*' or '/'
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	default:
		panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
	}
}

// call represents expression function invocation, for example six(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	default:
		panic(fmt.Sprintf("unsupported function call: %q", c.fn))
	}
}
