package apicpa

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"bytes"
)

var Conf = map[string]string{}

func SetCredentialsFromEnv() {
	Conf = map[string]string{
        "CPA_API_URI"                : "CPA_API_URI",
    	"CPA_AUTH_URL"               : "CPA_AUTH_URL",
    	"CPA_SERVICE_URL"            : "CPA_SERVICE_URL",
    	"CPA_COLLECTION_URL"         : "CPA_COLLECTION_URL",
    	"CPA_COLLECTION_DATA_URL"    : "CPA_COLLECTION_DATA_URL",
    	"CPA_COLLECTION_SERVICE_URL" : "CPA_COLLECTION_SERVICE_URL",
    	"POST_ACCESS_TOKEN"          : "POST_ACCESS_TOKEN",
    	"ID_PUBLIC_SERVICE"          : "ID_PUBLIC_SERVICE",
    	"ID_PRIVATE_SERVICE"         : "ID_PRIVATE_SERVICE",
	}
    for _, value := range Conf {
        Conf[value] = os.Getenv(value)
    }
}

func getAuthRequest() (result *http.Request, err error) {
	authUrl := Conf["CPA_API_URI"] + Conf["CPA_AUTH_URL"]
	data := url.Values{}
	data.Set("client_id", Conf["ID_PUBLIC_SERVICE"])
	data.Add("client_secret", Conf["ID_PRIVATE_SERVICE"])

	if result, err = http.NewRequest("POST", authUrl, bytes.NewBufferString(data.Encode())); err == nil {
		result.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		result.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	}
	return
}

func authService() (result *http.Response, err error) {

	req, err := getAuthRequest()
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}

func Authenticate() (result string, err error) {

	resp, err := authService()
	if err != nil {
		return "", err
	} else if resp.StatusCode != 200 {
		// return "", getAPIErrors(resp)
		return "", nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var receiver JSONContentSingleData
	err = json.Unmarshal(body, &receiver)
	if err != nil {
		return "", errors.New("Service response is unidentified")
	}
	var token CPAToken
	err = mapstructure.Decode(receiver.Data.Attributes, &token)
	if err != nil {
		return "", errors.New("Service response is unidentified")
	}
	return token.Access_token, nil
}
