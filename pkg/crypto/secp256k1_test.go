package crypto

import (
	"bytes"
	"testing"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func TestCompressP256K1Pubkey(t *testing.T) {
	var pubkey [64]byte
	for i := 0; i < 200; i++ {
		prvkey, err := ethcrypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		prvkey.PublicKey.X.FillBytes(pubkey[:32])
		prvkey.PublicKey.Y.FillBytes(pubkey[32:])
		want := ethcrypto.CompressPubkey(&prvkey.PublicKey)
		got := CompressP256k1Pubkey(pubkey)
		if !bytes.Equal(got, want) {
			t.Fatal("not equal")
		}
	}
}
