package envy

import (
	"os"
	"strings"
	"testing"
)

func TestFileOrDefault(t *testing.T) {
	actual, err := fileOrDefault("")
	if err != nil {
		t.Error(err)
	}
	expected := ".env"

	if actual != expected {
		t.Errorf("\nActual: %s\nExpected: %s", actual, expected)
	}

	err = os.Remove(actual)
	if err != nil {
		t.Error("Couldn't remove file")
	}
}

func TestParseFile(t *testing.T) {
	testReader := strings.NewReader("FOO=BAR\n#COMMENT=TRUE\nFOO_FOO=BAR-BAR\nFOO-FOO=BAR-BAR\nFOO.FOO=BAR-BAR")

	kvPair, err := ParseFile(testReader)
	if err != nil {
		t.Error(err)
	}

	t.Log(kvPair)
}

func TestParseLine(t *testing.T) {
	lines := []string{"FOO=BAR", "FOO_FOO=BAR-BAR", "FOO-FOO=BAR-BAR", "FOO.FOO=BAR-BAR"}
	for _, v := range lines {
		kv, err := ParseLine(v)
		if err != nil {
			t.Error(err)
		}
		t.Log(kv.key, kv.value)
	}
}
