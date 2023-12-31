/*
Apodeixis notifications-router-svc

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NotificatorServiceKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificatorServiceKey{}

// NotificatorServiceKey struct for NotificatorServiceKey
type NotificatorServiceKey struct {
	Type string `json:"type"`
}

type _NotificatorServiceKey NotificatorServiceKey

// NewNotificatorServiceKey instantiates a new NotificatorServiceKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificatorServiceKey(type_ string) *NotificatorServiceKey {
	this := NotificatorServiceKey{}
	this.Type = type_
	return &this
}

// NewNotificatorServiceKeyWithDefaults instantiates a new NotificatorServiceKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificatorServiceKeyWithDefaults() *NotificatorServiceKey {
	this := NotificatorServiceKey{}
	return &this
}

// GetType returns the Type field value
func (o *NotificatorServiceKey) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificatorServiceKey) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificatorServiceKey) SetType(v string) {
	o.Type = v
}

func (o NotificatorServiceKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificatorServiceKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *NotificatorServiceKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNotificatorServiceKey := _NotificatorServiceKey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificatorServiceKey)

	if err != nil {
		return err
	}

	*o = NotificatorServiceKey(varNotificatorServiceKey)

	return err
}

type NullableNotificatorServiceKey struct {
	value *NotificatorServiceKey
	isSet bool
}

func (v NullableNotificatorServiceKey) Get() *NotificatorServiceKey {
	return v.value
}

func (v *NullableNotificatorServiceKey) Set(val *NotificatorServiceKey) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificatorServiceKey) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificatorServiceKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificatorServiceKey(val *NotificatorServiceKey) *NullableNotificatorServiceKey {
	return &NullableNotificatorServiceKey{value: val, isSet: true}
}

func (v NullableNotificatorServiceKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificatorServiceKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
