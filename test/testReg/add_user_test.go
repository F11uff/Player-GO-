package testReg

import (
	"github.com/golang/mock/gomock"
	"player/test/testReg/mocks"
	"testing"
)

type MockRow struct {
	err error
}

func (m *MockRow) Scan(dest ...interface{}) error {
	return m.err
}

func TestFindUserForEmailMocks_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mocks.NewMockDBMocks(ctrl)
	mockRow := &MockRow{err: nil}

	user := UserRegistrationMocks{Email: "test@example.com"}

	mock.EXPECT().QueryRow(gomock.Any(), user.Email).Return(mockRow).Times(1)

	reg, err := user.FindUserForEmailMocks(mock)

	if err != nil {
		t.Fatalf("FindUserForEmailMocks err:%v", err)
	}

	if reg != true {
		t.Fatalf("FindUserForEmailMocks err:%v", reg)
	}
}
