package main

type Program struct {
	Statements []*Statement `@@*`
}

type Statement struct {
	Let   *LetStatement   `( @@ `
	Print *PrintStatement `| @@ ) ";"`
}

type LetStatement struct {
	Variable string      `"let" @Ident`
	Value    *Expression `"=" @@`
}

type PrintStatement struct {
	Value *Expression `"print" @@`
}

type Expression struct {
	Expression *AdditionExpression `@@`
}

type AdditionExpression struct {
	Lhs *TermExpression `@@`
	Op  *string         `[ @"+"`
	Rhs *TermExpression `  @@ ]`
}

type TermExpression struct {
	Variable *string `@Ident`
	Number   *int    `| @Number`
}
