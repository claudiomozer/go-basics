package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exists")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "payed using cash") {
		t.Error("The Cash payment method message wasn't correct")
	}
	t.Log("Log:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Fatal("A payment method of type 'Debit Card' must exists")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "payed using debit card") {
		t.Error("The Cash payment method message wasn't correct")
	}
	t.Log("Log:", msg)
}

func TestGetNonExistentPaymentMethod(t *testing.T) {
	_, err := GetPaymentMethod(99)

	if err == nil {
		t.Error("A payment method with id 99 must no exists")
	}
	t.Log("Log:", err)
}
