package model

import (
	"loco/domain"

	"github.com/beego/beego/orm"
)

/*
CREATE TABLE transactions (
    id INT NOT NULL AUTO_INCREMENT,
    transactionId INT NOT NULL,
    `type` ENUM('', 'shopping', 'travel', 'food', 'groceries') DEFAULT '',
    amount FLOAT,
    parentId INT DEFAULT NULL,
    createdOn DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedOn DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)
*/

type Transactions struct {
	Id              int     `orm:"column(id);auto_increment"`
	TransactionId   int     `orm:"column(transactionId);"`
	TransactionType string  `orm:"column(transactionType)"`
	Amount          float64 `orm:"column(amount)"`
	ParentId        int     `orm:"column(parentId)"`
}

func init() {
	orm.RegisterModel(new(Transactions))
}

func (t *Transactions) Insert(input *domain.Transaction) error {
	o := orm.NewOrm()
	_, err := o.Insert(t.domainToModel(input))
	return err
}

func (t *Transactions) GetTransaction(input *domain.Transaction) (*domain.Transaction, error) {
	o := orm.NewOrm()
	err := o.QueryTable(new(Transactions)).Filter("transactionId", input.TransactionId).One(t)
	return t.modelToDomain(t), err
}

func (t *Transactions) GetTransactionsByType(input *domain.Transaction) (output []*domain.Transaction, err error) {
	o := orm.NewOrm()

	var allTxn []*Transactions
	_, err = o.QueryTable(new(Transactions)).Filter("transactionType", input.TransactionType).All(&allTxn)

	for _, txn := range allTxn {
		output = append(output, txn.modelToDomain(txn))
	}
	return output, err
}

func (t *Transactions) GetTransactionsByParentId(input *domain.Transaction) (output []*domain.Transaction, err error) {
	o := orm.NewOrm()

	var allTxn []*Transactions
	_, err = o.QueryTable(new(Transactions)).Filter("parentId", input.ParentId).All(&allTxn)

	for _, txn := range allTxn {
		output = append(output, txn.modelToDomain(txn))
	}
	return output, err
}

func (t *Transactions) modelToDomain(in *Transactions) (output *domain.Transaction) {
	if in == nil {
		return
	}

	output = &domain.Transaction{
		Id:              in.Id,
		TransactionId:   in.TransactionId,
		TransactionType: in.TransactionType,
		Amount:          in.Amount,
		ParentId:        in.ParentId,
	}
	return
}

func (t *Transactions) domainToModel(in *domain.Transaction) (output *Transactions) {
	if in == nil {
		return
	}

	output = &Transactions{
		Id:              in.Id,
		TransactionId:   in.TransactionId,
		TransactionType: in.TransactionType,
		Amount:          in.Amount,
		ParentId:        in.ParentId,
	}
	return
}
