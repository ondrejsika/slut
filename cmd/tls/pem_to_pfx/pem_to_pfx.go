package parse

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	tls_cmd "github.com/sikalabs/slu/cmd/tls"
	pkcs12 "software.sslmate.com/src/go-pkcs12"

	"github.com/spf13/cobra"
)

var FlagInputKey string
var FlagInputCert string
var FlagInputCaCert string
var FlagOutput string
var FlagPassword string

var Cmd = &cobra.Command{
	Use:   "pem-to-pfx",
	Short: "Convert PEM (cer/key) to PFX",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		tlsCert, err := tls.LoadX509KeyPair(FlagInputCert, FlagInputKey)
		if err != nil {
			log.Fatal(err)
		}
		cert, err := x509.ParseCertificate(tlsCert.Certificate[0])
		if err != nil {
			log.Fatal(err)
		}

		var caCerts []*x509.Certificate
		if FlagInputCaCert != "" {
			f, err := os.ReadFile(FlagInputCaCert)
			if err != nil {
				log.Fatal(err)
			}
			block, _ := pem.Decode(f)
			if block == nil {
				panic("failed to parse certificate PEM")
			}
			ca, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				log.Fatal(err)
			}
			caCerts = append(caCerts, ca)
		}

		pfxBytes, err := pkcs12.Encode(rand.Reader, tlsCert.PrivateKey, cert, caCerts, FlagPassword)

		if err != nil {
			log.Fatal(err)
		}

		// see if pfxBytes valid
		_, _, _, err = pkcs12.DecodeChain(pfxBytes, pkcs12.DefaultPassword)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(
			FlagOutput,
			pfxBytes,
			os.ModePerm,
		)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	tls_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagOutput,
		"out",
		"o",
		"",
		"Output cert.pfx",
	)
	Cmd.MarkFlagRequired("out")

	Cmd.Flags().StringVarP(
		&FlagInputCert,
		"in-cert",
		"c",
		"",
		"Input cert.pem",
	)
	Cmd.MarkFlagRequired("in-cert")
	Cmd.Flags().StringVarP(
		&FlagInputKey,
		"in-key",
		"k",
		"",
		"Input key.pem",
	)
	Cmd.MarkFlagRequired("in-key")
	Cmd.Flags().StringVarP(
		&FlagInputCaCert,
		"in-ca",
		"a",
		"",
		"Input ca-cert.pem",
	)
	Cmd.Flags().StringVarP(
		&FlagPassword,
		"password",
		"p",
		pkcs12.DefaultPassword,
		"Password for output PFX file",
	)
}
