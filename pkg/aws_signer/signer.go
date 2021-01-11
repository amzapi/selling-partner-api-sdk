package aws_signer

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/google/uuid"
)

const (
	ServiceName = "execute-api"
)

// Config
type Config struct {
	AccessKeyID string //AWS IAM User Access Key Id
	SecretKey   string //AWS IAM User Secret Key
	Region      string //AWS Region
	RoleArn     string //AWS IAM Role ARN

	// Deprecated: Use WithRoleSessionName instead.
	RoleSessionName string //AWS IAM Role Session Name
}

// AWSSigV4Signer
type AWSSigV4Signer struct {
	cfg                    Config
	aws4Signer             *v4.Signer
	awsStsCredentials      *sts.Credentials
	awsStsAssumeRoleOutput *sts.AssumeRoleOutput
	awsSession             *session.Session
	newRoleSessionNameFn   NewRoleSessionNameFn
}

// AWSSigV4SignerOption
type AWSSigV4SignerOption func(*AWSSigV4Signer) error

type NewRoleSessionNameFn func() string

func WithRoleSessionName(fn NewRoleSessionNameFn) AWSSigV4SignerOption {
	return func(s *AWSSigV4Signer) error {
		s.newRoleSessionNameFn = fn
		return nil
	}
}

func NewSigner(cfg Config, opts ...AWSSigV4SignerOption) (*AWSSigV4Signer, error) {

	s := AWSSigV4Signer{
		cfg: cfg,
	}

	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}

	newSession, err := session.NewSession(
		&aws.Config{Credentials: credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretKey, "")},
	)

	if err != nil {
		return nil, err
	}

	s.cfg = cfg
	s.awsSession = newSession

	return &s, nil
}

func (s *AWSSigV4Signer) autoRefreshCredentials() error {

	roleSessionName := uuid.New().String()

	if s.newRoleSessionNameFn != nil {
		roleSessionName = s.newRoleSessionNameFn()
	}

	role, err := sts.New(s.awsSession).AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         aws.String(s.cfg.RoleArn),
		RoleSessionName: aws.String(roleSessionName),
	})

	if err != nil {
		return err
	}

	s.awsStsAssumeRoleOutput = role

	signer := v4.NewSigner(credentials.NewStaticCredentials(
		*role.Credentials.AccessKeyId,
		*role.Credentials.SecretAccessKey,
		*role.Credentials.SessionToken),
		func(s *v4.Signer) {
			s.DisableURIPathEscaping = true
		},
	)

	if signer != nil {
		s.aws4Signer = signer
	}

	return nil
}

func (s *AWSSigV4Signer) Sign(r *http.Request) error {
	if s.aws4Signer == nil || s.aws4Signer.Credentials.IsExpired() {
		if err := s.autoRefreshCredentials(); err != nil {
			return err
		}
	}
	r.Header.Add("X-Amz-Security-Token", *s.awsStsAssumeRoleOutput.Credentials.SessionToken)
	t := time.Now().UTC()
	var body io.ReadSeeker
	if r.Body != nil {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(payload))
		body = bytes.NewReader(payload)
	}
	_, err := s.aws4Signer.Sign(r, body, ServiceName, s.cfg.Region, t)
	return err
}
