package models

import (
	db "../dbs"
	"log"
	"strconv"
)

type Member struct {
	Id        int    `json:"id" form:"id"`
	Account string `json:"account" form:"account"`
	Password  string `json:"password" form:"password"`
}

// 增
func (m *Member) AddMember() (id int64, err error) {
	res, err := db.Conns.Exec("INSERT INTO bao_member (account, password) VALUES (?, ?)", m.Account, m.Password)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

// 删
func DeleteMember(id int) (n int64, err error) {
	n = 0
	rs, err := db.Conns.Exec("DELETE FROM bao_member WHERE id=?", id)
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
func (m *Member) UpdateMember(id int) (n int64, err error) {
	res, err := db.Conns.Prepare("UPDATE bao_member SET account=?,password=? WHERE id=?")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}
	rs, err := res.Exec(m.Account, m.Password, m.Id)
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
func MemberList(page, pageSize int, filters ...interface{}) (lists []Member, count int64, err error) {
	lists = make([]Member, 0)  // 初始化数据
	where := "WHERE 1=1"
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 3 {
			where = where + " AND " + filters[k].(string) + filters[k+1].(string) + filters[k+2].(string)
		}
	}
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	rows, err := db.Conns.Query("SELECT id, account, password FROM bao_member " + where + " LIMIT " + limit)
	defer rows.Close()
	if err != nil {
		return
	}
	count = 0
	for rows.Next() {
		var member Member
		rows.Scan(&member.Id, &member.Account, &member.Password)
		lists = append(lists, member)
		count++
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
