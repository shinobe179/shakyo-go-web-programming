package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	// 証明書で使うシリアルナンバーの生成
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	// 証明書のサブジェクト部を設定する
	subject := pkix.Name{
		Organization:     []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:       "Go Web Programming",
	}

	// certificate構造体で証明書の構成を設定する
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:     subject,
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}

	// RSAキーペアを生成する
	// https://golang.org/pkg/crypto/rsa/#GenerateKey
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// DER形式証明書のバイトデータのスライスを生成する
	// https://golang.org/pkg/crypto/x509/#CreateCertificate
	// func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv interface{}) (cert []byte, err error)
	// The parameter pub is the public key of the signee and priv is the private key of the signer.
	// output: (cert []byte, err error)
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)

	// encoding/pemを使って証明書データを符号化(エンコード)してファイルにする
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// 生成したプライベート鍵を符号化してファイルにする
	// https://golang.org/pkg/crypto/x509/#MarshalPKCS1PrivateKey
	// MarshalPKCS1PrivateKey converts a private key to ASN.1 DER encoded form.
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
