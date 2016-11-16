 package apicpa

 import (
 	"bytes"
 	"encoding/csv"
 	"encoding/json"
 	"errors"
 	"fmt"
 	"github.com/zemirco/uid"
 	"net/http"
 	//"net/url"
 	"os"
    "reflect"
    "strconv"
 )

 func setFromCsv(record []string, model reflect.Value) (res interface{}, err error) {
     v := reflect.Indirect(model)
     for i := 0; i < v.NumField() && i < len(record); i++ {
         switch v.Field(i).Kind() {
         case reflect.String:
             v.Field(i).SetString(record[i])
         case reflect.Int, reflect.Int64:
             val, err := strconv.Atoi(record[i])
             if err != nil {
                 return nil, err
             }
             v.Field(i).SetInt(int64(val))
         default:
             fmt.Println("Unrecognized type")
             continue
         }
     }
     return v.Interface(), nil
 }

 func postCollectionData(postEntry CPAPostModel, collectionID string) (err error) {
 	postUrl := Conf["CPA_API_URI"] + Conf["CPA_COLLECTION_URL"] + collectionID + "/" + Conf["CPA_COLLECTION_DATA_URL"]
 	b, err := json.Marshal(postEntry)
 	resp, err := http.Post(postUrl, "application/json;charset=utf-8", bytes.NewReader(b))
 	if err != nil {
 		return err
 	}
 	if resp.StatusCode == 400 {
 		fmt.Println("Post to collections failed:" + getAPIErrors(resp).Error())
 		return errors.New("Failed to post")
 	}
 	fmt.Println("Post request DONE")
 	return
 }



 func PostCollectionFromCSV(collectionID string, accessToken string, filename string, model interface{}, separator rune) (err error) {
 	file, err := os.Open(filename)
 	if err != nil {
 		return err
 	}
 	defer file.Close()

 	reader := csv.NewReader(file)
 	reader.Comma = separator
 	lines, err := reader.ReadAll()
 	if err != nil {
 		fmt.Println(err)
 		return err
 	}
 	postEntry := CPAPostModel{Data: JSONData{Attributes: model}}
 	postEntry.Access_token = accessToken
 	postEntry.Data.Type = "donnees"
 	for _, line := range lines {
 		postEntry.Data.Id = uid.New(10)
        intPtr := reflect.New(reflect.TypeOf(model))
        toSet, err := setFromCsv(line, intPtr)
        if err != nil {
            return err
        }
        var tbl []interface{} = make([]interface{}, 1)
        tbl[0] = toSet
        postEntry.Data.Attributes = tbl
 		if err != nil {
 			break
 		}
 		err = postCollectionData(postEntry, collectionID)
 		if err != nil {
 			break
 		}
 	}
 	return
 }
