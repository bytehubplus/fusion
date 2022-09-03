package did

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type DID struct {
	Scheme           string `json:"scheme"`
	Method           string `json:"method"`
	MethodSpecificID string `json:"methodSpecificID"`
}

func (did *DID) MarshalJSON() ([]byte, error) {
	return json.Marshal(did.String())
}

func (did *DID) UnmarshalJSON(data []byte) error {
	d, err := Parse(string(data))
	if err != nil {
		return err
	}
	did.Method = d.Method
	did.Scheme = d.Scheme
	did.MethodSpecificID = d.MethodSpecificID
	return nil
}

// String returns a string representation of the DID
func (d *DID) String() string {
	if d.Scheme != "" && d.Method != "" && d.MethodSpecificID != "" {
		return ""
	}
	return fmt.Sprintf("%s:%s:%s", d.Scheme, d.Method, d.MethodSpecificID)
}

// Parse parses a did string to DID struct
func Parse(did string) (*DID, error) {
	// const idChar = `a-zA-Z0-9.-_`
	// const methodChar = `a-z0-9`
	//
	// regex := fmt.Sprintf(`^did:[a-z][%s]+:(:+|[:%s]+)*[%%:%s]+[^:]$`, methodChar, methodChar, idChar)

	const idchar = `a-zA-Z0-9-_\.`
	regex := fmt.Sprintf(`^did:[a-z0-9]+:(:+|[:%s]+)*[%%:%s]+[^:]$`, idchar, idchar)

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
	DID      DID                 `json:"did,omitempty"`
	Path     string              `json:"path,omitempty"`
	Queries  map[string][]string `json:"queries,omitempty"`
	Fragment string              `json:"fragment,omitempty"`
}

func (du *DIDURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(du.String())
}

func (du *DIDURL) UnmarshalJSON(data []byte) error {
	didUrl, err := ParseDIDURL(string(data))
	if err != nil {
		return err
	}
	du.DID = didUrl.DID
	du.Path = didUrl.Path
	du.Queries = didUrl.Queries
	du.Fragment = didUrl.Fragment
	return nil
}

func (du *DIDURL) Relative() bool {
	return du.DID.String() == "" || du.DID.String() == "::"
}

// Parses a string into DIDURL
func ParseDIDURL(didURL string) (*DIDURL, error) {
	didStop := strings.IndexAny(didURL, "/?#")
	did := didURL
	pathQueryFragment := ""
	if didStop != -1 {
		did = didURL[:didStop]
		pathQueryFragment = did[didStop:]
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

func (du *DIDURL) String() string {
	result := du.DID.String()
	if du.Path != "" {
		result = fmt.Sprintf("%s/%s", result, du.Path)
	}

	if len(du.Queries) > 0 {
		result = fmt.Sprintf("%s?", result)
		for k, v := range du.Queries {
			result = fmt.Sprintf("%s%s=%s", result, k, v)
		}
	}

	if du.Fragment != "" {
		result = fmt.Sprintf("%s#%s", result, du.Fragment)
	}

	return result
}
