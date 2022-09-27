// The AGPLv3 License (AGPLv3)

// Copyright (c) 2022 ZHAO Zhenhua <zhao.zhenhua@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package did

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Document struct {
	Context              []Context
	ID                   DID
	AlsoKnownas          []string
	Controller           []DID
	VerificationMethod   []VerificationMethod
	Authentication       []Verification
	AssertionMethod      []Verification
	KeyAgreement         []Verification
	CapabilityInvocation []Verification
	CapabilityDelegation []Verification
	Service              []Service
}

type rawDocument struct {
	Context              []string             `json:"context,omitempty"`
	ID                   string               `json:"id"`
	AlsoKnownas          []string             `json:"alsoKnownas,omitempty"`
	Controller           []string             `json:"controller,omitempty"`
	VerificationMethod   []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication       []Verification       `json:"authentication,omitempty"`
	AssertionMethod      []Verification       `json:"assertionMethod,omitempty"`
	KeyAgreement         []Verification       `json:"keyAgreement,omitempty"`
	CapabilityInvocation []Verification       `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation []Verification       `json:"capabilityDelegation,omitempty"`
	Service              []Service            `json:"service,omitempty"`
}

func (doc *Document) UnmarshalJSON(data []byte) error {
	var rawDoc = &rawDocument{}
	err := json.Unmarshal(data, &rawDoc)
	if err != nil {
		return fmt.Errorf("invalid DID document format: %s", err)
	}

	for _, ctx := range rawDoc.Context {
		url, err := url.Parse(ctx)
		if err != nil {
			return fmt.Errorf("invalid context: %s", err)
		}

		doc.Context = append(doc.Context, Context{*url})
	}

	did, err := Parse(rawDoc.ID)
	if err != nil {
		return fmt.Errorf("invalid DID: %s", err)
	}

	doc.ID = *did

	for _, controller := range rawDoc.Controller {
		did, err := Parse(controller)
		if err != nil {
			return fmt.Errorf("invalid DID: %s", err)
		}

		doc.Controller = append(doc.Controller, *did)
	}

	doc.VerificationMethod = rawDoc.VerificationMethod
	doc.Authentication = rawDoc.Authentication
	doc.AssertionMethod = rawDoc.AssertionMethod
	doc.KeyAgreement = rawDoc.KeyAgreement
	doc.CapabilityInvocation = rawDoc.CapabilityInvocation
	doc.CapabilityDelegation = rawDoc.CapabilityDelegation
	doc.Service = rawDoc.Service
	return nil
}

// type rawVerificationMethod struct {
// 	ID                 string       `json:"id"`
// 	Controller         string       `json:"controller"`
// 	Type               string       `json:"type"`
// 	PublicKeyJwk       PublicKeyJwk `json:"publicKeyJwk,omitempty"`
// 	PublicKeyMultibase string       `json:"publicKeyMultibase,omitempty"`
// }

type VerificationMethod struct {
	ID                 DIDURL       `json:"id,omitempty"`
	Controller         DID          `json:"controller,omitempty"`
	Type               string       `json:"type,omitempty"`
	PublicKeyJwk       PublicKeyJwk `json:"publicKeyJwk,omitempty"`
	PublicKeyMultibase string       `json:"publicKeyMultibase,omitempty"`
}

// func (vm *VerificationMethod) UnmarshalJSON(data []byte) error {
// 	var rvm rawVerificationMethod
// 	err := json.Unmarshal(data, &rvm)
// 	if err != nil {
// 		return err
// 	}
// 	didUrl, err := ParseDIDURL(rvm.ID)
// 	if err != nil {
// 		return err
// 	}
// 	vm.ID = *didUrl
// 	did, err := Parse(rvm.Controller)
// 	if err != nil {
// 		return err
// 	}
// 	vm.Controller = *did

// 	vm.Type = rvm.Type
// 	vm.PublicKeyJwk = rvm.PublicKeyJwk
// 	vm.PublicKeyMultibase = rvm.PublicKeyMultibase
// 	return nil
// }

type PublicKeyJwk struct {
	Crv string `json:"crv,omitempty"`
	X   string `json:"x,omitempty"`
	Kty string `json:"kty,omitempty"`
	Kid string `json:"kid,omitempty"`
}

// https://www.w3.org/TR/did-core/#referring-to-verification-methods
// a verfication either contains a referring id or a verification method.
type Verification struct {
	VerificationMethod VerificationMethod
	Id                 DIDURL
	// a VerficationRelationship is ether embeded or referenced
	// Embedded has higher priority
	Embedded bool
	//https://www.w3.org/TR/did-core/#relative-did-urls
	// not implemented yet
	Relative bool
}

func (vr *Verification) UnmarshalJSON(bytes []byte) error {
	var i interface{}
	err := json.Unmarshal(bytes, &i)
	if err != nil {
		return fmt.Errorf("unmarshal VerificationRelationship failed: %s", err)
	}

	switch v := i.(type) {
	//referring verification method. it's a DID URL
	case string:
		didUrl, err := ParseDIDURL(v)
		if err != nil {
			return fmt.Errorf("invalid DID URL: %s", err)
		}
		vr.Id = *didUrl
	//embedded verfication method. it's a verification method
	case map[string]interface{}:
		vm := VerificationMethod{}

		_, found := v["id"]
		if found {
			didUrl, _ := ParseDIDURL(v["id"].(string))
			vm.ID = *didUrl
		}
		_, found = v["type"]
		if found {
			vm.Type = v["type"].(string)
		}
		_, found = v["controller"]
		if found {
			didUrl, err := ParseDIDURL(v["controller"].(string))
			if err != nil {
				return fmt.Errorf("parse DID URL failed: %s", err)
			}
			vm.Controller = didUrl.DID
		}
		_, found = v["publicKeyHex"]
		if found {
			vm.PublicKeyMultibase = v["publicKeyHex"].(string)
		}
		var jwk PublicKeyJwk
		_, found = v["publickeyjwk"]
		if found {
			err = json.Unmarshal(v["publickeyjwk"].([]byte), &jwk)
			if err != nil {
				return fmt.Errorf("unmarshal publick key jwk failed: %s", err)
			}
			vm.PublicKeyJwk = jwk
		}
		vr.VerificationMethod = vm
	default:
		return fmt.Errorf("invalid verfication")
	}
	return nil
}

type URI struct {
	url.URL
}

func (u URI) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *URI) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}

	url, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("could not parse URI: %w", err)
	}
	u.URL = *url
	return nil
}

type PublicKey struct {
}

type Service struct {
	ID              URI      `json:"id"`
	Type            []string `json:"type"`
	ServiceEndpoint []string `json:"serviceEndpoint"`
	// ServiceEndpoint interface{} `json:"serviceEndpoint"`
}

// func (s Service) MarshalJSON() ([]byte, error) {
// 	type alias Service
// 	tmp := alias(s)
// 	if data, err := json.Marshal(tmp); err != nil {
// 		return nil, err
// 	} else {
// 		return marshal.NormalizeDocument(data, marshal.Unplural("serviceEndpointKey"))
// 	}
// }

// func (s *Service) UnmarshalJSON(data []byte) error {
// 	pluralContext := marshal.Plural("@context")
// 	normalizedData, err := marshal.NormalizeDocument(data, pluralContext, marshal.PluralValueOrMap("serviceEndpointKey"))
// 	if err != nil {
// 		return err
// 	}
// 	type alias Service
// 	var result alias
// 	if err := json.Unmarshal(normalizedData, &result); err != nil {
// 		return err
// 	}
// 	*s = (Service)(result)
// 	return nil
// }

// // Unmarshal unmarshalls the service endpoint into a domain-specific type.
// func (s Service) UnmarshalServiceEndpoint(target interface{}) error {
// 	var valueToMarshal interface{}
// 	if asSlice, ok := s.ServiceEndpoint.([]interface{}); ok && len(asSlice) == 1 {
// 		valueToMarshal = asSlice[0]
// 	} else {
// 		valueToMarshal = s.ServiceEndpoint
// 	}
// 	if asJSON, err := json.Marshal(valueToMarshal); err != nil {
// 		return err
// 	} else {
// 		return json.Unmarshal(asJSON, target)
// 	}
// }

// func (s *Service) UnmarshalJSON(bytes []byte) error {
// 	var item map[string]interface{}
// 	err := json.Unmarshal(bytes, &item)
// 	if err != nil {
// 		return fmt.Errorf("unmarshal service failed: %s", err)
// 	}
// 	for field, value := range item {
// 		switch strings.ToLower(field) {
// 		case "id":
// 			// uri, _ := url.Parse(value.(string))
// 			// s.ID = URI{*uri}
// 			var url url.URL
// 			json.Unmarshal(value.([]byte), &url)
// 			// json.Unmarshal(value.([]byte), s.ID)
// 		case "type":
// 			s.Type = append(s.Type, value.(string))
// 			// json.Unmarshal(value.([]byte), s.Type)
// 		case "serviceendpoint":
// 			switch value.(type) {
// 			case string:
// 				// s.ServiceEndpoint = append(s.ServiceEndpoint, value.(string))
// 				s.ServiceEndpoint = value
// 				// case map[string]interface{}:
// 				// 	var vm VerificationMethod
// 				// 	err = json.Unmarshal([]byte(fmt.Sprintf("%v", value)), &vm)
// 				// 	if err != nil {
// 				// 		return fmt.Errorf("unmarshal verfication method failed: %s", err)
// 				// 	}
// 				// 	s.ServiceEndpoint = append(s.ServiceEndpoint, value.(string))
// 			}

// 		}
// 	}
// 	return nil
// }

type Context struct {
	url.URL
}

func (ctx Context) MarshalJSON() ([]byte, error) {
	return json.Marshal(ctx.String())
}

func (ctx *Context) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}

	parsedUrl, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("could not parse URI: %w", err)
	}
	ctx.URL = *parsedUrl
	return nil
}

func ParseDocument(bytes []byte) (*Document, error) {
	var doc Document
	if err := json.Unmarshal(bytes, &doc); err != nil {
		return nil, fmt.Errorf("JSON marshalling of document failed: %w", err)
	}

	return &doc, nil
}
