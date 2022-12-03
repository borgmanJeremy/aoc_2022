package main

import (
	"testing"
)

func TestToNumber(t *testing.T) {
	if toNumber('a') != 1 {
		t.Error("Conversion was incorrect got: ", toNumber('a'), "want: ", 1)
	}

	if toNumber('z') != 26 {
		t.Error("Conversion was incorrect got: ", toNumber('z'), "want: ", 26)
	}

	if toNumber('A') != 27 {
		t.Error("Conversion was incorrect got: ", toNumber('A'), "want: ", 27)
	}

	if toNumber('Z') != 52 {
		t.Error("Conversion was incorrect got: ", toNumber('Z'), "want: ", 52)
	}
}
