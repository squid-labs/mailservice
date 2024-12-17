package eventapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Unmarshal returns Event from RawEvent.
func Unmarshal(raw RawEvent) (*Event, error) {
	// Convert RawEvent (assumed to be map[string]interface{}) to JSON
	rawData, err := json.Marshal(raw)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal raw event: %s", err)
	}

	// Decode JSON into Event
	var event Event
	if err := json.Unmarshal(rawData, &event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal raw event: %s", err)
	}

	return &event, nil
}

// FixAndValidate validates the event and fixes default values where needed.
func (event *Event) FixAndValidate(language string) (*Record, error) {
	// Convert event.Record (assumed to be map[string]interface{}) to JSON
	recordData, err := json.Marshal(event.Record)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal record: %s", err)
	}

	// Decode JSON into Record
	var record Record
	if err := json.Unmarshal(recordData, &record); err != nil {
		return nil, fmt.Errorf("failed to unmarshal record: %s", err)
	}

	// Validate record fields
	if record.User == nil {
		return nil, errors.New("event: record.user is nil")
	}

	if record.User.UID == "" {
		return nil, errors.New("event: record.user.uid is empty")
	}

	if record.User.Email == "" {
		return nil, errors.New("event: record.user.email is empty")
	}

	// Set default language if missing
	if record.Language == "" {
		record.Language = language
	}

	return &record, nil
}
