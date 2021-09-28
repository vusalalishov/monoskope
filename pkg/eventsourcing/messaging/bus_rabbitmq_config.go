// Copyright 2021 Monoskope Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package messaging

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/finleap-connect/monoskope/pkg/eventsourcing/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	DefaultExchangeName = "m8_events" // name of the monoskope exchange
	CACertPath          = "/etc/eventstore/certs/buscerts/ca.crt"
	TLSCertPath         = "/etc/eventstore/certs/buscerts/tls.crt"
	TLSKeyPath          = "/etc/eventstore/certs/buscerts/tls.key"
)

type RabbitEventBusConfig struct {
	name             string // Name of the client, required
	url              string // Connection string, required
	routingKeyPrefix string // Prefix for routing of messages
	exchangeName     string // Name of the exchange to initialize/use
	amqpConfig       *amqp.Config
}

// NewRabbitEventBusConfig creates a new RabbitEventBusConfig with defaults.
func NewRabbitEventBusConfig(name, url, routingKeyPrefix string) (*RabbitEventBusConfig, error) {
	if routingKeyPrefix == "" {
		routingKeyPrefix = "m8"
	}

	uri, err := amqp.ParseURI(url)
	if err != nil {
		return nil, err
	}

	conf := &RabbitEventBusConfig{
		name:             name,
		url:              url,
		routingKeyPrefix: routingKeyPrefix,
		exchangeName:     DefaultExchangeName,
		amqpConfig:       &amqp.Config{},
	}

	if uri.Scheme == "amqps" {
		if err := conf.configureTLS(); err != nil {
			return nil, err
		}
	}

	return conf, nil
}

// ConfigureTLS adds the configuration for TLS secured connection/auth
func (conf *RabbitEventBusConfig) configureTLS() error {
	var err error
	caCertPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(CACertPath)
	if err != nil {
		return err
	}
	caCertPool.AppendCertsFromPEM(ca)

	conf.amqpConfig.TLSClientConfig = &tls.Config{
		RootCAs:              caCertPool,
		GetClientCertificate: getClientCertificate,
	}
	conf.amqpConfig.SASL = []amqp.Authentication{&CertAuth{}}

	return nil
}

// getClientCertificate returns the loaded certificate for use by
// the TLSConfig fields getClientCertificate.
func getClientCertificate(hello *tls.CertificateRequestInfo) (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(TLSCertPath, TLSKeyPath)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

// Validate validates the configuration
func (conf *RabbitEventBusConfig) Validate() error {
	if conf.name == "" {
		return errors.ErrConfigNameRequired
	}
	if conf.url == "" {
		return errors.ErrConfigUrlRequired
	}
	return nil
}

// CertAuth for RabbitMQ-auth-mechanism-ssl.
type CertAuth struct {
}

func (me *CertAuth) Mechanism() string {
	return "EXTERNAL"
}

func (me *CertAuth) Response() string {
	return "\000*\000*"
}
