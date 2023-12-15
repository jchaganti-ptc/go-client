/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27678-64d64396ca66
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// OAuthFlow struct for OAuthFlow
type OAuthFlow struct {
	AuthorizationUrl *string                           `json:"authorizationUrl,omitempty"`
	Extensions       map[string]map[string]interface{} `json:"extensions,omitempty"`
	RefreshUrl       *string                           `json:"refreshUrl,omitempty"`
	Scopes           *OAuthFlowScopes                  `json:"scopes,omitempty"`
	TokenUrl         *string                           `json:"tokenUrl,omitempty"`
}

// NewOAuthFlow instantiates a new OAuthFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthFlow() *OAuthFlow {
	this := OAuthFlow{}
	return &this
}

// NewOAuthFlowWithDefaults instantiates a new OAuthFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthFlowWithDefaults() *OAuthFlow {
	this := OAuthFlow{}
	return &this
}

// GetAuthorizationUrl returns the AuthorizationUrl field value if set, zero value otherwise.
func (o *OAuthFlow) GetAuthorizationUrl() string {
	if o == nil || o.AuthorizationUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationUrl
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil || o.AuthorizationUrl == nil {
		return nil, false
	}
	return o.AuthorizationUrl, true
}

// HasAuthorizationUrl returns a boolean if a field has been set.
func (o *OAuthFlow) HasAuthorizationUrl() bool {
	if o != nil && o.AuthorizationUrl != nil {
		return true
	}

	return false
}

// SetAuthorizationUrl gets a reference to the given string and assigns it to the AuthorizationUrl field.
func (o *OAuthFlow) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *OAuthFlow) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *OAuthFlow) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *OAuthFlow) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetRefreshUrl returns the RefreshUrl field value if set, zero value otherwise.
func (o *OAuthFlow) GetRefreshUrl() string {
	if o == nil || o.RefreshUrl == nil {
		var ret string
		return ret
	}
	return *o.RefreshUrl
}

// GetRefreshUrlOk returns a tuple with the RefreshUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetRefreshUrlOk() (*string, bool) {
	if o == nil || o.RefreshUrl == nil {
		return nil, false
	}
	return o.RefreshUrl, true
}

// HasRefreshUrl returns a boolean if a field has been set.
func (o *OAuthFlow) HasRefreshUrl() bool {
	if o != nil && o.RefreshUrl != nil {
		return true
	}

	return false
}

// SetRefreshUrl gets a reference to the given string and assigns it to the RefreshUrl field.
func (o *OAuthFlow) SetRefreshUrl(v string) {
	o.RefreshUrl = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAuthFlow) GetScopes() OAuthFlowScopes {
	if o == nil || o.Scopes == nil {
		var ret OAuthFlowScopes
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetScopesOk() (*OAuthFlowScopes, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuthFlow) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given OAuthFlowScopes and assigns it to the Scopes field.
func (o *OAuthFlow) SetScopes(v OAuthFlowScopes) {
	o.Scopes = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *OAuthFlow) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthFlow) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *OAuthFlow) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *OAuthFlow) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

func (o OAuthFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationUrl != nil {
		toSerialize["authorizationUrl"] = o.AuthorizationUrl
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.RefreshUrl != nil {
		toSerialize["refreshUrl"] = o.RefreshUrl
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.TokenUrl != nil {
		toSerialize["tokenUrl"] = o.TokenUrl
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthFlow struct {
	value *OAuthFlow
	isSet bool
}

func (v NullableOAuthFlow) Get() *OAuthFlow {
	return v.value
}

func (v *NullableOAuthFlow) Set(val *OAuthFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthFlow(val *OAuthFlow) *NullableOAuthFlow {
	return &NullableOAuthFlow{value: val, isSet: true}
}

func (v NullableOAuthFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
