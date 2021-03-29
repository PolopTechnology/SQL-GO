package main

import (
     "fmt"
     "strings"
)

type Keyword string

const (
	SelectKeyword     Keyword = "select"
	FromKeyword       Keyword = "from"
	AsKeyword         Keyword = "as"
	TableKeyword      Keyword = "table"
	CreateKeyword     Keyword = "create"
	DropKeyword       Keyword = "drop"
	InsertKeyword     Keyword = "insert"
	IntoKeyword       Keyword = "into"
	ValuesKeyword     Keyword = "values"
	IntKeyword        Keyword = "int"
	TextKeyword       Keyword = "text"
	BoolKeyword       Keyword = "boolean"
	WhereKeyword      Keyword = "where"
	AndKeyword        Keyword = "and"
	OrKeyword         Keyword = "or"
	TrueKeyword       Keyword = "true"
	FalseKeyword      Keyword = "false"
	UniqueKeyword     Keyword = "unique"
	IndexKeyword      Keyword = "index"
	OnKeyword         Keyword = "on"
	PrimarykeyKeyword Keyword = "primary key"
	NullKeyword       Keyword = "null"
	LimitKeyword      Keyword = "limit"
	OffsetKeyword     Keyword = "offset"
)

type Symbol string

const (
	SemicolonSymbol  Symbol = ";"
	AsteriskSymbol   Symbol = "*"
	CommaSymbol      Symbol = ","
	LeftParenSymbol  Symbol = "("
	RightParenSymbol Symbol = ")"
	EqSymbol         Symbol = "="
	NeqSymbol        Symbol = "<>"
	NeqSymbol2       Symbol = "!="
	ConcatSymbol     Symbol = "||"
	PlusSymbol       Symbol = "+"
	LtSymbol         Symbol = "<"
	LteSymbol        Symbol = "<="
	GtSymbol         Symbol = ">"
	GteSymbol        Symbol = ">="
)
