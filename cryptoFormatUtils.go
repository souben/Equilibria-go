package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash"
	"log"

	"github.com/souben/equi/crypto"
)

const HASH_SIZE = 32

type Block struct {
	MajorVersion uint8
	MinorVersion uint8
	Timestamp    uint64
	PrevId       crypto.Hash
	Nonce        uint32
	//minerTx Transaction
	TxHashes []crypto.Hash
}

type blobData struct {
	data []byte
}

func get_block_longhash(b Block, res *hash.Hash) bool {
	bLocal := b
	//bd := getBlockHashing(b);

	if bLocal.MajorVersion < 6 {
		// hash
		return true
	} else {
		// hash
		return false
	}
	return false
}

func getTxTreeHash(b *Block) crypto.Hash {
	txsIds := make([]crypto.Hash, 0)
	var h crypto.Hash
	txsIds = append(txsIds, h)
	for _, th := range b.TxHashes {
		txsIds = append(txsIds, th)
	}
	return getTxTreeHashHelper(&txsIds)
}

func getTxTreeHashHelper(txsIds *[]crypto.Hash) crypto.Hash {
	var h crypto.Hash
	getTxTreeHashHelperTxsIds(txsIds, &h)
	return h
}

func getTxTreeHashHelperTxsIds(txIds *[]crypto.Hash, h *crypto.Hash) {
	crypto.TreeHash(txIds, uint64(len(*txIds)), h)
}

func GetBlockHashingBlob(b *Block) blobData {

	var network bytes.Buffer
	var blobdata blobData
	enc := gob.NewEncoder(&network)

	// serilaize block
	err := enc.Encode((*b).MajorVersion)
	err = enc.Encode((*b).MinorVersion)
	err = enc.Encode((*b).Timestamp)
	err = enc.Encode((*b).PrevId)
	err = enc.Encode((*b).Nonce)
	if err != nil {
		log.Fatal("error while encoding block data", err)
	}

	// serialize transactions
	treeRootHash := getTxTreeHash(b)
	fmt.Println(treeRootHash)
	err = enc.Encode(treeRootHash.Data)
	err = enc.Encode(len(treeRootHash.Data))
	if err != nil {
		log.Fatal("error while encoding root hash", err)
	}
	// add transactions
	err = enc.Encode(1 + len((*b).TxHashes))
	if err != nil {
		log.Fatal("error while encoding number of transactions", err)
	}
	blobdata.data = network.Bytes()
	return blobdata
}
