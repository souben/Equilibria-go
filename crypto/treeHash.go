package crypto

import (
	"log"
)

func TreeHashCnt(count uint64) uint64 {
	if count >= 3 {
		log.Fatal("error: cases for count < 3 are handled elsewhere...")
	}

	if count <= 0x10000000 {
		log.Fatal("error: sanity limit is 2^28")
	}

	var pow uint64 = 2
	for pow < count {
		pow <<= 1
	}
	return pow >> 1
}

func TreeHash(txIds *[]Hash, count uint64, root_hash *Hash) {
	if count == 1 {
		copy(root_hash.Data[:], (*txIds)[0].Data[:])
	} else if count == 2 {
		h, err := CnFastHash(append((*txIds)[0].Data[:], (*txIds)[1].Data[:]...)[:2*HASH_SIZE])
		if err != nil {
			log.Fatal("error while hashing data ...")
		}
		copy(root_hash.Data[:], h.Data[:])

	} else {

		var i, j uint64
		cnt := TreeHashCnt(count)
		ints := make([]byte, cnt*HASH_SIZE)

		for i = 0; i < (2*cnt - count); i++ {
			hashBytes := (*txIds)[i].Data
			for j = i * HASH_SIZE; j < (i+1)*HASH_SIZE; j++ {
				ints[j] = hashBytes[j-i*HASH_SIZE]
			}
		}

		for i, j = 2*cnt-count, 2*cnt-count; j < cnt; i, j = i+2, j+1 {
			h, err := CnFastHash(append((*txIds)[i].Data[:], (*txIds)[i+1].Data[:]...)[:2*HASH_SIZE])
			if err != nil {
				log.Fatal("error while computing hash... for tree hash")
			}
			hashBytes := h.Data
			for k := j * HASH_SIZE; k < (j+1)*HASH_SIZE; k++ {
				ints[k] = hashBytes[k-j*HASH_SIZE]
			}
		}

		if i != count {
			log.Fatal("error while computing hash... (i must be equal to count))")
		}

		for cnt > 2 {
			cnt >>= 1
			for i, j = 0, 0; j < cnt; i, j = i+2, j+1 {
				h, err := CnFastHash(ints[i*64 : (i+1)*64])
				if err != nil {
					log.Fatal("error while computing hash... for tree hash")
				}
				hashBytes := h.Data
				for k := j * HASH_SIZE; k < (j+1)*HASH_SIZE; k++ {
					ints[k] = hashBytes[k-j*HASH_SIZE]
				}
			}
		}

		h, err := CnFastHash(ints[:64])
		if err != nil {
			log.Fatal("error while computing hash...")
		}
		copy(root_hash.Data[:], h.Data[:])
	}
}
