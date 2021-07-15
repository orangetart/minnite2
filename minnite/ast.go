package minnite

type Program struct {
	Statements []*Statement `@@*`
}

type Statement struct {
	Let   *LetStatement   `( @@ `
	If    *IfStatement    `| @@ `
  For   *ForStatement   `| @@ `
  Switch*SwitchStatement `| @@ `
  Case*CaseStatement     `| @@ `
  Return*ReturnStatement `| @@ `
  Summon *SummonStatement`| @@ `
  Print *PrintStatement  `| @@ ) ";"`
}

type BlockStatement struct {
	Body []*Statement `"{" @@* "}"`
}

type LetStatement struct {
	Variable string      `"let" @Ident`
	Value    *Expression `"=" @@`
}

type PrintStatement struct {
	Value *Expression `"print"@@`
}

type IfStatement struct {
	Cond *Expression `"if" @@`
	Con  *Statement  `@@`
	Alt  *Statement  `@@`
}

type SwitchStatement struct {
	Va  *Expression `"switch" @@`
}

type CaseStatement struct {
	CN   *Expression `"case" @@`
  Num  *Statement  `@@`
}

type ForStatement struct {
	Init *Statement `"for" @@`
	Crit  *Expression  `@@`
	Proc  *Statement  `@@`
  Chan  *Statement  `@@`
}

type SummonStatement struct {
  Summon *Expression `"call" @@`
}

type ReturnStatement struct {
	Value *Expression `"return" @@`
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
	Expression *Expression         `( "(" @@ ")" )`
	Function   *FunctionExpression `| @@`
  Call       *CallExpression     `| @@`
	Variable   *string             `| @Ident`
	Number     *int                `| @Number`
}

type FunctionExpression struct {
	Body *BlockStatement `"func" "(" ")" @@`
}

type CallExpression struct {
	Name *string `@Ident "(" ")"`
}
