package parser

const (
    EOF = -1
)

var elements = []rune{'月', '火', '水', '木', '金', '土', '日'}
var holiday = '祝'

type Scanner struct {
    src    []rune
    offset int
}

func (s *Scanner) Init(src string) {
    s.src = []rune(src)
}

func (s *Scanner) Scan() (token int, literal rune) {
	switch ch := s.peek(); {
	case isWeekElement(ch):
		token = WEEKELEMENT
		literal = s.peek()
		s.skipWeekChars()
	case ch == holiday:
		token, literal = s.checkHoliday(ch)
	default:
		switch ch {
		case -1:
			token = EOF
		case '~', '･':
			token = int(ch)
			literal = ch
		}
		s.next()
	}
	return
}

func (s *Scanner) peek() rune {
    if !s.reachEOF() {
        return s.src[s.offset]
    } else {
        return -1
    }
}

func (s *Scanner) next() {
    if !s.reachEOF() {
        s.offset++
    }
}

func (s *Scanner) reachEOF() bool {
    return len(s.src) <= s.offset
}

func isWeekElement(ch rune) bool {
    return stringInSlice(ch, elements)
}

func stringInSlice(a rune, list []rune) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func (s *Scanner) skipWeekChars() {
	s.next()
	if s.peek() == '曜' {
		s.next()
		if s.peek() == '日' {
			s.next()
		}
	}
}

func (s *Scanner) checkHoliday(ch rune) (int, rune) {
	s.next()
	if s.peek() == '前' {
		s.next()
		if s.peek() == '日' {
			s.next()
		}
		return PREHOLIDAY, '前'
	} else if s.peek() == '日' {
		s.next()
		if s.peek() == '前' {
			s.next()
			return PREHOLIDAY, '前'
		}
	}
	return HOLIDAY, '祝'
}
