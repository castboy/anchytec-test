package intergration

import (
	"anchytec-test/mock"
	"anchytec-test/mock/mock"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
)

func TestMyThing(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock_mock.NewMockMyOperation(mockCtrl)
	mockObj.EXPECT().Add(1, 1).Return(2)
	mockObj.EXPECT().Put(3)
	mockObj.EXPECT().RtnErr().Return(errors.New("sfasf"))

	ch := make(chan bool)
	first := mockObj.EXPECT().RtnChan().Return(ch)
	second := mockObj.EXPECT().RtnChan().Return(nil)
	mockObj.EXPECT().RtnChan().Return(ch).After(first).After(second)

	mockObj.EXPECT().Add(1, 2).Do(func(arg0, arg1 int){
		fmt.Printf("%d + %d = %d\n", arg0, arg1, arg0+arg1)
	})

	mockObj.EXPECT().Ptr(1, gomock.Any()).SetArg(1, 3).Return(4)

	mock.AddOper(mockObj)
}
