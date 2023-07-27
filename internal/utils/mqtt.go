package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func NewNanoMQClient(addr string, clientid string) (mqtt.Client,error) {
	opts := mqtt.NewClientOptions().AddBroker(addr).SetClientID(clientid)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	//opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	//opts.SetWill("/device/dead", clientid, 1, false)
	//tlsconfig := NewTLSConfig()
	//opts.SetTLSConfig(tlsconfig)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Error("connect error:", token.Error())
		return nil, token.Error()
	}
	log.Infof("nano mqtt connect %s success!",addr)
	return c,nil
}
//var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
//	fmt.Printf("TOPIC: %s\n", msg.Topic())
//	fmt.Printf("MSG: %s\n", msg.Payload())
//}

func NewTLSConfig() *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to default openssl CA bundle.
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("./cert/ca.pem")
	if err != nil {
		fmt.Println("0. read file error, game over!!")
		panic("TLS File Open Err!!")
	}
	certpool.AppendCertsFromPEM(pemCerts)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: false,
		// Certificates = list of certs client sends to server.
		// Certificates: []tls.Certificate{cert},
	}
}
