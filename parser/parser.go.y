%{
// ヘッダー
package parser

type Expression interface{}
type Token struct {
    token   int
    literal rune
}
type Expr struct {
  literal rune
}
%}
%union{
    token Token
    expr  Expression
}
%type<expr> expr
%token<token> ELEMENT
%%
// 規則部
expr
    : ELEMENT
    {
        $$ = Expr{literal: $1.literal}
        yylex.(*Lexer).result = $$
    }
%%
// ユーザ定義部
type Lexer struct {
    s *Scanner
    result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
    token, literal := l.s.Scan()
    if token == EOF {
        return 0
    }
    lval.token = Token{token: token, literal: literal}
    return token
}

func (l *Lexer) Error(e string) {
    panic(e)
}

func Parse(s *Scanner) Expression {
    l := Lexer{s: s}
    if yyParse(&l) != 0 {
        panic("Parse error")
    }
    return l.result
}