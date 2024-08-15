package domain

import "strings"

const AllTransactionTypes = "shopping,travel,food,groceries"

type Transaction struct {
	Id              int     `json:"id"`
	TransactionId   int     `json:"transaction_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"type"`
	ParentId        int     `json:"parent_id"`
}

type TransactionResponse struct {
	Amount   int    `json:"amount"`
	Type     string `json:"type"`
	ParentId int    `json:"parentId"`
}

type TransactionListResponse struct {
	TransactionsList []int `json:"transactionsList"`
}

func (t *Transaction) IsValidRequest() bool {
	return t.TransactionId != 0 && t.Amount != 0 && t.TransactionType != ""
}

func (t *Transaction) IsValidTransactionTypes() bool {
	return strings.Contains(AllTransactionTypes, t.TransactionType)
}
