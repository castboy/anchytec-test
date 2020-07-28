package error

type encodeError interface{
	appendCallFunc(f ...string)
	getOriginErr() error
	encodeCallFunc() string
	encodeComment() string

	getErrCodeMsg() string
	getErrCodeSubMsg() string
	getErrCodeSubSubMsg() string
}