package pki

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	vault "github.com/hashicorp/vault/api"
	"golang.org/x/crypto/sha3"
	"log"
	"main/config"
	"time"
)

type PKI struct {
	PublicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	client     *vault.Client
}

func NewPKI(cfg config.Config) (*PKI, error) {
	publicKeyExists, privateKeyExists := false, false

	vaultConfig := vault.DefaultConfig()
	vaultConfig.Address = cfg.Vault.Address

	client, err := vault.NewClient(vaultConfig)
	if err != nil {
		return nil, err
	}

	client.SetToken(cfg.Vault.Token)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	public, err := client.KVv2(cfg.Vault.MountPath).Get(ctx, cfg.Vault.PublicPath)
	if err == nil {
		publicKeyExists = true
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	private, err := client.KVv2(cfg.Vault.MountPath).Get(ctx, cfg.Vault.PrivatePath)
	if err == nil {
		privateKeyExists = true
	}

	p := &PKI{
		client: client,
	}

	if !publicKeyExists && !privateKeyExists {
		log.Printf("generating new keys")

		err = p.generateKeys()
		if err != nil {
			return nil, err
		}

		err = p.writePublicKey(cfg.Vault.MountPath, cfg.Vault.PublicPath)
		if err != nil {
			return nil, err
		}

		err = p.writePrivateKey(cfg.Vault.MountPath, cfg.Vault.PrivatePath)
		if err != nil {
			return nil, err
		}
	} else if !publicKeyExists && privateKeyExists {
		return nil, errors.New("public key not found")
	} else if publicKeyExists && !privateKeyExists {
		return nil, errors.New("private key not found")
	} else {
		log.Printf("reading keys")

		valuePublic, ok := public.Data[cfg.Vault.PublicPath].(string)
		if !ok {
			return nil, errors.New("public key not found")
		}

		data, err := base64.StdEncoding.DecodeString(valuePublic)
		if err != nil {
			return nil, err
		}

		err = p.readPublicKey(data)
		if err != nil {
			return nil, err
		}

		valuePrivate, ok := private.Data[cfg.Vault.PrivatePath].(string)
		if !ok {
			return nil, errors.New("private key not found")
		}

		data, err = base64.StdEncoding.DecodeString(valuePrivate)
		if err != nil {
			return nil, err
		}

		err = p.readPrivateKey(data)
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

func (receiver *PKI) writePublicKey(mountPath, publicPath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	block := &pem.Block{
		Type:  "BEGIN RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(receiver.PublicKey),
	}

	data := map[string]interface{}{
		publicPath: base64.StdEncoding.EncodeToString(pem.EncodeToMemory(block)),
	}

	_, err := receiver.client.KVv2(mountPath).Put(ctx, publicPath, data)
	if err != nil {
		return err
	}

	return nil
}

func (receiver *PKI) writePrivateKey(mountPath, privatePath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	block := &pem.Block{
		Type:  "BEGIN RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(receiver.privateKey),
	}

	data := map[string]interface{}{
		privatePath: base64.StdEncoding.EncodeToString(pem.EncodeToMemory(block)),
	}

	_, err := receiver.client.KVv2(mountPath).Put(ctx, privatePath, data)
	if err != nil {
		return err
	}

	return nil
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
