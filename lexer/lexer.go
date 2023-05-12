package lexer

type Lexer struct {
	input        string
	position     int  // points to current char
	readPosition int  // points after current char
	ch           byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	return l
}
