package usecase

import (
	"testing"

	datastoremock "github.com/OLIENTTECH/backend-challenges/mock/datastore"
	repositorymock "github.com/OLIENTTECH/backend-challenges/mock/repository"
	"github.com/OLIENTTECH/backend-challenges/pkg/log"
	gomock "go.uber.org/mock/gomock"
)

type testFixture struct {
	ds       *datastoremock.MockDataStore
	userRepo *repositorymock.MockUser
	logger   *log.Logger
}

func newTestFixture(t *testing.T) *testFixture {
	t.Helper()
	ctrl := gomock.NewController(t)

	return &testFixture{
		ds:       datastoremock.NewMockDataStore(ctrl),
		userRepo: repositorymock.NewMockUser(ctrl),
		logger:   log.New(),
	}
}

func newUser(f *testFixture) *userUsecase {
	return &userUsecase{
		ds:     f.ds,
		logger: f.logger,
	}
}
