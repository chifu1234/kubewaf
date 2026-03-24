/*
Copyright 2025 Buzz-IT GmbH.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

// Condition represents one matching condition within a SecRule (or SecAction).
// A condition consists of variables/collections to inspect, an operator to apply,
// and optional transformations.
type Condition struct {
	// Variables to evaluate in this condition.
	Variables []Variable `json:"variables,omitempty" yaml:"variables,omitempty"`

	// Collections to evaluate (specialized variable groups).
	Collections []Collection `json:"collections,omitempty" yaml:"collections,omitempty"`

	// Operator defines the comparison to perform on the variable value(s).
	Operator Operator `json:"operator,omitempty" yaml:"operator,omitempty"`

	// Transformations to apply to the variable values before operator evaluation.
	// +optional
	Transformations []Transformation `json:"transformations,omitempty" yaml:"transformations,omitempty"`

	// AlwaysMatch is used for SecAction style unconditional rules.
	AlwaysMatch bool `json:"always-match,omitempty" yaml:"always-match,omitempty"`

	// Script is the path for Lua or other script-based conditions (advanced).
	Script string `json:"script,omitempty" yaml:"script,omitempty"`
}
