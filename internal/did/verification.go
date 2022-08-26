package did

import (
	"errors"
	"time"
)

var VerificationMethodType = map[string]string{
	"VerifiableCondition2021":           "VerifiableCondition2021",
	"JsonWebKey2020":                    "JsonWebKey2020",
	"EcdsaSecp256k1VerificationKey2019": "EcdsaSecp256k1VerificationKey2019",
	"EcdsaSecp256k1VerificationKey2018": "EcdsaSecp256k1VerificationKey2018",
	"Bls12381G1Key2020":                 "Bls12381G1Key2020",
	"Bls12381G2Key2020":                 "Bls12381G2Key2020",
	"PgpVerificationKey2021":            "PgpVerificationKey2021",
	"RsaVerificationKey2018":            "RsaVerificationKey2018",
	"X25519KeyAgreementKey2019":         "X25519KeyAgreementKey2019",
	"EcdsaSecp256k1RecoveryMethod2020":  "EcdsaSecp256k1RecoveryMethod2020",
}

type DocumentMetadata struct {
	Created       *time.Time `json:"created,omitempty"`
	Updated       *time.Time `json:"updated,omitempty"`
	Deactivated   bool       `json:"deactivated,omitempty"`
	NextUpdate    *time.Time `json:"nextUpdate,omitempty"`
	VersionId     string     `json:"versionId,omitempty"`
	NextVersionID string     `json:"nextVersionID,omitempty"`
	EquivalentID  []DID      `json:"equivalentID,omitempty"`
	CanonicalID   DID        `json:"canonicalID,omitempty"`
}

var (
	ErrResolutionInvalidDId                 = errors.New("invalidDid")
	ErrResolutionInvalidDidUrl              = errors.New("invalidDidUrl")
	ErrResolutionNotFound                   = errors.New("notFound")
	ErrResolutionRepresentationNotSupported = errors.New("representationNotSupported")
	ErrResolutionInternalError              = errors.New("representationNotSupported")
	ErrResolutionInvalidPublicKey           = errors.New("invalidPublicKey")
	ErrResolutionInvalidPublicKeyLength     = errors.New("invalidPublicKeyLength")
	ErrResolutionInvalidPublicKeyType       = errors.New("invalidPublicKeyType")
	ErrResolutionUnsupportedPublicKeyType   = errors.New("unsupportedPublicKeyType")
)

// const (
// 	InvalidDID
// )

type Resolution struct {
}

type ResolutionMetadata struct {
	ContentType string `json:"contentType,omitempty"`
	Error       string `json:"error,omitempty"`
}

type Metadata struct {
	Accept string `json:"accept,omitempty"`
	Error  string `json:"error,omitempty"`
}

func (r *Resolution) ResolveRepresentation(did string, resolutionOptions Metadata) (ResolutionMetadata, []byte, DocumentMetadata) {
	return ResolutionMetadata{}, []byte(""), DocumentMetadata{}
}

func (r *Resolution) Resolve(did string, resolutionOptions Metadata) (ResolutionMetadata, Document, DocumentMetadata) {
	return ResolutionMetadata{}, Document{}, DocumentMetadata{}
}
