package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Message struct {
	message string
	time    string
	sender  string
}

var db *sql.DB

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	} else {
		fmt.Println("database connected")
	}
	a := 0
	for {
		showMenu()
		fmt.Print(":")
		fmt.Scan(&a)
		switch a {
		case 1:
			register()
		case 2:
			login()
		case 3:
			return
		default:
			fmt.Println("please input again!")
		}
	}
}

func showMenu() {
	fmt.Println("**********")
	fmt.Println("* 1.注册 *")
	fmt.Println("* 2.登陆 *")
	fmt.Println("* 3.退出 *")
	fmt.Println("**********")
}

func initDB() (err error) {
	dsn := "root:root@tcp(localhost:3306)/lv3database"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func register() {
	u := ""
	p := ""
	for {
		fmt.Print("username:")
		fmt.Scan(&u)
		sqlStr1 := "select username from userdata where username=? "
		row := db.QueryRow(sqlStr1, u)
		err := row.Scan(&u)
		if err != nil {
			fmt.Println("name is available")
			fmt.Print("password:")
			fmt.Scan(&p)
			sqlStr2 := "insert into userdata values(?,?)"
			_, err := db.Exec(sqlStr2, u, p)
			if err != nil {
				fmt.Println("Exec err:", err)
				return
			}
			fmt.Println("register ok!")
			return
		}
		fmt.Println("name existed!")
	}
}

func login() {
	u := ""
	p := ""
	fmt.Print("username:")
	fmt.Scan(&u)
	fmt.Print("password:")
	fmt.Scan(&p)
	sqlStr := "select username,password from userdata where username=? and password=?"
	row := db.QueryRow(sqlStr, u, p)
	err := row.Scan(&u, &p)
	if err != nil {
		fmt.Println("username or password is wrong!")
		return
	}
	fmt.Println(u, " login ok！")
	ShowMessages(u)
	a := 0
	for {
		fmt.Println("**********")
		fmt.Println("* 1.回复 *")
		fmt.Println("* 0.注销 *")
		fmt.Println("**********")
		fmt.Scan(&a)
		switch a {
		case 1:
			r := ""
			m := ""
			fmt.Print("receiver:")
			fmt.Scan(&r)
			fmt.Print("message:")
			fmt.Scan(&m)
			send(u, r, m)
		case 0:
			return
		default:
			fmt.Println("please input again!")
		}
	}

	return

}
func ShowMessages(username string) {
	sqlStr := "select time, sender, message from chattable where receiver=?"
	rows, err := db.Query(sqlStr, username)
	if err != nil {
		fmt.Println("Query err:", err)
		return
	}
	defer rows.Close()
	var massage Message
	for rows.Next() {
		err = rows.Scan(&massage.time, &massage.sender, &massage.message)
		if err != nil {
			fmt.Println("Scan err:", err)
			return
		}
		fmt.Println("----------------------------------------")
		fmt.Println("   time: ", massage.time)
		fmt.Println("   from: ", massage.sender)
		fmt.Println("message: ", massage.message)
		fmt.Println("----------------------------------------")
	}
}
func send(sender string, receiver string, message string) {
	sqlStr := "insert into chattable(sender,receiver,time,message) values(?,?,?,?)"
	t := time.Now()
	_, err := db.Exec(sqlStr, sender, receiver, strconv.Itoa(t.Hour())+strconv.Itoa(t.Minute())+strconv.Itoa(t.Second()), message)
	if err != nil {
		fmt.Println("receiver not exist!")
		return
	}
	fmt.Println("message sent successfully!")
}
