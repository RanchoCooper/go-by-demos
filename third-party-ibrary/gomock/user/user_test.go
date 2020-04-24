package user

import (
	"testing"
	
	"github.com/golang/mock/gomock"

	"github.com/RanchoCooper/go-by-demos/third-party-ibrary/gomock/mock"
)

func TestUser_GetUserInfo(t *testing.T) {
	// 创建mock控件
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	
	var id int64 = 1
	
	// 创建mock实例
	mockMale := male_mock.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(nil),
	)
	
	user := NewUser(mockMale)
	err := user.GetUserInfo(id)
	if err != nil {
		t.Errorf("user.GetUserInfo err: %v", err)
	}
}

