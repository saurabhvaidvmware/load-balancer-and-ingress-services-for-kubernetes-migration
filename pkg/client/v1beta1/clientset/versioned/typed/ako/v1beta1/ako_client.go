/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/apis/ako/v1beta1"
	"github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/client/v1beta1/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type AkoV1beta1Interface interface {
	RESTClient() rest.Interface
	AviInfraSettingsGetter
	HTTPRulesGetter
	HostRulesGetter
}

// AkoV1beta1Client is used to interact with features provided by the ako.vmware.com group.
type AkoV1beta1Client struct {
	restClient rest.Interface
}

func (c *AkoV1beta1Client) AviInfraSettings() AviInfraSettingInterface {
	return newAviInfraSettings(c)
}

func (c *AkoV1beta1Client) HTTPRules(namespace string) HTTPRuleInterface {
	return newHTTPRules(c, namespace)
}

func (c *AkoV1beta1Client) HostRules(namespace string) HostRuleInterface {
	return newHostRules(c, namespace)
}

// NewForConfig creates a new AkoV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*AkoV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &AkoV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new AkoV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AkoV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AkoV1beta1Client for the given RESTClient.
func New(c rest.Interface) *AkoV1beta1Client {
	return &AkoV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AkoV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
