package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	privateKey,err :=  crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyInBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyInBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot parse to public key")
	}

	publicKeyInBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyInBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyInBytes[1:])
	fmt.Println(hash.Sum(nil)[12:])
}