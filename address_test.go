package main

import (
	"testing"
)

type AddressTestCase struct {
	addressString []string
	expected      []Address
}

var testCases AddressTestCase = AddressTestCase{
	[]string{
		"Vasudev Sathish Kamath <kamathvasudev@gmail.com>",
		"Vasudev Kamath <vasudev@copyninja.info>",
	},
	[]Address{
		Address{
			"Vasudev",
			"Sathish Kamath",
			"kamathvasudev@gmail.com",
		},
		Address{
			"Vasudev",
			"Kamath",
			"vasudev@copyninja.info",
		},
	},
}

func TestNewAddress(t *testing.T) {
	for i, test := range testCases.addressString {
		a := NewAddress(test)
		e := testCases.expected[i]

		if a.FirstName != e.FirstName {
			t.Fatalf("First Name not same (%s!=%s)",
				a.FirstName, e.FirstName)
		}

		if a.LastName != e.LastName {
			t.Fatalf("Last Name not same (%s != %s)",
				a.LastName, e.LastName)
		}

		if a.Email != e.Email {
			t.Fatalf("Email not same (%s != %s)",
				a.Email, e.Email)
		}
	}
}
