package main

import (
	"encoding/pem"
	"fmt"
	"os"

	"github.com/google/go-attestation/attest"
)

func main() {
	tpm, _ := attest.OpenTPM(nil)
	eks, _ := tpm.EKs()
	ek := eks[0]

	os.WriteFile("Certificate.crt", ek.Certificate.Raw, 0600)
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "Certificate",
		Bytes: ek.Certificate.Raw,
	})

	fmt.Printf("%s", certPem)
}
