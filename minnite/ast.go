package minnite

type Program struct {
	Statements []*Statement `@@*`
}

type Statement struct {
	Let   *LetStatement   `( @@ `
	If    *IfStatement    `| @@ `
  For   *ForStatement   `| @@ `
	Print *PrintStatement `| @@ ) ";"`
}

type LetStatement struct {
	Variable string      `"let" @Ident`
	Value    *Expression `"=" @@`
}

type PrintStatement struct {
	Value *Expression `"print" @@`
}

type IfStatement struct {
	Cond *Expression `"if" @@`
	Con  *Statement  `@@`
	Alt  *Statement  `@@`
}

type ForStatement struct {
	Init *Statement `"for" @@`
	Crit  *Expression  `@@`
	Proc  *Statement  `@@`
  Chan  *Statement  `@@`
}

type Expression struct {
	Expression *ComparisonExpression `@@`
}

type ComparisonExpression struct {
	Lhs *AdditionExpression `@@`
	Op  *string             `[ @Operator`
	Rhs *AdditionExpression `  @@ ]`
}

type AdditionExpression struct {
	Lhs *MultiplicationExpression `@@`
	Rhs []*OpAdditionExpression   `@@*`
}

type OpAdditionExpression struct {
	Op  *string                   `@("+" | "-")`
	Mul *MultiplicationExpression `  @@`
}

type MultiplicationExpression struct {
	Lhs *TermExpression               `@@`
	Rhs []*OpMultiplicationExpression `@@*`
}

type OpMultiplicationExpression struct {
	Op   *string         `@("*" | "/"|"%")`
	Term *TermExpression `@@`
}

type TermExpression struct {
	Variable   *string     `@Ident`
	Number     *int        `| @Number`
	Expression *Expression `| ( "(" @@ ")" )`
}
