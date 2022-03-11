package automapper_test

import (
	"testing"

	"github.com/rafiulgits/automapper"
)

func TestMapping(t *testing.T) {
	type A struct {
		Name  string
		Phone string
	}

	type B struct {
		Name  string
		Phone string
		Email string
	}

	a := &A{Name: "Tony", Phone: "12355"}
	b := &B{}

	automapper.Map(a, b, func(src *A, dst *B) {
		dst.Email = "hello@world.com"
	})

	if b.Name != a.Name || b.Phone != a.Phone {
		t.Error("failed to map by reflaction")
	}

	if b.Email != "hello@world.com" {
		t.Error("failed to map by callback")
	}
}
