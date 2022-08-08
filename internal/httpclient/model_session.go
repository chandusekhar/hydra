/*
Ory Hydra API

Documentation for all of Ory Hydra's APIs.

API version:
Contact: hi@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Session struct for Session
type Session struct {
	AllowedTopLevelClaims []string               `json:"allowed_top_level_claims,omitempty"`
	ClientId              *string                `json:"client_id,omitempty"`
	ConsentChallenge      *string                `json:"consent_challenge,omitempty"`
	ExcludeNotBeforeClaim *bool                  `json:"exclude_not_before_claim,omitempty"`
	ExpiresAt             *map[string]time.Time  `json:"expires_at,omitempty"`
	Extra                 map[string]interface{} `json:"extra,omitempty"`
	Headers               *Headers               `json:"headers,omitempty"`
	IdTokenClaims         *IDTokenClaims         `json:"id_token_claims,omitempty"`
	Kid                   *string                `json:"kid,omitempty"`
	Subject               *string                `json:"subject,omitempty"`
	Username              *string                `json:"username,omitempty"`
}

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession() *Session {
	this := Session{}
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetAllowedTopLevelClaims returns the AllowedTopLevelClaims field value if set, zero value otherwise.
func (o *Session) GetAllowedTopLevelClaims() []string {
	if o == nil || o.AllowedTopLevelClaims == nil {
		var ret []string
		return ret
	}
	return o.AllowedTopLevelClaims
}

// GetAllowedTopLevelClaimsOk returns a tuple with the AllowedTopLevelClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetAllowedTopLevelClaimsOk() ([]string, bool) {
	if o == nil || o.AllowedTopLevelClaims == nil {
		return nil, false
	}
	return o.AllowedTopLevelClaims, true
}

// HasAllowedTopLevelClaims returns a boolean if a field has been set.
func (o *Session) HasAllowedTopLevelClaims() bool {
	if o != nil && o.AllowedTopLevelClaims != nil {
		return true
	}

	return false
}

// SetAllowedTopLevelClaims gets a reference to the given []string and assigns it to the AllowedTopLevelClaims field.
func (o *Session) SetAllowedTopLevelClaims(v []string) {
	o.AllowedTopLevelClaims = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Session) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Session) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Session) SetClientId(v string) {
	o.ClientId = &v
}

// GetConsentChallenge returns the ConsentChallenge field value if set, zero value otherwise.
func (o *Session) GetConsentChallenge() string {
	if o == nil || o.ConsentChallenge == nil {
		var ret string
		return ret
	}
	return *o.ConsentChallenge
}

// GetConsentChallengeOk returns a tuple with the ConsentChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetConsentChallengeOk() (*string, bool) {
	if o == nil || o.ConsentChallenge == nil {
		return nil, false
	}
	return o.ConsentChallenge, true
}

// HasConsentChallenge returns a boolean if a field has been set.
func (o *Session) HasConsentChallenge() bool {
	if o != nil && o.ConsentChallenge != nil {
		return true
	}

	return false
}

// SetConsentChallenge gets a reference to the given string and assigns it to the ConsentChallenge field.
func (o *Session) SetConsentChallenge(v string) {
	o.ConsentChallenge = &v
}

// GetExcludeNotBeforeClaim returns the ExcludeNotBeforeClaim field value if set, zero value otherwise.
func (o *Session) GetExcludeNotBeforeClaim() bool {
	if o == nil || o.ExcludeNotBeforeClaim == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeNotBeforeClaim
}

// GetExcludeNotBeforeClaimOk returns a tuple with the ExcludeNotBeforeClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetExcludeNotBeforeClaimOk() (*bool, bool) {
	if o == nil || o.ExcludeNotBeforeClaim == nil {
		return nil, false
	}
	return o.ExcludeNotBeforeClaim, true
}

// HasExcludeNotBeforeClaim returns a boolean if a field has been set.
func (o *Session) HasExcludeNotBeforeClaim() bool {
	if o != nil && o.ExcludeNotBeforeClaim != nil {
		return true
	}

	return false
}

// SetExcludeNotBeforeClaim gets a reference to the given bool and assigns it to the ExcludeNotBeforeClaim field.
func (o *Session) SetExcludeNotBeforeClaim(v bool) {
	o.ExcludeNotBeforeClaim = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Session) GetExpiresAt() map[string]time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret map[string]time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetExpiresAtOk() (*map[string]time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Session) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given map[string]time.Time and assigns it to the ExpiresAt field.
func (o *Session) SetExpiresAt(v map[string]time.Time) {
	o.ExpiresAt = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Session) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetExtraOk() (map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Session) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *Session) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *Session) GetHeaders() Headers {
	if o == nil || o.Headers == nil {
		var ret Headers
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetHeadersOk() (*Headers, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *Session) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given Headers and assigns it to the Headers field.
func (o *Session) SetHeaders(v Headers) {
	o.Headers = &v
}

// GetIdTokenClaims returns the IdTokenClaims field value if set, zero value otherwise.
func (o *Session) GetIdTokenClaims() IDTokenClaims {
	if o == nil || o.IdTokenClaims == nil {
		var ret IDTokenClaims
		return ret
	}
	return *o.IdTokenClaims
}

// GetIdTokenClaimsOk returns a tuple with the IdTokenClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetIdTokenClaimsOk() (*IDTokenClaims, bool) {
	if o == nil || o.IdTokenClaims == nil {
		return nil, false
	}
	return o.IdTokenClaims, true
}

// HasIdTokenClaims returns a boolean if a field has been set.
func (o *Session) HasIdTokenClaims() bool {
	if o != nil && o.IdTokenClaims != nil {
		return true
	}

	return false
}

// SetIdTokenClaims gets a reference to the given IDTokenClaims and assigns it to the IdTokenClaims field.
func (o *Session) SetIdTokenClaims(v IDTokenClaims) {
	o.IdTokenClaims = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *Session) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *Session) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *Session) SetKid(v string) {
	o.Kid = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Session) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Session) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Session) SetSubject(v string) {
	o.Subject = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Session) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Session) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Session) SetUsername(v string) {
	o.Username = &v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedTopLevelClaims != nil {
		toSerialize["allowed_top_level_claims"] = o.AllowedTopLevelClaims
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ConsentChallenge != nil {
		toSerialize["consent_challenge"] = o.ConsentChallenge
	}
	if o.ExcludeNotBeforeClaim != nil {
		toSerialize["exclude_not_before_claim"] = o.ExcludeNotBeforeClaim
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.IdTokenClaims != nil {
		toSerialize["id_token_claims"] = o.IdTokenClaims
	}
	if o.Kid != nil {
		toSerialize["kid"] = o.Kid
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}