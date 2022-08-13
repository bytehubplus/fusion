package did

import (
	"encoding/json"
	"fmt"
)

type Document struct {
	ID                   DID                  `json:"id,omitempty"`
	AlsoKnownas          []string             `json:"alsoKnownas,omitempty"`
	Controller           []DID                `json:"controller,omitempty"`
	VerificationMethod   []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication       string               `json:"authentication,omitempty"`
	AssertionMethod      string               `json:"assertionMethod,omitempty"`
	KeyAgreement         string               `json:"keyAgreement,omitempty"`
	CapabilityInvocation string               `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation string               `json:"capabilityDelegation,omitempty"`
	Service              string               `json:"service,omitempty"`
}

type VerificationMethod struct {
}

type PublicKey struct {
}

type Service struct {
	ID              DID
	Type            string
	ServiceEndpoint string
}

func ParseDocument(rj []byte) (*Document, error) {
	raw := &Document{}
	err := json.Unmarshal(rj, &raw)
	if err != nil {
		return nil, fmt.Errorf("JSON marshalling of  document failed: %w", err)
	} else {
		if raw == nil {
			return nil, fmt.Errorf("document payload is not provided")
		}
	}

	doc := &Document{
		ID:                   raw.ID,
		AlsoKnownas:          raw.AlsoKnownas,
		Controller:           raw.Controller,
		VerificationMethod:   raw.VerificationMethod,
		Authentication:       raw.Authentication,
		KeyAgreement:         raw.KeyAgreement,
		CapabilityInvocation: raw.CapabilityInvocation,
		CapabilityDelegation: raw.CapabilityDelegation,
	}

	return doc, nil
}
