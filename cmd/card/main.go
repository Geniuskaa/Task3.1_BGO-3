package main

import (
	"fmt"
	"github.com/Geniuskaa/Task3.1_BGO-3/pkg/card"
	"time"
)

func main() {
	card1 := &card.Card{
		Id:           1,
		Issuer:       "VISA",
		Currency:     "RUB",
		Balance:      145_500_00,
		Number:       4924_2873_3819_3841,
		Icon:         "https://someIcon.com",
		Transactions: []*card.Transaction{},
	}

	card1.Transactions = append(card1.Transactions, &card.Transaction{
		Id:     25,
		Amount: 115_000_00,
		MCC:    "5000", // Обозначим этот MCC как пополнение
		Date:   time.Now().Unix(),
		Status: "completed",
	})

	card.AddTransaction(card1, &card.Transaction{ // Пример использования функции AddTransaction
		Id:     26,
		Amount: 12_000_00,
		MCC:    "6501", // Этот MCC будет связан с супермаркетами
		Date:   time.Now().Unix(),
		Status: "Completed",
	})

	card.AddTransaction(card1, &card.Transaction{
		Id:     27,
		Amount: 10_000_00,
		MCC:    "6501",
		Date:   time.Now().Unix(),
		Status: "Completed",
	})

	someMccCodes := []string{"6555", "6501", "5000"}
	supermarketExpenses := card.SumByMCC(card1.Transactions, someMccCodes)
	fmt.Println(supermarketExpenses["6555"])
	fmt.Println(supermarketExpenses["6501"]) // Вывод в копейках
	fmt.Println(supermarketExpenses["5000"])


	category := card.TranslateMCC(card1.Transactions[0].MCC)
	fmt.Println(category)

}
