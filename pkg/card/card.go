package card

type Transaction struct {
	Id int64
	Amount int64
	MCC string
	Date int64
	Status string
}

type Card struct {
	Id int64
	Issuer string
	Currency string
	Balance int64
	Number int64
	Icon string
	Transactions []*Transaction  // Слайс будет состоять из ссылок на элементы типа Transaction ?
}

func AddTransaction(card *Card, transaction *Transaction) {
	(*card).Transactions = append((*card).Transactions, transaction)
}

func SumByMCC(transactions []*Transaction, mcc []string) map[string]int64 {
	sumOfMcc := make(map[string]int64,len(mcc))
	for _, m := range mcc {
		for _, transaction := range transactions {
			if (transaction.MCC == m){
				sumOfMcc[m] += transaction.Amount
			}
		}
	}
	return sumOfMcc
}


