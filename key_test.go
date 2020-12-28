package main

import "testing"
import "crypto/ecdsa"
import "crypto/rand"
import "crypto/elliptic"
import "crypto/sha256"
import "golang.org/x/crypto/ripemd160"
import "math/big"
import "github.com/ethereum/go-ethereum/crypto/secp256k1"
import "github.com/btcsuite/btcutil/base58"

//import "crypto/x509"
import "fmt"

/*
Input Scripts- ScriptSig: PUSHDATA(71)[3044022069f0c9da12ce7b002c68d25e4445191b149627987367fdec4e4b80e442379f3902202109994c46d8c6afd8e94542d07fa8a07da234f7537c4a1d639d2dcdd32f9c5b01] PUSHDATA(33)[03c26e5ff8c75d6272b2d9cd3a882c20315d440251deb79873b75b690d0b71f4cf]

Output Scripts- DUP HASH160 PUSHDATA(20)[264cf7a09b68a436bafc4d7281743d7f1c721ded] EQUALVERIFY CHECKSIG HASH160 PUSHDATA(20)[72cf56ba8b7312ae658debf033d88e4370b9a8f5] EQUAL
*/


// https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses#How_to_create_Bitcoin_Address

// Attempting to generate a Bitcoin address based on above URL. The final address is 1PMycacnJaSqwwJqjawXBErnLsZ7RkXUAs
func TestBitcoinAddress(t *testing.T) {
  // t.Skip()
  privKey := "18e14a7b6a307f426a94f8114701e7c8e774e7f9a47e2c2035db29a206321725"
  fmt.Printf("Step 0: %s\n", privKey)

  // Load private key from string, add 0x02 to public key x coord end up with 33 bytes
  var e ecdsa.PrivateKey
  e.PublicKey.Curve = secp256k1.S256()
  e.D, _ = new(big.Int).SetString(privKey, 16)
  e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
  pubKey33 := append([]byte{02}, e.PublicKey.X.Bytes()...)
  fmt.Printf("Step 1: %x\n", pubKey33)

  // Hash 33 byte pubKey with sha256
  hash256 := sha256.Sum256(pubKey33)
  fmt.Printf("Step 2: %x\n", hash256)

  // RIPEMD60 hash
  h := ripemd160.New()
  h.Write(hash256[:])
  hash160 := h.Sum(nil)
  fmt.Printf("Step 3: %x\n", hash160)

  // Version byte 0x00 main network
  step4 := append([]byte{0}, hash160...)
  fmt.Printf("Step 4: %x\n", step4)

  // Checksum sha256 twice
  hash1 := sha256.Sum256(step4)
  fmt.Printf("Step 5: %x\n", hash1)

    // Checksum sha256 twice
    hash2 := sha256.Sum256(hash1[:])
    fmt.Printf("Step 6: %x\n", hash2)

  // First 4 bytes are checksum
  checksum := hash2[:4]
  fmt.Printf("Step 7: %x\n", checksum)

  // Add checksum
  address := append(step4, checksum...)
  fmt.Printf("Step 8: %x\n", address)

  // Finally Base58 encode
  addressBTC := base58.Encode(address)
  fmt.Printf("Step 9: %s\n", addressBTC)
}

func TestAddress(t *testing.T) {
  t.Skip()
    var e ecdsa.PrivateKey
    privateKeyString := "4a1419236a2c0af55b59cdb4d71af961dce8fc781ca2a86572575bd79c433d5e"
    e.PublicKey.Curve = elliptic.P256()
    e.D, _ = new(big.Int).SetString(privateKeyString, 16)
    fmt.Printf("%x", e.D);
    e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
    fmt.Printf("%x", e.Public())
    //privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    //fmt.Printf("%x", privateKey.D)
    //x509.MarshalECPrivateKey(privateKey)
    //fmt.Printf("Public Key X: %x\n", privateKey.PublicKey.X.Bytes())
}

func TestSign(t *testing.T) {
  t.Skip("")

  privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

  fmt.Printf("%d\n", privateKey.D)
  fmt.Printf("%d\n", privateKey.PublicKey.X)

  message := "Test message"
  hash := sha256.Sum256([]byte(message))

  r, s, _ := ecdsa.Sign(rand.Reader, privateKey, hash[:])
  fmt.Printf("%d\n", r)

  verify := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)

  if !verify {
    t.Errorf("signed message: %s not verified\n", message)
  } else {
    fmt.Printf("message: %s signature verified\n", message)
  }

}