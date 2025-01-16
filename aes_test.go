package crypto4go

import (
	"testing"
)

//import (
//	"testing"
//	"fmt"
//	"encoding/hex"
//)
//
//var plain = "use this method to release shared resources, save user data, invalidate timers, and store enough application state information to restore your application to its current state in case it is terminated later."
//var key = "11111111111111111111111111111111"
//var iv  = "1111111111111111"
//
//func Test_AES_CBC(t *testing.T) {
//	var a = []byte(plain)
//	var r, _ = AESCBCEncrypt(a, []byte(key), []byte(iv))
//	fmt.Println("AES CBC Encrypt: ", hex.EncodeToString(r))
//
//	r, _ = AESCBCDecrypt(r, []byte(key), []byte(iv))
//	fmt.Println("AES CBC Decrypt: ", string(r))
//}
//
//func Test_AES_CFB(t *testing.T) {
//	var a = []byte(plain)
//	var r, _ = AESCFBEncrypt(a, []byte(key), []byte(iv))
//	fmt.Println("AES CFB Encrypt: ", hex.EncodeToString(r))
//
//	r, _ = AESCFBDecrypt(r, []byte(key), []byte(iv))
//	fmt.Println("AES CFB Decrypt: ", string(r))
//}

var (
	kPlaintext = "test data"
	kKey128    = []byte("test-key-aes-128")                 // AES-128，长度为 16
	kKey192    = []byte("test-key-aes-192-0000000")         // AES-192，长度为 24
	kKey256    = []byte("test-key-aes-256-000000000000000") // AES-256，长度为 32
)

func TestAESGCMEncrypt_128(t *testing.T) {
	ciphertext, err := AESGCMEncrypt([]byte(kPlaintext), kKey128)
	if err != nil {
		t.Fatal(err)
	}

	plaintext, err := AESGCMDecrypt(ciphertext, kKey128)
	if err != nil {
		t.Fatal(err)
	}

	if string(plaintext) != kPlaintext {
		t.Fatal("加解密失败")
	}
}
