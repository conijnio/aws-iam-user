package adapters

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
)

var (
	ClientError                  = &testtools.StubError{Err: errors.New("ClientError")}
	CallerIdentityIAMUserJohnDoe = testtools.Stub{
		OperationName: "GetCallerIdentity",
		Input:         &sts.GetCallerIdentityInput{},
		Output: &sts.GetCallerIdentityOutput{
			Account: aws.String("111122223333"),
			Arn:     aws.String("arn:aws:sts::111122223333:user/john.doe"),
			UserId:  aws.String("AIDAXXXXXXXXXXXXXXXXX"),
		},
	}
	CallerIdentityIAMRoleDeveloper = testtools.Stub{
		OperationName: "GetCallerIdentity",
		Input:         &sts.GetCallerIdentityInput{},
		Output: &sts.GetCallerIdentityOutput{
			Account: aws.String("111122223333"),
			Arn:     aws.String("arn:aws:sts::111122223333:assumed-role/Developer/john.doe"),
			UserId:  aws.String("AIDAXXXXXXXXXXXXXXXXX"),
		},
	}
	CallerIdentityInvalidARN = testtools.Stub{
		OperationName: "GetCallerIdentity",
		Input:         &sts.GetCallerIdentityInput{},
		Output: &sts.GetCallerIdentityOutput{
			Account: aws.String("111122223333"),
			Arn:     aws.String("Invalid ARN"),
			UserId:  aws.String("AIDAXXXXXXXXXXXXXXXXX"),
		},
	}
	CallerIdentityClientFailure = testtools.Stub{
		OperationName: "GetCallerIdentity",
		Input:         &sts.GetCallerIdentityInput{},
		Error:         ClientError,
	}
	SingleAccessKeys = testtools.Stub{
		OperationName: "ListAccessKeys",
		Input:         &iam.ListAccessKeysInput{UserName: aws.String("john.doe")},
		Output: &iam.ListAccessKeysOutput{
			AccessKeyMetadata: []types.AccessKeyMetadata{
				{AccessKeyId: aws.String("MyAccessKey")},
			},
		},
	}
	TwoAccessKeys = testtools.Stub{
		OperationName: "ListAccessKeys",
		Input:         &iam.ListAccessKeysInput{UserName: aws.String("john.doe")},
		Output: &iam.ListAccessKeysOutput{
			AccessKeyMetadata: []types.AccessKeyMetadata{
				{AccessKeyId: aws.String("MyAccessKey")},
				{AccessKeyId: aws.String("MyOtherAccessKey")},
			},
		},
	}
	CreateAccessKey = testtools.Stub{
		OperationName: "CreateAccessKey",
		Input:         &iam.CreateAccessKeyInput{UserName: aws.String("john.doe")},
		Output: &iam.CreateAccessKeyOutput{
			AccessKey: &types.AccessKey{
				AccessKeyId:     aws.String("MyNewAccessKey"),
				SecretAccessKey: aws.String("MyNewSecretAccessKey"),
				Status:          "Active",
			},
		},
	}
	CreateAccessKeyFailure = testtools.Stub{
		OperationName: "CreateAccessKey",
		Input:         &iam.CreateAccessKeyInput{UserName: aws.String("john.doe")},
		Error:         ClientError,
	}
	UpdateAccessKey = testtools.Stub{
		OperationName: "UpdateAccessKey",
		Input: &iam.UpdateAccessKeyInput{
			UserName:    aws.String("john.doe"),
			AccessKeyId: aws.String("MyAccessKey"),
			Status:      "Inactive",
		},
		Output: &iam.UpdateAccessKeyOutput{},
	}
	DeleteAccessKey = testtools.Stub{
		OperationName: "DeleteAccessKey",
		Input: &iam.DeleteAccessKeyInput{
			UserName:    aws.String("john.doe"),
			AccessKeyId: aws.String("MyAccessKey"),
		},
		Output: &iam.DeleteAccessKeyOutput{},
	}
)
