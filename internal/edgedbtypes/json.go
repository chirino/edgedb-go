// This source file is part of the EdgeDB open source project.
//
// Copyright EdgeDB Inc. and the EdgeDB authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edgedbtypes

import "encoding/json"

// Json represents a json value.
type Json = json.RawMessage

// NewOptionalJson is a convenience function for creating an OptionalJson
// with its value set to v.
func NewOptionalJson(v json.RawMessage) OptionalJson {
	o := OptionalJson{}
	o.Set(v)
	return o
}

// OptionalJson is an optional json. Optional types must be used for out
// parameters when a shape field is not required.
type OptionalJson struct {
	val   json.RawMessage
	isSet bool
}

// Get returns the value and a boolean indicating if the value is present.
func (o OptionalJson) Get() (json.RawMessage, bool) { return o.val, o.isSet }

// Value returns the value or the zero value if not set.
func (o OptionalJson) Value() json.RawMessage { return o.val }

// Set sets the value.
func (o *OptionalJson) Set(val json.RawMessage) {
	if val == nil {
		o.Unset()
		return
	}

	o.val = val
	o.isSet = true
}

// Unset marks the value as missing.
func (o *OptionalJson) Unset() {
	o.val = nil
	o.isSet = false
}

// MarshalJSON returns o marshaled as json.
func (o OptionalJson) MarshalJSON() (json.RawMessage, error) {
	if o.isSet {
		return json.Marshal(o.val)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON unmarshals bytes into *o.
func (o *OptionalJson) UnmarshalJSON(value json.RawMessage) error {
	if value[0] == 0x6e { // null
		o.Unset()
		return nil
	}

	if err := json.Unmarshal(value, &o.val); err != nil {
		return err
	}
	o.isSet = true

	return nil
}
