package certs

import (
	"airport-app-backend/config"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

// Generates a self-signed TLS certificate. Will overwrite existing files.
func GenerateSelfSignedCertAndKey(certFilePath string, keyFilePath string, certValidityDays int) {
	log.Info().
		Str("cert-file-path", certFilePath).
		Str("key-file-path", keyFilePath).
		Msg("Generating new self-signed TLS certificate and key pair")

	if len(config.TlsCertificateHosts) < 0 {
		log.Fatal().Msg("Missing required host values for certificate")
	}

	privKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		log.Fatal().Err(err).Msg("Error when generating private key")
	}

	notBefore := time.Now()
	log.Info().Int("validity-days", certValidityDays).Msg("Setting self signed certificate validity")
	notAfter := notBefore.Add(time.Hour * 24 * time.Duration(certValidityDays))

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate certificate serial number")
	}

	certTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Self-signed"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	hosts := strings.Split(config.TlsCertificateHosts, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			certTemplate.IPAddresses = append(certTemplate.IPAddresses, ip)
		} else {
			certTemplate.DNSNames = append(certTemplate.DNSNames, h)
		}
	}

	// Generate certificate and save to file
	derBytes, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, privKey.Public(), privKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create certificate")
	}

	certOut, err := os.Create(certFilePath)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to open '%s' for writing", certFilePath)
	}

	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		log.Fatal().Err(err).Msgf("Failed to write data to '%s'", certFilePath)
	}

	if err := certOut.Close(); err != nil {
		log.Fatal().Err(err).Msgf("Error closing '%s'", certFilePath)
	}

	log.Info().Msg("Finished writing certificate to file")

	// Marshall and save private key to file
	keyOut, err := os.OpenFile(keyFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to open '%s' for writing", keyFilePath)
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to marshall private key")
	}

	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		log.Fatal().Err(err).Msgf("Failed to write data to '%s'", keyFilePath)
	}

	if err := keyOut.Close(); err != nil {
		log.Fatal().Err(err).Msgf("Error closing '%s'", keyFilePath)
	}

	log.Info().Msg("Finished writing private key to file")
}
