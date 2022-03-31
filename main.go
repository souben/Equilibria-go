package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/souben/equi/crypto"
)

func main() {

	hash := crypto.Hash{}
	hash_array := [32]byte{}
	data, err := hex.DecodeString("85bb9128c170896673aa1b47f2c7d238f77b6c6f06cd7f25b399747d5015577e")
	copy(hash_array[:], data)
	hash.SetData(hash_array)
	b := Block{
		1,
		9,
		1541014386,
		hash,
		1960719487,
		//minerTx Transaction
		[]crypto.Hash{},
	}
	blobdata := GetBlockHashingBlob(&b)
	h, err := crypto.CnFastHash(blobdata.data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(h.ToString())
}
