package model

import "gopkg.in/gorp.v1"

// Along 沿線情報
type Along struct {
	ID   int    `db:"along_id"`
	Name string `db:"along_name"`
}

// GetAlongs 沿線情報を取得します
func GetAlongs(tx *gorp.Transaction) ([]Along, error) {

	alongs, err := selectToAlongs(tx)
	if err != nil {
		return alongs, err
	}

	return alongs, nil
}

// selectToAlongs 沿線情報を検索します
func selectToAlongs(tx *gorp.Transaction) ([]Along, error) {
	var alongs []Along
	_, err := tx.Select(&alongs, `
		select
		  id along_id,
		  name along_name
		from
		  alongs
		order by
		  id
	`)
	if err != nil {
		return alongs, err
	}

	return alongs, nil
}
