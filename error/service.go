package error

type encodeError interface{
	encodeErrCode() string
	encodeComment() string
	getCallFunc() []string
	getOriginErr() error
}
