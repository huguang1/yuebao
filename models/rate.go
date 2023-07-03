package models

import (
	db "../dbs"
	"log"
	"strconv"
)

type Rate struct {
	Id        int    `json:"id" form:"id"`
	InterestRate string `json:"interest_rate" form:"interest_rate"`
}

// 增
func (r *Rate) AddRate() (id int64, err error) {
	res, err := db.Conns.Exec("INSERT INTO bao_rate (interest_rate) VALUES (?)", r.InterestRate)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

// 删
func DeleteRate(id int) (n int64, err error) {
	n = 0
	rs, err := db.Conns.Exec("DELETE FROM bao_rate WHERE id=?", id)
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
func (r *Rate) UpdateRate(id int) (n int64, err error) {
	res, err := db.Conns.Prepare("UPDATE bao_rate SET interest_rate=? WHERE id=?")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}
	rs, err := res.Exec(r.InterestRate, r.Id)
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
func RateList(page, pageSize int, filters ...interface{}) (lists []Rate, count int64, err error) {
	lists = make([]Rate, 0)  // 初始化数据
	where := "WHERE 1=1"
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 3 {
			where = where + " AND " + filters[k].(string) + filters[k+1].(string) + filters[k+2].(string)
		}
	}
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	rows, err := db.Conns.Query("SELECT id, interest_rate FROM bao_rate " + where + " LIMIT " + limit)
	defer rows.Close()
	if err != nil {
		return
	}
	count = 0
	for rows.Next() {
		var rate Rate
		rows.Scan(&rate.Id, &rate.InterestRate)
		lists = append(lists, rate)
		count++
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
