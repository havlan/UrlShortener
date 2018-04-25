package main

import "testing"

func TestBaseEncoding(t *testing.T) {
	exp_decoded := "jS"
	actual := baseUEncode(1000)
	if actual != exp_decoded {
		t.Errorf("Error in base encoding, got %s wanted %s. ", actual, exp_decoded)
	}
}
func TestBaseDecoding(t *testing.T) {
	exp_decoded := 25602
	actual := baseUDecode("abc")
	if actual != exp_decoded {
		t.Errorf("Error in base decoding, got %d wanted %d. ", actual, exp_decoded)
	}
}
