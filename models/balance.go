package models

import (
	db "../dbs"
	"log"
	"strconv"
)

type Balance struct {
	Id        int    `json:"id" form:"id"`
	Balance string `json:"balance" form:"balance"`
	TransferIn  string `json:"transfer_in" form:"transfer_in"`
	TransferOut  string `json:"transfer_out" form:"transfer_out"`
	Interest  string `json:"interest" form:"interest"`
	Member  string `json:"member" form:"member"`
	IsCompute  int `json:"is_compute" form:"is_compute"`

}

// 增
func (b *Balance) AddBalance() (id int64, err error) {
	res, err := db.Conns.Exec("INSERT INTO bao_balance (balance, transfer_in, transfer_out, interest, member, is_compute," +
		" create_time) VALUES (?, ?, ?, ?, ?, ?)", b.Balance, b.TransferIn, b.TransferOut, b.Interest, b.Member, b.IsCompute)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

// 删
func DeleteBalance(id int) (n int64, err error) {
	n = 0
	rs, err := db.Conns.Exec("DELETE FROM bao_balance WHERE id=?", id)
	if err != nil {
		log.Fatalln(err)
		return
	}
	n, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

// 改
func (b *Balance) UpdateBalance(id int) (n int64, err error) {
	res, err := db.Conns.Prepare("UPDATE bao_balance SET balance, transfer_in, transfer_out, interest, member, is_compute" +
		" WHERE id=?")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}
	rs, err := res.Exec(b.Balance, b.TransferIn, b.TransferOut, b.Interest, b.Member, b.IsCompute, b.Id)
	if err != nil {
		log.Fatal(err)
	}
	n, err = rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return
}

// 查
func BalanceList(page, pageSize int, filters ...interface{}) (lists []Balance, count int64, err error) {
	lists = make([]Balance, 0)  // 初始化数据
	where := "WHERE 1=1"
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 3 {
			where = where + " AND " + filters[k].(string) + filters[k+1].(string) + filters[k+2].(string)
		}
	}
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	rows, err := db.Conns.Query("SELECT id, balance, transfer_in, transfer_out, interest, member, is_compute FROM bao_balance " + where + " LIMIT " + limit)
	defer rows.Close()
	if err != nil {
		return
	}
	count = 0
	for rows.Next() {
		var balance Balance
		rows.Scan(&balance.Id, &balance.Balance, &balance.TransferIn, &balance.TransferOut, &balance.Interest, &balance.Member, &balance.IsCompute)
		lists = append(lists, balance)
		count++
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
