package address

import (
	"testing"
)

type testScenario struct {
	inputAddress   string
	expectedResult string
}

func TestAddressType(t *testing.T) {

	scenarios := []testScenario{
		{"street alpha", "street"},
		{"road 64", "road"},
		{"ROAD 64", "road"},
		{"alley downtown", "alley"},
		{"highway  cruzer", "highway"},
		{"rua dos programadores", "Invalid type"},
		{"", "Invalid type"},
	}

	for _, scenario := range scenarios {
		result := AddressType(scenario.inputAddress)
		if result != scenario.expectedResult {
			t.Errorf("Test failed! Expected %s and received %s", scenario.expectedResult, result)
		}
	}
}
