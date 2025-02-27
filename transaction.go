package braintree

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gametimesf/braintree-go/customfields"
)

type TransactionStatus string

const (
	TransactionStatusAuthorizationExpired   TransactionStatus = "authorization_expired"
	TransactionStatusAuthorizing            TransactionStatus = "authorizing"
	TransactionStatusAuthorized             TransactionStatus = "authorized"
	TransactionStatusGatewayRejected        TransactionStatus = "gateway_rejected"
	TransactionStatusFailed                 TransactionStatus = "failed"
	TransactionStatusProcessorDeclined      TransactionStatus = "processor_declined"
	TransactionStatusSettled                TransactionStatus = "settled"
	TransactionStatusSettlementConfirmed    TransactionStatus = "settlement_confirmed"
	TransactionStatusSettlementDeclined     TransactionStatus = "settlement_declined"
	TransactionStatusSettlementPending      TransactionStatus = "settlement_pending"
	TransactionStatusSettling               TransactionStatus = "settling"
	TransactionStatusSubmittedForSettlement TransactionStatus = "submitted_for_settlement"
	TransactionStatusVoided                 TransactionStatus = "voided"
	TransactionStatusUnrecognized           TransactionStatus = "unrecognized"
)

type TransactionSource string

const (
	TransactionSourceRecurringFirst TransactionSource = "recurring_first"
	TransactionSourceRecurring      TransactionSource = "recurring"
	TransactionSourceMOTO           TransactionSource = "moto"
	TransactionSourceMerchant       TransactionSource = "merchant"
	TransactionSourceUnscheduled    TransactionSource = "unscheduled"
)

type PaymentInstrumentType string

const (
	PaymentInstrumentTypeAndroidPayCard   PaymentInstrumentType = "android_pay_card"
	PaymentInstrumentTypeApplePayCard     PaymentInstrumentType = "apple_pay_card"
	PaymentInstrumentTypeCreditCard       PaymentInstrumentType = "credit_card"
	PaymentInstrumentTypeMasterpassCard   PaymentInstrumentType = "masterpass_card"
	PaymentInstrumentTypePaypalAccount    PaymentInstrumentType = "paypal_account"
	PaymentInstrumentTypeVenmoAccount     PaymentInstrumentType = "venmo_account"
	PaymentInstrumentTypeVisaCheckoutCard PaymentInstrumentType = "visa_checkout_card"
)

type Transaction struct {
	XMLName                      string                    `xml:"transaction"`
	Id                           string                    `xml:"id"`
	Status                       TransactionStatus         `xml:"status"`
	Type                         string                    `xml:"type"`
	CurrencyISOCode              string                    `xml:"currency-iso-code"`
	Amount                       *Decimal                  `xml:"amount"`
	OrderId                      string                    `xml:"order-id"`
	PaymentMethodToken           string                    `xml:"payment-method-token"`
	PaymentMethodNonce           string                    `xml:"payment-method-nonce"`
	MerchantAccountId            string                    `xml:"merchant-account-id"`
	PlanId                       string                    `xml:"plan-id"`
	SubscriptionId               string                    `xml:"subscription-id"`
	SubscriptionDetails          *SubscriptionDetails      `xml:"subscription"`
	CreditCard                   *CreditCard               `xml:"credit-card"`
	Customer                     *Customer                 `xml:"customer"`
	BillingAddress               *Address                  `xml:"billing"`
	ShippingAddress              *Address                  `xml:"shipping"`
	TaxAmount                    *Decimal                  `xml:"tax-amount"`
	TaxExempt                    bool                      `xml:"tax-exempt"`
	DeviceData                   string                    `xml:"device-data"`
	ServiceFeeAmount             *Decimal                  `xml:"service-fee-amount,attr"`
	CreatedAt                    *time.Time                `xml:"created-at"`
	UpdatedAt                    *time.Time                `xml:"updated-at"`
	DisbursementDetails          *DisbursementDetails      `xml:"disbursement-details"`
	RefundId                     string                    `xml:"refund-id"`
	RefundIds                    *[]string                 `xml:"refund-ids>item"`
	RefundedTransactionId        *string                   `xml:"refunded-transaction-id"`
	ProcessorResponseCode        ProcessorResponseCode     `xml:"processor-response-code"`
	ProcessorResponseText        string                    `xml:"processor-response-text"`
	ProcessorResponseType        ProcessorResponseType     `xml:"processor-response-type"`
	ProcessorAuthorizationCode   string                    `xml:"processor-authorization-code"`
	SettlementBatchId            string                    `xml:"settlement-batch-id"`
	EscrowStatus                 EscrowStatus              `xml:"escrow-status"`
	PaymentInstrumentType        PaymentInstrumentType     `xml:"payment-instrument-type"`
	ThreeDSecureInfo             *ThreeDSecureInfo         `xml:"three-d-secure-info,omitempty"`
	PayPalDetails                *PayPalDetails            `xml:"paypal"`
	VenmoAccountDetails          *VenmoAccountDetails      `xml:"venmo-account"`
	AndroidPayDetails            *AndroidPayDetails        `xml:"android-pay-card"`
	ApplePayDetails              *ApplePayDetails          `xml:"apple-pay"`
	AdditionalProcessorResponse  string                    `xml:"additional-processor-response"`
	RiskData                     *RiskData                 `xml:"risk-data"`
	Descriptor                   *Descriptor               `xml:"descriptor"`
	Channel                      string                    `xml:"channel"`
	CustomFields                 customfields.CustomFields `xml:"custom-fields"`
	AVSErrorResponseCode         AVSResponseCode           `xml:"avs-error-response-code"`
	AVSPostalCodeResponseCode    AVSResponseCode           `xml:"avs-postal-code-response-code"`
	AVSStreetAddressResponseCode AVSResponseCode           `xml:"avs-street-address-response-code"`
	CVVResponseCode              CVVResponseCode           `xml:"cvv-response-code"`
	GatewayRejectionReason       GatewayRejectionReason    `xml:"gateway-rejection-reason"`
	PurchaseOrderNumber          string                    `xml:"purchase-order-number"`
	Disputes                     []*Dispute                `xml:"disputes>dispute"`
	AuthorizationExpiresAt       *time.Time                `xml:"authorization-expires-at"`
}

type TransactionRequest struct {
	XMLName                  string                      `xml:"transaction"`
	CustomerID               string                      `xml:"customer-id,omitempty"`
	Type                     string                      `xml:"type,omitempty"`
	Amount                   *Decimal                    `xml:"amount"`
	OrderId                  string                      `xml:"order-id,omitempty"`
	SharedPaymentMethodToken string                      `xml:"shared-payment-method-token,omitempty"`
	PaymentMethodToken       string                      `xml:"payment-method-token,omitempty"`
	PaymentMethodNonce       string                      `xml:"payment-method-nonce,omitempty"`
	MerchantAccountId        string                      `xml:"merchant-account-id,omitempty"`
	PlanId                   string                      `xml:"plan-id,omitempty"`
	CreditCard               *CreditCard                 `xml:"credit-card,omitempty"`
	Customer                 *CustomerRequest            `xml:"customer,omitempty"`
	BillingAddress           *Address                    `xml:"billing,omitempty"`
	ShippingAddress          *Address                    `xml:"shipping,omitempty"`
	TaxAmount                *Decimal                    `xml:"tax-amount,omitempty"`
	TaxExempt                bool                        `xml:"tax-exempt,omitempty"`
	DeviceData               string                      `xml:"device-data,omitempty"`
	Options                  *TransactionOptions         `xml:"options,omitempty"`
	ServiceFeeAmount         *Decimal                    `xml:"service-fee-amount,attr,omitempty"`
	RiskData                 *RiskDataRequest            `xml:"risk-data,omitempty"`
	Descriptor               *Descriptor                 `xml:"descriptor,omitempty"`
	Channel                  string                      `xml:"channel,omitempty"`
	CustomFields             customfields.CustomFields   `xml:"custom-fields,omitempty"`
	PurchaseOrderNumber      string                      `xml:"purchase-order-number,omitempty"`
	TransactionSource        TransactionSource           `xml:"transaction-source,omitempty"`
	LineItems                TransactionLineItemRequests `xml:"line-items,omitempty"`
}

type CreateTransactionRiskContextRequest struct {
	SenderAccountId   string `json:"sender_account_id"`   //value example:A12345N343
	SenderFirstName   string `json:"sender_first_name"`   //value example:John
	SenderLastName    string `json:"sender_last_name"`    //value example:Smith
	SenderEmail       string `json:"sender_email"`        //value example:john@demo.com
	SenderPhone       string `json:"sender_phone"`        //value example: (402) 555 5555
	SenderCountryCode string `json:"sender_country_code"` //value example:US
	SenderCreatedDate string `json:"sender_create_date"`  //value example:2023-09-06T14:38:41Z

	// Delivery Information – Required for intangible transactions only; otherwise, optional.
	DeliveryMethod string `json:"dg_delivery_method"` //value example: {email, phone, venue_pickup, kiosk_pickup} // TODO modify type to enum

	// Merchant Custom Data - These fields are optional. Include them only if applicable.
	// CustomStrDataOne Free text field (suggested data: Date of the event, in ISO 8601)
	CustomStrDataOne string `json:"cd_string_one"`
	// CustomStrDataTwo
	// Free text field (suggested data: Type of event; for example, music, arts_and_theater, family,
	// sports, miscellaneous, clubs, special_events, fairs_and_exhibitions, festivals, comedy.)
	CustomStrDataTwo string `json:"cd_string_two"`
	// CustomIntDataOne
	// Free number field (suggested data: The total number of tickets that the buyer purchased within
	// a single transaction. )
	CustomIntDataOne int `json:"cd_int_one"`
}

func (c *CreateTransactionRiskContextRequest) Request() map[string]interface{} {
	return map[string]interface{}{
		"riskContext": map[string]interface{}{
			"fields": c.getFields(true),
		},
	}
}

func (c *CreateTransactionRiskContextRequest) getFields(skipEmpty bool) []Field {
	if c == nil {
		return []Field{}
	}

	fields, _ := getFiles(c, "json", skipEmpty)
	return fields
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

func getFiles(s any, tag string, skipEmpty bool) ([]Field, error) {
	rt := reflect.TypeOf(s)
	rv := reflect.ValueOf(s)

	switch rt.Kind() {
	case reflect.Pointer:
		rt = reflect.ValueOf(s).Elem().Type()
		if rt.Kind() != reflect.Struct {
			return nil, fmt.Errorf("expected struct type, got %s", rt.Kind().String())
		}
		rv = reflect.ValueOf(s).Elem()
	case reflect.Struct:
	default:
		return nil, fmt.Errorf("expected struct type, got %s", rt.Kind().String())
	}

	var fields = make([]Field, 0, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)

		value, exist := f.Tag.Lookup(tag)
		if exist {
			if skipEmpty && isZeroOfUnderlyingType(rv.Field(i).Interface()) {
				continue
			}
			value = strings.Split(value, ",")[0] // use split to ignore tag "options" like omitempty, etc.
			fields = append(fields, Field{
				Name:  value,
				Value: rv.Field(i).Interface(),
			})
		}
	}

	return fields, nil
}

type Field struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type GraphQLRawResponse[T any] struct {
	Data   map[string]T `json:"data"`
	Errors interface{}  `json:"errors"`
}

type CreateTransactionRiskContextResult struct {
	ClientMetadataId string `json:"clientMetadataId"`
}

type TransactionRefundRequest struct {
	XMLName string   `xml:"transaction"`
	Amount  *Decimal `xml:"amount"`
	OrderID string   `xml:"order-id,omitempty"`
}

func (t *Transaction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type typeWithNoFunctions Transaction
	if err := d.DecodeElement((*typeWithNoFunctions)(t), &start); err != nil {
		return err
	}
	if t.SubscriptionDetails != nil &&
		t.SubscriptionDetails.BillingPeriodStartDate == "" &&
		t.SubscriptionDetails.BillingPeriodEndDate == "" {
		t.SubscriptionDetails = nil
	}
	return nil
}

// TODO: not all transaction fields are implemented yet, here are the missing fields (add on demand)
//
// <transaction>
//   <voice-referral-number nil="true"></voice-referral-number>
//   <status-history type="array">
//     <status-event>
//       <timestamp type="datetime">2013-10-07T17:26:14Z</timestamp>
//       <status>authorized</status>
//       <amount>7.00</amount>
//       <user>eaigner</user>
//       <transaction-source>Recurring</transaction-source>
//     </status-event>
//     <status-event>
//       <timestamp type="datetime">2013-10-07T17:26:14Z</timestamp>
//       <status>submitted_for_settlement</status>
//       <amount>7.00</amount>
//       <user>eaigner</user>
//       <transaction-source>Recurring</transaction-source>
//     </status-event>
//     <status-event>
//       <timestamp type="datetime">2013-10-08T07:06:38Z</timestamp>
//       <status>settled</status>
//       <amount>7.00</amount>
//       <user nil="true"></user>
//       <transaction-source></transaction-source>
//     </status-event>
//   </status-history>
//   <plan-id>bronze</plan-id>
//   <subscription-id>jqsydb</subscription-id>
//   <subscription>
//     <billing-period-end-date type="date">2013-11-06</billing-period-end-date>
//     <billing-period-start-date type="date">2013-10-07</billing-period-start-date>
//   </subscription>
//   <add-ons type="array"/>
//   <discounts type="array"/>
//   <descriptor>
//     <name nil="true"></name>
//     <phone nil="true"></phone>
//   </descriptor>
//   <recurring type="boolean">true</recurring>
// </transaction>

type Transactions struct {
	Transaction []*Transaction `xml:"transaction"`
}

type TransactionOptions struct {
	SubmitForSettlement              bool                                   `xml:"submit-for-settlement,omitempty"`
	StoreInVault                     bool                                   `xml:"store-in-vault,omitempty"`
	StoreInVaultOnSuccess            bool                                   `xml:"store-in-vault-on-success,omitempty"`
	AddBillingAddressToPaymentMethod bool                                   `xml:"add-billing-address-to-payment-method,omitempty"`
	StoreShippingAddressInVault      bool                                   `xml:"store-shipping-address-in-vault,omitempty"`
	HoldInEscrow                     bool                                   `xml:"hold-in-escrow,omitempty"`
	TransactionOptionsPaypalRequest  *TransactionOptionsPaypalRequest       `xml:"paypal,omitempty"`
	SkipAdvancedFraudChecking        bool                                   `xml:"skip_advanced_fraud_checking,omitempty"`
	ThreeDSecure                     *TransactionOptionsThreeDSecureRequest `xml:"three-d-secure,omitempty"`
}

type TransactionOptionsPaypalRequest struct {
	CustomField       string
	PayeeEmail        string
	Description       string
	SupplementaryData map[string]string
}

func (r TransactionOptionsPaypalRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if r.CustomField != "" {
		if err := e.EncodeElement(r.CustomField, xml.StartElement{Name: xml.Name{Local: "custom-field"}}); err != nil {
			return err
		}
	}
	if r.PayeeEmail != "" {
		if err := e.EncodeElement(r.PayeeEmail, xml.StartElement{Name: xml.Name{Local: "payee-email"}}); err != nil {
			return err
		}
	}
	if r.Description != "" {
		if err := e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}}); err != nil {
			return err
		}
	}
	if len(r.SupplementaryData) > 0 {
		start := xml.StartElement{Name: xml.Name{Local: "supplementary-data"}}
		if err := e.EncodeToken(start); err != nil {
			return err
		}
		for k, v := range r.SupplementaryData {
			if err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: k}}); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(start.End()); err != nil {
			return err
		}
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	if err := e.Flush(); err != nil {
		return err
	}

	return nil
}

type TransactionOptionsThreeDSecureRequest struct {
	Required bool `xml:"required"`
}

type TransactionSearchResult struct {
	TotalItems int
	TotalIDs   []string

	CurrentPageNumber int
	PageSize          int
	Transactions      []*Transaction
}

type RiskData struct {
	ID       string `xml:"id"`
	Decision string `xml:"decision"`
}

type RiskDataRequest struct {
	CustomerBrowser string `xml:"customer-browser"`
	CustomerIP      string `xml:"customer-ip"`
}

type SubscriptionDetails struct {
	BillingPeriodStartDate string `xml:"billing-period-start-date"`
	BillingPeriodEndDate   string `xml:"billing-period-end-date"`
}
