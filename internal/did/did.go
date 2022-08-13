package did

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type DID struct {
	Scheme           string
	Method           string
	MethodSpecificID string
}

// String returns a string representation of the DID
func (d *DID) String() string {
	return fmt.Sprintf("%s:%s:%s", d.Scheme, d.Method, d.MethodSpecificID)
}

// Parse parses a did string to DID struct
func Parse(did string) (*DID, error) {
	const idChar = `a-zA-Z0-9.-_`
	const methodChar = `a-z0-9`

	regex := fmt.Sprintf(`^did:[a-z]+:(:+|[:%s]+)*[%%:%s]+[^:]$`, methodChar, idChar)

	result, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex=%s %w", regex, err)
	}

	if !result.MatchString(did) {
		return nil, fmt.Errorf("invalid did: %s. ", did)
	}

	parts := strings.SplitN(did, ":", 3)
	return &DID{
		Scheme:           "did",
		Method:           parts[1],
		MethodSpecificID: parts[2],
	}, nil
}

// https://www.w3.org/TR/2022/REC-did-core-20220719/#did-url-syntax
// did-url = did path-abempty [ "?" query ] [ "#" fragment ]
type DIDURL struct {
	DID
	Path     string
	Queries  map[string][]string
	Fragment string
}

// Parses a string into DIDURL
func ParseDIDURL(didURL string) (*DIDURL, error) {
	position := strings.Index(didURL, "?/#")
	did := didURL
	pathQueryFragment := ""
	if position != -1 {
		did = didURL[:position]
		pathQueryFragment = did[position:]
	}

	redDID, err := Parse(did)
	if err != nil {
		return nil, err
	}

	if pathQueryFragment == "" {
		return &DIDURL{
			DID:     *redDID,
			Queries: map[string][]string{},
		}, nil
	}

	hasPath := pathQueryFragment[0] == '/'
	if !hasPath {
		pathQueryFragment = "/" + pathQueryFragment
	}
	urlParts, err := url.Parse(pathQueryFragment)
	if err != nil {
		return nil, fmt.Errorf("failed to parse path, query, and fragment components of DID URL: %w", err)
	}

	result := &DIDURL{
		DID:      *redDID,
		Queries:  urlParts.Query(),
		Fragment: urlParts.Fragment,
	}
	if hasPath {
		result.Path = urlParts.Path
	}
	return result, nil
}

type Context interface {
}

// todo
type Document struct {
	Context              Context
	ID                   string
	AlsoKnownAs          string
	Controller           DID
	VerificationMethod   VerificationMethod
	Authentication       VerificationRelationship
	AssertionMethod      VerificationRelationship
	KeyAgreement         VerificationRelationship
	CapabilityInvocation VerificationRelationship
	CapabilityDelegation VerificationRelationship
	Service              []Service
}

type DocumentResolution struct {
	Context     Context
	DIDDocument Document
}

// MethodMetadata method metadata.
type MethodMetadata struct {
	// UpdateCommitment is update commitment key.
	UpdateCommitment string `json:"updateCommitment,omitempty"`
	// RecoveryCommitment is recovery commitment key.
	RecoveryCommitment string `json:"recoveryCommitment,omitempty"`
	// Published is published key.
	Published bool `json:"published,omitempty"`
	// AnchorOrigin is anchor origin.
	AnchorOrigin string `json:"anchorOrigin,omitempty"`
	// UnpublishedOperations unpublished operations
	UnpublishedOperations []*ProtocolOperation `json:"unpublishedOperations,omitempty"`
	// PublishedOperations published operations
	PublishedOperations []*ProtocolOperation `json:"publishedOperations,omitempty"`
}

// ProtocolOperation info.
type ProtocolOperation struct {
	// Operation is operation request.
	Operation string `json:"operation,omitempty"`
	// ProtocolVersion is protocol version.
	ProtocolVersion int `json:"protocolVersion,omitempty"`
	// TransactionNumber is transaction number.
	TransactionNumber int `json:"transactionNumber,omitempty"`
	// TransactionTime is transaction time.
	TransactionTime int64 `json:"transactionTime,omitempty"`
	// Type is type of operation.
	Type string `json:"type,omitempty"`
	// AnchorOrigin is anchor origin.
	AnchorOrigin string `json:"anchorOrigin,omitempty"`
	// CanonicalReference is canonical reference
	CanonicalReference string `json:"canonicalReference,omitempty"`
	// EquivalentReferences is equivalent references
	EquivalentReferences []string `json:"equivalentReferences,omitempty"`
}

// DocumentMetadata document metadata.
type DocumentMetadata struct {
	// VersionID is version ID key.
	VersionID string `json:"versionId,omitempty"`
	// Deactivated is deactivated flag key.
	Deactivated bool `json:"deactivated,omitempty"`
	// CanonicalID is canonical ID key.
	CanonicalID string `json:"canonicalId,omitempty"`
	// EquivalentID is equivalent ID array.
	EquivalentID []string `json:"equivalentId,omitempty"`
	// Method is used for method metadata within did document metadata.
	Method *MethodMetadata `json:"method,omitempty"`
}

// processingMeta include info how to process the doc.
type processingMeta struct {
	baseURI string
}

// https://www.w3.org/TR/2022/REC-did-core-20220719/#verification-method-properties
type VerificationMethod struct {
	ID                 string
	Controller         string
	Type               string
	Value              []byte
	PublicKeyJwk       map[string]interface{}
	PublicKeyMultibase string
}

// Service DID doc service.
type Service struct {
	ID              string      `json:"id"`
	Type            string      `json:"type"`
	ServiceEndpoint interface{} `json:"serviceEndpoint"`
}

// VerificationRelationship defines a verification relationship between DID subject and a verification method.
type VerificationRelationship int

const (
	// VerificationRelationshipGeneral is a special case of verification relationship: when a verification method
	// defined in Verification is not used by any Verification.
	VerificationRelationshipGeneral VerificationRelationship = iota

	// Authentication defines verification relationship.
	Authentication

	// AssertionMethod defines verification relationship.
	AssertionMethod

	// CapabilityDelegation defines verification relationship.
	CapabilityDelegation

	// CapabilityInvocation defines verification relationship.
	CapabilityInvocation

	// KeyAgreement defines verification relationship.
	KeyAgreement
)

// Verification authentication verification.
type Verification struct {
	VerificationMethod VerificationMethod
	Relationship       VerificationRelationship
	Embedded           bool
}

type rawDoc struct {
	Context              Context                  `json:"@context,omitempty"`
	ID                   string                   `json:"id,omitempty"`
	AlsoKnownAs          []interface{}            `json:"alsoKnownAs,omitempty"`
	VerificationMethod   []map[string]interface{} `json:"verificationMethod,omitempty"`
	PublicKey            []map[string]interface{} `json:"publicKey,omitempty"`
	Service              []map[string]interface{} `json:"service,omitempty"`
	Authentication       []interface{}            `json:"authentication,omitempty"`
	AssertionMethod      []interface{}            `json:"assertionMethod,omitempty"`
	CapabilityDelegation []interface{}            `json:"capabilityDelegation,omitempty"`
	CapabilityInvocation []interface{}            `json:"capabilityInvocation,omitempty"`
	KeyAgreement         []interface{}            `json:"keyAgreement,omitempty"`
	Created              *time.Time               `json:"created,omitempty"`
	Updated              *time.Time               `json:"updated,omitempty"`
	Proof                []interface{}            `json:"proof,omitempty"`
}

// Proof is cryptographic proof of the integrity of the DID Document.
type Proof struct {
	Type         string
	Created      *time.Time
	Creator      string
	ProofValue   []byte
	Domain       string
	Nonce        []byte
	ProofPurpose string
	relativeURL  bool
}

// didKeyResolver implements public key resolution for DID public keys.
type didKeyResolver struct {
	PubKeys []VerificationMethod
}

// DocOption provides options to build DID Doc.
type DocOption func(opts *Document)
