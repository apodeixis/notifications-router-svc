package data

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/apodeixis/notifications-router-svc/resources"
)

type Message resources.Message

func (m Message) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Message) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	return json.Unmarshal(source, m)
}

type TemplateMessageAttributes struct {
	Payload *json.RawMessage `json:"payload"`
	Locale  *string          `json:"locale"`
	Files   []string         `json:"files"`
}
