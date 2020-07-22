package core

import "fmt"

func encodeErrMsg(e encodeError) string {
	return fmt.Sprintf("ERROR_CODE_MSG: %s, ERROR_ORIGIN: %v, COMMENT: %s, CALL_FUNC: %s", e.encodeErrCode(), e.getOriginErr(), e.encodeComment(), e.getCallFunc())
}
