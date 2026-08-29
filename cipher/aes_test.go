package cipher

import "testing"

func TestAesDefaultStringRoundTrip(t *testing.T) {
	ciphertext, err := AesEncDefaultStr("Hello World")
	if err != nil {
		t.Fatalf("encrypt: %v", err)
	}

	plaintext, err := AesDecDefaultStr(ciphertext)
	if err != nil {
		t.Fatalf("decrypt: %v", err)
	}
	if plaintext != "Hello World" {
		t.Fatalf("plaintext = %q, want %q", plaintext, "Hello World")
	}
}

func TestAesDecRejectsInvalidCiphertext(t *testing.T) {
	if _, err := AesDec([]byte{1}, DefaultAesKey); err == nil {
		t.Fatal("AesDec accepted a partial block")
	}
}
