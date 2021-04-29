package service

import (
	"fmt"
	"time"

	"github.com/exange_rate_service/repository"
	"github.com/exange_rate_service/repository/model"
	exchageRateCrawler "github.com/exange_rate_service/service/crawler"
	"github.com/roylee0704/gron"
)

func Run() {
	updateExchangeRate()

	c := gron.New()
	c.AddFunc(gron.Every(10*time.Minute), updateExchangeRate)
	c.Start()
}

func updateExchangeRate() {
	fmt.Println("schedule run at:", time.Now().Format(time.RFC3339))

	data := exchageRateCrawler.GetExchangeRate()

	repo := new(repository.ExchangeRateRepository)
	repo.Init()
	for _, item := range data {
		var dao model.ExchangeRateDao
		dao.Name = item.Name
		dao.CashBuyingRate = item.CashBuyingRate
		dao.CashSellingRate = item.CashSellingRate
		dao.SignBuyingRate = item.SignBuyingRate
		dao.SignSellingRate = item.SignSellingRate
		repo.InsertOrUpdate(dao)
	}
}
