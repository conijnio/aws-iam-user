package adapters

import (
	"errors"
	"github.com/conijnio/aws-iam-user/pkg/models"
	"testing"
)

type MockAdapter struct {
	LoadUserFailure bool
}

func (u *MockAdapter) LoadUser() (*models.User, error) {
	if u.LoadUserFailure {
		return &models.User{}, errors.New("failed")
	}
	return &models.User{}, nil
}

func (u *MockAdapter) RotateCredentials(user *models.User) error {
	return nil
}

func TestAdapterRegistration(t *testing.T) {
	adapter := &MockAdapter{}
	RegisterAdapter("eu-west-1", "default", adapter)

	if res := HasAdapter("eu-west-1", "default"); !res {
		t.Errorf("Expected True received %t", res)
	}

	if GetAdapter("eu-west-1", "default") != adapter {
		t.Errorf("Expected the given adapter to be returned")
	}
}

func TestSecondAdapterRegistrationDifferentProfile(t *testing.T) {
	adapter := &MockAdapter{}
	RegisterAdapter("eu-west-1", "default", adapter)

	if res := HasAdapter("eu-west-1", "default"); !res {
		t.Errorf("Expected True received %t", res)
	}

	if GetAdapter("eu-west-1", "default") != adapter {
		t.Errorf("Expected the given adapter to be returned")
	}
}

func TestThirdAdapterRegistrationDifferentProfile(t *testing.T) {
	adapter := &MockAdapter{}
	RegisterAdapter("eu-west-1", "default", adapter)

	if res := HasAdapter("eu-west-1", "default"); !res {
		t.Errorf("Expected True received %t", res)
	}

	if GetAdapter("eu-west-1", "default") != adapter {
		t.Errorf("Expected the given adapter to be returned")
	}
}

func TestHasUnknownProfileAdapter(t *testing.T) {
	if res := HasAdapter("eu-west-1", "non-existing-adapter"); res {
		t.Errorf("Expected False received %t", res)
	}
}

func TestHasUnknownRegionAdapter(t *testing.T) {
	if res := HasAdapter("eu-unknown-1", "non-existing-adapter"); res {
		t.Errorf("Expected False received %t", res)
	}
}
