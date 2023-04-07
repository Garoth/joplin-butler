package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Garoth/joplin-butler/config"
	"github.com/Garoth/joplin-butler/types"
	"github.com/Garoth/joplin-butler/utils"
)

var (
	AUTH_TOKEN = ""
	API_TOKEN  = ""
)

func Auth() error {
	config := &config.Config{}
	_ = config.Load()
	if config.APIToken != "" && config.AuthToken != "" {
		registerConfig(config)
		out, err := utils.GetPath("notes")
		if err != nil {
			log.Println("WARN: auth token seemed to not work, retrying:", err, out)
		} else {
			log.Println("Loaded API Key:", config.APIToken)
			return nil
		}
	}

	resp, err := utils.PostPath("auth", "")
	if err != nil {
		return err
	}
	var authResp types.Auth
	if err := json.Unmarshal([]byte(resp), &authResp); err != nil {
		return err
	}
	if len(authResp.AuthToken) <= 1 {
		return fmt.Errorf("Empty auth token from API")
	}

	log.Println("Please switch to the Joplin app and confirm API access")
	for {
		time.Sleep(time.Millisecond * 250)
		check, err := checkAuth(authResp.AuthToken)
		if err != nil {
			return err
		}

		switch check.Status {
		case types.AuthCheckWaiting:
			continue
		case types.AuthCheckRejected:
			return fmt.Errorf("Auth was rejected by user, giving up")
		case types.AuthCheckAccepted:
			log.Println("Thank you for confirming, using token:", authResp.AuthToken)
			config.AuthToken = authResp.AuthToken
			config.APIToken = check.Token
			registerConfig(config)
			if err := config.Save(); err != nil {
				log.Println("WARN:", err)
			}
			return nil
		}
	}

	return nil
}

func registerConfig(config *config.Config) {
	AUTH_TOKEN = config.AuthToken
	utils.AUTH_TOKEN = AUTH_TOKEN
	API_TOKEN = config.APIToken
	utils.API_TOKEN = API_TOKEN
}

func checkAuth(authToken string) (*types.AuthCheck, error) {
	resp, err := utils.GetPath("auth/check?auth_token=" + authToken)
	if err != nil {
		return nil, err
	}
	var authStatusResp types.AuthCheck
	if err := json.Unmarshal([]byte(resp), &authStatusResp); err != nil {
		return nil, err
	}
	return &authStatusResp, nil
}
