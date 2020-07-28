package error

type encodeError interface{
	appendCallFunc(f ...string)
	getOriginErr() error
	encodeCallFunc() string
	encodeComment() string

	errCodeMsgor
}

type errCodeMsgor interface {
	getErrCodeMsg() string
	getErrCodeSubMsg() string
	getErrCodeSubSubMsg() string
}