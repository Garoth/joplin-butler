package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const (
	SERVER_PATH  = "http://localhost"
	PORT_LOWEST  = 41184
	PORT_HIGHEST = 41194
)

var (
	FOUND_PORT = 0
	AUTH_TOKEN = ""
	API_TOKEN  = ""
)

func init() {
	for port := PORT_LOWEST; port <= PORT_HIGHEST; port++ {
		resp, err := getPortPath(port, "ping")
		if err != nil {
			continue
		}
		if resp != "JoplinClipperServer" {
			log.Printf("WARN: ignoring server at %s:%d because "+
				"it responded with '%s' instead of 'JoplinClipperServer'",
				SERVER_PATH, port, resp)
			continue
		}
		FOUND_PORT = port
		break
	}

	if FOUND_PORT == 0 {
		errMsg := fmt.Sprintf("Could not find Joplin data API on %s ports %d-%d",
			SERVER_PATH, PORT_LOWEST, PORT_HIGHEST)
		panic(errMsg)
	}
}

func RegisterAuthToken(token string) {
	AUTH_TOKEN = token
}

func GetPath(path string) (string, error) {
	return getPortPath(FOUND_PORT, path)
}

func DeletePath(path string) (string, error) {
	fullPath := appendAPITokenToPath(fmt.Sprintf("%s:%d/%s", SERVER_PATH, FOUND_PORT, path))
	req, err := http.NewRequest("DELETE", fullPath, nil)
	if err != nil {
		return "", err
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func FileFormPath(path, title, filepath string) (string, error) {
	// Open the file to be uploaded
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a new buffer to store the request body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a new form field for the file
	filePart, err := writer.CreateFormFile("data", file.Name())
	if err != nil {
		return "", err
	}

	// Copy the file contents to the form field
	if _, err := io.Copy(filePart, file); err != nil {
		return "", err
	}

	// Add the JSON payload as a form field
	props := map[string]string{
		"title": title,
	}
	propsJSON, err := json.Marshal(props)
	if err != nil {
		return "", err
	}
	err = writer.WriteField("props", string(propsJSON))
	if err != nil {
		return "", err
	}
	writer.Close()

	req, err := http.NewRequest("POST",
		appendAPITokenToPath(fmt.Sprintf("%s:%d/%s", SERVER_PATH, FOUND_PORT, path)),
		body)

	if err != nil {
		return "", err
	}

	// Set the Content-Type header to "multipart/form-data"
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var respBody bytes.Buffer
	_, err = io.Copy(&respBody, res.Body)
	if err != nil {
		return "", err
	}

	return respBody.String(), nil
}

func PostPath(path, query string) (string, error) {
	data := strings.NewReader(query)
	fullPath := appendAPITokenToPath(fmt.Sprintf("%s:%d/%s", SERVER_PATH, FOUND_PORT, path))
	resp, err := http.Post(fullPath,
		"application/json", data)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getPortPath(port int, path string) (string, error) {
	fullPath := appendAPITokenToPath(fmt.Sprintf("%s:%d/%s", SERVER_PATH, port, path))
	resp, err := http.Get(fullPath)

	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func appendAPITokenToPath(path string) string {
	if API_TOKEN == "" {
		return path
	}
	return AppendQueryStringToPath(path, "token", API_TOKEN)
}

func AppendQueryStringToPath(path, query string, value any) string {
	startChar := "?"
	if strings.Contains(path, "?") {
		startChar = "&"
	}
	return fmt.Sprintf("%s%s%s=%v", path, startChar, query, value)
}

func AppendQueryStringsToPath(path string, queries map[string]any) string {
	for query, value := range queries {
		path = AppendQueryStringToPath(path, query, value)
	}
	return path
}
