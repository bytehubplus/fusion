package did

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Document struct {
	Context              []Context                 `json:"@context,omitempty"`
	ID                   DID                       `json:"id,omitempty"`
	AlsoKnownas          []VerficationRelationship `json:"alsoKnownas,omitempty"`
	Controller           []DID                     `json:"controller,omitempty"`
	VerificationMethod   []VerificationMethod      `json:"verificationMethod,omitempty"`
	Authentication       []VerficationRelationship `json:"authentication,omitempty"`
	AssertionMethod      []VerficationRelationship `json:"assertionMethod,omitempty"`
	KeyAgreement         []VerficationRelationship `json:"keyAgreement,omitempty"`
	CapabilityInvocation []VerficationRelationship `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation []VerficationRelationship `json:"capabilityDelegation,omitempty"`
	Service              []Service                 `json:"service,omitempty"`
	Created              *time.Time                `json:"created,omitempty"`
	Updated              *time.Time                `json:"update,omitempty"`
}

type VerificationMethod struct {
	ID                 string       `json:"id,omitempty"`
	Controller         DID          `json:"controller,omitempty"`
	Type               string       `json:"type,omitempty"`
	PublicKeyJwk       PublicKeyJwk `json:"publicKeyJwk,omitempty"`
	PublicKeyMultibase string       `json:"publicKeyMultibase,omitempty"`
}

type PublicKeyJwk struct {
	Crv string `json:"crv,omitempty"`
	X   string `json:"x,omitempty"`
	Kty string `json:"kty,omitempty"`
	Kid string `json:"kid,omitempty"`
}

type VerficationRelationship struct {
	VerificationMethod VerificationMethod
	Id                 DIDURL
}

func (vr *VerficationRelationship) String() string {
	vmEmpty := VerificationMethod{}
	if vr.VerificationMethod != vmEmpty {
		return fmt.Sprint(vr.VerificationMethod)
	} else {
		return fmt.Sprint(vr.Id.String())
	}
}

func (vr *VerficationRelationship) UnmarshalJSON(bytes []byte) error {
	var i interface{}
	err := json.Unmarshal(bytes, &i)
	if err != nil {
		return fmt.Errorf("unmarshal VerificationRelationship failed: %s", err)
	}

	switch v := i.(type) {
	case string:
		didUrl, err := ParseDIDURL(v)
		if err != nil {
			return fmt.Errorf("invalid DID URL: %s", err)
		}
		vr.Id = *didUrl
	case map[string]interface{}:
		_, found := v["id"]
		if found {
			vr.VerificationMethod.ID = v["id"].(string)
		}
		_, found = v["type"]
		if found {
			vr.VerificationMethod.Type = v["type"].(string)
		}
		_, found = v["controller"]
		if found {
			didUrl, err := ParseDIDURL(v["controller"].(string))
			if err != nil {
				return fmt.Errorf("parse DID URL failed: %s", err)
			}
			vr.VerificationMethod.Controller = didUrl.DID
		}
		_, found = v["publicKeyHex"]
		if found {
			vr.VerificationMethod.PublicKeyMultibase = v["publicKeyHex"].(string)
		}
		var jwk PublicKeyJwk
		_, found = v["publickeyjwk"]
		if found {
			err = json.Unmarshal(v["publickeyjwk"].([]byte), &jwk)
			if err != nil {
				return fmt.Errorf("unmarshal publick key jwk failed: %s", err)
			}
			vr.VerificationMethod.PublicKeyJwk = jwk
		}
	}
	return nil
}

// type VerficationRelationships []VerficationRelationship

type PublicKey struct {
}

type Service struct {
	ID              URI                       `json:"id,omitempty"`
	Type            string                    `json:"type,omitempty"`
	ServiceEndpoint []VerficationRelationship `json:"serviceEndpoint,omitempty"`
}

func (s *Service) UnmarshalJSON(bytes []byte) error {
	var item map[string]interface{}
	err := json.Unmarshal(bytes, &item)
	if err != nil {
		return fmt.Errorf("unmarshal service failed: %s", err)
	}

	for field, value := range item {
		switch strings.ToLower(field) {
		case "id":
			uri, _ := url.Parse(value.(string))
			s.ID = URI{*uri}

		case "type":
			s.Type = value.(string)

		case "serviceendpoint":
			// v := value.(type)
			switch value.(type) {
			case string:
				didUrl, err := ParseDIDURL(value.(string))
				if err != nil {
					return fmt.Errorf("parse DID URL failed: %s", err)
				}
				s.ServiceEndpoint = append(s.ServiceEndpoint, VerficationRelationship{Id: *didUrl})
			case map[string]interface{}:
				var vm VerificationMethod
				err = json.Unmarshal([]byte(fmt.Sprintf("%v", value)), &vm)
				if err != nil {
					return fmt.Errorf("unmarshal verfication method failed: %s", err)
				}
				s.ServiceEndpoint = append(s.ServiceEndpoint, VerficationRelationship{VerificationMethod: vm})
			}

		}
	}

	return nil
}

type Context struct {
	url.URL
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
