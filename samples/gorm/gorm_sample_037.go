package gorm

import (
	"fmt"
	"log"
)

/*
	HasMany(1対多)の関連先のテーブルを検索条件にしてレコードを取得する
*/
func GormSample037() {
	fmt.Println("gorm_sample_037")

	db := DbInit()
	AutoMigrate(tables003...)

	records := []Rich{}

	//銀行名(CreditCard.Bank) が Mizuho のクレジットカードを持つ Richレコードを全て取得
	ret := db.Preload("CreditCards").Joins("JOIN credit_cards ON credit_cards.rich_id = riches.id").
	Where("credit_cards.bank = ?", "Mizuho").Find(&records)
	if ret.Error != nil {
		log.Fatal(ret.Error)
	}
	fmt.Println(records)
	fmt.Println("RowsAffected:", ret.RowsAffected)
}
