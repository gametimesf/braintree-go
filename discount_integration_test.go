//go:build integration
// +build integration

package braintree

import (
	"context"
	"testing"
)

func TestDiscounts(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	discounts, err := testGateway.Discount().All(ctx)

	if err != nil {
		t.Error(err)
	} else if len(discounts) != 1 {
		t.Fatalf("expected to retrieve 1 discount, retrieved %d", len(discounts))
	}

	discount := discounts[0]

	t.Log(discount)

	if discount.Id != "test_discount" {
		t.Fatalf("expected Id to be %s, was %s", "test_discount", discount.Id)
	} else if discount.Amount.Cmp(NewDecimal(1000, 2)) != 0 {
		t.Fatalf("expected Amount to be %s, was %s", NewDecimal(1000, 2), discount.Amount)
	} else if discount.Kind != ModificationKindDiscount {
		t.Fatalf("expected Kind to be %s, was %s", ModificationKindDiscount, discount.Kind)
	} else if discount.Name != "test_discount_name" {
		t.Fatalf("expected Name to be %s, was %s", "test_discount_name", discount.Name)
	} else if discount.NeverExpires != true {
		t.Fatalf("expected NeverExpires to be %t, was %t", true, discount.NeverExpires)
	} else if discount.CurrentBillingCycle != 0 {
		t.Fatalf("expected current billing cycle to be %d, was %d", 0, discount.CurrentBillingCycle)
	} else if discount.Description != "A test discount" {
		t.Fatalf("expected Description to be %s, was %s", "A test discount", discount.Description)
	}
}
