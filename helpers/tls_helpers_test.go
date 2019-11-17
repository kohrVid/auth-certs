package helpers

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	pwd, _              = os.Getwd()
	certFileName string = "test.crt"
	keyFileName  string = "test.key"
)

func TestTlsCertsCertFile(t *testing.T) {
	expectedCert := filepath.Join(pwd, certFileName)
	cert, _ := TlsCerts(".", certFileName, keyFileName)

	if cert != expectedCert {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			cert,
			expectedCert,
		)
	}
}

func TestTlsCertsKeyFile(t *testing.T) {
	expectedKey := filepath.Join(pwd, keyFileName)
	_, key := TlsCerts(".", certFileName, keyFileName)

	if key != expectedKey {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			key,
			expectedKey,
		)
	}
}
