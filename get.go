package apicpa

import (
	"encoding/json"
	"errors"
	"strconv"
)

func GetCollectionDataID(collectionID string, accessToken string) (collectionDataID []string, err error) {
	collectionDataID = []string{}
	npage := 1
	total_pages := 0

	for ok := true; ok; ok = (npage <= total_pages) {
		collectionDataInfoUrl := Conf["CPA_API_URI"] + Conf["CPA_COLLECTION_URL"] + collectionID + "/" + Conf["CPA_COLLECTION_DATA_URL"] + Conf["POST_ACCESS_TOKEN"] + accessToken
		collectionDataInfoUrl += "&page[size]=200&page[number]=" + strconv.Itoa(npage)
		data, err := get(collectionDataInfoUrl)
		if err != nil {
			return nil, err
		}
		var content JSONContent
		err = json.Unmarshal(data, &content)
		if err != nil {
			return nil, errors.New("Service response is unidentified")
		}
		for _, value := range content.Data {
			collectionDataID = append(collectionDataID, value.Id)
		}
		npage++
		total_pages = content.Meta.Total_pages
	}
	return
}

func GetData(dataID string, collectionID string, accessToken string) (result JSONContentSingleData, err error) {
	dataUrl := Conf["CPA_API_URI"] + Conf["CPA_COLLECTION_URL"] + collectionID + "/" + Conf["CPA_COLLECTION_DATA_URL"] + dataID + "/" + Conf["POST_ACCESS_TOKEN"] + accessToken
	data, err := get(dataUrl)
	if err != nil {
		return result, errors.New("Data doesn't exist")
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("Service response is unidentified")
	}
	return
}

func GetDataWithQuery(query QueryModel, collectionID string, accessToken string) (result JSONContent, err error) {
	dataUrl := Conf["CPA_API_URI"] + Conf["CPA_COLLECTION_URL"] + collectionID + "/" + Conf["CPA_COLLECTION_DATA_URL"] + Conf["POST_ACCESS_TOKEN"] + accessToken
	for key, value := range query.Filter {
		dataUrl += "&filter[data." + key + "]=" + value
	}
	for key, value := range query.Page {
		dataUrl += "&page[" + key + "]=" + value
	}
	data, err := get(dataUrl)
	if err != nil || data == nil {
		return result, errors.New("Data doesn't exist:" + err.Error())
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("Parsing error")
	}
	return
}
