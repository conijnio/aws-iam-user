package adapters

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/conijnio/aws-iam-user/pkg/models"
	log "github.com/sirupsen/logrus"
	"regexp"
)

type IUserAdapter interface {
	LoadUser() (*models.User, error)
}

type UserAdapter struct {
	stsClient *sts.Client
}

func LoadUserAdapter(region string, profile string) IUserAdapter {
	prepareRegionMap(region)

	if !HasAdapter(region, profile) {
		RegisterAdapter(region, profile, New(region, profile))
	}

	return GetAdapter(region, profile)
}

func New(profile string, region string) *UserAdapter {
	proxy := &UserAdapter{}
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile), config.WithRegion(region))

	if err == nil {
		proxy.stsClient = sts.NewFromConfig(cfg)
	}

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

	log.Debug("Resolved the user:", user)

	return user, nil
}

func (u *UserAdapter) resolveFromArn(arn string) (string, string) {
	log.Debug("Resolving the user `Name` based on the ARN")
	regex := regexp.MustCompile(`arn:aws:\w+::\d+:(user|assumed-role)/(.*/|)(.*)`)
	matches := regex.FindStringSubmatch(arn)

	if len(matches) == 4 {
		return matches[1], matches[3]
	}

	return "", ""
}
