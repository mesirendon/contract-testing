package usecase

import (
	"errors"
	"testing"

	"github.com/mesirendon/contract-testing/consumer/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGreetUser_GreetUser(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
		want    string
		mocker  func(userService *mockUserGetter)
	}{
		{
			name: "error getting the user from the service",
			args: args{
				id: "10",
			},
			wantErr: "fetching user from user service: service error",
			mocker: func(userService *mockUserGetter) {
				userService.EXPECT().GetUser(mock.AnythingOfType("int")).
					Return(model.User{}, errors.New("service error"))
			},
		},
		{
			name: "successfully getting the user with id 10",
			args: args{
				id: "10",
			},
			want: "Hello John Doe!",
			mocker: func(userService *mockUserGetter) {
				userService.EXPECT().GetUser(10).
					Return(
						model.User{
							FirstName: "John",
							LastName:  "Doe",
							Username:  "drwho",
							Type:      "user",
							ID:        10,
						},
						nil,
					)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := newMockUserGetter(t)
			tt.mocker(us)

			uc := NewGreetUser(us)

			got, err := uc.GreetUser(tt.args.id)

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
