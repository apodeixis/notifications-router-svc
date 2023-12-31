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

// checks if the NotificationAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationAllOfAttributes{}

// NotificationAllOfAttributes struct for NotificationAllOfAttributes
type NotificationAllOfAttributes struct {
	Topic     string  `json:"topic"`
	Message   Message `json:"message"`
	Channel   *string `json:"channel,omitempty"`
	CreatedAt int64   `json:"created_at"`
}

type _NotificationAllOfAttributes NotificationAllOfAttributes

// NewNotificationAllOfAttributes instantiates a new NotificationAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationAllOfAttributes(topic string, message Message, createdAt int64) *NotificationAllOfAttributes {
	this := NotificationAllOfAttributes{}
	this.Topic = topic
	this.Message = message
	this.CreatedAt = createdAt
	return &this
}

// NewNotificationAllOfAttributesWithDefaults instantiates a new NotificationAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationAllOfAttributesWithDefaults() *NotificationAllOfAttributes {
	this := NotificationAllOfAttributes{}
	return &this
}

// GetTopic returns the Topic field value
func (o *NotificationAllOfAttributes) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *NotificationAllOfAttributes) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value
func (o *NotificationAllOfAttributes) SetTopic(v string) {
	o.Topic = v
}

// GetMessage returns the Message field value
func (o *NotificationAllOfAttributes) GetMessage() Message {
	if o == nil {
		var ret Message
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotificationAllOfAttributes) GetMessageOk() (*Message, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotificationAllOfAttributes) SetMessage(v Message) {
	o.Message = v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *NotificationAllOfAttributes) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAllOfAttributes) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *NotificationAllOfAttributes) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *NotificationAllOfAttributes) SetChannel(v string) {
	o.Channel = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NotificationAllOfAttributes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NotificationAllOfAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NotificationAllOfAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

func (o NotificationAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["topic"] = o.Topic
	toSerialize["message"] = o.Message
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *NotificationAllOfAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"topic",
		"message",
		"created_at",
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

	varNotificationAllOfAttributes := _NotificationAllOfAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationAllOfAttributes)

	if err != nil {
		return err
	}

	*o = NotificationAllOfAttributes(varNotificationAllOfAttributes)

	return err
}

type NullableNotificationAllOfAttributes struct {
	value *NotificationAllOfAttributes
	isSet bool
}

func (v NullableNotificationAllOfAttributes) Get() *NotificationAllOfAttributes {
	return v.value
}

func (v *NullableNotificationAllOfAttributes) Set(val *NotificationAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationAllOfAttributes(val *NotificationAllOfAttributes) *NullableNotificationAllOfAttributes {
	return &NullableNotificationAllOfAttributes{value: val, isSet: true}
}

func (v NullableNotificationAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
