package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

//获取服务端证书配置
func GetServerCreds() credentials.TransportCredentials {
	//服务端加载证书
	//creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_passwd.key")
	//if err != nil {
	//	log.Fatal(err)
	//}
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: 				 []tls.Certificate{cert}, //服务端证书
		ClientAuth:                  tls.RequireAndVerifyClientCert, //双向验证
		ClientCAs:                   certPool,
	})
	return creds
}

//获取客户端证书配置
func GetClientCreds() credentials.TransportCredentials {
	//客户端加载证书
	//creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "leixiaotian")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//客户端加载证书,双向验证
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: 				 []tls.Certificate{cert}, //客户端证书
		ServerName:                  "localhost",
		RootCAs:                     certPool,
	})
	return creds
}
