package pki

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/sha3"
	"log"
	"main/config"
	"os"
)

type PKI struct {
	PublicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func New(cfg config.Config) (*PKI, error) {
	publicKeyExists, privateKeyExists := false, false

	public, err := os.ReadFile(cfg.PKI.PublicKey)
	if err == nil {
		publicKeyExists = true
	}

	private, err := os.ReadFile(cfg.PKI.PrivateKey)
	if err == nil {
		privateKeyExists = true
	}

	p := &PKI{}

	if !publicKeyExists && !privateKeyExists {
		log.Printf("generating new keys")

		err = p.generateKeys()
		if err != nil {
			return nil, err
		}

		err = p.writePublicKey(cfg.PKI.PublicKey)
		if err != nil {
			return nil, err
		}

		err = p.writePrivateKey(cfg.PKI.PrivateKey)
		if err != nil {
			return nil, err
		}
	} else if !publicKeyExists && privateKeyExists {
		return nil, errors.New("public key not found")
	} else if publicKeyExists && !privateKeyExists {
		return nil, errors.New("private key not found")
	} else {
		log.Printf("reading keys")

		err = p.readPublicKey(public)
		if err != nil {
			return nil, err
		}

		err = p.readPrivateKey(private)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (receiver *PKI) generateKeys() error {
	private, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}
	receiver.PublicKey = &private.PublicKey
	receiver.privateKey = private
	return nil
}

func (receiver *PKI) writePublicKey(path string) error {
	block := &pem.Block{
		Type:  "BEGIN RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(receiver.PublicKey),
	}

	return os.WriteFile(path, pem.EncodeToMemory(block), 0600)
}

func (receiver *PKI) writePrivateKey(path string) error {
	block := &pem.Block{
		Type:  "BEGIN RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(receiver.privateKey),
	}

	return os.WriteFile(path, pem.EncodeToMemory(block), 0600)
}

func (receiver *PKI) readPublicKey(data []byte) error {
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "BEGIN RSA PUBLIC KEY" {
		return errors.New("failed to decode PEM block containing public key")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return err
	}

	receiver.PublicKey = publicKey
	return nil
}

func (receiver *PKI) readPrivateKey(data []byte) error {
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "BEGIN RSA PRIVATE KEY" {
		return errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	receiver.privateKey = privateKey
	return nil
}

func (receiver *PKI) Encrypt(message string) (string, error) {
	encrypted, err := rsa.EncryptOAEP(sha3.New512(), rand.Reader, receiver.PublicKey, []byte(message), nil)
	return hex.EncodeToString(encrypted), err
}

func (receiver *PKI) Decrypt(message string) (string, error) {
	encrypted, err := hex.DecodeString(message)
	if err != nil {
		return "", err
	}

	decrypted, err := rsa.DecryptOAEP(sha3.New512(), rand.Reader, receiver.privateKey, encrypted, nil)
	return string(decrypted), err
}
