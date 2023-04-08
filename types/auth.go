package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Auth struct {
	AuthToken string `json:"auth_token"`
}

type AuthCheckStatus int

const (
	AuthCheckWaiting = iota
	AuthCheckAccepted
	AuthCheckRejected
)

type AuthCheck struct {
	Status AuthCheckStatus
	Token  string
}

func (me *AuthCheck) UnmarshalJSON(b []byte) error {
	tmpType := struct {
		Status string
		Token  string
	}{}
	err := json.Unmarshal(b, &tmpType)
	if err != nil {
		return err
	}
	str := strings.ToLower(tmpType.Status)
	me.Token = tmpType.Token

	switch str {
	case "waiting":
		me.Status = AuthCheckWaiting
	case "accepted":
		me.Status = AuthCheckAccepted
	case "rejected":
		me.Status = AuthCheckRejected
	default:
		return fmt.Errorf("Unknown AuthCheckStatus '%s'", str)
	}

	return nil
}
