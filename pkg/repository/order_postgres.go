package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	wbtechl0 "wb-tech-l0"
)

type OrderPostgres struct {
	db    *sqlx.DB
	cache *Cache
}

func NewOrderPostgres(db *sqlx.DB, cache *Cache) *OrderPostgres {
	return &OrderPostgres{db: db, cache: cache}
}

func (r *OrderPostgres) GetById(OrderUid string) (wbtechl0.Order, error) {
	logrus.Printf("Service is trying to find the order(%s)...\n", OrderUid)
	return r.cache.GetOrder(OrderUid)
}

func (r *OrderPostgres) CreateOrder(order wbtechl0.Order) error {
	logrus.Printf("Db is adding an new order(%s)...\n", order.OrderUid)
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	createOrderQuery := fmt.Sprintf("INSERT INTO \"%s\" (%s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", orderTable, orderTableFields)
	createDeliveryQuery := fmt.Sprintf("INSERT INTO %s (orderuid, %s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", deliveryTable, deliveryTableFields)
	createPaymentQuery := fmt.Sprintf("INSERT INTO %s (orderuid, %s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", paymentTable, paymentTableFields)
	createItemQuery := fmt.Sprintf("INSERT INTO %s (orderuid, %s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", itemTable, itemTableFields)

	id := order.OrderUid
	_, err = tx.Exec(createOrderQuery, id, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.ShardKey, order.SmId, order.DateCreated, order.OofShard)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(createDeliveryQuery, id, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(createPaymentQuery, id, order.Payment.Transaction, order.Payment.RequestId, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range order.Items {
		_, err = tx.Exec(createItemQuery, id, item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	logrus.Printf("Cache is adding an new order(%s)...\n", order.OrderUid)
	r.cache.AddOrder(order)

	logrus.Printf("Order(%s) was added to DataBase and Cache.\n", order.OrderUid)
	return tx.Commit()
}

func (r *OrderPostgres) UploadCache() error {
	logrus.Println("Service is trying to upload cache...")
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	var orders []wbtechl0.Order
	getOrdersQuery := fmt.Sprintf("SELECT * FROM \"%s\"", orderTable)
	getDeliveryQuery := fmt.Sprintf("SELECT %s FROM %s WHERE orderuid = $1", deliveryTableFields, deliveryTable)
	getPaymentQuery := fmt.Sprintf("SELECT %s FROM %s WHERE orderuid = $1", paymentTableFields, paymentTable)
	getItemsQuery := fmt.Sprintf("SELECT %s FROM %s WHERE orderuid = $1", itemTableFields, itemTable)

	err = tx.Select(&orders, getOrdersQuery)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, order := range orders {
		var delivery wbtechl0.Delivery
		err = tx.Get(&delivery, getDeliveryQuery, order.OrderUid)

		if err != nil {
			tx.Rollback()
			return err
		}

		var payment wbtechl0.Payment
		err = tx.Get(&payment, getPaymentQuery, order.OrderUid)

		if err != nil {
			tx.Rollback()
			return err
		}

		var items []wbtechl0.Item
		err = tx.Select(&items, getItemsQuery, order.OrderUid)

		if err != nil {
			tx.Rollback()
			return err
		}

		order.Delivery = delivery
		order.Payment = payment
		order.Items = items

		r.cache.AddOrder(order)
	}

	logrus.Println("Upload of cache completed!")
	return tx.Commit()
}
