// Code generated by rsagen, DO NOT EDIT.

package mtproto

import (
	"crypto/rsa"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gotd/td/internal/crypto"
)

func TestVendoredKeys(t *testing.T) {
	fingerprint := func(pubkey *rsa.PublicKey) string {
		return fmt.Sprintf("%08x", uint64(crypto.RSAFingerprint(pubkey)))
	}
	expected := []string{
		"c3b42b026ce86b21",
		"bc35f3509f7b7a5",
		"15ae5fa8b5529542",
		"aeae98e13cd7f94f",
		"5a181b2235057d98",
	}
	assert.Len(t, vendoredKeys, len(expected))
	for i, pubkey := range vendoredKeys {
		assert.Equal(t, expected[i], fingerprint(pubkey))
	}
}
