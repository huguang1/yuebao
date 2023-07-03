package models

import (
	db "../dbs"
	"log"
	"strconv"
)

type Record struct {
	Id        int    `json:"id" form:"id"`
	Type string `json:"type" form:"type"`
	Amount  string `json:"amount" form:"amount"`
	Member  string `json:"member" form:"member"`
}

// 增
func (r *Record) AddRecord() (id int64, err error) {
	res, err := db.Conns.Exec("INSERT INTO bao_record (type, amount, member) VALUES (?, ?, ?)", r.Type, r.Amount, r.Member)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

// 删
func DeleteRecord(id int) (n int64, err error) {
	n = 0
	rs, err := db.Conns.Exec("DELETE FROM bao_record WHERE id=?", id)
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
func (r *Record) UpdateRecord(id int) (n int64, err error) {
	res, err := db.Conns.Prepare("UPDATE bao_record SET type=?,amount=?,member=? WHERE id=?")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}
	rs, err := res.Exec(r.Type, r.Amount, r.Member, r.Id)
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
func RecordList(page, pageSize int, filters ...interface{}) (lists []Record, count int64, err error) {
	lists = make([]Record, 0)  // 初始化数据
	where := "WHERE 1=1"
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 3 {
			where = where + " AND " + filters[k].(string) + filters[k+1].(string) + filters[k+2].(string)
		}
	}
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	rows, err := db.Conns.Query("SELECT id, type, amount, member FROM bao_record " + where + " LIMIT " + limit)
	defer rows.Close()
	if err != nil {
		return
	}
	count = 0
	for rows.Next() {
		var record Record
		rows.Scan(&record.Id, &record.Type, &record.Amount, &record.Member)
		lists = append(lists, record)
		count++
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
