package model

import "gopkg.in/gorp.v1"

// GenerousWelfare 沿線情報
type GenerousWelfare struct {
	ID   int    `db:"generousWelfare_id" json:"id"`
	Name string `db:"generousWelfare_name" json:"name"`
}

// GetGenerousWelfare 福利厚生情報を取得します
func GetGenerousWelfare(tx *gorp.Transaction) ([]GenerousWelfare, error) {

	generousWelfares, err := selectToGenerousWelfares(tx)
	if err != nil {
		return generousWelfares, err
	}

	return generousWelfares, nil
}

// selectToGenerousWelfares 福利厚生情報を検索します
func selectToGenerousWelfares(tx *gorp.Transaction) ([]GenerousWelfare, error) {
	var generousWelfares []GenerousWelfare
	_, err := tx.Select(&generousWelfares, `
		select
		  id generousWelfare_id,
		  name generousWelfare_name
		from
		  generous_welfares
		order by
		  id
	`)
	if err != nil {
		return generousWelfares, err
	}

	return generousWelfares, nil
}
