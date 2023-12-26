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

// checks if the CreateNotificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNotificationRequest{}

// CreateNotificationRequest struct for CreateNotificationRequest
type CreateNotificationRequest struct {
	Data CreateNotification `json:"data"`
}

type _CreateNotificationRequest CreateNotificationRequest

// NewCreateNotificationRequest instantiates a new CreateNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNotificationRequest(data CreateNotification) *CreateNotificationRequest {
	this := CreateNotificationRequest{}
	this.Data = data
	return &this
}

// NewCreateNotificationRequestWithDefaults instantiates a new CreateNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNotificationRequestWithDefaults() *CreateNotificationRequest {
	this := CreateNotificationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *CreateNotificationRequest) GetData() CreateNotification {
	if o == nil {
		var ret CreateNotification
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetDataOk() (*CreateNotification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateNotificationRequest) SetData(v CreateNotification) {
	o.Data = v
}

func (o CreateNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateNotificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varCreateNotificationRequest := _CreateNotificationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNotificationRequest)

	if err != nil {
		return err
	}

	*o = CreateNotificationRequest(varCreateNotificationRequest)

	return err
}

type NullableCreateNotificationRequest struct {
	value *CreateNotificationRequest
	isSet bool
}

func (v NullableCreateNotificationRequest) Get() *CreateNotificationRequest {
	return v.value
}

func (v *NullableCreateNotificationRequest) Set(val *CreateNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNotificationRequest(val *CreateNotificationRequest) *NullableCreateNotificationRequest {
	return &NullableCreateNotificationRequest{value: val, isSet: true}
}

func (v NullableCreateNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}