package adapters

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/conijnio/aws-iam-user/pkg/config"
	"github.com/conijnio/aws-iam-user/pkg/models"
	log "github.com/sirupsen/logrus"
	"regexp"
)

type IUserAdapter interface {
	LoadUser() (*models.User, error)
	RotateCredentials(*models.User) error
}

type UserAdapter struct {
	config    config.IConfig
	stsClient *sts.Client
	iamClient *iam.Client
	profile   string
	region    string
}

func LoadUserAdapter(region string, profile string) IUserAdapter {
	prepareRegionMap(region)

	if !HasAdapter(region, profile) {
		RegisterAdapter(region, profile, New(region, profile))
	}

	return GetAdapter(region, profile)
}

func New(region string, profile string) *UserAdapter {
	log.Debugf("Initialize a new client using the %s profile and the %s region", profile, region)
	proxy := &UserAdapter{profile: profile, region: region}
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithSharedConfigProfile(profile), awsConfig.WithRegion(region))

	if err == nil {
		proxy.stsClient = sts.NewFromConfig(cfg)
		proxy.iamClient = iam.NewFromConfig(cfg)
	}

	proxy.config = config.New(proxy.profile)

	return proxy
}

func (u *UserAdapter) LoadUser() (*models.User, error) {
	user := &models.User{}

	log.Debug("Fetching caller identity using the AWS API")
	identity, err := u.stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})

	if err != nil {
		log.Errorf("Failed: %s", err)
		return user, err
	}

	user.Account = *identity.Account
	user.UserId = *identity.UserId
	user.Arn = *identity.Arn
	user.Type, user.Name = u.resolveFromArn(user.Arn)

	log.Debugf("Resolved the user: %s", user)

	return user, nil
}

func (u *UserAdapter) resolveFromArn(arn string) (string, string) {
	regex := regexp.MustCompile(`arn:aws:\w+::\d+:(user|assumed-role)/(.*/|)(.*)`)
	matches := regex.FindStringSubmatch(arn)

	if len(matches) == 4 {
		return matches[1], matches[3]
	} else {
		log.Warningf("Failed to parse: %s", arn)
	}

	return "", ""
}

func (u *UserAdapter) canRotate(user *models.User) (bool, error) {
	log.Debugf("Lookup the current access keys for %s", user.Name)
	listAccessKeyOutput, err := u.iamClient.ListAccessKeys(context.TODO(), &iam.ListAccessKeysInput{UserName: aws.String(user.Name)})
	log.Debugf("Lookup the current access keys for %d", len(listAccessKeyOutput.AccessKeyMetadata))

	if len(listAccessKeyOutput.AccessKeyMetadata) != 1 {
		message := fmt.Sprintf("there are already %d keys for this user so rotation is not possible", len(listAccessKeyOutput.AccessKeyMetadata))
		err = errors.New(message)
	}

	return len(listAccessKeyOutput.AccessKeyMetadata) == 1, err
}

func (u *UserAdapter) createAccessKey(user *models.User) error {
	if res, err := u.canRotate(user); !res {
		return err
	}

	createAccessKeyOutput, err := u.iamClient.CreateAccessKey(
		context.TODO(),
		&iam.CreateAccessKeyInput{UserName: aws.String(user.Name)})

	if err != nil {
		return err
	}

	log.Debugf("Create new access key: %s", *createAccessKeyOutput.AccessKey.AccessKeyId)

	return u.config.UpdateKeys(*createAccessKeyOutput.AccessKey.AccessKeyId, *createAccessKeyOutput.AccessKey.SecretAccessKey)
}

func (u *UserAdapter) disableAccessKey(user *models.User) error {
	updateAccessKeyInput := &iam.UpdateAccessKeyInput{
		UserName:    aws.String(user.Name),
		AccessKeyId: aws.String(u.config.AccessKeyId()),
		Status:      "Inactive",
	}

	log.Debugf("Disable old access key id: %s", u.config.AccessKeyId())
	_, err := u.iamClient.UpdateAccessKey(context.TODO(), updateAccessKeyInput)

	return err
}

func (u *UserAdapter) removeAccessKey(user *models.User) error {
	if err := u.disableAccessKey(user); err != nil {
		return err
	}

	deleteAccessKeyInput := &iam.DeleteAccessKeyInput{
		UserName:    aws.String(user.Name),
		AccessKeyId: aws.String(u.config.AccessKeyId()),
	}

	log.Debugf("Delete old access key id: %s", u.config.AccessKeyId())
	_, err := u.iamClient.DeleteAccessKey(context.TODO(), deleteAccessKeyInput)

	return err
}

func (u *UserAdapter) RotateCredentials(user *models.User) error {
	log.Debugf("Rotate credentials for %s", user.Name)

	if err := u.createAccessKey(user); err != nil {
		return err
	}

	if _, err := u.LoadUser(); err != nil {
		log.Errorf("New credentials are not working properly: %s", err)
		return u.config.Rollback(err)
	}

	return u.removeAccessKey(user)
}
