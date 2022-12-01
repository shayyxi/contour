// Copyright Project Contour Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package featuretests

// kubernetes helpers

import (
	v1 "k8s.io/api/core/v1"
	networking_v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func IngressBackend(svc *v1.Service) *networking_v1.IngressBackend {
	return &networking_v1.IngressBackend{
		Service: &networking_v1.IngressServiceBackend{
			Name: svc.Name,
			Port: networking_v1.ServiceBackendPort{
				Number: svc.Spec.Ports[0].Port,
			},
		},
	}
}

// nolint:revive
const (
	// CERTIFICATE generated by
	// openssl genrsa -out example-key.pem 2048
	// openssl req -new -x509 -days 18250 -key example-key.pem -sha256 -subj "/CN=www.example.com" -out example.pem
	CERTIFICATE = `-----BEGIN CERTIFICATE-----
MIIDFzCCAf+gAwIBAgIUZULFakfIJl0qaJXAVPCz2nzvB38wDQYJKoZIhvcNAQEL
BQAwGjEYMBYGA1UEAwwPd3d3LmV4YW1wbGUuY29tMCAXDTIyMDgxOTExMDkxNVoY
DzIwNzIwODA2MTEwOTE1WjAaMRgwFgYDVQQDDA93d3cuZXhhbXBsZS5jb20wggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDS9S4d/ea6wqiib8UeyHMptoks
w+q2DNuF75NQHLh5Z2rUnE/N8/KVhpIx81QdId1maWS0b3392hCRRFY3sDlMpRk/
1uoQgzLdk8pjw1JiqoDpvTiKZsADmVuUcCdHLNEzYtcLWBv0VyyNyE5pdrnVnMbx
w1aiQ8w2lcBCJQ8Y4DAc5oKlvBu49aUsvfFZwjL6Cr1qafQiYylqQcz7zqGBYjXc
iMzN+4fE1XQlw1iy6XmVZiHQr8Sb7EBI+g0iJapgNv7tBunzywSvAYK8N42QQOll
1sKEVf7thoNEmJTIUFo6m57Fys7LQ/B8in5JwBU+1FjqNWLJ1Gj+zIc93oc3AgMB
AAGjUzBRMB0GA1UdDgQWBBS/BZ2Uu1Y0//Um8bOqyyWz9LnPvzAfBgNVHSMEGDAW
gBS/BZ2Uu1Y0//Um8bOqyyWz9LnPvzAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3
DQEBCwUAA4IBAQCTZ6ZDi7aU7NjZdGWNLrRCBEt+FcD+mdvtRcaSp2K7m+WObnWl
rDM7V/s8ziu8ffwfwEbaBKYVLO7Mww8ke0WclBp1sq6A5AWy1sBCQYJuPCdOJNY0
fLaZObhUSQvNGw1wAXkgczrsOa/5QII356UsLiqhninXWYTMvNehab4+QW6Dldqo
EyxKgX2Ls984ZN5CDvvXfRnkeQW1/K705ReZq8qmtmCwU5wHYy0IoJGNapeX45VY
6s2n5I5CpH4L9Ua4NLgqphjC/QYK4q71GHTZD89mfTsmE+0flgFDS+wrv5SusK8u
CY2iW9j8VptZU8LVs9FrhgecEtfXbTA3MeSo
-----END CERTIFICATE-----`

	CERTIFICATE_WITH_TEXT = CERTIFICATE + "\t\r\n"

	RSA_PRIVATE_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA0vUuHf3musKoom/FHshzKbaJLMPqtgzbhe+TUBy4eWdq1JxP
zfPylYaSMfNUHSHdZmlktG99/doQkURWN7A5TKUZP9bqEIMy3ZPKY8NSYqqA6b04
imbAA5lblHAnRyzRM2LXC1gb9FcsjchOaXa51ZzG8cNWokPMNpXAQiUPGOAwHOaC
pbwbuPWlLL3xWcIy+gq9amn0ImMpakHM+86hgWI13IjMzfuHxNV0JcNYsul5lWYh
0K/Em+xASPoNIiWqYDb+7Qbp88sErwGCvDeNkEDpZdbChFX+7YaDRJiUyFBaOpue
xcrOy0PwfIp+ScAVPtRY6jViydRo/syHPd6HNwIDAQABAoIBAEf36RHGSu6v9gPk
iaUk0VULtuSUuf/9hu68es873Rtd0q5R3U/vx3SHglyUHMALi5Kipf6AgsUVnc1R
OPCqqAGj2WdUFGopuDKrdsJuIi8S6APVz/I3d45CxWFwmZXIjl4vfBmcp3zGOKbu
DQIhxOhBIgXclDOrWYHNuNdX+TyMr9cBe9bPtsWmN6Wl8V26kLzboQ3bIsM9taL+
sizVNrLL4U7oqbsKVhmI1m78+VaucQeK+XuyvuT6toOxY/Mgidcp4Mwq06aOTFug
N/u+VyXNen3Vk6/kMTxbaM9aQ05cnOrKqrgJAhiFGb+lcqvvGlwGxNqGcBABYkqf
ejhNdEECgYEA/3shOSl4GmEGZZXTARo0bDnelSpSJgmlStBeAGraq+lm/4IrTCCB
2tmJkEs6uiCEH7vkbbI/lLlhNjM2gKDovP3UwTQHGfiHlcIYYICMIpbp7+Ar9+ex
4KPXmgqqhTKC4xQYGoUE6INlpZcQ4blA4AtQbhgcC4Cp8FD/QbXp0xECgYEA02Ll
EAxzHZBNoK/5oPL0LiDUW/zPdOWCQCW8oGW5gDUvCLQ/lntegL8Qzubt7PZ4xgeg
m2ENTDcp1Zfn9s0T1V2T9Sba8gShUCvm9nLVCj4OQ3lwDufwHFuhKFpj1BVnHD5y
9yhXfyrFgvhamepEjq6LZUCrL1HxZqgOezv6JccCgYBqM1oFNArcFFcfZV+YRrdi
AdBX+4a4jyvp5KIe1ExgSB7rucWb2KuCOQmpNMyN0LR7qJR1UTKC9WjGqhVO9RSq
c228fo8xKZHbHBscCnO2cTt/3pUIcYUM167pNuPZiLzF/nVimMcIjI51fk2jN2oT
eECP82+9DFgYMONbAm7XsQKBgH/Y3ztenDz0Ks8Vv3/FkUNY3bco5vwHV0ieyj+k
ZpYRFHpKMe88fEKXzH2mk53uz8rNkCiJgTZoYqfpcQUGsYkpSLRLpL4daMcJVm4V
s523PH84sjqBsuojzQuP57K8oxkk9/ld79VctAprVLikRISbMnmxrBc5kywIVoHY
G4m/AoGBAPV0wuytURR3CbPHz28wrKRh/xnHrW+Fp3ooRfj8Pr/4zDVsqqNz2gA3
yVb6kAEpg2ON9NOWSfoUfC+THitOnRm8pKL1QL7oiq/2+s3IiK+jSevEU7TUfjio
1LwtUqv1MbKdv7TgkU0YQ99iLocWF4F4oWF6AX86/BL9y7gcbE0y
-----END RSA PRIVATE KEY-----`
)

func Secretdata(cert, key string) map[string][]byte {
	return map[string][]byte{
		v1.TLSCertKey:       []byte(cert),
		v1.TLSPrivateKeyKey: []byte(key),
	}
}

func Endpoints(ns, name string, subsets ...v1.EndpointSubset) *v1.Endpoints {
	return &v1.Endpoints{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Subsets: subsets,
	}
}

func Ports(eps ...v1.EndpointPort) []v1.EndpointPort {
	return eps
}

func Port(name string, port int32) v1.EndpointPort {
	return v1.EndpointPort{
		Name:     name,
		Port:     port,
		Protocol: "TCP",
	}
}

func Addresses(ips ...string) []v1.EndpointAddress {
	var addrs []v1.EndpointAddress
	for _, ip := range ips {
		addrs = append(addrs, v1.EndpointAddress{IP: ip})
	}
	return addrs
}
