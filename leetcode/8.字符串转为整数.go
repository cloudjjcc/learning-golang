package leetcode

import "math"

func myAtoi(s string) int {
	d := NewDFA1()
	for _, ch := range s {
		if !d.input(byte(ch)) {
			break
		}
	}
	return d.ans * d.sign
}

type state1 byte

const (
	stateStart state1 = iota
	stateSigned
	stateInNumber
	stateEnd
)

type dfa1 struct {
	sign  int
	ans   int
	state state1
}

var stateTable = map[state1][4]state1{
	stateStart:    {stateStart, stateSigned, stateInNumber, stateEnd},
	stateSigned:   {stateEnd, stateEnd, stateInNumber, stateEnd},
	stateInNumber: {stateEnd, stateEnd, stateInNumber, stateEnd},
	stateEnd:      {stateEnd, stateEnd, stateEnd, stateEnd},
}

func NewDFA1() *dfa1 {
	return &dfa1{
		sign:  1,
		ans:   0,
		state: stateStart,
	}
}
func (d *dfa1) input(ch byte) bool {
	switch d.state = stateTable[d.state][d.charType(ch)]; d.state {
	case stateStart:
		return true
	case stateSigned:
		if ch == '-' {
			d.sign = -1
		}
	case stateInNumber:
		d.ans = d.ans*10 + int(ch-'0')
		if d.sign == -1 && d.ans*-1 < math.MinInt32 {
			d.ans = -math.MinInt32
			return false
		}
		if d.sign == 1 && d.ans*d.sign > math.MaxInt32 {
			d.ans = math.MaxInt32
			return false
		}
	default:
		return false
	}
	return true
}
func (d *dfa1) charType(ch byte) int {
	switch {
	case ch == ' ':
		return 0
	case ch == '+' || ch == '-':
		return 1
	case ch >= '0' && ch <= '9':
		return 2
	default:
		return 3
	}
}
