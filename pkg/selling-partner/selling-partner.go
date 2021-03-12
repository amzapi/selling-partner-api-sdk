package selling_partner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/service/sts"
)

type AccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type Config struct {
	ClientID     string //SP-API
	ClientSecret string //SP-API
	RefreshToken string //
	AccessKeyID  string //AWS IAM User Access Key Id
	SecretKey    string //AWS IAM User Secret Key
	Region       string //AWS Region
	RoleArn      string //AWS IAM Role ARN
}

func (o Config) IsValid() (bool, error) {
	if o.RefreshToken == "" {
		return false, errors.New("refresh token is required")
	}
	if o.ClientID == "" {
		return false, errors.New("client id is required")
	}
	if o.ClientSecret == "" {
		return false, errors.New("client secret is required")
	}
	if o.AccessKeyID == "" {
		return false, errors.New("aws iam user access key id is required")
	}
	if o.SecretKey == "" {
		return false, errors.New("aws iam user secret key is required")
	}
	if o.RoleArn == "" {
		return false, errors.New("aws iam role arn is required")
	}
	if doesMatch, err := regexp.MatchString("^(eu-west-1|us-east-1|us-west-2)$", o.Region); !doesMatch || err != nil {
		return false, errors.New("region should be one of eu-west-1, us-east-1, or us-west-2")
	}
	return true, nil
}

type SellingPartner struct {
	cfg               *Config
	accessToken       string
	accessTokenExpiry time.Time
	aws4Signer        *v4.Signer
	awsStsCredentials *sts.Credentials
	awsSession        *session.Session
}

func NewSellingPartner(cfg *Config) (*SellingPartner, error) {
	if isValid, err := cfg.IsValid(); !isValid {
		return nil, err
	}

	newSession, err := session.NewSession(
		&aws.Config{Credentials: credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretKey, "")},
	)

	if err != nil {
		return nil, errors.New("NewSellingPartner call failed with error " + err.Error())
	}

	sp := &SellingPartner{}
	sp.cfg = cfg
	sp.awsSession = newSession

	return sp, nil
}

func (s *SellingPartner) RefreshToken() error {

	reqBody, _ := json.Marshal(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": s.cfg.RefreshToken,
		"client_id":     s.cfg.ClientID,
		"client_secret": s.cfg.ClientSecret,
	})

	resp, err := http.Post(
		"https://api.amazon.com/auth/o2/token",
		"application/json",
		bytes.NewBuffer(reqBody))

	if err != nil {
		return errors.New("RefreshToken call failed with error " + err.Error())
	}

	defer resp.Body.Close()

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("RefreshToken read response with error " + err.Error())
	}

	theResp := &AccessTokenResponse{}

	if err = json.Unmarshal(respBodyBytes, theResp); err != nil {
		return errors.New("RefreshToken response parse failed. Body: " + string(respBodyBytes))
	}

	if theResp.AccessToken != "" {
		s.accessToken = theResp.AccessToken
		s.accessTokenExpiry = time.Now().UTC().Add(time.Duration(theResp.ExpiresIn) * time.Second) //set expiration time
	} else if theResp.Error != "" {
		return errors.New(fmt.Sprintf("RefreshToken failed with code %s, description %s", theResp.Error, theResp.ErrorDescription))
	} else {
		return errors.New(fmt.Sprintf("RefreshToken failed with unknown reason. Body: %s", string(respBodyBytes)))
	}

	return nil
}

func (s *SellingPartner) RefreshCredentials() error {

	roleSessionName := uuid.New().String()

	role, err := sts.New(s.awsSession).AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         aws.String(s.cfg.RoleArn),
		RoleSessionName: aws.String(roleSessionName),
	})

	if err != nil {
		return errors.New("RefreshCredentials call failed with error " + err.Error())
	}

	if role == nil || role.Credentials == nil {
		return errors.New("AssumeRole call failed in return")
	}

	s.awsStsCredentials = role.Credentials

	s.aws4Signer = v4.NewSigner(credentials.NewStaticCredentials(
		*role.Credentials.AccessKeyId,
		*role.Credentials.SecretAccessKey,
		*role.Credentials.SessionToken),
		func(s *v4.Signer) {
			s.DisableURIPathEscaping = true
		},
	)

	return nil
}

// expiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expirations due to client-server time mismatches.
const expiryDelta = 1 * time.Minute

func (s *SellingPartner) SignRequest(r *http.Request) error {

	if s.accessToken == "" ||
		s.accessTokenExpiry.IsZero() ||
		s.accessTokenExpiry.Round(0).Add(-expiryDelta).Before(time.Now().UTC()) {
		if err := s.RefreshToken(); err != nil {
			return fmt.Errorf("cannot refresh token. Error: %s", err.Error())
		}
	}

	if s.aws4Signer == nil ||
		s.awsStsCredentials == nil ||
		s.aws4Signer.Credentials.IsExpired() ||
		s.awsStsCredentials.Expiration.IsZero() ||
		s.awsStsCredentials.Expiration.Round(0).Add(-expiryDelta).Before(time.Now().UTC()) {
		if err := s.RefreshCredentials(); err != nil {
			return fmt.Errorf("cannot refresh role credentials. Error: %s", err.Error())
		}
	}

	var body io.ReadSeeker
	if r.Body != nil {
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(payload))
		body = bytes.NewReader(payload)
	}

	r.Header.Add("X-Amz-Access-Token", s.accessToken)

	_, err := s.aws4Signer.Sign(r, body, "execute-api", s.cfg.Region, time.Now().UTC())

	return err
}
