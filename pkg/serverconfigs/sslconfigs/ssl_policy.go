package sslconfigs

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"golang.org/x/net/http2"
)

// TLSVersion TLS Version
type TLSVersion = string

// TLSCipherSuite Cipher Suites
type TLSCipherSuite = string

// SSLPolicy SSL配置
type SSLPolicy struct {
	Id   int64 `yaml:"id" json:"id"`     // ID
	IsOn bool  `yaml:"isOn" json:"isOn"` // 是否开启

	CertRefs       []*SSLCertRef     `yaml:"certRefs" json:"certRefs"`
	Certs          []*SSLCertConfig  `yaml:"certs" json:"certs"`
	ClientAuthType SSLClientAuthType `yaml:"clientAuthType" json:"clientAuthType"` // 客户端认证类型
	ClientCARefs   []*SSLCertRef     `yaml:"clientCARefs" json:"clientCARefs"`     // 客户端认证CA证书引用
	ClientCACerts  []*SSLCertConfig  `yaml:"clientCACerts" json:"clientCACerts"`   // 客户端认证CA

	MinVersion       TLSVersion       `yaml:"minVersion" json:"minVersion"`             // 支持的最小版本
	CipherSuitesIsOn bool             `yaml:"cipherSuitesIsOn" json:"cipherSuitesIsOn"` // 是否自定义加密算法套件
	CipherSuites     []TLSCipherSuite `yaml:"cipherSuites" json:"cipherSuites"`         // 加密算法套件

	HSTS         *HSTSConfig `yaml:"hsts" json:"hsts"`                 // HSTS配置
	HTTP2Enabled bool        `yaml:"http2Enabled" json:"http2Enabled"` // 是否启用HTTP2

	nameMapping map[string]*tls.Certificate // dnsName => cert

	minVersion   uint16
	cipherSuites []uint16

	clientCAPool *x509.CertPool

	tlsConfig *tls.Config
}

// Init 校验配置
func (this *SSLPolicy) Init() error {
	this.nameMapping = map[string]*tls.Certificate{}

	// certs
	var certs = []tls.Certificate{}
	for _, cert := range this.Certs {
		err := cert.Init()
		if err != nil {
			return err
		}
		certs = append(certs, *cert.CertObject())
		for _, dnsName := range cert.DNSNames {
			this.nameMapping[dnsName] = cert.CertObject()
		}
	}

	// CA certs
	for _, cert := range this.ClientCACerts {
		err := cert.Init()
		if err != nil {
			return err
		}
		certs = append(certs, *cert.CertObject())
		for _, dnsName := range cert.DNSNames {
			this.nameMapping[dnsName] = cert.CertObject()
		}
	}

	// min version
	this.convertMinVersion()

	// cipher suite categories
	this.initCipherSuites()

	// hsts
	if this.HSTS != nil {
		err := this.HSTS.Init()
		if err != nil {
			return err
		}
	}

	// tls config
	this.tlsConfig = &tls.Config{}
	cipherSuites := this.TLSCipherSuites()
	if !this.CipherSuitesIsOn || len(cipherSuites) == 0 {
		cipherSuites = nil
	}

	nextProto := []string{}
	if this.HTTP2Enabled {
		nextProto = []string{http2.NextProtoTLS}
	}
	this.tlsConfig = &tls.Config{
		Certificates:   certs,
		MinVersion:     this.TLSMinVersion(),
		CipherSuites:   cipherSuites,
		GetCertificate: nil,
		ClientAuth:     GoSSLClientAuthType(this.ClientAuthType),
		ClientCAs:      this.CAPool(),
		NextProtos:     nextProto,
	}

	return nil
}

// TLSMinVersion 取得最小版本
func (this *SSLPolicy) TLSMinVersion() uint16 {
	return this.minVersion
}

// TLSCipherSuites 套件
func (this *SSLPolicy) TLSCipherSuites() []uint16 {
	return this.cipherSuites
}

// MatchDomain 校验是否匹配某个域名
func (this *SSLPolicy) MatchDomain(domain string) (cert *tls.Certificate, ok bool) {
	cert, ok = this.nameMapping[domain]
	if ok {
		return
	}

	for name, cert := range this.nameMapping {
		if configutils.MatchDomain(name, domain) {
			return cert, true
		}
	}
	return nil, false
}

// FirstCert 取得第一个证书
func (this *SSLPolicy) FirstCert() *tls.Certificate {
	for _, cert := range this.Certs {
		return cert.CertObject()
	}
	return nil
}

// CAPool CA证书Pool，用于TLS对客户端进行认证
func (this *SSLPolicy) CAPool() *x509.CertPool {
	return this.clientCAPool
}

func (this *SSLPolicy) TLSConfig() *tls.Config {
	return this.tlsConfig
}
