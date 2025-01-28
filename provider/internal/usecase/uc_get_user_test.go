package usecase

import (
	"testing"

	"github.com/mesirendon/contract-testing/provider/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGetUser_GetUser(t *testing.T) {
	type fields struct {
		db map[int]model.User
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr string
	}{
		{
			name: "invalid id",
			fields: fields{
				db: map[int]model.User{},
			},
			args: args{
				id: "12bc",
			},
			want:    model.User{},
			wantErr: "converting id",
		},
		{
			name: "user not found",
			fields: fields{
				db: map[int]model.User{},
			},
			args: args{
				id: "10",
			},
			wantErr: "not found",
		},
		{
			name: "ok",
			fields: fields{
				db: map[int]model.User{
					10: {
						FirstName: "John",
						LastName:  "Doe",
						Username:  "drwho",
						Type:      "user",
						ID:        10,
					},
				},
			},
			args: args{
				id: "10",
			},
			want: model.User{
				FirstName: "John",
				LastName:  "Doe",
				Username:  "drwho",
				Type:      "user",
				ID:        10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewGetUser(tt.fields.db)

			got, err := uc.GetUser(tt.args.id)

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
