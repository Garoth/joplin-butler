package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Error struct {
	ErrorRaw   string `json:"error"`
	ErrorLines []string
}

func NewError(jsonStr string) (*Error, error) {
	me := &Error{}
	if err := json.Unmarshal([]byte(jsonStr), me); err != nil {
		return nil, err
	}
	if strings.TrimSpace(me.ErrorRaw) == "" {
		return nil, fmt.Errorf("not an error json")
	}
	parts := strings.Split(me.ErrorRaw, "\n")
	me.ErrorLines = []string{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			me.ErrorLines = append(me.ErrorLines, part)
		}
	}
	return me, nil
}

func (me *Error) Error() string {
	return strings.Trim(
		strings.ReplaceAll(
			me.ErrorLines[0], "Internal Server Error: ", ""), ": ")
}

// Checks both API level and JSON-based error messages, merges them together
func CheckError(json string, err error) error {
	if err != nil {
		return err
	}
	jsonErr, err := NewError(json)
	if err == nil {
		return fmt.Errorf("%s", jsonErr.Error())
	}
	return nil
}
