package apicpa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func doRequest(reqUrl string, reqType string) (result *http.Response, err error) {
	req, err := http.NewRequest(reqType, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}

func getAPIErrors(resp *http.Response) (err error) {
	if resp == nil {
		return errors.New("Service error response is unidentified")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Service error response is unidentified")
	}
	defer resp.Body.Close()
	var f interface{}
	json.Unmarshal(body, &f)
	errorsResp, ok := f.(map[string]interface{})["errors"].([]interface{})
	if ok == false {
		return errors.New("Service error response is unidentified")
	}
	errMessage := "Response read with " + strconv.Itoa(len(errorsResp)) + " errors:\n"
	for _, val := range errorsResp {
		errMessage += val.(string) + "_\n"
	}
	return errors.New(errMessage)
}

func get(reqUrl string) (data []byte, err error) {
	resp, err := doRequest(reqUrl, "GET")
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 && resp.StatusCode != 206 {
		fmt.Println("URL:", reqUrl)
		return nil, getAPIErrors(resp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Service response is unidentified")
	}
	defer resp.Body.Close()
	return body, nil
}
