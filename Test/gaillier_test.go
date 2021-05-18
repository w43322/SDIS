package Test

import (
	"crypto/rand"
	"math/big"
	"testing"

	"prototype1.0.0/gaillier"
)

func TestAdd(t *testing.T) {
	case1 := new(big.Int).SetInt64(1)
	case2 := new(big.Int).SetInt64(1)

	pub, priv, err := gaillier.GenerateKeyPair(rand.Reader, 512)

	if err != nil {
		t.Errorf("Error Generating Keypair")
	}
	//Encrypt
	encCase1, err1 := gaillier.Encrypt(pub, case1.Bytes())
	encCase2, err2 := gaillier.Encrypt(pub, case2.Bytes())

	if err1 != nil || err2 != nil {
		t.Errorf("Error Encrypting Integers")
	}

	res := gaillier.Add(pub, encCase1, encCase2)

	corr := new(big.Int).SetInt64(2)

	decRes, err := gaillier.Decrypt(priv, res)
	if err != nil {
		t.Errorf("Failed to Decrypt Result got %v want %v with Error : %v", decRes, corr, err)
	}

	resB := new(big.Int).SetBytes(decRes)

	if resB.Cmp(corr) != 0 {
		t.Errorf("Failed to Add two ciphers got %v want %v", resB, corr)
	}

}

func TestAddConstant(t *testing.T) {

	k := new(big.Int).SetInt64(10)
	c := new(big.Int).SetInt64(32)

	pub, priv, err := gaillier.GenerateKeyPair(rand.Reader, 102)

	if err != nil {
		t.Errorf("Failed to generated keypair %v", err)
	}

	//encrypt
	encC, err := gaillier.Encrypt(pub, c.Bytes())
	if err != nil {
		t.Errorf("Failed to encrypt c")
	}
	res := gaillier.AddConstant(pub, encC, k.Bytes())

	decRes, err := gaillier.Decrypt(priv, res)
	if err != nil {
		t.Errorf("Failed to decrypt result")
	}

	result := new(big.Int).SetBytes(decRes)
	corr := new(big.Int).SetInt64(42)
	if result.Cmp(corr) != 0 {
		t.Errorf("Error Add Constant function want %d , got %d", corr, result)
	}

}

func TestMul(t *testing.T) {

	k := new(big.Int).SetInt64(10)
	c := new(big.Int).SetInt64(32)

	pub, priv, err := gaillier.GenerateKeyPair(rand.Reader, 102)

	if err != nil {
		t.Errorf("Failed to generated keypair %v", err)
	}

	//encrypt
	encC, err := gaillier.Encrypt(pub, c.Bytes())
	if err != nil {
		t.Errorf("Failed to encrypt c")
	}
	res := gaillier.Mul(pub, encC, k.Bytes())

	decRes, err := gaillier.Decrypt(priv, res)
	if err != nil {
		t.Errorf("Failed to decrypt result")
	}

	result := new(big.Int).SetBytes(decRes)
	corr := new(big.Int).SetInt64(320)
	if result.Cmp(corr) != 0 {
		t.Errorf("Error Mul function want %d , got %d", corr, result)
	}
}

func BenchmarkEncrypt(b *testing.B) {
	pubKey, _, _ := gaillier.GenerateKeyPair(rand.Reader, 2048)
	val := new(big.Int).SetInt64(int64(1234)).Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := gaillier.Encrypt(pubKey, val)
		if err != nil {
			b.Fail()
		}
	}
}
