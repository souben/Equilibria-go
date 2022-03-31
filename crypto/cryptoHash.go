package crypto

import (
	"encoding/hex"
	"hash"
)

const HASH_SIZE = 32

type Hash struct {
	Data [HASH_SIZE]byte
}

func (h Hash) ToString() string {
	return hex.EncodeToString(h.Data[:])
}

func (h *Hash) SetData(d [HASH_SIZE]byte) {
	(*h).Data = d
}

func CnFastHash(data []byte) (Hash, error) {
	var h hash.Hash = New256()
	_, err := h.Write(data)
	if err != nil {
		return Hash{}, err
	} else {
		var b [HASH_SIZE]byte
		copy(b[:], h.Sum(nil)) // need to copy slice data to array, // need to improve
		return Hash{b}, err
	}
}
