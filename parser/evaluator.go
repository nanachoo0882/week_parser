package parser

import "container/ring"

var r = ring.New(len(weekElements))
var weekElements = [7]rune{'月', '火', '水', '木', '金', '土', '日'}

func init() {

	for i := 0; i < len(weekElements); i++ {
		r.Value = weekElements[i]

		r = r.Next()
	}
}

func Eval(e Expression) map[rune]bool {
	var weekExists = map[rune]bool{
		'月': false,
		'火': false,
		'水': false,
		'木': false,
		'金': false,
		'土': false,
		'日': false,
		'祝': false,
		'前': false,
	}
	return calc(e, weekExists)
}

func calc(e Expression, weekExists map[rune]bool) map[rune]bool {
	switch e.(type) {
	case BinOpExpr:
		switch e.(BinOpExpr).operator {
		case '~':
			left := e.(BinOpExpr).left.(WeekExpr).literal
			right := e.(BinOpExpr).right.(WeekExpr).literal
			weekExists = createWeekRange(left, right, weekExists)
		case '･':
			weekExists = calc(e.(BinOpExpr).left, weekExists)
			weekExists = calc(e.(BinOpExpr).right, weekExists)
		}
	case WeekExpr:
		weekExists[e.(WeekExpr).literal] = true
	}
	return weekExists
}

func Error(e string) {
	panic(e)
}

func createWeekRange(first rune, last rune, weekExists map[rune]bool) map[rune]bool {
	i := 0
	for {
		if i == len(weekElements) {
			panic("not exists")
		}
		if r.Value.(rune) == first {
			break
		}
		i++
		r = r.Next()
	}
	j := 0
	for {
		if j == len(weekElements) {
			panic("not exists")
		}
		weekExists[r.Value.(rune)] = true
		if r.Value.(rune) == last {
			break
		}
		j++
		r = r.Next()
	}
	return weekExists
}