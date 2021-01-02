package inst

import (
	"reflect"
	"testing"
)

func TestFormSupportedSuffixes(t *testing.T) {
	cases := []struct {
		Form   Form
		Expect []Suffixes
	}{
		{
			Form: Form{},
			Expect: []Suffixes{
				{},
			},
		},
		{
			Form: Form{
				Broadcast: true,
			},
			Expect: []Suffixes{
				{},
				{BCST},
			},
		},
		{
			Form: Form{
				EmbeddedRounding: true,
			},
			Expect: []Suffixes{
				{},
				{RN_SAE},
				{RZ_SAE},
				{RD_SAE},
				{RU_SAE},
			},
		},
		{
			Form: Form{
				SuppressAllExceptions: true,
			},
			Expect: []Suffixes{
				{},
				{SAE},
			},
		},
		{
			Form: Form{
				Zeroing: true,
			},
			Expect: []Suffixes{
				{},
				{"Z"},
			},
		},
		{
			Form: Form{
				EmbeddedRounding: true,
				Zeroing:          true,
			},
			Expect: []Suffixes{
				{},
				{RN_SAE},
				{RZ_SAE},
				{RD_SAE},
				{RU_SAE},
				{Z},
				{RN_SAE, Z},
				{RZ_SAE, Z},
				{RD_SAE, Z},
				{RU_SAE, Z},
			},
		},
	}
	for _, c := range cases {
		got := c.Form.SupportedSuffixes()
		if !reflect.DeepEqual(c.Expect, got) {
			t.Errorf("%v.SupportedSuffixes() = %v; expect %v", c.Form, got, c.Expect)
		}
	}
}

func TestActionValidate(t *testing.T) {
	invalid := []Action{
		M,
		R | M | W,
		R | M,
	}
	for _, a := range invalid {
		if err := a.Validate(); err == nil {
			t.Errorf("action %q is invalid but passed validation", a)
		}
	}
}
