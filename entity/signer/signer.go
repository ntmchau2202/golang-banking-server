package signer

import (
	"crypto/ecdsa"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer struct {
	privateKey *ecdsa.PrivateKey
}

const (
	prefix string = "\x19Ethereum Signed Message:\n"
)

func NewSigner(privKey string) (s *Signer, err error) {
	fmt.Println("Going to parse privKey:", privKey)
	pK, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return
	}
	pubKey := hexutil.Encode(crypto.FromECDSAPub(&pK.PublicKey))
	fmt.Println("public key:", pubKey)
	addr := crypto.Keccak256Hash(crypto.FromECDSAPub(&pK.PublicKey)).String()
	// addr = addr[24:]
	fmt.Println("Address:", addr)
	s = &Signer{
		privateKey: pK,
	}
	return
}

func (s *Signer) createMessage(msg string) (actualMsg string) {
	actualMsg = prefix + strconv.FormatInt(int64(len(msg)), 10) + msg
	return
}

func (s *Signer) Sign(msg string) (encodedSignature string, err error) {
	actualMsg := s.createMessage(msg)
	fmt.Println("actual message:", actualMsg)
	hash := crypto.Keccak256Hash([]byte(actualMsg))
	fmt.Println("hash:", hash)
	signature, err := crypto.Sign(hash.Bytes(), s.privateKey)
	if err != nil {
		return
	}
	fmt.Println("signature:", hexutil.Encode(signature))
	return hexutil.Encode(signature), nil
}
