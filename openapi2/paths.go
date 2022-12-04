package openapi2

import (
	"encoding/json"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type Paths struct {
	om *orderedmap.OrderedMap[string, *PathItem]
}

// MarshalJSON returns the JSON encoding of Paths.
func (paths *Paths) MarshalJSON() ([]byte, error) {
	if paths == nil || paths.om == nil {
		return []byte("{}"), nil
	}
	return paths.om.MarshalJSON()
}

// UnmarshalJSON sets Paths to a copy of data.
func (paths *Paths) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &paths.om)
}

func (paths *Paths) Value(key string) *PathItem {
	// if paths == nil || paths.om == nil {
	// 	return nil
	// }
	return paths.om.Value(key)
}

func (paths *Paths) Set(key string, value *PathItem) {
	// if paths != nil || paths.om != nil {
	_, _ = paths.om.Set(key, value)
	// }
}

func (paths *Paths) Len() int {
	if paths == nil || paths.om == nil {
		return 0
	}
	return paths.om.Len()
}

func (paths *Paths) Iter() *pathsKV {
	if paths == nil || paths.om == nil {
		return nil
	}
	return (*pathsKV)(paths.om.Oldest())
}

type pathsKV orderedmap.Pair[string, *PathItem] //FIXME: pub?

func (pair *pathsKV) Next() *pathsKV {
	ompair := (*orderedmap.Pair[string, *PathItem])(pair)
	return (*pathsKV)(ompair.Next())
}

// NewPathsWithCapacity builds a paths object of the given capacity.
func NewPathsWithCapacity(cap int) *Paths {
	return &Paths{om: orderedmap.New[string, *PathItem](cap)}
}

// NewPaths builds a paths object with path items in insertion order.
func NewPaths(opts ...NewPathsOption) *Paths {
	paths := NewPathsWithCapacity(len(opts))
	for _, opt := range opts {
		opt(paths)
	}
	return paths
}

// NewPathsOption describes options to NewPaths func
type NewPathsOption func(*Paths)

// WithPath adds paths as an option to NewPaths
func WithPath(path string, pathItem *PathItem) NewPathsOption {
	return func(paths *Paths) { paths.Set(path, pathItem) }
}
