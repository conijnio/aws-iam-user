package adapters

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
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
)
