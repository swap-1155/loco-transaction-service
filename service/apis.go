package service

import (
	"fmt"
	"loco/domain"
	"loco/util/model"
	"net/http"

	"github.com/spf13/cast"
)

func GetTransactionById(request *domain.Transaction) (response map[string]interface{}, err error) {

	response = make(map[string]interface{})
	response["status"] = http.StatusText(http.StatusInternalServerError)

	transaction, err := new(model.Transactions).GetTransaction(request)
	if err != nil {
		response["error"] = err.Error()
		fmt.Println(err)
		return
	}

	response["status"] = http.StatusText(http.StatusOK)
	response["model"] = transaction

	return
}

func PostTransactionById(request *domain.Transaction) (response map[string]interface{}, err error) {

	response = make(map[string]interface{})
	response["status"] = http.StatusText(http.StatusInternalServerError)

	if !request.IsValidRequest() {
		response["status"] = http.StatusText(http.StatusBadRequest)
		response["error"] = "Data missing in request"
		response["data"] = request

		fmt.Println(err)
		return
	}

	err = new(model.Transactions).Insert(request)
	if err != nil {
		response["error"] = err.Error()
		fmt.Println(err)
		return
	}

	response["status"] = http.StatusText(http.StatusOK)

	return
}

func GetTransactionsByType(request *domain.Transaction) (response map[string]interface{}, err error) {

	response = make(map[string]interface{})
	response["status"] = http.StatusText(http.StatusInternalServerError)

	if !request.IsValidTransactionTypes() {
		response["status"] = http.StatusText(http.StatusBadRequest)
		response["error"] = "invalid transaction type: " + request.TransactionType

		fmt.Println(err)
		return
	}

	transactions, err := new(model.Transactions).GetTransactionsByType(request)
	if err != nil {
		response["error"] = err.Error()
		fmt.Println(err)
		return
	}

	if transactions == nil {
		response["status"] = "No data available"
		return
	}

	response["status"] = http.StatusText(http.StatusOK)

	transactionListResponse := &domain.TransactionListResponse{}
	for _, txn := range transactions {
		transactionListResponse.TransactionsList = append(transactionListResponse.TransactionsList, txn.TransactionId)
	}
	response["model"] = transactionListResponse

	return
}

func GetTotalTransactionAmount(request *domain.Transaction) (response map[string]interface{}, err error) {

	response = make(map[string]interface{})
	response["status"] = http.StatusText(http.StatusInternalServerError)

	transaction, err := new(model.Transactions).GetTransaction(request)
	if err != nil {
		response["error"] = err.Error()
		fmt.Println(err)
		return
	}

	var childTransactionAmount float64
	transaction.ParentId = transaction.TransactionId

	childTransactionAmount, err = GetChildTransactionAmount(transaction)
	if err != nil {
		return
	}

	response["status"] = http.StatusText(http.StatusOK)
	response["sum"] = cast.ToFloat64(fmt.Sprintf("%.2f", transaction.Amount+childTransactionAmount))

	return
}

func GetChildTransactionAmount(request *domain.Transaction) (amount float64, err error) {

	transactions, err := new(model.Transactions).GetTransactionsByParentId(request)
	if err != nil {
		return
	}

	// var wg sync.WaitGroup

	for _, txn := range transactions {
		// wg.Add(1)

		// go func(txn *domain.Transaction) {
		// 	defer wg.Done()

		var childTransactionAmount float64
		txn.ParentId = txn.TransactionId

		childTransactionAmount, err = GetChildTransactionAmount(txn)
		if err != nil {
			return
		}

		amount = amount + txn.Amount + childTransactionAmount

		// }(txn)
	}

	// wg.Wait()

	return
}
