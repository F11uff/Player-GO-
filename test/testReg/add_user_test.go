package testReg

import (
	"github.com/golang/mock/gomock"
	mocks "player/test/testReg/mockReg"
	"testing"
)

type MockDB struct {
	email string
	err   error
}

func (m *MockDB) QueryRow(query string, args ...interface{}) (string, error) {
	return m.email, m.err
}

func TestFindUserForEmailMocks_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	mock := mocks.NewMockDBMocks(ctrl)
	mockDB := &MockDB{email: "test@example.com", err: nil}

	user := UserRegistrationMocks{Email: "test@example.com"}

	mock.EXPECT().QueryRow(gomock.Any(), user.Email).Return(mockDB.email, mockDB.err).Times(1)

	reg, err := user.FindUserForEmailMocks(mock)

	if err != nil {
		t.Fatalf("FindUserForEmailMocks err:%v", err)
	}

	if !reg {
		t.Fatalf("FindUserForEmailMocks expected true, got %v", reg)
	}
}
