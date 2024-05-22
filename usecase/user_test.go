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

const (
	ShopID1 = "01F9ZG3ZZW8Y3VW0KR1H7ZE84T"
	ShopID2 = "01F9ZG3XJ90TPTKBK9FJGHK4QY"
	ShopID3 = "01F9ZG3TQM2X7VMP8Z9M7P0TZ2"
)

func Test_user_Login(t *testing.T) {
	t.Parallel()
	lastLoginedAt := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	shopData := map[string]*model.Shop{
		ShopID1: {ID: ShopID1, Name: "ショップ名1"},
		ShopID2: {ID: ShopID2, Name: "ショップ名2"},
		ShopID3: {ID: ShopID3, Name: "ショップ名3"},
	}

	tests := []struct {
		name     string
		input    *input.LoginUserDTO
		setup    func(t *testing.T, f *testFixture)
		want     *output.UserDTO
		wantCode cerror.Code
	}{
		{
			name: "success",
			input: &input.LoginUserDTO{
				ShopID:   ShopID2,
				Email:    "test1@example.com",
				Password: "307170ea-b13d-474d-82d0-5a35f04af8b0",
			},
			setup: func(t *testing.T, f *testFixture) {
				t.Helper()
				f.ds.EXPECT().User().Return(f.userRepo)
				f.ds.EXPECT().Shop().Return(f.shopRepo).AnyTimes()
				f.shopRepo.EXPECT().Get(gomock.Any(), ShopID2).Return(shopData[ShopID2], nil).AnyTimes()
				f.userRepo.EXPECT().Login(gomock.Any(), ShopID2, "test1@example.com", "307170ea-b13d-474d-82d0-5a35f04af8b0").
					Return(&model.User{
						ID:            "01HTDPT94BX2YC8AY75T5M9W6X",
						ShopID:        ShopID2,
						Name:          "ユーザー名1",
						Email:         "test1@example.com",
						Password:      "307170ea-b13d-474d-82d0-5a35f04af8b0",
						LastLoginedAt: bun.NullTime{Time: lastLoginedAt},
					}, nil)
			},
			want: &output.UserDTO{
				User: output.User{
					ID:            "01HTDPT94BX2YC8AY75T5M9W6X",
					ShopID:        ShopID2,
					Name:          "ユーザー名1",
					Email:         "test1@example.com",
					Password:      "307170ea-b13d-474d-82d0-5a35f04af8b0",
					LastLoginedAt: &lastLoginedAt,
					CreatedAt:     &time.Time{},
					UpdatedAt:     &time.Time{},
				},
				Shop: output.Shop{
					ID:        ShopID2,
					Name:      "ショップ名2",
					CreatedAt: &time.Time{},
					UpdatedAt: &time.Time{},
				},
			},
			wantCode: cerror.OK,
		},
		{
			name: "failed to login user",
			input: &input.LoginUserDTO{
				ShopID:   ShopID1,
				Email:    "invalid@example.com",
				Password: "invalidpassword",
			},
			setup: func(t *testing.T, f *testFixture) {
				t.Helper()
				f.ds.EXPECT().User().Return(f.userRepo)
				f.ds.EXPECT().Shop().Return(f.shopRepo).AnyTimes()
				f.shopRepo.EXPECT().Get(gomock.Any(), ShopID1).Return(shopData[ShopID1], nil).AnyTimes()
				f.userRepo.EXPECT().Login(gomock.Any(), ShopID1, "invalid@example.com", "invalidpassword").Return(nil, cerror.New("invalid credentials", cerror.WithPostgreSQLCode()))
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
			got, err := user.Login(context.Background(), tt.input)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.wantCode, cerror.GetCode(err))
		})
	}
}

func Test_user_List(t *testing.T) {
	t.Parallel()
	lastLoginedAt := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	shopData := map[string]*model.Shop{
		ShopID1: {ID: ShopID1, Name: "ショップ名1"},
		ShopID2: {ID: ShopID2, Name: "ショップ名2"},
		ShopID3: {ID: ShopID3, Name: "ショップ名3"},
	}

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
				f.ds.EXPECT().Shop().Return(f.shopRepo).AnyTimes()
				f.shopRepo.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(
					func(_ context.Context, shopID string) (*model.Shop, error) {
						if shop, exists := shopData[shopID]; exists {
							return shop, nil
						}

						return nil, cerror.New("shop not found")
					},
				).AnyTimes()

				f.userRepo.EXPECT().List(context.Background()).
					Return([]*model.User{
						{
							ID:            "01HTDPT94BX2YC8AY75T5M9W6X",
							ShopID:        ShopID2,
							Name:          "ユーザー名1",
							Email:         "test1@example.com",
							Password:      "307170ea-b13d-474d-82d0-5a35f04af8b0",
							LastLoginedAt: bun.NullTime{Time: lastLoginedAt},
						},
						{
							ID:            "01HTDPT94BF4CPVA9XMTBT09HP",
							ShopID:        ShopID1,
							Name:          "ユーザー名2",
							Email:         "test2@example.com",
							Password:      "e28f0a3e-28d7-4657-958e-1d20577c69ae",
							LastLoginedAt: bun.NullTime{Time: lastLoginedAt},
						},
						{
							ID:            "01HTDPT94BN5TAQ59Z4KWGR86Y",
							ShopID:        ShopID1,
							Name:          "ユーザー名3",
							Email:         "test3@example.com",
							Password:      "08e71f5c-4f30-4c5c-b755-a693ae4b7270",
							LastLoginedAt: bun.NullTime{Time: lastLoginedAt},
						},
					}, nil)
			},
			want: &output.ListUsers{
				Users: []*output.UserDTO{
					{
						User: output.User{
							ID:            "01HTDPT94BX2YC8AY75T5M9W6X",
							ShopID:        ShopID2,
							Name:          "ユーザー名1",
							Email:         "test1@example.com",
							Password:      "307170ea-b13d-474d-82d0-5a35f04af8b0",
							LastLoginedAt: &lastLoginedAt,
							CreatedAt:     &time.Time{},
							UpdatedAt:     &time.Time{},
						},
						Shop: output.Shop{
							ID:        ShopID2,
							Name:      "ショップ名2",
							CreatedAt: &time.Time{},
							UpdatedAt: &time.Time{},
						},
					},
					{
						User: output.User{
							ID:            "01HTDPT94BF4CPVA9XMTBT09HP",
							ShopID:        ShopID1,
							Name:          "ユーザー名2",
							Email:         "test2@example.com",
							Password:      "e28f0a3e-28d7-4657-958e-1d20577c69ae",
							LastLoginedAt: &lastLoginedAt,
							CreatedAt:     &time.Time{},
							UpdatedAt:     &time.Time{},
						},
						Shop: output.Shop{
							ID:        ShopID1,
							Name:      "ショップ名1",
							CreatedAt: &time.Time{},
							UpdatedAt: &time.Time{},
						},
					},
					{
						User: output.User{
							ID:            "01HTDPT94BN5TAQ59Z4KWGR86Y",
							ShopID:        ShopID1,
							Name:          "ユーザー名3",
							Email:         "test3@example.com",
							Password:      "08e71f5c-4f30-4c5c-b755-a693ae4b7270",
							LastLoginedAt: &lastLoginedAt,
							CreatedAt:     &time.Time{},
							UpdatedAt:     &time.Time{},
						},
						Shop: output.Shop{
							ID:        ShopID1,
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
				f.ds.EXPECT().Shop().Return(f.shopRepo).AnyTimes()
				f.shopRepo.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(
					func(_ context.Context, shopID string) (*model.Shop, error) {
						if shop, exists := shopData[shopID]; exists {
							return shop, nil
						}

						return nil, cerror.New("shop not found")
					},
				).AnyTimes()

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

	shopData := map[string]*model.Shop{
		ShopID1: {ID: ShopID1, Name: "ショップ名1"},
		ShopID2: {ID: ShopID2, Name: "ショップ名2"},
		ShopID3: {ID: ShopID3, Name: "ショップ名3"},
	}

	tests := []struct {
		name     string
		input    *input.CreateUserDTO
		setup    func(t *testing.T, f *testFixture)
		wantCode cerror.Code
	}{
		{
			name: "success",
			input: &input.CreateUserDTO{
				ShopID: ShopID2,
				Name:   "user1",
				Email:  "test@example.com",
			},
			setup: func(t *testing.T, f *testFixture) {
				t.Helper()
				f.ds.EXPECT().User().Return(f.userRepo)
				f.ds.EXPECT().Shop().Return(f.shopRepo).AnyTimes()
				f.shopRepo.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(
					func(_ context.Context, shopID string) (*model.Shop, error) {
						if shop, exists := shopData[shopID]; exists {
							return shop, nil
						}

						return nil, cerror.New("shop not found")
					},
				).AnyTimes()

				f.userRepo.EXPECT().Create(context.Background(), gomock.Any()).Return(nil)
			},
			wantCode: cerror.OK,
		},
		{
			name: "failed to create user",
			input: &input.CreateUserDTO{
				ShopID: ShopID2,
				Name:   "user2",
				Email:  "test@example.com",
			},
			setup: func(t *testing.T, f *testFixture) {
				t.Helper()
				f.ds.EXPECT().User().Return(f.userRepo)
				f.ds.EXPECT().Shop().Return(f.shopRepo).AnyTimes()
				f.shopRepo.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(
					func(_ context.Context, shopID string) (*model.Shop, error) {
						if shop, exists := shopData[shopID]; exists {
							return shop, nil
						}

						return nil, cerror.New("shop not found")
					},
				).AnyTimes()

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
