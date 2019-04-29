package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Errorf("singleton instance is nil")
	}

	currentCount := counter1.Increase()
	if currentCount != 1 {
		t.Errorf("unexpected counter value, expected to be 1; got %d", currentCount)
	}

	// instance a second counter
	counter2 := GetInstance()

	// we expect that both counters are the same
	if counter2 != counter1 {
		t.Errorf("expected the counter2 to be the same as counter1 but got a different instance")
	}
}
