package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type stu struct {
	name string
	id   int
	age  int
}

var db *sql.DB

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	} else {
		fmt.Println("connected")
	}

	//QueryRow()
	//DeleteRow()
	//AddRow()
	QueryRows()
}

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:root@tcp(localhost:3306)/mydatabase"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func QueryRow() {
	sqlStr := "select * from mytable where id=1"
	row := db.QueryRow(sqlStr)
	var stu1 stu
	err := row.Scan(&stu1.name, &stu1.id, &stu1.age)
	if err != nil {
		fmt.Println("scan err")
		return
	}
	fmt.Println(stu1)

}
func DeleteRow() {
	sqlStr := "delete from mytable where id=2"
	res, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(affected, "处发生改变")
	}
}
func AddRow() {
	sqlStr := "insert into mytable(name, id, age) values('xiaote2', 2, 18)"
	res, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(affected, "处发生改变")
	}
}
func QueryRows() {
	sqlStr := "select * from mytable where name like 'xiaote_'"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("query err:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var stu0 stu
		err := rows.Scan(&stu0.name, &stu0.id, &stu0.age)
		if err != nil {
			fmt.Println("scan err:", err)
		}
		fmt.Println(stu0)
	}

}
func update(username string, p string) {
	sqlStr := "update userdata set password=? where username=?"
	_, err := db.Exec(sqlStr, p, username)
	if err != nil {
		fmt.Println("Exec err: ", err)
		return
	}
	fmt.Println("changed!")
}
func query(username string, password string) {
	sqlStr := "select username,password from userdata where username=? and password=?"
	row := db.QueryRow(sqlStr, username, password)
	err := row.Scan(&username, &password)
	if err != nil {
		fmt.Println("用户名或密码错误")
		return
	}
	fmt.Println(username, "登陆成功！")
}

func add(username string, password string) {
	sqlStr := "insert into userdata values(?,?)"
	_, err := db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Println("Exec err:", err)
		return
	}
}
