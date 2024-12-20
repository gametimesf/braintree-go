package braintree

import (
	"encoding/xml"
	"time"
)

type VenmoAccount struct {
	XMLName           xml.Name       `xml:"venmo-account"`
	CustomerId        string         `xml:"customer-id"`
	Token             string         `xml:"token"`
	Username          string         `xml:"username"`
	VenmoUserID       string         `xml:"venmo-user-id"`
	SourceDescription string         `xml:"source-description"`
	ImageURL          string         `xml:"image-url"`
	CreatedAt         *time.Time     `xml:"created-at"`
	UpdatedAt         *time.Time     `xml:"updated-at"`
	Subscriptions     *Subscriptions `xml:"subscriptions"`
	Default           bool           `xml:"default"`
}

type VenmoAccounts struct {
	VenmoAccount []*VenmoAccount `xml:"venmo-account"`
}

func (v *VenmoAccounts) PaymentMethods() []PaymentMethod {
	if v == nil {
		return nil
	}
	var paymentMethods []PaymentMethod
	for _, va := range v.VenmoAccount {
		paymentMethods = append(paymentMethods, va)
	}
	return paymentMethods
}

func (v *VenmoAccount) GetCustomerId() string {
	return v.CustomerId
}

func (v *VenmoAccount) GetToken() string {
	return v.Token
}

func (v *VenmoAccount) IsDefault() bool {
	return v.Default
}

func (v *VenmoAccount) GetImageURL() string {
	return v.ImageURL
}

// AllSubscriptions returns all subscriptions for this venmo account, or nil if none present.
func (v *VenmoAccount) AllSubscriptions() []*Subscription {
	if v.Subscriptions != nil {
		subs := v.Subscriptions.Subscription
		if len(subs) > 0 {
			a := make([]*Subscription, 0, len(subs))
			a = append(a, subs...)
			return a
		}
	}
	return nil
}
