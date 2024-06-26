// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type ExpenseCategory string

const (
	ExpenseCategoryValue0 ExpenseCategory = "Супермаркеты"
	ExpenseCategoryValue1 ExpenseCategory = "Развлечение"
	ExpenseCategoryValue2 ExpenseCategory = "Спорт"
	ExpenseCategoryValue3 ExpenseCategory = "Красота"
	ExpenseCategoryValue4 ExpenseCategory = "Медицина"
	ExpenseCategoryValue5 ExpenseCategory = "Фастфуд"
	ExpenseCategoryValue6 ExpenseCategory = "Рестораны"
	ExpenseCategoryValue7 ExpenseCategory = "Другое"
)

func (e *ExpenseCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ExpenseCategory(s)
	case string:
		*e = ExpenseCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for ExpenseCategory: %T", src)
	}
	return nil
}

type NullExpenseCategory struct {
	ExpenseCategory ExpenseCategory
	Valid           bool // Valid is true if ExpenseCategory is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullExpenseCategory) Scan(value interface{}) error {
	if value == nil {
		ns.ExpenseCategory, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ExpenseCategory.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullExpenseCategory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ExpenseCategory), nil
}

type Expense struct {
	ID         int64
	Owner      int64
	Wallet     int64
	WalletName string
	Currency   string
	Value      float64
	Name       string
	
	Category   ExpenseCategory
	Date       time.Time
	CreatedAt  time.Time
}

type User struct {
	ID               int64
	Username         string
	HashedPassword   string
	FullName         string
	Email            string
	PasswordChangeAt time.Time
	CreatedAt        time.Time
}

type Wallet struct {
	ID        int64
	Owner     int64
	Name      string
	Balance   sql.NullFloat64
	Currency  string
	CreatedAt time.Time
}
