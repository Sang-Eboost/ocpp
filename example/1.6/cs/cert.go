package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	logg"log"
)

func main() {
	// Đường dẫn đến các file cert và private key
	certFile := "/etc/letsencrypt/live/car-uat-ocpp.eboost.vn/cert.pem"
	keyFile := "/etc/letsencrypt/live/car-uat-ocpp.eboost.vn/privkey.pem"
	caCertFile := "/etc/letsencrypt/live/car-uat-ocpp.eboost.vn/chain.pem"

	// Tải chứng chỉ và khóa riêng
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		logg.Fatalf("Không thể tải cert/key pair: %v", err)
	}

	// Tạo pool CA và thêm chứng chỉ CA
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		logg.Fatalf("Không thể đọc CA cert: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Cấu hình TLS với chứng chỉ và khóa
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	// In ra thông tin thành công
	fmt.Println("Đã tải thành công chứng chỉ và khóa riêng!")
	fmt.Printf("TLS Config: %+v\n", tlsConfig)
}
