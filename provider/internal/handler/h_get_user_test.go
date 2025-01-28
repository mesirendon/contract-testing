package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mesirendon/contract-testing/provider/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUser_GetUser(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		header     string
		body       string
		mocker     func(userGetter *mockUserGetter)
	}{
		{
			name:       "internal server error",
			statusCode: http.StatusInternalServerError,
			header:     "application/json",
			mocker: func(userGetter *mockUserGetter) {
				userGetter.EXPECT().GetUser(mock.AnythingOfType("string")).
					Return(
						model.User{},
						errors.New("user getter error"),
					)
			},
		},
		{
			name:       "user not found",
			statusCode: http.StatusNotFound,
			header:     "application/json",
			mocker: func(userGetter *mockUserGetter) {
				userGetter.EXPECT().GetUser("10").
					Return(
						model.User{},
						errors.New("not found"),
					)
			},
		},
		{
			name:       "ok",
			statusCode: http.StatusOK,
			header:     "application/json",
			body:       `{"firstName":"John","lastName":"Doe","username":"drwho","type":"user","id":10}`,
			mocker: func(userGetter *mockUserGetter) {
				userGetter.EXPECT().GetUser("10").
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
			ug := newMockUserGetter(t)
			tt.mocker(ug)

			h := NewGetUser(ug)

			req := httptest.NewRequest(http.MethodGet, "/user/10", nil)
			rr := httptest.NewRecorder()

			h.GetUser(rr, req)

			assert.Equal(t, tt.statusCode, rr.Code)
			assert.Equal(t, tt.header, rr.Header().Get("Content-Type"))
			assert.Equal(t, tt.body, strings.TrimSpace(rr.Body.String()))
		})
	}
}
