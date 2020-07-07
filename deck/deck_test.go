package deck

import "testing"

func TestNew(t *testing.T) {
	//setup
	d, _ := New()

	expected := 54
	actual := len(d)

	if actual != expected {
		t.Errorf("expected deck length was %v but actual shows %v", expected, actual)
	}
}
