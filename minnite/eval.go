 package minnite

import "fmt"

var cn Value

func (p *Program) Eval(ctx *Context) Value {
result := NewVoid()

	for _, stmt := range p.Statements {
		result = stmt.Eval(ctx)
	}

	return result
}

func (s *Statement) Eval(ctx *Context) Value {
	switch {
	case s.Let != nil:
		return s.Let.Eval(ctx)
	case s.Print != nil:
		return s.Print.Eval(ctx)
	case s.If != nil:
		return s.If.Eval(ctx)
  case s.For != nil:
    return s.For.Eval(ctx)
  case s.Return != nil:
		return s.Return.Eval(ctx)
  case s.Switch != nil:
		return s.Switch.Eval(ctx)
  case s.Case != nil:
		return s.Case.Eval(ctx)
  case s.Summon != nil:
		return s.Summon.Eval(ctx)
	}
  
	panic("unreachable")
}

func (s *BlockStatement) Eval(ctx *Context) Value {
	result := NewVoid()

	for _, stmt := range s.Body {
		result = stmt.Eval(ctx)

		if stmt.Return != nil {
			break
		}
	}

	return result
}

func (s *LetStatement) Eval(ctx *Context) Value {
	ctx.AddVariable(s.Variable, s.Value.Eval(ctx))
	return NewVoid()
}

func (s *PrintStatement) Eval(ctx *Context) Value {
	fmt.Println(s.Value.Eval(ctx))
  return NewVoid()
}

func (s *IfStatement) Eval(ctx *Context) Value {
	cond := s.Cond.Eval(ctx).(Boolean)

	if cond {
		s.Con.Eval(ctx)
	} else {
		s.Alt.Eval(ctx)
	}

	return NewVoid()
}

func (s *SwitchStatement) Eval(ctx *Context) Value {
  cn=s.Va.Eval(ctx)
	return NewVoid()
}

func (s *CaseStatement) Eval(ctx *Context) Value {
  if s.CN.Eval(ctx)==cn{
    s.Num.Eval(ctx)
  }
	return NewVoid()
}

func (s *ForStatement) Eval(ctx *Context) Value {
	crit := s.Crit.Eval(ctx).(Boolean)

for s.Init.Eval(ctx);crit;s.Chan.Eval(ctx){
  s.Proc.Eval(ctx)
  if crit==false{
break
  }
}
	return NewVoid()
}

func (s *SummonStatement) Eval(ctx *Context) Value {
  fmt.Println(s.Summon.Eval(ctx))
	return NewVoid()
}

func (s *ReturnStatement) Eval(ctx *Context) Value {
	return s.Value.Eval(ctx)
}

func (e *Expression) Eval(ctx *Context) Value {
	return e.Expression.Eval(ctx)
}

func (e *ComparisonExpression) Eval(ctx *Context) Value {
	if e.Op == nil {
		return e.Lhs.Eval(ctx)
	}

	lhs := e.Lhs.Eval(ctx).(Integer)
	rhs := e.Rhs.Eval(ctx).(Integer)
	var result bool
	switch *e.Op {
	case "==":
		result = lhs == rhs
  case "!=":
    result =lhs != rhs
  case "<":
		result = lhs < rhs
	case "<=":
		result = lhs <= rhs
	case ">":
		result = lhs > rhs
	case ">=":
		result = lhs >= rhs
	}

	return NewBoolean(result)
}

func (e *AdditionExpression) Eval(ctx *Context) Value {
	if len(e.Rhs) == 0 {
		return e.Lhs.Eval(ctx)
	}

	lhs := e.Lhs.Eval(ctx).(Integer)
	for _, rhs := range e.Rhs {
		op := *rhs.Op
		rhs := rhs.Mul.Eval(ctx).(Integer)
		switch op {
		case "+":
			lhs += rhs
		case "-":
			lhs -= rhs

		}
	}

	return lhs
}

func (e *MultiplicationExpression) Eval(ctx *Context) Value {
	if len(e.Rhs) == 0 {
		return e.Lhs.Eval(ctx)
	}

	lhs := e.Lhs.Eval(ctx).(Integer)
	for _, rhs := range e.Rhs {
		op := *rhs.Op
		rhs := rhs.Term.Eval(ctx).(Integer)
		switch op {
		case "*":
			lhs *= rhs
		case "/":
			lhs /= rhs
    case "%":
			lhs %= rhs
		}
	}

	return lhs
}

func (t *TermExpression) Eval(ctx *Context) Value {
	switch {
	case t.Variable != nil:
		return ctx.FindVariable(*t.Variable)
	case t.Number != nil:
		return NewInteger(*t.Number)
  case t.Expression != nil:
		return t.Expression.Eval(ctx)
  case t.Function != nil:
		return t.Function.Eval(ctx)
  case t.Call != nil:
		return t.Call.Eval(ctx)
    
    
	}

	panic("unreachable")
}

func (f *FunctionExpression) Eval(ctx *Context) Value {
	return NewFunction(f.Body)
}

func (c *CallExpression) Eval(ctx *Context) Value {
	f := ctx.FindVariable(*c.Name).(*Function)
	return f.Body.Eval(ctx)
}
