package imgconverter

import "testing"

var extensionTests = []struct {
	s          string
	optionType string
	out        bool
}{
	{".jpeg", "input", true},
	{".jpeg", "output", true},
	{".gif", "input", true},
	{".gif", "output", false},
	{".test", "input", false},
}

func TestIsValidExtension(t *testing.T) {
	for _, tt := range extensionTests {
		returnedValue := IsValidExtension(tt.s, tt.optionType)
		if tt.out != returnedValue {
			t.Errorf("(%s, %s), want %T, but returned %T", tt.s, tt.optionType, tt.out, returnedValue)
		}
	}
}
