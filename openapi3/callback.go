package openapi3

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/jsonpointer"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type Callbacks struct {
	om *orderedmap.OrderedMap[string, *CallbackRef]
}

// MarshalJSON returns the JSON encoding of Callbacks.
func (c *Callbacks) MarshalJSON() ([]byte, error) {
	return c.om.MarshalJSON()
}

// UnmarshalJSON sets Callbacks to a copy of data.
func (c *Callbacks) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &c.om)
}

func (c *Callbacks) Value(key string) *CallbackRef {
	// if c == nil || c.om == nil {
	// 	return nil
	// }
	return c.om.Value(key)
}

func (c *Callbacks) Set(key string, value *CallbackRef) {
	// if c != nil || c.om != nil {
	_, _ = c.om.Set(key, value)
	// }
}

func (c *Callbacks) Len() int {
	if c == nil || c.om == nil {
		return 0
	}
	return c.om.Len()
}

func (c *Callbacks) Iter() *callbacksKV {
	if c == nil || c.om == nil {
		return nil
	}
	return (*callbacksKV)(c.om.Oldest())
}

type callbacksKV orderedmap.Pair[string, *CallbackRef] //FIXME: pub?

func (pair *callbacksKV) Next() *callbacksKV {
	ompair := (*orderedmap.Pair[string, *CallbackRef])(pair)
	return (*callbacksKV)(ompair.Next())
}

// NewCallbacksWithCapacity builds a callbacks object of the given capacity.
func NewCallbacksWithCapacity(cap int) *Callbacks {
	return &Callbacks{om: orderedmap.New[string, *CallbackRef](cap)}
}

var _ jsonpointer.JSONPointable = (*Callbacks)(nil)

// JSONLookup implements https://pkg.go.dev/github.com/go-openapi/jsonpointer#JSONPointable
func (c *Callbacks) JSONLookup(token string) (interface{}, error) {
	ref := c.Value(token)
	if ref == nil {
		return nil, fmt.Errorf("object has no field %q", token)
	}

	if ref.Ref != "" {
		return &Ref{Ref: ref.Ref}, nil
	}
	return ref.Value, nil
}

// Callback is specified by OpenAPI/Swagger standard version 3.
// See https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.3.md#callback-object
type Callback struct {
	om *orderedmap.OrderedMap[string, *PathItem]
}

// MarshalJSON returns the JSON encoding of Callback.
func (callback *Callback) MarshalJSON() ([]byte, error) {
	return callback.om.MarshalJSON()
}

// UnmarshalJSON sets Callback to a copy of data.
func (callback *Callback) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &callback.om)
}

func (callback *Callback) Value(key string) *PathItem {
	// if callback == nil || callback.om == nil {
	// 	return nil
	// }
	return callback.om.Value(key)
}

func (callback *Callback) Set(key string, value *PathItem) {
	// if callback != nil || callback.om != nil {
	_, _ = callback.om.Set(key, value)
	// }
}

func (callback *Callback) Len() int {
	if callback == nil || callback.om == nil {
		return 0
	}
	return callback.om.Len()
}

func (callback *Callback) Iter() *callbackKV {
	if callback == nil || callback.om == nil {
		return nil
	}
	return (*callbackKV)(callback.om.Oldest())
}

type callbackKV orderedmap.Pair[string, *PathItem] //FIXME: pub?

func (pair *callbackKV) Next() *callbackKV {
	ompair := (*orderedmap.Pair[string, *PathItem])(pair)
	return (*callbackKV)(ompair.Next())
}

// Validate returns an error if Callback does not comply with the OpenAPI spec.
func (callback *Callback) Validate(ctx context.Context, opts ...ValidationOption) error {
	ctx = WithValidationOptions(ctx, opts...)

	for pair := callback.Iter(); pair != nil; pair = pair.Next() {
		v := pair.Value
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}
	return nil
}
