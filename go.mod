module github.com/numenory/selling-partner-api-sdk

go 1.16

require (
	github.com/aws/aws-sdk-go v1.38.47
	github.com/deepmap/oapi-codegen v1.7.0
	github.com/getkin/kin-openapi v0.62.0
	github.com/google/uuid v1.2.0
	github.com/labstack/echo/v4 v4.3.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	gopkg.me/selling-partner-api-sdk v0.0.0-20210312054314-cf9bdde067d8
)

replace gopkg.me/selling-partner-api-sdk v0.1.0 => github.com/numenory/selling-partner-api-sdk v0.1.0
