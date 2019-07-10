package myError

import (
	"fmt"
	"runtime"
)

type Error struct {
	msg   string
	stack []*ErrorStack
}

type ErrorStack struct {
	file string
	f    string
	line int
}

func (e *Error) Error() string {
	stacks := "{"
	for i := range e.stack {
		stack := fmt.Sprintf("[file: %s, func: %s, line: %d]", e.stack[i].file, e.stack[i].f, e.stack[i].line)
		stacks += stack
	}
	stacks += "}"

	return fmt.Sprintf("{err: %s, stack: %+v}", e.msg, stacks)
}

func New(msg string) *Error {
	stack := NewStack()
	return &Error{msg: msg, stack: []*ErrorStack{stack}}
}

func Errorf(format string, args ...interface{}) *Error {
	stack := NewStack()
	return &Error{msg: fmt.Sprintf(format, args), stack: []*ErrorStack{stack}}
}

func Wrapf(err error, format string, args ...interface{}) *Error {
	s := fmt.Sprintf(format, args)
	return Wrap(err, s)
}

func Wrap(err error, msg ...string) *Error {
	stack := NewStack()
	if err == nil {
		if len(msg) == 0 {
			return &Error{stack: []*ErrorStack{stack}}
		}
		return &Error{msg: msg[0], stack: []*ErrorStack{stack}}
	}

	switch e := err.(type) {
	case *Error:
		e.stack = append(e.stack, stack)
		if len(msg) != 0 {
			e.msg = msg[0] + ":: " + e.msg
		}
		return e

	default:
		if len(msg) != 0 {
			m := msg[0] + ":: " + e.Error()
			return &Error{msg: m, stack: []*ErrorStack{stack}}
		}
		return &Error{msg: e.Error(), stack: []*ErrorStack{stack}}
	}

	return nil
}

func NewStack() *ErrorStack {
	pc, file, line, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	return &ErrorStack{
		file: file,
		f:    f.Name(),
		line: line,
	}
}
