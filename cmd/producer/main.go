package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	wbtechl0 "wb-tech-l0"
)

func main() {
	logrus.Println("Producer started to work...")

	logrus.Println("Initializing the config...")
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	logrus.Println("Initialization completed.")

	sc, err := stan.Connect("test-cluster", "test-client", stan.NatsURL(viper.GetString("nats.uri")))

	if err != nil {
		log.Fatal(err)
	}
	logrus.Println("Connected to nats server.")

	defer sc.Close()

	go func() {
		logrus.Println("Sending orders...")
		tik := time.NewTicker(time.Second * 5)
		for range tik.C {
			order := CreateOrder()
			orderJSON, err := json.Marshal(order)

			if err != nil {
				log.Fatal(err)
			}

			logrus.Println("Sending message to orders topic...")
			logrus.Printf("Order Id: %s", order.OrderUid)

			if err = sc.Publish("orders", orderJSON); err != nil {
				log.Fatal(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Sending process is finished")

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func CreateOrder() wbtechl0.Order {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyz"
	numbers := "1234567890"
	letAndNum := "abcdefghijklmnopqrstuvwxyz1234567890"

	name := randomStr(letters, 4)
	delivery := wbtechl0.Delivery{
		Name:    name + " " + name + "ov",
		Phone:   "+97" + randomStr(numbers, 9),
		Zip:     randomStr(numbers, 7),
		City:    randomStr(letters, 9),
		Address: randomStr(letters, 9) + randomStr(numbers, 3),
		Region:  randomStr(letters, 7),
		Email:   randomStr(letAndNum, 7) + "@gmail.com",
	}

	payment := wbtechl0.Payment{
		Transaction:  randomStr(letAndNum, 20),
		RequestId:    "",
		Currency:     randomStr(letters, 3),
		Provider:     randomStr(letters, 6),
		Amount:       rand.Intn(10000),
		PaymentDt:    int64(1000 + rand.Intn(9000)),
		Bank:         randomStr(letters, 7),
		DeliveryCost: 1000 + rand.Intn(9000),
		GoodsTotal:   100 + rand.Intn(900),
		CustomFee:    rand.Intn(10000),
	}

	item := wbtechl0.Item{
		ChrtId:      10000 + rand.Intn(90000),
		TrackNumber: randomStr(letters, 14),
		Price:       100 + rand.Intn(900),
		Rid:         randomStr(letAndNum, 20),
		Name:        randomStr(letters, 10),
		Sale:        rand.Intn(100),
		Size:        randomStr(numbers, 1),
		TotalPrice:  rand.Intn(100000),
		NmId:        10000 + rand.Intn(90000),
		Brand:       randomStr(letters, 10),
		Status:      rand.Intn(100),
	}

	timeOrder, _ := time.Parse(time.RFC3339, "2021-11-26T06:22:19Z")
	order := wbtechl0.Order{
		TrackNumber:       randomStr(letters, 14),
		Entry:             randomStr(letters, 4),
		Locale:            randomStr(letters, 2),
		InternalSignature: "",
		CustomerId:        randomStr(letters, 4),
		DeliveryService:   randomStr(letters, 4),
		ShardKey:          randomStr(numbers, 1),
		SmId:              rand.Intn(100),
		DateCreated:       timeOrder,
		OofShard:          randomStr(numbers, 1),
	}

	order.OrderUid = uuid.New().String()
	order.Delivery = delivery
	order.Payment = payment
	order.Items = append(order.Items, item)

	return order
}

func randomStr(str string, lenStr int) string {
	var b strings.Builder
	for i := 0; i < lenStr; i++ {
		ch := string(str[rand.Intn(len(str))])
		b.WriteString(ch)
	}
	return b.String()
}
