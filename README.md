# API-CPA / BAASILE.io

Go package for connection to api cpa / Baasile.io with the help of @elojah

------------------------------------------------------------------------------------------------------

# BEFORE USE

```bash
echo "export CPA_API_URI=https://api-cpa.herokuapp.com/api/v1/ && export CPA_AUTH_URL=oauth/token/ && export CPA_SERVICE_URL=services/ && export CPA_COLLECTION_URL=collections/ && export CPA_COLLECTION_DATA_URL=relationships/donnees/ && export CPA_COLLECTION_SERVICE_URL=relationships/collections/ && export POST_ACCESS_TOKEN="?access_token=" && export ID_PUBLIC_SERVICE="public_token" && export ID_PRIVATE_SERVICE="prive_token"" >> $HOME/.zshrc && source ~/.zshrc
```

package name : apicpa

------------------------------------------------------------------------------------------------------

# USE

apicpa.SetCredentialsFromEnv() ==> To use before Auth

accessToken, err := apicpa.Authenticate() ==> Authenticate generate token from Env credential

collectionsName := []string{"name_of_collection"}

collectionsID, err := apicpa.GetCollectionID(collectionsName, accessToken)

Getting collectionId by his name, using the accessToken previously generate from Authenticate(),
mapped with the name of the collection

------------------------------------------------------------------------------------------------------

# GET


CollectionsDataId, err := apicpa.GetCollectionDataID(CollectionsID["name_of_collection"], accessToken)

Getting the Id of each data from the name of collection in type []string

oneData, err := apicpa.GetData("id_of_my_data", CollectionsID["name_of_collection"], accessToken)

Getting one Object from the api in type JSONContentSingleData

filter := make(map[string]string)
filter["Name_of_field"] = "Value_of_field"
query := apicpa.QueryModel{Filter:filter}
multyContent, err := apicpa.GetDataWithQuery(query, CollectionsID["name_of_collection"], accessToken)

Getting multiple value with a query of type field=value, for the moment this is the only query implemented
I'll work on it later

multyContent type is JSONContentData

------------------------------------------------------------------------------------------------------

# POST

err = apicpa.PostCollectionFromCSV(CollectionsID["name_of_collection"], accessToken, "./path_to/my_file.csv", CustomModel{}, separator)

v := url.Values{}
v.Set("field1", "value1")
v.Set("field2", "value2")
v.Set("field3", "1234")
id, err = apicpa.PostCollectionFromUrl(CollectionsID["name_of_collection"], accessToken, v, CustomModel{})

Allow to post data on the api directly from a csv file

CustomModel is a type struct with the same name_field and type field than the collection model on the api

It will only work for the moment for int and string field

If the field type in the model is int, it will auto change the string "1234" into a int before pushing to
the DB

separator is a rune type, corresponding with the separator in the csv (usually ',' or ';')

PostCollectionFromUrl will return the id created from the push into the DB

--------------------------------------------------------------------------------------------------------

# MODEL

type JSONContent struct {
	Meta JSONMetaPagination `json:"meta"`
	Data []JSONData         `json:"data"`
}

type JSONContentSingleData struct {
	Meta JSONMetaPagination `json:"meta"`
	Data JSONData           `json:"data"`
}

type JSONData struct {
	Id         string      `json:"id"`
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
	Links      struct {
		Self  string `json:"self"`
		First string `json:"first"`
	} `json:"links"`
	Meta struct {
		Creation     string `json:"creation"`
		Modification string `json:"modification"`
		Version      int    `json:"version"`
	}
}

type JSONMetaPagination struct {
	Total       int `json:"total"`
	Total_pages int `json:"total_pages"`
	Offset      int `json:"offset"`
	Limit       int `json:"limit"`
	Count       int `json:"count"`
}

type QueryModel struct {
	Filter map[string]string `json:"filter"`
	Page   map[string]string `json:"page"`
}

----------------------------------------------------------------------------------------------------------

# Enjoy!
