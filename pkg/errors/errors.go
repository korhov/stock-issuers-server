package errors

import (
	"context"
	"fmt"
	"io"
)

var logKeys []string

func SetLogKeys(keys []string) {
	logKeys = keys
}

type ContextError interface {
	GetContext() context.Context
	SetContext(ctx context.Context)
	Error() string
	Value(key interface{}) interface{}
	ValueInt(key interface{}, def int) int
	ValueString(key interface{}, def string) string
	ValueBool(key interface{}, def bool) bool
}

type contextError struct {
	msg string
	err error
	ctx context.Context
	*stack
}

func (err *contextError) Error() string {
	return err.msg
}

func (err *contextError) GetContext() context.Context {
	return err.ctx
}

func (err *contextError) SetContext(ctx context.Context) {
	err.ctx = ctx
}

func (err *contextError) Value(key interface{}) interface{} {
	return err.ctx.Value(key)
}

func (err *contextError) ValueInt(key interface{}, def int) int {
	val, ok := err.ctx.Value(key).(int)

	if !ok {
		return def
	}

	return val
}

func (err *contextError) ValueString(key interface{}, def string) string {
	val, ok := err.ctx.Value(key).(string)

	if !ok {
		return def
	}

	return val
}

func (err *contextError) ValueBool(key interface{}, def bool) bool {
	val, ok := err.ctx.Value(key).(bool)

	if !ok {
		return def
	}

	return val
}

func (err *contextError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, err.msg+"\n") // nolint
			_, _ = fmt.Fprintf(s, "%+v", err.err)  // nolint
			err.stack.Format(s, verb)

			for _, k := range logKeys {
				v, _ := err.ctx.Value(k).(string)
				_, _ = fmt.Fprintf(s, "\n%s:%s", k, v) // nolint
				fmt.Println(k, v)
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, err.msg) // nolint
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", err.msg) // nolint
	}
}

func New(message string) error {
	return &contextError{
		msg:   message,
		ctx:   context.TODO(),
		stack: callers(),
	}
}

func NewWithValues(message string, values ...interface{}) error {
	ctx := fillContext(context.TODO(), values...)
	return &contextError{
		msg:   message,
		ctx:   ctx,
		stack: callers(),
	}
}

func Wrap(err error, message string) error {
	ctx := context.TODO()

	// raise up context
	if ctxErr, ok := err.(ContextError); ok {
		ctx = ctxErr.GetContext()
	}

	return &contextError{
		msg:   message,
		err:   err,
		ctx:   ctx,
		stack: callers(),
	}
}

func WrapWithValues(err error, message string, values ...interface{}) error {
	ctx := context.TODO()

	// raise up context
	if ctxErr, ok := err.(ContextError); ok {
		ctx = ctxErr.GetContext()
	}

	ctx = fillContext(ctx, values...)

	return &contextError{
		msg:   message,
		err:   err,
		ctx:   ctx,
		stack: callers(),
	}
}

func WithValues(err error, values ...interface{}) error {
	if ctxErr, ok := err.(ContextError); ok {
		ctx := ctxErr.GetContext()
		for i := 0; i < len(values); i += 2 {
			ctx = context.WithValue(ctx, values[i], values[i+1])
		}

		ctxErr.SetContext(ctx)
		return ctxErr
	}

	ctx := fillContext(context.TODO(), values...)

	return &contextError{
		msg:   err.Error(),
		err:   err,
		ctx:   ctx,
		stack: callers(),
	}
}

func fillContext(ctx context.Context, values ...interface{}) context.Context {
	for i := 0; i < len(values); i += 2 {
		ctx = context.WithValue(ctx, values[i], values[i+1])
	}

	return ctx
}
