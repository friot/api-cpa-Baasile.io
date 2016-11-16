package apicpa

import (
)

type CPACollectionModel struct {
	Nom                        string `json:"nom"`
	Description                string `json:"description"`
	Tableau_de_donnees         bool   `json:"tableau_de_donnees"`
	Jeton_fc_lecture_ecriture  bool   `json:"jeton_fc_lecture_ecriture"`
	Jeton_fc_lecture_seulement bool   `json:"jeton_fc_lecture_seulement"`
}

type CPAPostModel struct {
	Access_token string   `json:"access_token"`
	Data         JSONData `json:"data"`
}

type CPAToken struct {
	Access_token string `json:"access_token"`
}

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
