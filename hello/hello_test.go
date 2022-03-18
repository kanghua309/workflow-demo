package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions 4" {
		t.Errorf("Greet() = %s; Expected Hello GitHub actions!", result)
	}
}
