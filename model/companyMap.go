package model

import (
	"database/sql"
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
	Languages []Language `json:"languages"`
	Alongs    []Along    `json:"alongs"`
	GenerousWelfares []GenerousWelfare `json:"generousWelfares"`
}

// Company 企業情報
type Company struct {
	CompanyID    int     `db:"company_id"`
	CompanyName  string  `db:"company_name"`
	X            float32 `db:"x"`
	Y            float32 `db:"y"`
	LanguageID   int     `db:"language_id"`
	LanguageName string  `db:"language_name"`
	AlongID      int     `db:"along_id"`
	AlongName    string  `db:"along_name"`
	GenerousWelfareID sql.NullInt64 `db:"generous_welfare_id"`
}

// GetCompanyMaps 企業マップを検索します
func GetCompanyMaps(tx *gorp.Transaction) ([]CompanyMap, error) {

	var result []CompanyMap

	// 企業情報を検索
	companies, err := selectToCompanyMap(tx)
	if err != nil {
		return result, err
	}

	// 検索結果の言語情報をまとめる
	var languages = make(map[int][]Language)
	for _, c := range companies {
		if _, ok := languages[c.CompanyID]; ok {
			isContain := false
			for _, lang := range languages[c.CompanyID] {
				if c.LanguageID == lang.ID {
					isContain = true
					break
				}
			}
			if !isContain {
				languages[c.CompanyID] = append(languages[c.CompanyID], Language{ID: c.LanguageID, Name: c.LanguageName})
			}
		} else {
			languages[c.CompanyID] = []Language{Language{ID: c.LanguageID, Name: c.LanguageName}}
		}
	}

	// 検索結果の沿線情報をまとめる
	var alongs = make(map[int][]Along)
	for _, c := range companies {
		if _, ok := alongs[c.CompanyID]; ok {
			isContain := false
			for _, along := range alongs[c.CompanyID] {
				if c.AlongID == along.ID {
					isContain = true
					break
				}
			}
			if !isContain {
				alongs[c.CompanyID] = append(alongs[c.CompanyID], Along{ID: c.AlongID, Name: c.AlongName})
			}
		} else {
			alongs[c.CompanyID] = []Along{Along{ID: c.AlongID, Name: c.AlongName}}
		}
	}

	// 検索結果の福利厚生情報をまとめる
	var generousWelfares = make(map[int][]GenerousWelfare)
	for _, c := range companies {

		var generousWelfareID int

		// 福利厚生が存在しない場合、次の企業をチェック
		if c.GenerousWelfareID.Valid {
			generousWelfareID = int(c.GenerousWelfareID.Int64);
		} else {
			continue;
		}

		if _, ok := generousWelfares[c.CompanyID]; ok {

			isContain := false
			for _, generousWelfare := range generousWelfares[c.CompanyID] {
				if generousWelfareID == generousWelfare.ID {
					isContain = true
					break
				}
			}
			if !isContain {
				generousWelfares[c.CompanyID] = append(generousWelfares[c.CompanyID], GenerousWelfare{ID: generousWelfareID})
			}
		} else {
			generousWelfares[c.CompanyID] = []GenerousWelfare{GenerousWelfare{ID: generousWelfareID}}
		}
	}

	// クライアントへの返却用に整形
	var companyMaps = make(map[int]CompanyMap)
	for _, c := range companies {
		if _, ok := companyMaps[c.CompanyID]; !ok {

			// 福利厚生が存在しない場合、空の配列を返却する
			var generousWelfare = []GenerousWelfare{}
			if generousWelfares[c.CompanyID] != nil {
				generousWelfare = generousWelfares[c.CompanyID]
			}

			companyMaps[c.CompanyID] = CompanyMap{
				Name: c.CompanyName,
				Latlng: struct {
					Lat  float32 `json:"lat"`
					Lang float32 `json:"lang"`
				}{
					c.X,
					c.Y,
				},
				Languages: languages[c.CompanyID],
				Alongs:    alongs[c.CompanyID],
				GenerousWelfares: generousWelfare,
			}
		}

	}

	// TODO map→sliceの変換はそのうち、見直す
	for key := range companyMaps {
		result = append(result, companyMaps[key])
	}

	return result, nil
}

// 企業マップ情報を検索します
func selectToCompanyMap(tx *gorp.Transaction) ([]Company, error) {
	// TODO: 一回で取得するか、分割して取得してからマージ！？ 企業名とか緯度経度とか重複して大量のレコードが取得される
	var companies []Company
	_, err := tx.Select(&companies, `
		select
		  com.id company_id,
		  com.company_name,
		  X(loc.latlng) x,
		  Y(loc.latlng) y,
		  lan.id language_id,
		  lan.language_name,
		  alg.id along_id,
		  alg.name along_name,
		  cbf.generous_welfare_id
		from
		  companies com
		  inner join locations loc on com.id = loc.companies_id
		  inner join technologies tec on com.id = tec.company_id
		  inner join languages lan on tec.language_id = lan.id
		  inner join commuting cmu on com.id = cmu.company_id
		  inner join alongs alg on cmu.along_id = alg.id
		  left join company_benefits cbf on com.id = cbf.company_id
		order by
		  com.id
		`)
	if err != nil {
		return companies, err
	}

	return companies, nil
}
