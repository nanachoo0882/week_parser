%{
// ヘッダー
package parser

type Expression interface{}
type Token struct {
    token   int
    literal rune
}
type WeekExpr struct {
    literal rune
}
type BinOpExpr struct {
    left     Expression
    operator rune
    right    Expression
}
%}
%union{
    token Token
    expr  Expression
}
%type<expr> program
%type<expr> expr
%token<token> WEEKELEMENT HOLIDAY PREHOLIDAY

%left '･'
%left '~'
%%
// 規則部
program
    : expr
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }

expr
    : WEEKELEMENT
    {
        $$ = WeekExpr{literal: $1.literal}
    }
    | HOLIDAY
    {
        $$ = WeekExpr{literal: $1.literal}
    }
    | PREHOLIDAY
    {
        $$ = WeekExpr{literal: $1.literal}
    }
    | expr '･' expr
    {
        $$ = BinOpExpr{left: $1, operator: '･', right: $3}
    }
    | WEEKELEMENT '~' WEEKELEMENT
    {
        $$ = BinOpExpr{left: WeekExpr{literal: $1.literal}, operator: '~', right: WeekExpr{literal: $3.literal}}
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