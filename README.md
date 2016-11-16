# api-cpa
go package for connection to api cpa with the help of @elojah


```bash
echo "export CPA_API_URI=https://api-cpa.herokuapp.com/api/v1/ && export CPA_AUTH_URL=oauth/token/ && export CPA_SERVICE_URL=services/ && export CPA_COLLECTION_URL=collections/ && export CPA_COLLECTION_DATA_URL=relationships/donnees/ && export CPA_COLLECTION_SERVICE_URL=relationships/collections/ && export POST_ACCESS_TOKEN="?access_token=" && export ID_PUBLIC_SERVICE="public_token" && export ID_PRIVATE_SERVICE="prive_token"" >> $HOME/.zshrc && source ~/.zshrc
```

USE

apicpa.SetCredentialsFromEnv() ==> To use before Auth

accessToken, err := apicpa.Authenticate() ==> Authenticate generate token from Env credential

------------------------------------------------------------------------------------------------------

collectionsName := []string{"name_of_collection"}

collectionsID, err := apicpa.GetCollectionID(collectionsName, accessToken)

Getting collectionId by his name, using the accessToken previously generate from Authenticate(),
mapped with the name of the collection

------------------------------------------------------------------------------------------------------

CollectionsDataId, err := apicpa.GetCollectionDataID(CollectionsID["name_of_collection"], accessToken)

Getting the data from the name of collection

------------------------------------------------------------------------------------------------------

err = apicpa.PostCollectionFromCSV(CollectionsID["name_of_collection"], accessToken, "./path_to/my_file.csv", CustomModel{}, separator)

Allow to post data on the api directly from a csv file

CustomModel is a type struct with the same name_field than the collection model on the api

separator is a rune type, corresponding with the separator in the csv (usually ',' or ';')

Enjoy!
