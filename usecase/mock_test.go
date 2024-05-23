package usecase

import (
	"testing"

	gomock "go.uber.org/mock/gomock"

	datastoremock "github.com/OLIENTTECH/backend-challenges/mock/datastore"
	repositorymock "github.com/OLIENTTECH/backend-challenges/mock/repository"
	"github.com/OLIENTTECH/backend-challenges/pkg/log"
)

type testFixture struct {
	ds       *datastoremock.MockDataStore
	userRepo *repositorymock.MockUser
	shopRepo *repositorymock.MockShop
	logger   *log.Logger
}

func newTestFixture(t *testing.T) *testFixture {
	t.Helper()
	ctrl := gomock.NewController(t)

	return &testFixture{
		ds:       datastoremock.NewMockDataStore(ctrl),
		userRepo: repositorymock.NewMockUser(ctrl),
		shopRepo: repositorymock.NewMockShop(ctrl),
		logger:   log.New(),
	}
}

func newUser(f *testFixture) *userUsecase {
	return &userUsecase{
		ds:     f.ds,
		logger: f.logger,
	}
}
