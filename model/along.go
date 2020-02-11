package model

import "gopkg.in/gorp.v1"

// Along 沿線情報
type Along struct {
	ID   int    `db:"along_id"`
	Name string `db:"along_name"`
}

// GetAlongs 沿線情報を取得します
func GetAlongs(tx *gorp.Transaction) error {
	return nil
}
