// Package listingsItems provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package listingsItems

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Decimal defines model for Decimal.
type Decimal string

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList struct {
	Errors []Error `json:"errors"`
}

// FulfillmentAvailability defines model for FulfillmentAvailability.
type FulfillmentAvailability struct {

	// The code of the fulfillment network that will be used.
	FulfillmentChannelCode string `json:"fulfillmentChannelCode"`

	// The quantity of the item you are making available for sale.
	Quantity *int `json:"quantity,omitempty"`
}

// Issue defines model for Issue.
type Issue struct {

	// The names of the attributes associated with the issue, if applicable.
	AttributeNames *[]string `json:"attributeNames,omitempty"`

	// List of issue categories.
	//
	// Possible vales:
	//
	// * `INVALID_ATTRIBUTE` - Indicating an invalid attribute in the listing.
	//
	// * `MISSING_ATTRIBUTE` - Highlighting a missing attribute in the listing.
	//
	// * `INVALID_IMAGE` - Signifying an invalid image in the listing.
	//
	// * `MISSING_IMAGE` - Noting the absence of an image in the listing.
	//
	// * `INVALID_PRICE` - Pertaining to issues with the listing's price-related attributes.
	//
	// * `MISSING_PRICE` - Pointing out the absence of a price attribute in the listing.
	//
	// * `DUPLICATE` - Identifying listings with potential duplicate problems, such as this ASIN potentially being a duplicate of another ASIN.
	//
	// * `QUALIFICATION_REQUIRED` - Indicating that the listing requires qualification-related approval.
	Categories []string `json:"categories"`

	// An issue code that identifies the type of issue.
	Code string `json:"code"`

	// This field provides information about the enforcement actions taken by Amazon that affect the publishing or status of a listing. It also includes details about any associated exemptions.
	Enforcements *IssueEnforcements `json:"enforcements,omitempty"`

	// A message that describes the issue.
	Message string `json:"message"`

	// The severity of the issue.
	Severity string `json:"severity"`
}

// IssueEnforcementAction defines model for IssueEnforcementAction.
type IssueEnforcementAction struct {

	// The enforcement action name.
	//
	// Possible values:
	//
	// * `LISTING_SUPPRESSED` - This enforcement takes down the current listing item's buyability.
	//
	// * `ATTRIBUTE_SUPPRESSED` - An attribute's value on the listing item is invalid, which causes it to be rejected by Amazon.
	//
	// * `CATALOG_ITEM_REMOVED` - This catalog item is inactive on Amazon, and all offers against it in the applicable marketplace are non-buyable.
	//
	// * `SEARCH_SUPPRESSED` - This value indicates that the catalog item is hidden from search results.
	Action string `json:"action"`
}

// IssueEnforcements defines model for IssueEnforcements.
type IssueEnforcements struct {

	// List of enforcement actions taken by Amazon that affect the publishing or status of a listing.
	Actions []IssueEnforcementAction `json:"actions"`

	// Conveying the status of the listed enforcement actions and, if applicable, provides information about the exemption's expiry date.
	Exemption IssueExemption `json:"exemption"`
}

// IssueExemption defines model for IssueExemption.
type IssueExemption struct {

	// This field represents the timestamp, following the ISO 8601 format, which specifies the date when temporary exemptions, if applicable, will expire, and Amazon will begin enforcing the listed actions.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`

	// This field indicates the current exemption status for the listed enforcement actions. It can take values such as `EXEMPT`, signifying permanent exemption, `EXEMPT_UNTIL_EXPIRY_DATE` indicating temporary exemption until a specified date, or `NOT_EXEMPT` signifying no exemptions, and enforcement actions were already applied.
	Status string `json:"status"`
}

// Item defines model for Item.
type Item struct {

	// A JSON object containing structured listings item attribute data keyed by attribute name.
	Attributes *ItemAttributes `json:"attributes,omitempty"`

	// The fulfillment availability for the listings item.
	FulfillmentAvailability *[]FulfillmentAvailability `json:"fulfillmentAvailability,omitempty"`

	// The issues associated with the listings item.
	Issues *ItemIssues `json:"issues,omitempty"`

	// Offer details for the listings item.
	Offers *ItemOffers `json:"offers,omitempty"`

	// The vendor procurement information for the listings item.
	Procurement *[]ItemProcurement `json:"procurement,omitempty"`

	// A selling partner provided identifier for an Amazon listing.
	Sku string `json:"sku"`

	// Summary details of a listings item.
	Summaries *ItemSummaries `json:"summaries,omitempty"`
}

// ItemAttributes defines model for ItemAttributes.
type ItemAttributes struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ItemIdentifiers defines model for ItemIdentifiers.
type ItemIdentifiers []ItemIdentifiersByMarketplace

// ItemIdentifiersByMarketplace defines model for ItemIdentifiersByMarketplace.
type ItemIdentifiersByMarketplace struct {

	// Amazon Standard Identification Number (ASIN) of the listings item.
	Asin *string `json:"asin,omitempty"`

	// A marketplace identifier. Identifies the Amazon marketplace for the listings item.
	MarketplaceId *string `json:"marketplaceId,omitempty"`
}

// ItemImage defines model for ItemImage.
type ItemImage struct {

	// The height of the image in pixels.
	Height int `json:"height"`

	// The link, or URL, to the image.
	Link string `json:"link"`

	// The width of the image in pixels.
	Width int `json:"width"`
}

// ItemIssues defines model for ItemIssues.
type ItemIssues []Issue

// ItemOfferByMarketplace defines model for ItemOfferByMarketplace.
type ItemOfferByMarketplace struct {

	// The Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`

	// Type of offer for the listings item.
	OfferType string `json:"offerType"`

	// The number of Amazon Points offered with the purchase of an item, and their monetary value. Note that the `Points` element is only returned in Japan (JP).
	Points *Points `json:"points,omitempty"`

	// The currency type and amount.
	Price Money `json:"price"`
}

// ItemOffers defines model for ItemOffers.
type ItemOffers []ItemOfferByMarketplace

// ItemProcurement defines model for ItemProcurement.
type ItemProcurement struct {

	// The currency type and amount.
	CostPrice Money `json:"costPrice"`
}

// ItemSummaries defines model for ItemSummaries.
type ItemSummaries []ItemSummaryByMarketplace

// ItemSummaryByMarketplace defines model for ItemSummaryByMarketplace.
type ItemSummaryByMarketplace struct {

	// Amazon Standard Identification Number (ASIN) of the listings item.
	Asin string `json:"asin"`

	// Identifies the condition of the listings item.
	ConditionType *string `json:"conditionType,omitempty"`

	// The date the listings item was created in ISO 8601 format.
	CreatedDate time.Time `json:"createdDate"`

	// The fulfillment network stock keeping unit is an identifier used by Amazon fulfillment centers to identify each unique item.
	FnSku *string `json:"fnSku,omitempty"`

	// The name or title associated with an Amazon catalog item.
	ItemName string `json:"itemName"`

	// The date the listings item was last updated in ISO 8601 format.
	LastUpdatedDate time.Time `json:"lastUpdatedDate"`

	// The image for the listings item.
	MainImage *ItemImage `json:"mainImage,omitempty"`

	// A marketplace identifier. Identifies the Amazon marketplace for the listings item.
	MarketplaceId string `json:"marketplaceId"`

	// The Amazon product type of the listings item.
	ProductType string `json:"productType"`

	// Statuses that apply to the listings item.
	Status []string `json:"status"`
}

// ListingsItemPatchRequest defines model for ListingsItemPatchRequest.
type ListingsItemPatchRequest struct {

	// One or more JSON Patch operations to perform on the listings item.
	Patches []PatchOperation `json:"patches"`

	// The Amazon product type of the listings item.
	ProductType string `json:"productType"`
}

// ListingsItemPutRequest defines model for ListingsItemPutRequest.
type ListingsItemPutRequest struct {

	// A JSON object containing structured listings item attribute data keyed by attribute name.
	Attributes ListingsItemPutRequest_Attributes `json:"attributes"`

	// The Amazon product type of the listings item.
	ProductType string `json:"productType"`

	// The name of the requirements set for the provided data.
	Requirements *string `json:"requirements,omitempty"`
}

// ListingsItemPutRequest_Attributes defines model for ListingsItemPutRequest.Attributes.
type ListingsItemPutRequest_Attributes struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ListingsItemSubmissionResponse defines model for ListingsItemSubmissionResponse.
type ListingsItemSubmissionResponse struct {

	// Identity attributes associated with the item in the Amazon catalog, such as the ASIN.
	Identifiers *ItemIdentifiers `json:"identifiers,omitempty"`

	// Listings item issues related to the listings item submission.
	Issues *[]Issue `json:"issues,omitempty"`

	// A selling partner provided identifier for an Amazon listing.
	Sku string `json:"sku"`

	// The status of the listings item submission.
	Status string `json:"status"`

	// The unique identifier of the listings item submission.
	SubmissionId string `json:"submissionId"`
}

// Money defines model for Money.
type Money struct {

	// A decimal number with no loss of precision. Useful when precision loss is unnaceptable, as with currencies. Follows RFC7159 for number representation.
	Amount Decimal `json:"amount"`

	// Three-digit currency code in ISO 4217 format.
	CurrencyCode string `json:"currencyCode"`
}

// PatchOperation defines model for PatchOperation.
type PatchOperation struct {

	// Type of JSON Patch operation. Supported JSON Patch operations include add, replace, and delete. Refer to [JavaScript Object Notation (JSON) Patch](https://tools.ietf.org/html/rfc6902) for more information.
	Op string `json:"op"`

	// JSON Pointer path of the element to patch. Refer to [JavaScript Object Notation (JSON) Patch](https://tools.ietf.org/html/rfc6902) for more information.
	Path string `json:"path"`

	// JSON value to add, replace, or delete.
	Value []map[string]interface{} `json:"value,omitempty"`
}

// Points defines model for Points.
type Points struct {
	PointsNumber int `json:"pointsNumber"`
}

// DeleteListingsItemParams defines parameters for DeleteListingsItem.
type DeleteListingsItemParams struct {

	// A comma-delimited list of Amazon marketplace identifiers for the request.
	MarketplaceIds []string `json:"marketplaceIds"`

	// A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: `en_US`, `fr_CA`, `fr_FR`. Localized messages default to `en_US` when a localization is not available in the specified locale.
	IssueLocale *string `json:"issueLocale,omitempty"`
}

// GetListingsItemParams defines parameters for GetListingsItem.
type GetListingsItemParams struct {

	// A comma-delimited list of Amazon marketplace identifiers for the request.
	MarketplaceIds []string `json:"marketplaceIds"`

	// A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: `en_US`, `fr_CA`, `fr_FR`. Localized messages default to `en_US` when a localization is not available in the specified locale.
	IssueLocale *string `json:"issueLocale,omitempty"`

	// A comma-delimited list of data sets to include in the response. Default: `summaries`.
	IncludedData *[]string `json:"includedData,omitempty"`
}

// PatchListingsItemJSONBody defines parameters for PatchListingsItem.
type PatchListingsItemJSONBody ListingsItemPatchRequest

// PatchListingsItemParams defines parameters for PatchListingsItem.
type PatchListingsItemParams struct {

	// A comma-delimited list of Amazon marketplace identifiers for the request.
	MarketplaceIds []string `json:"marketplaceIds"`

	// A comma-delimited list of data sets to include in the response. Default: `issues`.
	IncludedData *[]string `json:"includedData,omitempty"`

	// The mode of operation for the request.
	Mode *string `json:"mode,omitempty"`

	// A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: `en_US`, `fr_CA`, `fr_FR`. Localized messages default to `en_US` when a localization is not available in the specified locale.
	IssueLocale *string `json:"issueLocale,omitempty"`
}

// PutListingsItemJSONBody defines parameters for PutListingsItem.
type PutListingsItemJSONBody ListingsItemPutRequest

// PutListingsItemParams defines parameters for PutListingsItem.
type PutListingsItemParams struct {

	// A comma-delimited list of Amazon marketplace identifiers for the request.
	MarketplaceIds []string `json:"marketplaceIds"`

	// A comma-delimited list of data sets to include in the response. Default: `issues`.
	IncludedData *[]string `json:"includedData,omitempty"`

	// The mode of operation for the request.
	Mode *string `json:"mode,omitempty"`

	// A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: `en_US`, `fr_CA`, `fr_FR`. Localized messages default to `en_US` when a localization is not available in the specified locale.
	IssueLocale *string `json:"issueLocale,omitempty"`
}

// PatchListingsItemRequestBody defines body for PatchListingsItem for application/json ContentType.
type PatchListingsItemJSONRequestBody PatchListingsItemJSONBody

// PutListingsItemRequestBody defines body for PutListingsItem for application/json ContentType.
type PutListingsItemJSONRequestBody PutListingsItemJSONBody

// Getter for additional properties for ItemAttributes. Returns the specified
// element and whether it was found
func (a ItemAttributes) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ItemAttributes
func (a *ItemAttributes) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ItemAttributes to handle AdditionalProperties
func (a *ItemAttributes) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ItemAttributes to handle AdditionalProperties
func (a ItemAttributes) MarshalJSON() ([]byte, error) {
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

// Getter for additional properties for ListingsItemPutRequest_Attributes. Returns the specified
// element and whether it was found
func (a ListingsItemPutRequest_Attributes) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ListingsItemPutRequest_Attributes
func (a *ListingsItemPutRequest_Attributes) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ListingsItemPutRequest_Attributes to handle AdditionalProperties
func (a *ListingsItemPutRequest_Attributes) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ListingsItemPutRequest_Attributes to handle AdditionalProperties
func (a ListingsItemPutRequest_Attributes) MarshalJSON() ([]byte, error) {
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
