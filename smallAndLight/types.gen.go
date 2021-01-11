// Package smallAndLight provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package smallAndLight

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional information that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList struct {
	Errors *[]Error `json:"errors,omitempty"`
}

// FeeLineItem defines model for FeeLineItem.
type FeeLineItem struct {
	FeeCharge MoneyType `json:"feeCharge"`

	// The type of fee charged to the seller.
	FeeType string `json:"feeType"`
}

// FeePreview defines model for FeePreview.
type FeePreview struct {

	// The Amazon Standard Identification Number (ASIN) value used to identify the item.
	Asin *string `json:"asin,omitempty"`

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// A list of the Small and Light fees for the item.
	FeeBreakdown *[]FeeLineItem `json:"feeBreakdown,omitempty"`
	Price        *MoneyType     `json:"price,omitempty"`
	TotalFees    *MoneyType     `json:"totalFees,omitempty"`
}

// Item defines model for Item.
type Item struct {

	// The Amazon Standard Identification Number (ASIN) value used to identify the item.
	Asin  string    `json:"asin"`
	Price MoneyType `json:"price"`
}

// MarketplaceId defines model for MarketplaceId.
type MarketplaceId string

// MoneyType defines model for MoneyType.
type MoneyType struct {

	// The monetary value.
	Amount *float32 `json:"amount,omitempty"`

	// The currency code in ISO 4217 format.
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// SellerSKU defines model for SellerSKU.
type SellerSKU string

// SmallAndLightEligibility defines model for SmallAndLightEligibility.
type SmallAndLightEligibility struct {

	// A marketplace identifier.
	MarketplaceId MarketplaceId `json:"marketplaceId"`

	// Identifies an item in the given marketplace. SellerSKU is qualified by the seller's SellerId, which is included with every operation that you submit.
	SellerSKU SellerSKU `json:"sellerSKU"`

	// The Small and Light eligibility status of the item.
	Status SmallAndLightEligibilityStatus `json:"status"`
}

// SmallAndLightEligibilityStatus defines model for SmallAndLightEligibilityStatus.
type SmallAndLightEligibilityStatus string

// List of SmallAndLightEligibilityStatus
const (
	SmallAndLightEligibilityStatus_ELIGIBLE     SmallAndLightEligibilityStatus = "ELIGIBLE"
	SmallAndLightEligibilityStatus_NOT_ELIGIBLE SmallAndLightEligibilityStatus = "NOT_ELIGIBLE"
)

// SmallAndLightEnrollment defines model for SmallAndLightEnrollment.
type SmallAndLightEnrollment struct {

	// A marketplace identifier.
	MarketplaceId MarketplaceId `json:"marketplaceId"`

	// Identifies an item in the given marketplace. SellerSKU is qualified by the seller's SellerId, which is included with every operation that you submit.
	SellerSKU SellerSKU `json:"sellerSKU"`

	// The Small and Light enrollment status of the item.
	Status SmallAndLightEnrollmentStatus `json:"status"`
}

// SmallAndLightEnrollmentStatus defines model for SmallAndLightEnrollmentStatus.
type SmallAndLightEnrollmentStatus string

// List of SmallAndLightEnrollmentStatus
const (
	SmallAndLightEnrollmentStatus_ENROLLED     SmallAndLightEnrollmentStatus = "ENROLLED"
	SmallAndLightEnrollmentStatus_NOT_ENROLLED SmallAndLightEnrollmentStatus = "NOT_ENROLLED"
)

// SmallAndLightFeePreviewRequest defines model for SmallAndLightFeePreviewRequest.
type SmallAndLightFeePreviewRequest struct {

	// A list of items for which to retrieve fee estimates (limit: 25).
	Items []Item `json:"items"`

	// A marketplace identifier.
	MarketplaceId MarketplaceId `json:"marketplaceId"`
}

// SmallAndLightFeePreviews defines model for SmallAndLightFeePreviews.
type SmallAndLightFeePreviews struct {

	// A list of fee estimates for the requested items. The order of the fee estimates will follow the same order as the items in the request, with duplicates removed.
	Data *[]FeePreview `json:"data,omitempty"`
}

// GetSmallAndLightEligibilityBySellerSKUParams defines parameters for GetSmallAndLightEligibilityBySellerSKU.
type GetSmallAndLightEligibilityBySellerSKUParams struct {

	// The marketplace for which the eligibility status is retrieved. NOTE: Accepts a single marketplace only.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// DeleteSmallAndLightEnrollmentBySellerSKUParams defines parameters for DeleteSmallAndLightEnrollmentBySellerSKU.
type DeleteSmallAndLightEnrollmentBySellerSKUParams struct {

	// The marketplace in which to remove the item from the Small and Light program. Note: Accepts a single marketplace only.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// GetSmallAndLightEnrollmentBySellerSKUParams defines parameters for GetSmallAndLightEnrollmentBySellerSKU.
type GetSmallAndLightEnrollmentBySellerSKUParams struct {

	// The marketplace for which the enrollment status is retrieved. Note: Accepts a single marketplace only.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// PutSmallAndLightEnrollmentBySellerSKUParams defines parameters for PutSmallAndLightEnrollmentBySellerSKU.
type PutSmallAndLightEnrollmentBySellerSKUParams struct {

	// The marketplace in which to enroll the item. Note: Accepts a single marketplace only.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// GetSmallAndLightFeePreviewJSONBody defines parameters for GetSmallAndLightFeePreview.
type GetSmallAndLightFeePreviewJSONBody SmallAndLightFeePreviewRequest

// GetSmallAndLightFeePreviewRequestBody defines body for GetSmallAndLightFeePreview for application/json ContentType.
type GetSmallAndLightFeePreviewJSONRequestBody GetSmallAndLightFeePreviewJSONBody