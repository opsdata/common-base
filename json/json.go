package json

import "encoding/json"

type RawMessage = json.RawMessage

var (
	// Marshal is exported by pkg/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by pkg/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by pkg/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by pkg/json package.
	NewEncoder = json.NewEncoder
)
