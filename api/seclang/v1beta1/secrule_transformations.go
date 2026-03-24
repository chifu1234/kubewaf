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

// Transformation represents a data transformation function applied to variables
// before the operator is evaluated. Transformations can be chained.
type Transformation string

const (
	UnknownTransformation Transformation = "unknownTransformation"
	Base64Decode          Transformation = "base64Decode"
	Base64DecodeExt       Transformation = "base64DecodeExt"
	Base64Encode          Transformation = "base64Encode"
	CmdLine               Transformation = "cmdLine"
	CompressWhitespace    Transformation = "compressWhitespace"
	EscapeSeqDecode       Transformation = "escapeSeqDecode"
	CssDecode             Transformation = "cssDecode"
	HexEncode             Transformation = "hexEncode"
	HexDecode             Transformation = "hexDecode"
	HtmlEntityDecode      Transformation = "htmlEntityDecode"
	JsDecode              Transformation = "jsDecode"
	Length                Transformation = "length"
	Lowercase             Transformation = "lowercase"
	Md5                   Transformation = "md5"
	None                  Transformation = "none"
	NormalisePath         Transformation = "normalisePath"
	NormalisePathWin      Transformation = "normalisePathWin"
	ParityEven7bit        Transformation = "parityEven7bit"
	ParityOdd7bit         Transformation = "parityOdd7bit"
	ParityZero7bit        Transformation = "parityZero7bit"
	RemoveComments        Transformation = "removeComments"
	RemoveCommentsChar    Transformation = "removeCommentsChar"
	RemoveNulls           Transformation = "removeNulls"
	RemoveWhitespace      Transformation = "removeWhitespace"
	ReplaceComments       Transformation = "replaceComments"
	ReplaceNulls          Transformation = "replaceNulls"
	Sha1                  Transformation = "sha1"
	SqlHexDecode          Transformation = "sqlHexDecode"
	Trim                  Transformation = "trim"
	TrimLeft              Transformation = "trimLeft"
	TrimRight             Transformation = "trimRight"
	Uppercase             Transformation = "uppercase"
	UrlEncode             Transformation = "urlEncode"
	UrlDecode             Transformation = "urlDecode"
	UrlDecodeUni          Transformation = "urlDecodeUni"
	Utf8toUnicode         Transformation = "utf8toUnicode"
)

// +kubebuilder:object:generate=false
// goverter:converter
// goverter:output:file ./convert/zz_generated_transformations.go
// goverter:enum:unknown @ignore
type TransformationMapper interface {
	Convert(source Transformation) types.Transformation
}

// +kubebuilder:object:generate=false
// goverter:converter
// goverter:output:file ./convert/zz_generated_transformations.go
// goverter:enum:unknown @ignore
type TransformationMapperCSR interface {
	Convert(source types.Transformation) Transformation
}
