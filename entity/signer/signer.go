package signer

import (
	"crypto/ecdsa"
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
	pK, err := crypto.HexToECDSA("c64a65031ae65fd012029834728af47249bca942d9653f7350f23ea83e72576f")
	if err != nil {
		return
	}
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
	hash := crypto.Keccak256Hash([]byte(actualMsg))

	signature, err := crypto.Sign(hash.Bytes(), s.privateKey)
	if err != nil {
		return
	}

	return hexutil.Encode(signature), nil
}
