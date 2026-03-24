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

import "github.com/coreruleset/crslang/types"

// Operator defines how to compare a variable/collection value against a pattern or value.
// The Negate field inverts the match (equivalent to ! in SecLang).
type Operator struct {
	Negate bool         `json:"negate,omitempty" yaml:"negate,omitempty"`
	Name   OperatorType `json:"name" yaml:"name"`
	Value  string       `json:"value,omitempty" yaml:"value,omitempty"`
}

// OperatorType is the type of comparison operator used in SecLang conditions.
// See https://coraza.io/docs/seclang/operators for full reference.
type OperatorType string

const (
	UnknownOperator      OperatorType = "unknownOperator"
	BeginsWith           OperatorType = "beginsWith"
	Contains             OperatorType = "contains"
	ContainsWord         OperatorType = "containsWord"
	DetectSQLi           OperatorType = "detectSQLi"
	DetectXSS            OperatorType = "detectXSS"
	EndsWith             OperatorType = "endsWith"
	Eq                   OperatorType = "eq"
	FuzzyHash            OperatorType = "fuzzyHash"
	Ge                   OperatorType = "ge"
	GeoLookup            OperatorType = "geoLookup"
	GsbLookup            OperatorType = "gsbLookup"
	Gt                   OperatorType = "gt"
	InspectFile          OperatorType = "inspectFile"
	IpMatchF             OperatorType = "ipMatchF"
	IpMatchFromFile      OperatorType = "ipMatchFromFile"
	IpMatch              OperatorType = "ipMatch"
	Le                   OperatorType = "le"
	Lt                   OperatorType = "lt"
	Pmf                  OperatorType = "pmf"
	PmFromFile           OperatorType = "pmFromFile"
	Pm                   OperatorType = "pm"
	Rbl                  OperatorType = "rbl"
	Rsub                 OperatorType = "rsub"
	Rx                   OperatorType = "rx"
	RxGlobal             OperatorType = "rxGlobal"
	StrEq                OperatorType = "strEq"
	StrMatch             OperatorType = "strMatch"
	UnconditionalMatch   OperatorType = "unconditionalMatch"
	ValidateByteRange    OperatorType = "validateByteRange"
	ValidateDTD          OperatorType = "validateDTD"
	ValidateHash         OperatorType = "validateHash"
	ValidateSchema       OperatorType = "validateSchema"
	ValidateUrlEncoding  OperatorType = "validateUrlEncoding"
	ValidateUtf8Encoding OperatorType = "validateUtf8Encoding"
	VerifyCC             OperatorType = "verifyCC"
	VerifyCPF            OperatorType = "verifyCPF"
	VerifySSN            OperatorType = "verifySSN"
	VerifySVNR           OperatorType = "verifySVNR"
	Within               OperatorType = "within"
)

// +kubebuilder:object:generate=false
// goverter:converter
// goverter:output:file ./convert/zz_generated_operator.go
// goverter:enum:unknown Eq
type OperatorMapper interface {
	Convert(source OperatorType) types.OperatorType
}

// +kubebuilder:object:generate=false
// goverter:converter
// goverter:output:file ./convert/zz_generated_operator.go
// goverter:enum:unknown @ignore
type OperatorReverseMapper interface {
	Convert(source types.OperatorType) OperatorType
}
