package error

import "fmt"

func encodeErrMsg(e encodeError) string {
	return fmt.Sprintf("ERROR_CODE_MSG: %s/%s/%s, ERROR_ORIGIN: %v, COMMENT: %s, CALL_FUNC: %s",
		e.getErrCodeMsg(), e.getErrCodeSubMsg(), e.getErrCodeSubSubMsg(), e.getOriginErr(), e.encodeComment(), e.encodeCallFunc())
}

func AppendCallFunc(e encodeError) {
	f := getCallFunc(appendDeep)
	e.appendCallFunc(f)
}
