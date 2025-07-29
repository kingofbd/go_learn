package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:wcq5201314@tcp(127.0.0.1:3306)/wcq?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	// 创建两张表
	db.AutoMigrate(Account{}, Transaction{})

	// 定义两个账户，并插入到表中
	a := Account{Balance: 280}
	b := Account{Balance: 100}
	db.Create(&a)
	db.Create(&b)

	// 进行转账
	transfer(db, &a, &b, 100)
	transfer(db, &a, &b, 100)
	transfer(db, &a, &b, 100)
}

type Account struct {
	ID      uint
	Balance int
}

type Transaction struct {
	ID            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        int
}

// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
// 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
func transfer(db *gorm.DB, a *Account, b *Account, amount int) error {
	// 使用事务
	return db.Transaction(func(tx *gorm.DB) error {
		// 检查a账户的余额是否满足转账需求
		res := tx.Debug().Model(&Account{}).Where("id = ? and balance >= ?", a.ID, amount).Update("balance", gorm.Expr("balance - ?", amount))
		if res.Error != nil {
			return res.Error
		}
		// 如果没查询到满足条件的账户，也要抛出异常
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		fmt.Println("账户", a.ID, "转出了", amount)

		// 走到这里，说明a账户已经扣款，那么b账户需要增加余额
		res2 := tx.Debug().Model(&Account{}).Where("id = ?", b.ID).Update("balance", gorm.Expr("balance + ?", amount))
		if res2.Error != nil {
			return res.Error
		}
		fmt.Println("账户", b.ID, "转入了", amount)

		// 对交易进行记录
		trans := Transaction{FromAccountId: a.ID, ToAccountId: b.ID, Amount: amount}
		res3 := tx.Debug().Create(&trans)
		if res3.Error != nil {
			return res3.Error
		}

		// 都没问题，提交事务
		return nil
	})
}
