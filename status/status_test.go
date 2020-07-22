package status_test

import (
	"testing"

	"github.com/ospokemon/api-go/status"
)

func TestStatus(t *testing.T) {
	status := status.New()

	if status.IsAsleep() {
		t.Fail()
	}
	if status.IsConfused() {
		t.Fail()
	}

	status.SetAsleep()

	if !status.IsAsleep() {
		t.Fail()
	}
	if status.IsConfused() {
		t.Fail()
	}

	status.SetConfused()

	if !status.IsAsleep() {
		t.Fail()
	}
	if !status.IsConfused() {
		t.Fail()
	}

	status.ClearConfused()

	if !status.IsAsleep() {
		t.Fail()
	}
	if status.IsConfused() {
		t.Fail()
	}
}
