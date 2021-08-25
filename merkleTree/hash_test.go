package merkleTree

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing hashes
func TestHash(t *testing.T) {

	tests := [][]string{
		{"first", "e0996a37c13d44c3b06074939d43fa3759bd32c1"},
		{"second", "352f7829a2384b001cc12b0c2613c756454a1f6a"},
		{"third", "34fb3300b9a77bebdc988ec3edd0d4a6a42a26f9"},
		{"fourth", "2db18e1d98e7ab7f49dea56027312c2d97b1a2e0"},
		{"fifth", "5ad43148c90a8f2d296fc74f9f327538e964d000"},
		{"sixth", "04053a7b8a6957822a1a10641c094af04adc071e"}}

	for _, val := range tests {
		hash := calculateHashFromBytes([]byte(val[0]))
		hashString := hex.EncodeToString(hash[:])

		assert.Equal(t, hashString, val[1])
	}

}
