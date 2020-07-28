package error

import "fmt"

func encodeErrCodeMsg(e errCodeMsgor) string {
	return fmt.Sprintf("%s/%s/%s", e.getErrCodeMsg(), e.getErrCodeSubMsg(), e.getErrCodeSubSubMsg())
}

func encodeErrMsg(e encodeError) string {
	return fmt.Sprintf("ERROR_CODE_MSG: %s, ERROR_ORIGIN: %v, COMMENT: %s, CALL_FUNC: %s", encodeErrCodeMsg(e), e.getOriginErr(), e.encodeComment(), e.encodeCallFunc())
}

func AppendCallFunc(e encodeError) {
	f := getCallFunc(appendDeep)
	e.appendCallFunc(f)
}
