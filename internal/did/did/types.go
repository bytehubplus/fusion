package did

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/xeipuuv/gojsonschema"
)

type KeyType string

// JsonWebKey2020 is a VerificationMethod type.
// https://w3c-ccg.github.io/lds-jws2020/
const JsonWebKey2020 = KeyType("JsonWebKey2020")

// ED25519VerificationKey2018 is the Ed25519VerificationKey2018 verification key type as specified here:
// https://w3c-ccg.github.io/lds-ed25519-2018/
const ED25519VerificationKey2018 = KeyType("Ed25519VerificationKey2018")

// ECDSASECP256K1VerificationKey2019 is the EcdsaSecp256k1VerificationKey2019 verification key type as specified here:
// https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/
const ECDSASECP256K1VerificationKey2019 = KeyType("EcdsaSecp256k1VerificationKey2019")

// RSAVerificationKey2018 is the RsaVerificationKey2018 verification key type as specified here:
// https://w3c-ccg.github.io/lds-rsa2018/
const RSAVerificationKey2018 = KeyType("RsaVerificationKey2018")

type ProofType string

// JsonWebSignature2020 is a Proof type.
// https://w3c-ccg.github.io/lds-jws2020
const JsonWebSignature2020 = ProofType("JsonWebSignature2020")

const (
	// ContextV1 of the DID document is the current V1 context name.
	ContextV1 = "https://www.w3.org/ns/did/v1"
	// ContextV1Old of the DID document representing the old/legacy V1 context name.
	ContextV1Old        = "https://w3id.org/did/v1"
	contextV011         = "https://w3id.org/did/v0.11"
	contextV12019       = "https://www.w3.org/2019/did/v1"
	jsonldType          = "type"
	jsonldID            = "id"
	jsonldPublicKey     = "publicKey"
	jsonldServicePoint  = "serviceEndpoint"
	jsonldRecipientKeys = "recipientKeys"
	jsonldRoutingKeys   = "routingKeys"
	jsonldPriority      = "priority"
	jsonldController    = "controller"
	jsonldOwner         = "owner"

	jsonldCreator        = "creator"
	jsonldCreated        = "created"
	jsonldProofValue     = "proofValue"
	jsonldSignatureValue = "signatureValue"
	jsonldDomain         = "domain"
	jsonldNonce          = "nonce"
	jsonldProofPurpose   = "proofPurpose"

	// various public key encodings.
	jsonldPublicKeyBase58    = "publicKeyBase58"
	jsonldPublicKeyMultibase = "publicKeyMultibase"
	jsonldPublicKeyHex       = "publicKeyHex"
	jsonldPublicKeyPem       = "publicKeyPem"
	jsonldPublicKeyjwk       = "publicKeyJwk"
)

var (
	schemaLoaderV1     = gojsonschema.NewStringLoader(schemaV1)     //nolint:gochecknoglobals
	schemaLoaderV011   = gojsonschema.NewStringLoader(schemaV011)   //nolint:gochecknoglobals
	schemaLoaderV12019 = gojsonschema.NewStringLoader(schemaV12019) //nolint:gochecknoglobals
)

// URI is a wrapper around url.URL to add json marshalling
type URI struct {
	url.URL
}

// MarshalText implements encoding.TextMarshaler
func (v URI) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v URI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func (v *URI) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	parsedUrl, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("could not parse URI: %w", err)
	}
	v.URL = *parsedUrl
	return nil
}

// ParseURI parses a raw URI. If it can't be parsed, an error is returned.
func ParseURI(input string) (*URI, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, err
	}

	return &URI{URL: *u}, nil
}

func MustParseURI(input string) URI {
	u, err := url.Parse(input)
	if err != nil {
		panic(err)
	}
	return URI{URL: *u}
}

func (v URI) String() string {
	return v.URL.String()
}
