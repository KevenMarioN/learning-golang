// Unit test
package addresses_test

import (
	. "test-introduction/addresses"
	"testing"
)

type sceneryTest struct {
	addressTest     string
	addressWaitType string
}

func TestTypeOfAddress(t *testing.T) {
	t.Parallel()
	testsScenery := []sceneryTest{
		{"Praça Brasil", "Address invalid!"},
		{"Rua Serra Negra", "Rua"},
		{"Avenida Salvador", "Avenida"},
		{"Estrada Interestadual", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA JB", "Avenida"},
		{"Viaduto Aquarela", "Address invalid!"},
	}
	for _, scenery := range testsScenery {
		result := TypeAddress(scenery.addressTest)
		if result != scenery.addressWaitType {
			t.Errorf("The type recived is different the wait! Wait: %s and recived: %s", scenery.addressWaitType, result)
		}
	}
}

func TestAlways(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Test Broken!")
	}
}
