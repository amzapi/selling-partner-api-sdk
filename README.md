# Amazon's Selling Partner API (SP-API) Golang SDK

[![Go Reference](https://pkg.go.dev/badge/gopkg.me/selling-partner-api-sdk.svg)](https://pkg.go.dev/gopkg.me/selling-partner-api-sdk)
[![Build Status](https://travis-ci.com/gopkg-dev/selling-partner-api-sdk.svg?branch=master)](https://travis-ci.com/gopkg-dev/selling-partner-api-sdk)

## Installation

~~~~
go get -u gopkg.me/selling-partner-api-sdk
~~~~

## Example

```go
package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"gopkg.me/selling-partner-api-sdk/pkg/aws_signer"
	"gopkg.me/selling-partner-api-sdk/pkg/oauth2"
	"gopkg.me/selling-partner-api-sdk/sellers"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func InitializeHeaders(req *http.Request, accessToken, requestID string) {
	req.Header.Add("x-amz-access-token", accessToken)
	req.Header.Add("x-amz-date", time.Now().UTC().Format("20060102T150405Z")
	req.Header.Add("X-Amzn-Requestid", requestID)
}

func main() {

	token, err := oauth2.RefreshAccessToken(oauth2.Config{
		ClientID:     "<ClientID>",
		ClientSecret: "<ClientSecret>",
		RefreshToken: "<RefreshToken>",
		GrantType:    "refresh_token",
		Scope:        "",
	})

	if err != nil {
		panic(err)
	}

	if token.Valid() == false {
		panic("error")
	}

	aws4Signer, err := aws_signer.NewSigner(aws_signer.Config{
		AccessKeyID: "<AWS IAM User Access Key Id>",
		SecretKey:   "<AWS IAM User Secret Key>",
		Region:      "<AWS Region>",
		RoleArn:     "<AWS IAM Role ARN>",
	}, aws_signer.WithRoleSessionName(func() string {
		return uuid.New().String()
	}))

	if err != nil {
		panic(err)
	}

	endpoint := "https://sellingpartnerapi-fe.amazon.com"

	seller, err := sellers.NewClientWithResponses(endpoint,
		sellers.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
			InitializeHeaders(req, token.AccessToken, uuid.New().String())
			err = aws4Signer.Sign(req)
			if err != nil {
				return errors.Wrap(err, "aws4Signer.Sign")
			}
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				return errors.Wrap(err, "DumpRequest Error")
			}
			log.Printf("DumpRequest = %s", dump)
			return nil
		}),
		sellers.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
			dump, err := httputil.DumpResponse(rsp, true)
			if err != nil {
				return errors.Wrap(err, "DumpResponse Error")
			}
			log.Printf("DumpResponse = %s", dump)
			return nil
		}),
	)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = seller.GetMarketplaceParticipationsWithResponse(ctx)

	if err != nil {
		panic(err)
	}
}

```
