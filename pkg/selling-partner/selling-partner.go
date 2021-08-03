package selling_partner

import (
	"bytes"
	"context"
	"encoding/json"
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
	ClientID             string    //SP-API
	ClientSecret         string    //SP-API
	AccessToken          string    //
	AccessTokenExpiry    time.Time //
	RefreshToken         string    //
	AccessKeyID          string    //AWS IAM User Access Key Id
	SecretKey            string    //AWS IAM User Secret Key
	Region               string    //AWS Region
	RoleArn              string    //AWS IAM Role ARN
	RoleCredentials      *sts.Credentials
	OnRefreshToken       func(string, time.Time) error
	OnRefreshCredentials func(*sts.Credentials) error
}

func (o Config) IsValid() (bool, error) {
	if o.RefreshToken == "" {
		return false, fmt.Errorf("refresh token is required")
	}
	if o.ClientID == "" {
		return false, fmt.Errorf("client id is required")
	}
	if o.ClientSecret == "" {
		return false, fmt.Errorf("client secret is required")
	}
	if o.AccessKeyID == "" {
		return false, fmt.Errorf("aws iam user access key id is required")
	}
	if o.SecretKey == "" {
		return false, fmt.Errorf("aws iam user secret key is required")
	}
	if o.RoleArn == "" {
		return false, fmt.Errorf("aws iam role arn is required")
	}
	if doesMatch, err := regexp.MatchString("^(eu-west-1|us-east-1|us-west-2)$", o.Region); !doesMatch || err != nil {
		return false, fmt.Errorf("region should be one of eu-west-1, us-east-1, or us-west-2")
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

type RequestBeforeFn func(ctx context.Context, req *http.Request) error

func NewSellingPartner(cfg *Config) (*SellingPartner, error) {
	if isValid, err := cfg.IsValid(); !isValid {
		return nil, err
	}

	newSession, err := session.NewSession(
		&aws.Config{Credentials: credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretKey, "")},
	)

	if err != nil {
		return nil, fmt.Errorf("NewSellingPartner call failed with error %w", err)
	}

	sp := &SellingPartner{
		cfg:        cfg,
		awsSession: newSession,
	}

	if cfg.AccessToken != "" && cfg.AccessTokenExpiry.After(time.Now()) {
		sp.accessToken = cfg.AccessToken
		sp.accessTokenExpiry = cfg.AccessTokenExpiry
	}

	if cfg.RoleCredentials != nil {
		sp.setRoleCredentials(cfg.RoleCredentials)
	}

	return sp, nil
}

func (s *SellingPartner) RefreshToken() error {
	_, err := s.makeTokenRequest(
		context.Background(),
		map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": s.cfg.RefreshToken,
			"client_id":     s.cfg.ClientID,
			"client_secret": s.cfg.ClientSecret,
		},
		nil,
	)

	return err
}

// GetTokensFromAuthorizationCode converts an authorization code (obtained with
// authorization/GetAuthorizationCode) to an access and refresh token.
func (s *SellingPartner) GetTokensFromAuthorizationCode(ctx context.Context, code string, beforeRequest RequestBeforeFn) (*AccessTokenResponse, error) {
	return s.makeTokenRequest(
		ctx,
		map[string]string{
			"grant_type":    "authorization_code",
			"client_id":     s.cfg.ClientID,
			"client_secret": s.cfg.ClientSecret,
			"code":          code,
		},
		beforeRequest,
	)
}

// GetMigrationAccessToken requests a new access token with the migration scope.
func (s *SellingPartner) GetMigrationAccessToken(ctx context.Context, beforeRequest RequestBeforeFn) (*AccessTokenResponse, error) {
	return s.makeTokenRequest(
		ctx,
		map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     s.cfg.ClientID,
			"client_secret": s.cfg.ClientSecret,
			"scope":         "sellingpartnerapi::migration",
		},
		beforeRequest,
	)
}

func (s *SellingPartner) makeTokenRequest(ctx context.Context, data map[string]string, beforeRequest RequestBeforeFn) (*AccessTokenResponse, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body; %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.amazon.com/auth/o2/token", bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to create a new request; %w", err)
	}

	req = req.WithContext(ctx)

	if beforeRequest != nil {
		if err := beforeRequest(ctx, req); err != nil {
			return nil, fmt.Errorf("failed to call beforeRequest callback; %w", err)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("token call failed with error %w", err)
	}

	defer resp.Body.Close()

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("RefreshToken read response with error %w", err)
	}

	theResp := &AccessTokenResponse{}

	if err = json.Unmarshal(respBodyBytes, theResp); err != nil {
		return nil, fmt.Errorf("RefreshToken response parse failed. Body: " + string(respBodyBytes))
	}

	if theResp.AccessToken != "" {
		s.accessToken = theResp.AccessToken
		s.accessTokenExpiry = time.Now().UTC().Add(time.Duration(theResp.ExpiresIn) * time.Second) //set expiration time
	} else if theResp.Error != "" {
		return nil, fmt.Errorf("RefreshToken failed with code %s, description %s", theResp.Error, theResp.ErrorDescription)
	} else {
		return nil, fmt.Errorf("RefreshToken failed with unknown reason. Body: %s", string(respBodyBytes))
	}

	if s.cfg.OnRefreshToken != nil {
		if err := s.cfg.OnRefreshToken(s.accessToken, s.accessTokenExpiry); err != nil {
			return nil, fmt.Errorf("Failed to call the OnRefreshToken callback with error %w", err)
		}
	}

	return theResp, nil
}

func (s *SellingPartner) RefreshCredentials() error {
	roleSessionName := uuid.New().String()

	role, err := sts.New(s.awsSession).AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         aws.String(s.cfg.RoleArn),
		RoleSessionName: aws.String(roleSessionName),
	})

	if err != nil {
		return fmt.Errorf("RefreshCredentials call failed with error %w", err)
	}

	if role == nil || role.Credentials == nil {
		return fmt.Errorf("AssumeRole call failed in return")
	}

	s.setRoleCredentials(role.Credentials)

	if s.cfg.OnRefreshCredentials != nil {
		if err := s.cfg.OnRefreshCredentials(role.Credentials); err != nil {
			return fmt.Errorf("Failed to call the OnRefreshCredentials callback with error %w", err)
		}
	}

	return nil
}

func (s *SellingPartner) setRoleCredentials(c *sts.Credentials) {
	s.awsStsCredentials = c

	s.aws4Signer = v4.NewSigner(credentials.NewStaticCredentials(
		*c.AccessKeyId,
		*c.SecretAccessKey,
		*c.SessionToken),
		func(s *v4.Signer) {
			s.DisableURIPathEscaping = true
		},
	)
}

// expiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expirations due to client-server time mismatches.
const expiryDelta = 1 * time.Minute

func (s *SellingPartner) SignRequest(r *http.Request, withAccessToken bool) error {
	if withAccessToken {
		if s.accessToken == "" ||
			s.accessTokenExpiry.IsZero() ||
			s.accessTokenExpiry.Round(0).Add(-expiryDelta).Before(time.Now().UTC()) {
			if err := s.RefreshToken(); err != nil {
				return fmt.Errorf("cannot refresh token. Error: %w", err)
			}
		}
	}

	if s.aws4Signer == nil ||
		s.awsStsCredentials == nil ||
		s.aws4Signer.Credentials.IsExpired() ||
		s.awsStsCredentials.Expiration.IsZero() ||
		s.awsStsCredentials.Expiration.Round(0).Add(-expiryDelta).Before(time.Now().UTC()) {
		if err := s.RefreshCredentials(); err != nil {
			return fmt.Errorf("cannot refresh role credentials. Error: %w", err)
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

	if withAccessToken {
		r.Header.Add("X-Amz-Access-Token", s.accessToken)
	}

	_, err := s.aws4Signer.Sign(r, body, "execute-api", s.cfg.Region, time.Now().UTC())

	return err
}
