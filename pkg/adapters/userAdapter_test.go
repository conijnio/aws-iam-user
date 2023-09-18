package adapters

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
	"testing"
)

func TestLoadAdapter(t *testing.T) {
	adapter := LoadUserAdapter("eu-west-1", "default")

	if adapter != GetAdapter("eu-west-1", "default") {
		t.Errorf("Expected the LoadAdapter and GetAdapter methods to return the same object")
	}
}

func TestLoadAdapterNewRegion(t *testing.T) {
	adapter := LoadUserAdapter("eu-central-1", "default")

	if adapter != GetAdapter("eu-central-1", "default") {
		t.Errorf("Expected the LoadAdapter and GetAdapter methods to return the same object")
	}
}

func TestUserIAMUserJohnDoe(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityIAMUserJohnDoe)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)
	user, _ := adapter.LoadUser()

	if user.IsRole() {
		t.Error("Expected IsRole() to return False")
	}

	if !user.IsUser() {
		t.Error("Expected IsUser() to return True")
	}

	testtools.ExitTest(stubber, t)
}

func TestUserIAMRoleDeveloper(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityIAMRoleDeveloper)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)
	user, err := adapter.LoadUser()

	if err != nil {
		t.Errorf("No error was expected but received: %s", err)
	}

	if !user.IsRole() {
		t.Error("Expected IsRole() to return True")
	}

	if user.IsUser() {
		t.Error("Expected IsUser() to return False")
	}

	testtools.ExitTest(stubber, t)
}

func TestUserWithInvalidARN(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityInvalidARN)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)
	user, _ := adapter.LoadUser()

	if user.IsRole() {
		t.Error("Expected IsRole() to return False")
	}

	if user.IsUser() {
		t.Error("Expected IsUser() to return False")
	}

	testtools.ExitTest(stubber, t)
}

func TestUserWithClientFailure(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityClientFailure)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)

	_, err := adapter.LoadUser()

	testtools.VerifyError(err, ClientError, t)

	testtools.ExitTest(stubber, t)
}

type MockConfig struct {
	failUpdate bool
}

func (c *MockConfig) AccessKeyId() string {
	return "MyAccessKey"
}

func (c *MockConfig) UpdateKeys(accessKey string, secretAccessKey string) error {
	if c.failUpdate {
		return errors.New("failed to update the credentials")
	}
	return nil
}

func (c *MockConfig) Rollback(err error) error {
	return err
}

func TestRotateCredentials(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityIAMUserJohnDoe)
	stubber.Add(SingleAccessKeys)
	stubber.Add(CreateAccessKey)
	stubber.Add(CallerIdentityIAMUserJohnDoe)
	stubber.Add(UpdateAccessKey)
	stubber.Add(DeleteAccessKey)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)
	adapter.iamClient = iam.NewFromConfig(*stubber.SdkConfig)

	adapter.config = &MockConfig{}

	user, _ := adapter.LoadUser()
	if err := adapter.RotateCredentials(user); err != nil {
		t.Error("expected RotateCredentials() not to return an error")
	}

	testtools.ExitTest(stubber, t)
}

func TestRotateCredentialsTwoAccessKeys(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityIAMUserJohnDoe)
	stubber.Add(TwoAccessKeys)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)
	adapter.iamClient = iam.NewFromConfig(*stubber.SdkConfig)

	adapter.config = &MockConfig{}

	user, _ := adapter.LoadUser()
	if err := adapter.RotateCredentials(user); err == nil {
		t.Error("expected RotateCredentials() to return an error")
	}

	testtools.ExitTest(stubber, t)
}

func TestRotateCredentialsCreateFailure(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityIAMUserJohnDoe)
	stubber.Add(SingleAccessKeys)
	stubber.Add(CreateAccessKeyFailure)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)
	adapter.iamClient = iam.NewFromConfig(*stubber.SdkConfig)

	adapter.config = &MockConfig{}

	user, _ := adapter.LoadUser()
	if err := adapter.RotateCredentials(user); err == nil {
		t.Error("expected RotateCredentials() to return an error")
	}

	testtools.ExitTest(stubber, t)
}

func TestRotateCredentialsConfigUpdateFailure(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityIAMUserJohnDoe)
	stubber.Add(SingleAccessKeys)
	stubber.Add(CreateAccessKey)
	stubber.Add(CallerIdentityClientFailure)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)
	adapter.iamClient = iam.NewFromConfig(*stubber.SdkConfig)

	adapter.config = &MockConfig{}

	user, _ := adapter.LoadUser()
	if err := adapter.RotateCredentials(user); err == nil {
		t.Error("expected RotateCredentials() to return an error")
	}

	testtools.ExitTest(stubber, t)
}

func TestRotateCredentialsNewCredentialFailure(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityIAMUserJohnDoe)
	stubber.Add(SingleAccessKeys)
	stubber.Add(CreateAccessKey)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)
	adapter.iamClient = iam.NewFromConfig(*stubber.SdkConfig)

	adapter.config = &MockConfig{failUpdate: true}

	user, _ := adapter.LoadUser()
	if err := adapter.RotateCredentials(user); err == nil {
		t.Error("expected RotateCredentials() to return an error")
	}

	testtools.ExitTest(stubber, t)
}
