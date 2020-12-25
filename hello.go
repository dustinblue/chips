
package main

import "fmt"
import "github.com/dustinblue/chips/src/lib"
import "crypto/sha256"
import "crypto/rand"
import "google.golang.org/protobuf/encoding/protojson"

// One coin is 100 million chips
const coin uint64 = 100000000;

func makeKey() ([]byte) {
    key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		// handle error here
	}

	return key
}

func main() {

    // Make a 256bit Private Key
    fmt.Printf("%b", makeKey())

    // coinbase is a special transaction that is allowed to pay out without referencing an input (UTXO)
    coinbase := chips.Transaction{}
    output := chips.Output{}
    output.Address = []byte{1,2,3,4,5}
    output.Amount = 50 * coin
    coinbase.Outputs = []*chips.Output{&output}

    genesis := chips.Block{}
    genesis.HashBlock = []byte{1,2,3,4}
    genesis.HashMerkleRoot = make([]byte, 0, 1)
    fmt.Printf("%x", sha256.Sum256(genesis.HashMerkleRoot))
    genesis.Transactions = []*chips.Transaction{&coinbase}
    fmt.Println("%s", genesis)

    m := protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: true}

    jsonData, _ := m.Marshal(&genesis)
    fmt.Println(string(jsonData))

    h := sha256.New()
    h.Write([]byte(genesis.HashBlock))
    fmt.Printf("%x", h.Sum(nil))
    fmt.Printf("%b", sha256.Sum256(genesis.HashBlock))
}
