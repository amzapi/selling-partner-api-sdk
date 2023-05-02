# Amazon's Selling Partner API (SP-API) Golang SDK

[![Go Reference](https://pkg.go.dev/badge/gopkg.me/selling-partner-api-sdk.svg)](https://pkg.go.dev/gopkg.me/selling-partner-api-sdk)

## Installation

~~~~
go get -u github.com/amzapi/selling-partner-api-sdk
~~~~

## Progress
                    
* [X] authorization ([authorization-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/authorization-api-model/authorization.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/authorization-api/authorization.md))
* [X] catalog ([catalog-items-api-model](https://github.com/amzn/selling-partner-api-docs/blob/main/references/catalog-items-api/catalogItemsV0.md) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/catalog-items-api/catalogItemsV0.md))
* [X] fbaInbound ([fulfillment-inbound-api-model](https://github.com/amzn/selling-partner-api-docs/blob/main/references/fulfillment-inbound-api/fulfillmentInboundV0.md) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/fulfillment-inbound-api/fulfillmentInboundV0.md))
* [X] fbaInventory ([fba-inventory-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/fba-inventory-api-model/fbaInventory.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/fba-inventory-api/fbaInventory.md))
* [X] fbaOutbound ([fulfillment-outbound-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/fulfillment-outbound-api-model/fulfillmentOutbound_2020-07-01.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/fulfillment-outbound-api/fulfillmentOutbound_2020-07-01.md))
* [X] feeds ([feeds-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/feeds-api-model/feeds_2020-09-04.json) [DOC](https://github.com/amzn/selling-partner-api-docs/tree/main/references/feeds-api))
* [X] fees ([product-fees-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/product-fees-api-model/productFeesV0.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/product-fees-api/productFeesV0.md))
* [X] finances ([finances-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/finances-api-model/financesV0.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/finances-api/financesV0.md))
* [X] merchantFulfillment ([merchant-fulfillment-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/merchant-fulfillment-api-model/merchantFulfillmentV0.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/merchant-fulfillment-api/merchantFulfillmentV0.md))
* [X] messaging ([messaging-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/messaging-api-model/messaging.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/messaging-api/messaging.md))
* [X] notifications ([notifications-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/notifications-api-model/notifications.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/notifications-api/notifications.md))
* [X] ordersV0 ([orders-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/orders-api-model/ordersV0.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/orders-api/ordersV0.md))
* [X] productPricing ([product-pricing-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/product-pricing-api-model/productPricingV0.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/product-pricing-api/productPricingV0.md))
* [X] reports ([reports-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/reports-api-model/reports_2020-09-04.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/reports-api/reports_2020-09-04.md))
* [X] sales ([sales-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/sales-api-model/sales.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/sales-api/sales.md))
* [X] sellers ([sellers-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/sellers-api-model/sellers.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/sellers-api/sellers.md))
* [X] service ([services-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/services-api-model/services.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/services-api/services.md))
* [X] shipping ([shipping-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/shipping-api-model/shipping.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/shipping-api/shipping.md))
* [ ] smallAndLight ([fba-small-and-light-api](https://github.com/amzn/selling-partner-api-models/blob/main/models/fba-small-and-light-api-model/fbaSmallandLight.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/fba-small-and-light-api/fbaSmallandLight.md))
* [X] solicitations ([solicitations-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/solicitations-api-model/solicitations.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/solicitations-api/solicitations.md))
* [X] uploads ([uploads-api-model](https://github.com/amzn/selling-partner-api-models/blob/main/models/uploads-api-model/uploads_2020-11-01.json) [DOC](https://github.com/amzn/selling-partner-api-docs/blob/main/references/uploads-api/uploads_2020-11-01.md))

## Example

```go
package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	
	sp "github.com/amzapi/selling-partner-api-sdk/pkg/selling-partner"
	"github.com/amzapi/selling-partner-api-sdk/sellers"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func main() {
	
	sellingPartner, err := sp.NewSellingPartner(&sp.Config{
		ClientID:     "<ClientID>",
		ClientSecret: "<ClientSecret>",
		RefreshToken: "<RefreshToken>",
		AccessKeyID: "<AWS IAM User Access Key Id>",
		SecretKey:   "<AWS IAM User Secret Key>",
		Region:      "<AWS Region>",
		RoleArn:     "<AWS IAM Role ARN>",
	})

	if err != nil {
		panic(err)
	}

	endpoint := "https://sellingpartnerapi-fe.amazon.com"

	seller, err := sellers.NewClientWithResponses(endpoint,
		sellers.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
			req.Header.Add("X-Amzn-Requestid", uuid.New().String()) //tracking requests
			err = sellingPartner.SignRequest(req)
			if err != nil {
				return errors.Wrap(err, "sign error")
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

#Report Decryption
Amazon specification of version `2020-09-04` returns encrypted reports. To decrypt the reports you could use [the Decrypt function of the decryption package](./pkg/decryption/decryptor.go).
### Test example
The test example uses `amzn.GetReportDocumentResponse` from [Amazon models](https://github.com/amzn/selling-partner-api-models/blob/main/models/reports-api-model/reports_2020-09-04.json#L137) and the 
Decrypt function to download, decrypt and dump report. 
```go
func TestSellingPartnerGetReportDocumentThirdParty(t *testing.T) {
	sellingPartner, err := sp.NewSellingPartner(&sp.Config{
		ClientID:     "***",
		ClientSecret: "***",
		RefreshToken: "***",
		AccessKeyID:  "***",
		SecretKey:    "***",
		Region:       "***",
		RoleArn:      "***",
	})

	if err != nil {
		t.Fatal("Failed to create NewSellingPartner: ", err)
	}

	report, err := reports.NewClientWithResponses(spiHost,
		reports.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
			err = sellingPartner.SignRequest(req)
			if err != nil {
				return errors.Wrap(err, "sign error")
			}
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				return errors.Wrap(err, "DumpRequest Error")
			}
			ioutil.WriteFile("test-samples/3dumpedThirdPartyReqGetReports.txt", dump, 0777)
			return nil
		}),
		reports.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
			dump, err := httputil.DumpResponse(rsp, true)
			if err != nil {
				return errors.Wrap(err, "DumpResponse Error")
			}
			ioutil.WriteFile("test-samples/3dumpedThirdPartyRespGetReports.txt", dump, 0777)
			return nil
		}),
	)
	if err != nil {
		t.Fatal("Failed to create NewClientWithResponses: ", err)
	}
	ctx := context.Background()
	reportDocumentId := "***"
	resp, err := report.GetReportDocument(ctx, reportDocumentId)
	if err != nil {
		t.Fatal("Failed to make GetReportDocument request: ", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		t.Fatal("Service returned a status that isn't 2xx: ", resp.StatusCode)
	}
	defer resp.Body.Close()
	var bodyData []byte
	resp.Body.Read(bodyData)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Failed to read GetReportDocument response body: ", err)
	}
	var reportDocumentResp amzn.GetReportDocumentResponse
	err = json.Unmarshal(content, &reportDocumentResp)
	if err != nil {
		t.Fatal("Failed to unmarshall GetReportDocument response: ", err)
	}
	if reportDocumentResp.Errors != nil {
		t.Fatal("Got errors in the GetReportDocument response: ", reportDocumentResp.Errors)
	}
	respWithDocContent, err := http.Get(reportDocumentResp.Payload.Url)
	if err != nil {
		t.Fatal("failed to make request to "+reportDocumentResp.Payload.Url+" : ", err)
	}
	if respWithDocContent.StatusCode < 200 || respWithDocContent.StatusCode > 299 {
		t.Fatal("Service returned a status that isn't 2xx: ", respWithDocContent.StatusCode)
	}
	defer respWithDocContent.Body.Close()

	reportContentBytes, err := ioutil.ReadAll(respWithDocContent.Body)
	if err != nil {
		t.Fatal("Failed to read Report content body: ", err)
	}
	ioutil.WriteFile("test-samples/3_encrypted_"+reportDocumentId+"_"+reportType+"_report.txt", reportContentBytes, 0777)
	decryptedFilecontent, err := decryption.Decrypt(reportDocumentResp.Payload.EncryptionDetails.Key, reportDocumentResp.Payload.EncryptionDetails.InitializationVector, reportContentBytes)
	if err != nil {
		t.Fatal("Failed to decrypt file content: ", err)
	}
	ioutil.WriteFile("test-samples/3_"+reportDocumentId+"_"+reportType+"_report.txt", decryptedFilecontent, 0777)
}
```
