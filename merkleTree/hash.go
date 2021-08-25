package merkleTree

import (
	"crypto/sha1"
	"encoding/hex"
	"hash/fnv"
)

/*
* A very simple implementation of hash with sha1
 */

type Hash [20]byte

func calculateHashFromBytes(s []byte) Hash {
	h := fnv.New32a()
	h.Write(s)
	return sha1.Sum(s)
}

func (h Hash) ToString() string {
	return hex.EncodeToString(h[:])
}
