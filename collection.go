package apicpa

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
)

func GetCollectionID(collectionName []string, accessToken string) (collectionsID map[string]string, err error) {

	collectionInfoUrl := Conf["CPA_API_URI"] + Conf["CPA_COLLECTION_URL"] + Conf["POST_ACCESS_TOKEN"] + accessToken
	data, err := get(collectionInfoUrl)
	if err != nil {
		return nil, errors.New("Can't get data from url")
	}
	var receiver JSONContent
	err = json.Unmarshal(data, &receiver)
	if err != nil {
		return nil, errors.New("Service global response is unidentified")
	}
	collectionsID = make(map[string]string)
	for _, value := range receiver.Data {
		id := value.Id
		var collection CPACollectionModel
		err := mapstructure.Decode(value.Attributes, &collection)
		if err != nil {
			return nil, errors.New("Service detail response is unidentified")
		}
		name := collection.Nom
		if stringInSlice(name, collectionName) {
			collectionsID[name] = id
		}
	}
	return
}
