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

// Collection represents a reference to a ModSecurity/Coraza collection in a rule.
// Collections are named variable containers (transient or persistent) that store
// data such as request arguments, headers, IP tracking, session state, etc.
// They are accessed via %{COLLECTION:KEY} syntax or actions like initcol, setvar,
// setsid, setuid, setrsc.
//
// ModSecurity: https://github.com/owasp-modsecurity/ModSecurity/wiki/Reference-Manual-%28v3.x%29#collections
// Coraza:     https://coraza.io/docs/seclang/variables
type Collection struct {
	// Name is the collection to target (e.g. IP, TX, ARGS, SESSION).
	Name CollectionName `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// Arguments are key=value pairs or selectors passed to the collection
	// (e.g. initcol:IP=%{REMOTE_ADDR} or setvar:TX.score=+5).
	Arguments []string `json:"arguments,omitempty" yaml:"arguments,omitempty" mapstructure:"arguments,omitempty"`

	// Excludes are keys to exclude from the collection (e.g. in ARGS or REQUEST_HEADERS).
	Excluded []string `json:"excludes,omitempty" yaml:"excludes,omitempty" mapstructure:"excludes,omitempty"`

	// Count, when true, treats the collection as a count operator (&COLLECTION).
	Count bool `json:"count,omitempty" yaml:"count,omitempty" mapstructure:"count,omitempty"`
}

// CollectionName specifies a ModSecurity/Coraza collection (variable set).
//
// Collections are used everywhere in rules: variable access, setvar, initcol,
// expirevar, geoLookup, etc. Only IP/SESSION/USER/GLOBAL/RESOURCE are persistent
// across requests; the rest are transient (per-transaction).
//
// ModSecurity: https://github.com/owasp-modsecurity/ModSecurity/wiki/Reference-Manual-%28v3.x%29#collections
// Coraza:     https://coraza.io/docs/seclang/variables
//
// Allowed values:
//
//   - ARGS                  : all request arguments (GET + POST)
//   - ARGS_GET / ARGS_POST  : query string or body parameters only
//   - ARGS_NAMES / ARGS_GET_NAMES / ARGS_POST_NAMES : parameter names only
//   - REQUEST_HEADERS / REQUEST_COOKIES : headers & cookies
//   - RESPONSE_HEADERS      : response headers
//   - TX                    : transient per-transaction data (scores, captures)
//   - IP / SESSION / USER / GLOBAL / RESOURCE : persistent collections
//   - GEO                   : geolocation data (after @geoLookup)
//   - FILES / MULTIPART_*   : uploaded files and multipart parts
//   - MATCHED_VARS / MATCHED_VARS_NAMES : last operator match
//   - RULE / PERF_RULES / ENV / XML : special/engine collections
//
// +kubebuilder:validation:Required
// +kubebuilder:validation:Enum=UNKNOWN_COLLECTION;ARGS;ARGS_GET;ARGS_GET_NAMES;ARGS_NAMES;ARGS_POST_NAMES;ARGS_POST;ENV;FILES;GEO;GLOBAL;IP;MATCHED_VARS_NAMES;MATCHED_VARS;MULTIPART_PART_HEADERS;PERF_RULES;REQUEST_COOKIES_NAMES;REQUEST_COOKIES;REQUEST_HEADERS_NAMES;REQUEST_HEADERS;RESOURCE;RESPONSE_HEADERS_NAMES;RESPONSE_HEADERS;RULE;SESSION;TX;USER;XML
type CollectionName string

const (

	// ARGS contains all request arguments (query string + POST body, including
	// form-urlencoded, JSON, and multipart when parsed).
	//
	// Most commonly used collection for input validation.
	//
	// ModSecurity: https://github.com/owasp-modsecurity/ModSecurity/wiki/Reference-Manual-%28v3.x%29#collections
	// Coraza:     https://coraza.io/docs/seclang/variables/#args
	ARGS CollectionName = "ARGS"

	// ARGS_GET contains only query-string (GET) parameters.
	//
	// ModSecurity: https://github.com/owasp-modsecurity/ModSecurity/wiki/Reference-Manual-%28v3.x%29#collections
	// Coraza:     https://coraza.io/docs/seclang/variables/#args_get
	ARGS_GET CollectionName = "ARGS_GET"

	// ARGS_GET_NAMES contains only the names of query-string parameters.
	ARGS_GET_NAMES CollectionName = "ARGS_GET_NAMES"

	// ARGS_NAMES contains the names of all request parameters (GET + POST).
	//
	// Commonly used for positive security model (whitelisting parameter names).
	ARGS_NAMES CollectionName = "ARGS_NAMES"

	// ARGS_POST_NAMES contains the names of POST body parameters only.
	ARGS_POST_NAMES CollectionName = "ARGS_POST_NAMES"

	// ARGS_POST contains only parameters from the request body (POST).
	ARGS_POST CollectionName = "ARGS_POST"

	// ENV contains environment variables set by ModSecurity/Coraza (via setenv)
	// or inherited from the web server.
	ENV CollectionName = "ENV"

	// FILES contains original client-supplied filenames of uploaded files.
	//
	// Only populated when SecRequestBodyAccess is On and files are present.
	FILES CollectionName = "FILES"

	// GEO is populated by the @geoLookup operator and contains country, city,
	// region, latitude, etc.
	GEO CollectionName = "GEO"

	// GLOBAL is a persistent collection shared across all transactions.
	//
	// Initialized with initcol without a key.
	GLOBAL CollectionName = "GLOBAL"

	// IP is a persistent collection keyed by client IP address.
	//
	// The most commonly used persistent collection (rate limiting, brute-force
	// protection, etc.). Initialized with initcol:IP=%{REMOTE_ADDR}.
	IP CollectionName = "IP"

	// MATCHED_VARS_NAMES contains the names of all variables that matched
	// in the most recent operator check.
	MATCHED_VARS_NAMES CollectionName = "MATCHED_VARS_NAMES"

	// MATCHED_VARS contains the values of all variables that matched in the
	// most recent operator check.
	MATCHED_VARS CollectionName = "MATCHED_VARS"

	// MULTIPART_PART_HEADERS contains headers from each multipart part.
	//
	// Available since ModSecurity v3.0.8.
	MULTIPART_PART_HEADERS CollectionName = "MULTIPART_PART_HEADERS"

	// PERF_RULES contains performance metrics for rule evaluation.
	PERF_RULES CollectionName = "PERF_RULES"

	// REQUEST_COOKIES_NAMES contains the names of request cookies.
	REQUEST_COOKIES_NAMES CollectionName = "REQUEST_COOKIES_NAMES"

	// REQUEST_COOKIES contains all request cookies.
	REQUEST_COOKIES CollectionName = "REQUEST_COOKIES"

	// REQUEST_HEADERS_NAMES contains the names of request headers.
	REQUEST_HEADERS_NAMES CollectionName = "REQUEST_HEADERS_NAMES"

	// REQUEST_HEADERS contains all request headers.
	REQUEST_HEADERS CollectionName = "REQUEST_HEADERS"

	// RESOURCE is a persistent collection keyed by a resource identifier
	// (set via setrsc).
	RESOURCE CollectionName = "RESOURCE"

	// RESPONSE_HEADERS_NAMES contains the names of response headers.
	RESPONSE_HEADERS_NAMES CollectionName = "RESPONSE_HEADERS_NAMES"

	// RESPONSE_HEADERS contains all response headers.
	RESPONSE_HEADERS CollectionName = "RESPONSE_HEADERS"

	// RULE is a special collection that gives access to the current rule's
	// metadata (id, rev, severity, msg, etc.).
	RULE CollectionName = "RULE"

	// SESSION is a persistent collection keyed by session ID (setsid).
	//
	// Requires a prior setsid action.
	SESSION CollectionName = "SESSION"

	// TX is the transient per-transaction collection.
	//
	// Used for anomaly scoring, capture groups (TX.0, TX.1, …), temporary
	// variables, etc. The most heavily used collection in rules.
	TX CollectionName = "TX"

	// USER is a persistent collection keyed by user identifier (setuid).
	USER CollectionName = "USER"

	// XML is the special collection for XPath expressions against parsed XML
	// request/response bodies.
	//
	// Requires ctl:requestBodyProcessor=XML.
	XML CollectionName = "XML"

	UNKNOWN_COLLECTION CollectionName = "UNKNOWN_COLLECTION"
)

// +kubebuilder:object:generate=false
// goverter:converter
// goverter:output:file ./convert/zz_generated_collections.go
// goverter:enum:unknown UNKNOWN_COLLECTION
type CollectionMapper interface {
	Convert(source CollectionName) types.CollectionName
}

// +kubebuilder:object:generate=false
// goverter:converter
// goverter:output:file ./convert/zz_generated_collections.go
// goverter:enum:unknown UNKNOWN_COLLECTION
type CollectionReverseMapper interface {
	Convert(source types.CollectionName) CollectionName
}
