package adapters

import (
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

	// Validate that all stubs are called
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

	// Validate that all stubs are called
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

	// Validate that all stubs are called
	testtools.ExitTest(stubber, t)
}

func TestUserWithClientFailure(t *testing.T) {
	stubber := testtools.NewStubber()
	stubber.Add(CallerIdentityClientFailure)

	adapter := New("eu-west-1", "default")
	adapter.stsClient = sts.NewFromConfig(*stubber.SdkConfig)

	_, err := adapter.LoadUser()

	testtools.VerifyError(err, ClientError, t)

	// Validate that all stubs are called
	testtools.ExitTest(stubber, t)
}
