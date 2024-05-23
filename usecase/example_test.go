package usecase

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// 	"github.com/uptrace/bun"
// 	"go.uber.org/mock/gomock"

// 	"github.com/OLIENTTECH/backend-challenges/domain/model"
// 	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
// 	"github.com/OLIENTTECH/backend-challenges/usecase/input"
// 	"github.com/OLIENTTECH/backend-challenges/usecase/output"
// )

// func Test_user_Get(t *testing.T) {
// 	t.Parallel()
// 	lastLoginedAt := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
// 	tests := []struct {
// 		name     string
// 		input    *input.GetUserDTO
// 		setup    func(t *testing.T, f *testFixture)
// 		want     *output.UserDTO
// 		wantCode cerror.Code
// 	}{
// 		{
// 			name: "success",
// 			input: &input.GetUserDTO{
// 				UserID: "5cb15c25-3ba0-4f8a-8928-5b2987075943",
// 			},
// 			setup: func(t *testing.T, f *testFixture) {
// 				t.Helper()
// 				f.ds.EXPECT().User().Return(f.userRepo)
// 				f.userRepo.EXPECT().Get(context.Background(), "5cb15c25-3ba0-4f8a-8928-5b2987075943").
// 					Return(&model.User{
// 						ID:         "11edf3a8-2264-d984-bd6f-0242ac120003",
// 						LoginID:    "0123456789",
// 						Password:   "P@ssw0rd",
// 						FamilyName: "logisco",
// 						GivenName:  "tarou",
// 						RoleID:     1,
// 						LastLoginedAt: bun.NullTime{
// 							Time: lastLoginedAt,
// 						},
// 						CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
// 						UpdatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
// 						DeletedAt: bun.NullTime{},
// 					}, nil)
// 			},
// 			want: &output.UserDTO{
// 				ID:            "11edf3a8-2264-d984-bd6f-0242ac120003",
// 				LoginID:       "0123456789",
// 				Password:      "P@ssw0rd",
// 				FamilyName:    "logisco",
// 				GivenName:     "tarou",
// 				Role:          "admin",
// 				LastLoginedAt: &lastLoginedAt,
// 			},
// 			wantCode: cerror.OK,
// 		},
// 		{
// 			name: "not found user",
// 			input: &input.GetUserDTO{
// 				UserID: "not-found-user-id",
// 			},
// 			setup: func(t *testing.T, f *testFixture) {
// 				t.Helper()
// 				f.ds.EXPECT().User().Return(f.userRepo)
// 				f.userRepo.EXPECT().Get(context.Background(), "not-found-user-id").
// 					Return(nil, cerror.New("dao: user not found", cerror.WithNotFoundCode()))
// 			},
// 			want:     nil,
// 			wantCode: cerror.NotFound,
// 		},
// 		{
// 			name: "failed to get user",
// 			input: &input.GetUserDTO{
// 				UserID: "test-id",
// 			},
// 			setup: func(t *testing.T, f *testFixture) {
// 				t.Helper()
// 				f.ds.EXPECT().User().Return(f.userRepo)
// 				f.userRepo.EXPECT().Get(context.Background(), "test-id").
// 					Return(nil, cerror.New("dao: failed to get user", cerror.WithPostgreSQLCode()))
// 			},
// 			want:     nil,
// 			wantCode: cerror.PostgreSQL,
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			f := newTestFixture(t)
// 			user := newUser(f)
// 			if tt.setup != nil {
// 				tt.setup(t, f)
// 			}
// 			got, err := user.Get(context.Background(), tt.input)
// 			require.Equal(t, tt.want, got)
// 			require.Equal(t, tt.wantCode, cerror.GetCode(err))
// 		})
// 	}
// }

// func Test_user_List(t *testing.T) {
// 	t.Parallel()
// 	lastLoginedAt := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
// 	tests := []struct {
// 		name     string
// 		setup    func(t *testing.T, f *testFixture)
// 		want     *output.ListUsers
// 		wantCode cerror.Code
// 	}{
// 		{
// 			name: "success",
// 			setup: func(t *testing.T, f *testFixture) {
// 				t.Helper()
// 				f.ds.EXPECT().User().Return(f.userRepo)
// 				f.userRepo.EXPECT().List(context.Background()).
// 					Return([]*model.User{
// 						{
// 							ID:         "11edf3a8-2264-d984-bd6f-0242ac120001",
// 							LoginID:    "0123456789",
// 							Password:   "P@ssw0rd",
// 							FamilyName: "logisco",
// 							GivenName:  "tarou1",
// 							RoleID:     1,
// 							LastLoginedAt: bun.NullTime{
// 								Time: lastLoginedAt,
// 							},
// 							CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
// 							UpdatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
// 							DeletedAt: bun.NullTime{},
// 						},
// 						{
// 							ID:         "11edf3a8-2264-d984-bd6f-0242ac120002",
// 							LoginID:    "0123456788",
// 							Password:   "P@ssw0rd",
// 							FamilyName: "logisco",
// 							GivenName:  "tarou2",
// 							RoleID:     1,
// 							LastLoginedAt: bun.NullTime{
// 								Time: lastLoginedAt,
// 							},
// 							CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
// 							UpdatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
// 							DeletedAt: bun.NullTime{},
// 						},
// 						{
// 							ID:         "11edf3a8-2264-d984-bd6f-0242ac120003",
// 							LoginID:    "0123456787",
// 							Password:   "P@ssw0rd",
// 							FamilyName: "logisco",
// 							GivenName:  "tarou3",
// 							RoleID:     2,
// 							LastLoginedAt: bun.NullTime{
// 								Time: lastLoginedAt,
// 							},
// 							CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
// 							UpdatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
// 							DeletedAt: bun.NullTime{},
// 						},
// 					}, nil)
// 			},
// 			want: &output.ListUsers{
// 				Users: []*output.UserDTO{
// 					{
// 						ID:            "11edf3a8-2264-d984-bd6f-0242ac120001",
// 						LoginID:       "0123456789",
// 						Password:      "P@ssw0rd",
// 						FamilyName:    "logisco",
// 						GivenName:     "tarou1",
// 						Role:          "admin",
// 						LastLoginedAt: &lastLoginedAt,
// 					},
// 					{
// 						ID:            "11edf3a8-2264-d984-bd6f-0242ac120002",
// 						LoginID:       "0123456788",
// 						Password:      "P@ssw0rd",
// 						FamilyName:    "logisco",
// 						GivenName:     "tarou2",
// 						Role:          "admin",
// 						LastLoginedAt: &lastLoginedAt,
// 					},
// 					{
// 						ID:            "11edf3a8-2264-d984-bd6f-0242ac120003",
// 						LoginID:       "0123456787",
// 						Password:      "P@ssw0rd",
// 						FamilyName:    "logisco",
// 						GivenName:     "tarou3",
// 						Role:          "general",
// 						LastLoginedAt: &lastLoginedAt,
// 					},
// 				},
// 			},
// 			wantCode: cerror.OK,
// 		},
// 		{
// 			name: "failed to get users",
// 			setup: func(t *testing.T, f *testFixture) {
// 				t.Helper()
// 				f.ds.EXPECT().User().Return(f.userRepo)
// 				f.userRepo.EXPECT().List(context.Background()).
// 					Return(nil, cerror.New("dao: failed to list users", cerror.WithPostgreSQLCode()))
// 			},
// 			want:     nil,
// 			wantCode: cerror.PostgreSQL,
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			f := newTestFixture(t)
// 			user := newUser(f)
// 			if tt.setup != nil {
// 				tt.setup(t, f)
// 			}
// 			got, err := user.List(context.Background())
// 			require.Equal(t, tt.want, got)
// 			require.Equal(t, tt.wantCode, cerror.GetCode(err))
// 		})
// 	}
// }

// func Test_user_Create(t *testing.T) {
// 	t.Parallel()
// 	tests := []struct {
// 		name     string
// 		input    *input.CreateUserDTO
// 		setup    func(t *testing.T, f *testFixture)
// 		wantCode cerror.Code
// 	}{
// 		{
// 			name: "success",
// 			input: &input.CreateUserDTO{
// 				LoginID:    "0123456789",
// 				Password:   "P@ssw0rd",
// 				FamilyName: "logisco",
// 				GivenName:  "tarou1",
// 				RoleID:     1,
// 			},
// 			setup: func(t *testing.T, f *testFixture) {
// 				t.Helper()
// 				f.ds.EXPECT().User().Return(f.userRepo)
// 				f.userRepo.EXPECT().Create(context.Background(), gomock.Any()).Return(nil)
// 			},
// 			wantCode: cerror.OK,
// 		},
// 		{
// 			name: "failed to create user",
// 			input: &input.CreateUserDTO{
// 				LoginID:    "0123456788",
// 				Password:   "P@ssw0rd",
// 				FamilyName: "logisco",
// 				GivenName:  "tarou2",
// 				RoleID:     1,
// 			},
// 			setup: func(t *testing.T, f *testFixture) {
// 				t.Helper()
// 				f.ds.EXPECT().User().Return(f.userRepo)
// 				f.userRepo.EXPECT().Create(context.Background(), gomock.Any()).Return(cerror.New("dao: failed to create user", cerror.WithPostgreSQLCode()))
// 			},
// 			wantCode: cerror.PostgreSQL,
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			f := newTestFixture(t)
// 			user := newUser(f)
// 			if tt.setup != nil {
// 				tt.setup(t, f)
// 			}
// 			_, err := user.Create(context.Background(), tt.input)
// 			require.Equal(t, tt.wantCode, cerror.GetCode(err))
// 		})
// 	}
// }
