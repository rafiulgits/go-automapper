package automapper_test

import (
	"testing"
	"time"

	"github.com/rafiulgits/go-automapper"
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

func TestArrayMapping(t *testing.T) {
	type A struct {
		Name string
	}

	type B struct {
		ID   int
		Name string
	}

	a1 := &A{Name: "One"}
	a2 := &A{Name: "Two"}

	aSlice := []*A{a1, a2}
	bSlice := []*B{}

	automapper.Map(aSlice, &bSlice, func(idx int, src *A, dst *B) {
		t.Log(idx, src, dst)
		dst.ID = idx + 1
	})

	if bSlice[0].ID != 1 && bSlice[1].ID != 2 {
		t.Error("slice map callback failed")
	}
}

func TestTimeMapping(t *testing.T) {
	type A struct {
		Time string
	}

	type B struct {
		Time time.Time
	}

	type C struct {
		ID   int
		Time time.Time
	}

	a := &A{Time: time.Now().Format(time.RFC3339)}
	b := &B{Time: time.Now()}
	c1 := &C{}
	c2 := &C{}

	automapper.Map(a, c1)
	automapper.Map(b, c2)

	if c1.Time.IsZero() {
		t.Error("time string to time object mapping failed")
	}

	if c2.Time.IsZero() {
		t.Error("time object to time object mapping failed")
	}

}
