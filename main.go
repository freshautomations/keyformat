package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/libs/tempfile"
	"github.com/cometbft/cometbft/privval"
	"os"
)

func main() {
	keyPtr := flag.String("key", "", "Decrypted key file (96 bytes)")
	outputPtr := flag.String("output", "", "Output file")
	softsignPtr := flag.Bool("softsign", false, "Output a softsign private key instead of priv_validator.json format")
	flag.Parse()

	key, err := os.ReadFile(*keyPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening key file: %v", err)
		os.Exit(1)
	}

	// Optionally, base64-encoded private keys are also accepted.
	data, err2 := base64.StdEncoding.DecodeString(string(key))
	if err2 != nil {
		data = key
	}

	// YubiHSM exported ed25519 keys are 96 bytes.
	// The first 32 bytes are unknown, maybe something to do with the original wrapping.
	// The second 32 bytes are the private key.
	// THe third 32 bytes are the public key.
	if len(data) != 96 {
		fmt.Fprintf(os.Stderr, "invalid key file length: %d (should be 96)", len(key))
		os.Exit(1)
	}

	if *softsignPtr {
		privKey := base64.StdEncoding.EncodeToString(data[32:64])
		err = tempfile.WriteFileAtomic(*outputPtr, []byte(privKey), 0600)
		if err != nil {
			panic(err)
		}
	} else {
		filePV := privval.NewFilePV(ed25519.PrivKey(data[32:]), *outputPtr, "")
		filePV.Key.Save()
	}
}
