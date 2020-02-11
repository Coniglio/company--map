package model

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// CompanyMap 企業マップ情報
type CompanyMap struct {
	Name   string `json:"name"`
	Latlng struct {
		Lat  float32 `json:"lat"`
		Lang float32 `json:"lang"`
	} `json:"latlng"`
	Languages []Language `json:"language"`
	Alongs    []Along    `json:"alongs"`
}

// Company 企業情報
type Company struct {
	CompanyID    int     `db:"company_id"`
	CompanyName  string  `db:"company_name"`
	X            float32 `db:"x"`
	Y            float32 `db:"y"`
	LanguageID   int     `db:"language_id"`
	LanguageName string  `db:"language_name"`
}

// GetCompanyMaps 企業マップを検索します
func GetCompanyMaps(tx *gorp.Transaction) ([]CompanyMap, error) {

	var companyMaps []CompanyMap

	// 企業情報を検索
	companies, err := selectToCompanyMap(tx)
	if err != nil {
		return companyMaps, err
	}

	// 検索結果の言語情報をまとめる
	var languages = make(map[int][]Language)
	for _, c := range companies {
		if _, ok := languages[c.CompanyID]; ok {
			languages[c.CompanyID] = append(languages[c.CompanyID], Language{ID: c.LanguageID, Name: c.LanguageName})
		} else {
			languages[c.CompanyID] = []Language{Language{ID: c.LanguageID, Name: c.LanguageName}}
		}
	}

	// クライアントへの返却用に整形

	for _, c := range companies {
		cmap := CompanyMap{
			Name: c.CompanyName,
			Latlng: struct {
				Lat  float32 `json:"lat"`
				Lang float32 `json:"lang"`
			}{
				c.X,
				c.Y,
			},
			Languages: languages[c.CompanyID],
			Alongs:    []Along{},
		}
		companyMaps = append(companyMaps, cmap)
	}

	return companyMaps, nil
}

func selectToCompanyMap(tx *gorp.Transaction) ([]Company, error) {
	var companies []Company
	_, err := tx.Select(&companies, `
		select
		  com.id company_id,
		  com.company_name,
		  X(loc.latlng) x,
		  Y(loc.latlng) y,
		  lan.id language_id,
		  lan.language_name
		from
		  companies com
		  inner join locations loc on com.id = loc.companies_id
		  inner join technologies tec on com.id = tec.company_id
		  inner join languages lan on tec.language_id = lan.id
		order by
		  com.id
		`)
	if err != nil {
		return companies, err
	}

	return companies, nil
}
