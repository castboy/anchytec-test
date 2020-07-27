package error

type encodeError interface{
	getOriginErr() error
	encodeErrCode() string
	encodeCallFunc() string
	encodeComment() string
}
