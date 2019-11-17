package parser

const (
    EOF = -1
)

type Scanner struct {
    src    []rune
    offset int
}

func (s *Scanner) Init(src string) {
    s.src = []rune(src)
}

func (s *Scanner) Scan() (token int, literal rune) {
    ch := s.peek()
    if ch == -1 {
        token = EOF
    } else {
        token = ELEMENT
        literal = ch
    }
    s.next()
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