package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// Displays 表示企業情報
type Displays struct {
	ID   int    `db:"company_id" json:"companyId"`
}

// GetDisplayCompanies 表示する企業情報を取得します
func GetDisplayCompanies(tx *gorp.Transaction, languages string, alongs string, generousWelfares string) ([]Displays, error) {
	displays, err := SelectToDisplays(tx, languages, alongs, generousWelfares)
	if err != nil {
		return displays, err
	}

	return displays, nil
}

// SelectToDisplays 沿線情報を検索します
func SelectToDisplays(tx *gorp.Transaction, languages string, alongs string, generousWelfares string) ([]Displays, error) {
	
	var whereLanguages string
	var whereAlongs string
	var whereGenerousWelfares string
	if (languages != "") {
		whereLanguages = fmt.Sprintf("tec.language_id in (%s)", languages)
	}
	if (alongs != "") {
		if (languages != "") {
			whereAlongs = fmt.Sprintf("and cmu.along_id in (%s)", alongs)
		} else {
			whereAlongs = fmt.Sprintf("cmu.along_id in (%s)", alongs)
		}
	}
	if (generousWelfares != "") {
		if (languages != "" || alongs != "") {
			whereGenerousWelfares = fmt.Sprintf("and ben.generous_welfare_id in (%s)", generousWelfares)
		} else {
			whereGenerousWelfares = fmt.Sprintf("ben.generous_welfare_id in (%s)", generousWelfares)
		}
	}

	var displays []Displays
	var sql = fmt.Sprintf(`
	select
  		cop.id company_id
	from
  		companies cop
  		left join technologies tec on cop.id = tec.company_id
  		left join commuting cmu on cop.id = cmu.company_id
  		left join company_benefits ben on cop.id = ben.company_id
	where
  		%s
  		%s
  		%s
	group by
  		company_id
	order by
  		company_id
	`, whereLanguages, whereAlongs, whereGenerousWelfares)

	_, err := tx.Select(&displays, sql)
	if err != nil {
		return displays, err
	}

	return displays, nil
}
