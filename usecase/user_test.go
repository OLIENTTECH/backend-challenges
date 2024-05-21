package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"go.uber.org/mock/gomock"

	"github.com/OLIENTTECH/backend-challenges/domain/model"
	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
	"github.com/OLIENTTECH/backend-challenges/usecase/input"
	"github.com/OLIENTTECH/backend-challenges/usecase/output"
)

func Test_user_List(t *testing.T) {
	t.Parallel()
	lastLoginedAt := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		setup    func(t *testing.T, f *testFixture)
		want     *output.ListUsers
		wantCode cerror.Code
	}{
		{
			name: "success",
			setup: func(t *testing.T, f *testFixture) {
				t.Helper()
				f.ds.EXPECT().User().Return(f.userRepo)
				f.userRepo.EXPECT().List(context.Background()).
					Return([]*model.User{
						{
							ID:            "01HTDPT94BX2YC8AY75T5M9W6X",
							ShopID:        "01F9ZG3XJ90TPTKBK9FJGHK4QY",
							Name:          "ユーザー名1",
							Email:         "test1@example.com",
							Password:      "307170ea-b13d-474d-82d0-5a35f04af8b0",
							IsShopManager: true,
							LastLoginedAt: bun.NullTime{Time: lastLoginedAt},
						},
						{
							ID:            "01HTDPT94BF4CPVA9XMTBT09HP",
							ShopID:        "01F9ZG3ZZW8Y3VW0KR1H7ZE84T",
							Name:          "ユーザー名2",
							Email:         "test2@example.com",
							Password:      "e28f0a3e-28d7-4657-958e-1d20577c69ae",
							IsShopManager: true,
							LastLoginedAt: bun.NullTime{
								Time: lastLoginedAt,
							},
						},
						{
							ID:            "01HTDPT94BN5TAQ59Z4KWGR86Y",
							ShopID:        "01F9ZG3ZZW8Y3VW0KR1H7ZE84T",
							Name:          "ユーザー名3",
							Email:         "test3@example.com",
							Password:      "08e71f5c-4f30-4c5c-b755-a693ae4b7270",
							IsShopManager: false,
							LastLoginedAt: bun.NullTime{
								Time: lastLoginedAt,
							},
						},
					}, nil)
			},
			want: &output.ListUsers{
				Users: []*output.UserDTO{
					{
						User: output.User{
							ID:            "01HTDPT94BX2YC8AY75T5M9W6X",
							ShopID:        "01F9ZG3XJ90TPTKBK9FJGHK4QY",
							Name:          "ユーザー名1",
							Email:         "test1@example.com",
							Password:      "307170ea-b13d-474d-82d0-5a35f04af8b0",
							IsShopManager: false,
							LastLoginedAt: &lastLoginedAt,
							CreatedAt:     &time.Time{},
							UpdatedAt:     &time.Time{},
						},
						Shop: output.Shop{
							ID:        "01F9ZG3XJ90TPTKBK9FJGHK4QY",
							Name:      "ショップ名2",
							CreatedAt: &time.Time{},
							UpdatedAt: &time.Time{},
						},
					},
					{
						User: output.User{
							ID:            "01HTDPT94BF4CPVA9XMTBT09HP",
							ShopID:        "01F9ZG3ZZW8Y3VW0KR1H7ZE84T",
							Name:          "ユーザー名2",
							Email:         "test2@example.com",
							Password:      "e28f0a3e-28d7-4657-958e-1d20577c69ae",
							LastLoginedAt: &lastLoginedAt,
							CreatedAt:     &time.Time{},
							UpdatedAt:     &time.Time{},
						},
						Shop: output.Shop{
							ID:        "01F9ZG3ZZW8Y3VW0KR1H7ZE84T",
							Name:      "ショップ名1",
							CreatedAt: &time.Time{},
							UpdatedAt: &time.Time{},
						},
					},
					{
						User: output.User{
							ID:            "01HTDPT94BN5TAQ59Z4KWGR86Y",
							ShopID:        "01F9ZG3ZZW8Y3VW0KR1H7ZE84T",
							Name:          "ユーザー名3",
							Email:         "test3@example.com",
							Password:      "08e71f5c-4f30-4c5c-b755-a693ae4b7270",
							LastLoginedAt: &lastLoginedAt,
							CreatedAt:     &time.Time{},
							UpdatedAt:     &time.Time{},
						},
						Shop: output.Shop{
							ID:        "01F9ZG3ZZW8Y3VW0KR1H7ZE84T",
							Name:      "ショップ名1",
							CreatedAt: &time.Time{},
							UpdatedAt: &time.Time{},
						},
					},
				},
			},
			wantCode: cerror.OK,
		},
		{
			name: "failed to get users",
			setup: func(t *testing.T, f *testFixture) {
				t.Helper()
				f.ds.EXPECT().User().Return(f.userRepo)
				f.userRepo.EXPECT().List(context.Background()).
					Return(nil, cerror.New("dao: failed to list users", cerror.WithPostgreSQLCode()))
			},
			want:     nil,
			wantCode: cerror.PostgreSQL,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := newTestFixture(t)
			user := newUser(f)
			if tt.setup != nil {
				tt.setup(t, f)
			}
			got, err := user.List(context.Background())
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantCode, cerror.GetCode(err))
		})
	}
}

func Test_users_Create(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    *input.CreateUserDTO
		setup    func(t *testing.T, f *testFixture)
		wantCode cerror.Code
	}{
		{
			name: "success",
			input: &input.CreateUserDTO{
				ShopID:        "01F9ZG3XJ90TPTKBK9FJGHK4QY",
				Name:          "user1",
				Email:         "test@example.com",
			},
			setup: func(t *testing.T, f *testFixture) {
				t.Helper()
				f.ds.EXPECT().User().Return(f.userRepo)
				f.userRepo.EXPECT().Create(context.Background(), gomock.Any()).Return(nil)
			},
			wantCode: cerror.OK,
		},
		{
			name: "failed to create user",
			input: &input.CreateUserDTO{
				ShopID:        "01F9ZG3XJ90TPTKBK9FJGHK4QY",
				Name:          "user2",
				Email:         "test@example.com",
			},
			setup: func(t *testing.T, f *testFixture) {
				t.Helper()
				f.ds.EXPECT().User().Return(f.userRepo)
				f.userRepo.EXPECT().Create(context.Background(), gomock.Any()).Return(cerror.New("dao: failed to create user", cerror.WithPostgreSQLCode()))
			},
			wantCode: cerror.PostgreSQL,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := newTestFixture(t)
			user := newUser(f)
			if tt.setup != nil {
				tt.setup(t, f)
			}
			_, err := user.Create(context.Background(), tt.input)
			require.Equal(t, tt.wantCode, cerror.GetCode(err))
		})
	}
}
