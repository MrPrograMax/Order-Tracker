package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	orderTable    = "order"
	deliveryTable = "delivery"
	paymentTable  = "payment"
	itemTable     = "item"

	orderTableFields    = "orderuid, tracknumber, entry, locale, internalsignature, customerid, deliveryservice, shardkey, smid, datecreated, oofshard"
	deliveryTableFields = "name, phone, zip, city, address, region, email"
	paymentTableFields  = "transaction, requestid, currency, provider, amount, paymentdt, bank, deliverycost, goodstotal, customfee"
	itemTableFields     = "chrtid, tracknumber, price, rid, name, sale, size, totalprice, nmid, brand, status"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
