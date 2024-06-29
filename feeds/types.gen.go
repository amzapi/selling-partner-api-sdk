// Package feeds provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package feeds

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// CancelFeedResponse defines model for CancelFeedResponse.
type CancelFeedResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`
}

// CreateFeedDocumentResponse defines model for CreateFeedDocumentResponse.
type CreateFeedDocumentResponse struct {

	// The identifier of the feed document.
	FeedDocumentId string `json:"feedDocumentId"`

	// The presigned URL for uploading the feed contents. This URL expires after 5 minutes.
	Url string `json:"url"`
}

// CreateFeedDocumentResult defines model for CreateFeedDocumentResult.
type CreateFeedDocumentResult struct {

	// Encryption details for required client-side encryption and decryption of document contents.
	EncryptionDetails FeedDocumentEncryptionDetails `json:"encryptionDetails"`

	// The identifier of the feed document.
	FeedDocumentId string `json:"feedDocumentId"`

	// The presigned URL for uploading the feed contents. This URL expires after 5 minutes.
	Url string `json:"url"`
}

// CreateFeedDocumentSpecification defines model for CreateFeedDocumentSpecification.
type CreateFeedDocumentSpecification struct {

	// The content type of the feed.
	ContentType string `json:"contentType"`
}

// CreateFeedResponse defines model for CreateFeedResponse.
type CreateFeedResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList        `json:"errors,omitempty"`
	Payload *CreateFeedResult `json:"payload,omitempty"`
}

// CreateFeedResult defines model for CreateFeedResult.
type CreateFeedResult struct {

	// The identifier for the feed. This identifier is unique only in combination with a seller ID.
	FeedId string `json:"feedId"`
}

// CreateFeedSpecification defines model for CreateFeedSpecification.
type CreateFeedSpecification struct {

	// Additional options to control the feed. These vary by feed type.
	FeedOptions *FeedOptions `json:"feedOptions,omitempty"`

	// The feed type.
	FeedType string `json:"feedType"`

	// The document identifier returned by the createFeedDocument operation. Encrypt and upload the feed document contents before calling the createFeed operation.
	InputFeedDocumentId string `json:"inputFeedDocumentId"`

	// A list of identifiers for marketplaces that you want the feed to be applied to.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList []Error

// Feed defines model for Feed.
type Feed struct {

	// The date and time when the feed was created, in ISO 8601 date time format.
	CreatedTime time.Time `json:"createdTime"`

	// The identifier for the feed. This identifier is unique only in combination with a seller ID.
	FeedId string `json:"feedId"`

	// The feed type.
	FeedType string `json:"feedType"`

	// A list of identifiers for the marketplaces that the feed is applied to.
	MarketplaceIds *[]string `json:"marketplaceIds,omitempty"`

	// The date and time when feed processing completed, in ISO 8601 date time format.
	ProcessingEndTime *time.Time `json:"processingEndTime,omitempty"`

	// The date and time when feed processing started, in ISO 8601 date time format.
	ProcessingStartTime *time.Time `json:"processingStartTime,omitempty"`

	// The processing status of the feed.
	ProcessingStatus string `json:"processingStatus"`

	// The identifier for the feed document. This identifier is unique only in combination with a seller ID.
	ResultFeedDocumentId *string `json:"resultFeedDocumentId,omitempty"`
}

// FeedDocument defines model for FeedDocument.
type FeedDocument struct {

	// If present, the feed document contents are compressed using the indicated algorithm.
	CompressionAlgorithm *string `json:"compressionAlgorithm,omitempty"`

	// Encryption details for required client-side encryption and decryption of document contents.
	EncryptionDetails FeedDocumentEncryptionDetails `json:"encryptionDetails"`

	// The identifier for the feed document. This identifier is unique only in combination with a seller ID.
	FeedDocumentId string `json:"feedDocumentId"`

	// A presigned URL for the feed document. This URL expires after 5 minutes.
	Url string `json:"url"`
}

// FeedDocumentEncryptionDetails defines model for FeedDocumentEncryptionDetails.
type FeedDocumentEncryptionDetails struct {

	// The vector to encrypt or decrypt the document contents using Cipher Block Chaining (CBC).
	InitializationVector string `json:"initializationVector"`

	// The encryption key used to encrypt or decrypt the document contents.
	Key string `json:"key"`

	// The encryption standard required to encrypt or decrypt the document contents.
	Standard string `json:"standard"`
}

// FeedList defines model for FeedList.
type FeedList []Feed

// FeedOptions defines model for FeedOptions.
type FeedOptions struct {
	AdditionalProperties map[string]string `json:"-"`
}

// GetFeedDocumentResponse defines model for GetFeedDocumentResponse.
type GetFeedDocumentResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList    `json:"errors,omitempty"`
	Payload *FeedDocument `json:"payload,omitempty"`
}

// GetFeedResponse defines model for GetFeedResponse.
type GetFeedResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList `json:"errors,omitempty"`
	Payload *Feed      `json:"payload,omitempty"`
}

// GetFeedsResponse defines model for GetFeedsResponse.
type GetFeedsResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// Returned when the number of results exceeds pageSize. To get the next page of results, call the getFeeds operation with this token as the only parameter.
	NextToken *string   `json:"nextToken,omitempty"`
	Payload   *FeedList `json:"payload,omitempty"`
}

// CreateFeedDocumentJSONBody defines parameters for CreateFeedDocument.
type CreateFeedDocumentJSONBody CreateFeedDocumentSpecification

// GetFeedsParams defines parameters for GetFeeds.
type GetFeedsParams struct {

	// A list of feed types used to filter feeds. When feedTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either feedTypes or nextToken is required.
	FeedTypes *[]string `json:"feedTypes,omitempty"`

	// A list of marketplace identifiers used to filter feeds. The feeds returned will match at least one of the marketplaces that you specify.
	MarketplaceIds *[]string `json:"marketplaceIds,omitempty"`

	// The maximum number of feeds to return in a single call.
	PageSize *int `json:"pageSize,omitempty"`

	// A list of processing statuses used to filter feeds.
	ProcessingStatuses *[]string `json:"processingStatuses,omitempty"`

	// The earliest feed creation date and time for feeds included in the response, in ISO 8601 format. The default is 90 days ago. Feeds are retained for a maximum of 90 days.
	CreatedSince *time.Time `json:"createdSince,omitempty"`

	// The latest feed creation date and time for feeds included in the response, in ISO 8601 format. The default is now.
	CreatedUntil *time.Time `json:"createdUntil,omitempty"`

	// A string token returned in the response to your previous request. nextToken is returned when the number of results exceeds the specified pageSize value. To get the next page of results, call the getFeeds operation and include this token as the only parameter. Specifying nextToken with any other parameters will cause the request to fail.
	NextToken *string `json:"nextToken,omitempty"`
}

// CreateFeedJSONBody defines parameters for CreateFeed.
type CreateFeedJSONBody CreateFeedSpecification

// CreateFeedDocumentRequestBody defines body for CreateFeedDocument for application/json ContentType.
type CreateFeedDocumentJSONRequestBody CreateFeedDocumentJSONBody

// CreateFeedRequestBody defines body for CreateFeed for application/json ContentType.
type CreateFeedJSONRequestBody CreateFeedJSONBody

// Getter for additional properties for FeedOptions. Returns the specified
// element and whether it was found
func (a FeedOptions) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for FeedOptions
func (a *FeedOptions) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for FeedOptions to handle AdditionalProperties
func (a *FeedOptions) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for FeedOptions to handle AdditionalProperties
func (a FeedOptions) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
