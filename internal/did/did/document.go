package did

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Document struct {
	Context []Context `json:"@context,omitempty"`
	ID      DID       `json:"id,omitempty"`
	// AlsoKnownas          VerficationRelationships `json:"alsoKnownas,omitempty"`
	// Controller           []DID                    `json:"controller,omitempty"`
	// VerificationMethod   VerificationMethods      `json:"verificationMethod,omitempty"`
	// Authentication       VerficationRelationships `json:"authentication,omitempty"`
	// AssertionMethod      VerficationRelationships `json:"assertionMethod,omitempty"`
	// KeyAgreement         VerficationRelationships `json:"keyAgreement,omitempty"`
	// CapabilityInvocation VerficationRelationships `json:"capabilityInvocation,omitempty"`
	// CapabilityDelegation VerficationRelationships `json:"capabilityDelegation,omitempty"`
	// Service              []Service                `json:"service,omitempty"`
	// Created              *time.Time               `json:"created,omitempty"`
	// Updated              *time.Time               `json:"update,omitempty"`
}

type VerificationMethod struct {
	ID           DID
	Controller   DID
	Type         string
	PublicKeyJwk map[string]interface{}
}

// A set of either Verification Method maps that conform to the rules in Verification Method properties) or strings that conform to the rules in 3.2 DID URL Syntax.
type VerificationMethods []*VerificationMethod

type VerficationRelationship struct {
	*VerificationMethod
	reference DID
}

type VerficationRelationships []VerficationRelationship

type PublicKey struct {
}

type Service struct {
	ID              URI      `json:"id"`
	Type            []string `json:"type"`
	ServiceEndpoint []URI    `json:"serviceEndpoint"`
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
