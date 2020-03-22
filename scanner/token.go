package scanner

type TokenKind int

const (
	Using                  TokenKind = iota // using
	Class                                   // class
	Identifier                              // any identifier
	Colon                                   // :
	Comma                                   // ,
	Semicolon                               // ;
	OpenBrace                               // {
	CloseBrace                              // }
	StringLiteral                           // "
	Namespace                               // namespace
	PublicModifier                          // public
	PrivateModifier                         // private
	Add                                     // +
	Minus                                   // -
	Mult                                    // *
	Divide                                  // /
	Var                                     // var
	Whitespace                              //
	OpenParens                              // (
	CloseParens                             // )
	Percent                                 // %
	OpenMultilineComments                   // /*
	CloseMultilineComments                  // */
	Dot                                     // .
	Return                                  // return
	Null                                    // null
	LambdaArrow                             // =>
	If                                      // if
	Equal                                   // ==
	OpenSingleComments                      // //
	Assign                                  // =
	OpenBracket                             // [
	CloseBracket                            // ]
	ProtectedModifier                       // protected
	False                                   // false
	True                                    // true
	New                                     // new
	IntegerLiteral                          // any int number
	RealLiteral                             // any real number
	EOF                                     // end of file
	Illegal                = -1             // just illegal token
)

func (tok TokenKind) String() string {
	switch tok {
	case Class:
		return "class"
	case Colon:
		return ":"
	case Comma:
		return ","
	case Semicolon:
		return ";"
	case OpenBrace:
		return "{"
	case CloseBrace:
		return "}"
	case StringLiteral:
		return "\""
	case Using:
		return "using"
	case Var:
		return "var"
	case Add:
		return "+"
	case Minus:
		return "-"
	case Mult:
		return "*"
	case Divide:
		return "/"
	case Namespace:
		return "namespace"
	case PublicModifier:
		return "public"
	case PrivateModifier:
		return "private"
	case OpenParens:
		return "("
	case CloseParens:
		return ")"
	case Percent:
		return "%"
	case OpenMultilineComments:
		return "/*"
	case CloseMultilineComments:
		return "*/"
	case Dot:
		return "."
	case Return:
		return "return"
	case Null:
		return "null"
	case LambdaArrow:
		return "=>"
	case If:
		return "if"
	case Equal:
		return "=="
	case OpenSingleComments:
		return "//"
	case Assign:
		return "="
	case OpenBracket:
		return "["
	case CloseBracket:
		return "]"
	case ProtectedModifier:
		return "protected"
	case False:
		return "false"
	case True:
		return "true"
	case New:
		return "new"
	}
	return ""
}

func (tok TokenKind) PrettyString() string {
	switch tok {
	case Class:
		return "CLASS"
	case Colon:
		return "COLON"
	case Comma:
		return "COMMA"
	case Semicolon:
		return "SEMICOLON"
	case OpenBrace:
		return "OPEN_BRACE"
	case CloseBrace:
		return "CLOSE_BRACE"
	case StringLiteral:
		return "STRING_LITERAL"
	case Using:
		return "USING"
	case Var:
		return "VAR"
	case Add:
		return "OP_SUM"
	case Minus:
		return "OP_MINUS"
	case Mult:
		return "OP_MULT"
	case Divide:
		return "OP_DIVIDE"
	case Namespace:
		return "NAMESPACE"
	case PublicModifier:
		return "PUBLIC_MODIFIER"
	case PrivateModifier:
		return "PRIVATE_MODIFIER"
	case OpenParens:
		return "OPEN_PARENS"
	case CloseParens:
		return "CLOSE_PARENS"
	case Percent:
		return "PERCENT"
	case OpenMultilineComments:
		return "OPEN_MULTILINE_COMMENTS"
	case CloseMultilineComments:
		return "CLOSE_MULTILINE_COMMENTS"
	case Dot:
		return "DOT"
	case Return:
		return "RETURN"
	case Null:
		return "NULL"
	case LambdaArrow:
		return "LAMBDA_ARROW"
	case If:
		return "IF_STMT"
	case Illegal:
		return "ILLEGAL"
	case Equal:
		return "EQUAL"
	case OpenSingleComments:
		return "OPEN_SINGLE_COMMENTS"
	case Assign:
		return "ASSIGN_OP"
	case OpenBracket:
		return "OPEN_BRACKET"
	case CloseBracket:
		return "CLOSE_BRACKET"
	case ProtectedModifier:
		return "PROTECTED_MODIFIER"
	case False:
		return "FALSE"
	case True:
		return "TRUE"
	case New:
		return "NEW"
	case IntegerLiteral:
		return "INTEGER_LITERAL"
	case RealLiteral:
		return "REAL_LITERAL"
	}
	return ""
}

type Token struct {
	Kind    TokenKind
	Literal string
}

func (t Token) String() string {
	if t.Kind != Identifier && t.Kind != Whitespace {
		return "Token{kind: " + t.Kind.PrettyString() + ", literal: " + t.Literal + "}"
	} else if t.Kind == Whitespace {
		return "Token{kind: WHITESPACE}"
	} else {
		return "Token{kind: IDENTIFIER, literal: " + t.Literal + "}"
	}
}
