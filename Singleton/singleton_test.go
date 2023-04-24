package creational

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Error("Instances are not equal!\n")
	}
}

func TestIncrement(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	instance1.Increment()
	if instance1.count != instance2.count {
		t.Error("The field of count must be the same for both instances\n")
	}

	instance2.Increment()
	if instance1.count != instance2.count {
		t.Error("The field of count must be the same for both instances\n")
	}
}
