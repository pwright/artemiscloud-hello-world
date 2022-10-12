/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v2alpha3

import (
	v2alpha3 "github.com/artemiscloud/activemq-artemis-operator/api/v2alpha3"
	rest "k8s.io/client-go/rest"
)

type BrokerV2alpha3Interface interface {
	RESTClient() rest.Interface
	ActiveMQArtemisesGetter
	ActiveMQArtemisAddressesGetter
}

// BrokerV2alpha3Client is used to interact with features provided by the broker.amq.io group.
type BrokerV2alpha3Client struct {
	restClient rest.Interface
}

func (c *BrokerV2alpha3Client) ActiveMQArtemises(namespace string) ActiveMQArtemisInterface {
	return newActiveMQArtemises(c, namespace)
}

func (c *BrokerV2alpha3Client) ActiveMQArtemisAddresses(namespace string) ActiveMQArtemisAddressInterface {
	return newActiveMQArtemisAddresses(c, namespace)
}

// NewForConfig creates a new BrokerV2alpha3Client for the given config.
func NewForConfig(c *rest.Config) (*BrokerV2alpha3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BrokerV2alpha3Client{client}, nil
}

// NewForConfigOrDie creates a new BrokerV2alpha3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BrokerV2alpha3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BrokerV2alpha3Client for the given RESTClient.
func New(c rest.Interface) *BrokerV2alpha3Client {
	return &BrokerV2alpha3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v2alpha3.GroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BrokerV2alpha3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
