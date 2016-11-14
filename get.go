/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   data.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/14 21:54:27 by hdezier           #+#    #+#             */
/*   Updated: 2016/10/15 17:17:31 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package apicpa

import (
	"encoding/json"
	"errors"
	"strconv"
)

/* GET
/api/v1/ID_COLLECTION/relationships/donnees
Toutes les données de la collection
Accès refusé si un jeton de connexion FranceConnect est requis pour cette collection

Paramètres d'entrée

    access_token

Retour
tableau_de_donnees == false

{
  "data": [{
    "id": ID_DONNEE,
    "type": "donnees",
    "attributes": {
      CHAMPS_DEFINIS_PAR_LE_SERVICE
    }
  }]
}

tableau_de_donnees == true

{
  "data": [{
    "id": ID_DONNEE,
    "type": "donnees",
    "attributes": [{
      CHAMPS_DEFINIS_PAR_LE_SERVICE
    }]
  }]
}

Status

    200 - Données trouvées
    400 - Paramètres d'entrée incorrect
    401 - Accès refusé
*/
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

/*
GET /api/v1/ID_COLLECTION/relationships/donnees/ID_DONNEE
Lire une donnée particulière

Paramètres d'entrée

access_token
fc_token (si jeton de connexion FranceConnect requis)
Retour

tableau_de_donnees == false
{
  "data": {
    "id": ID_DONNEE,
    "type": "donnees",
    "attributes": {
      CHAMPS_DEFINIS_PAR_LE_SERVICE
    }
  }
}
OU
tableau_de_donnees == true
{
  "data": {
    "id": ID_DONNEE,
    "type": "donnees",
    "attributes": [{
      CHAMPS_DEFINIS_PAR_LE_SERVICE
    }]
  }
}
Status

200 - Donnée trouvée
400 - Paramètres d'entrée incorrect
404 - Donnée introuvable
*/

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
