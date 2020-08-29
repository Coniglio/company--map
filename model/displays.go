package model

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// Displays 表示企業情報
type Displays struct {
	ID   int    `db:"company_id" json:"companyId"`
}

// GetDisplayCompanies 表示する企業情報を取得します
func GetDisplayCompanies(tx *gorp.Transaction) ([]Displays, error) {
	displays, err := SelectToDisplays(tx)
	if err != nil {
		return displays, err
	}

	return displays, nil
}

// SelectToDisplays 沿線情報を検索します
func SelectToDisplays(tx *gorp.Transaction) ([]Displays, error) {
	var displays []Displays
	var sql = `
	select
		cmp.id company_id
	from
		companies cmp
		inner join technologies tec on cmp.id = tec.company_id
		inner join commuting cmm on cmp.id = cmm.company_id
		inner join company_benefits cbf on cmp.id = cbf.company_id
	where
		%s
	order by
		id
	`

	// 検索条件をSQLに反映

	_, err := tx.Select(&displays, sql)
	if err != nil {
		return displays, err
	}

	return displays, nil
}
