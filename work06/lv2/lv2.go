package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

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
			forget()
		case 4:
			return
		default:
			fmt.Println("please input again!")
		}
	}
	//query("xiaote33", "123456")
	//register()
}

func forget() {
	u := ""
	q := ""
	a := ""
	resp := ""
	p := ""
	fmt.Println("username:")
	fmt.Scan(&u)
	sqlStr := "select question,answer from userdata where username=?"
	row := db.QueryRow(sqlStr, u)
	row.Scan(&q, &a)
	fmt.Println("Q:", q)
	fmt.Print("A:")
	fmt.Scan(&resp)
	if resp == a {
		fmt.Println("pass!")
		fmt.Print("new password:")
		fmt.Scan(&p)
		sqlStr := "update userdata set password=? where username=?"
		_, err := db.Exec(sqlStr, p, u)
		if err != nil {
			fmt.Println("Exec err: ", err)
			return
		}
		fmt.Println("password updated successfully!")
	} else {
		fmt.Println("NOT correct!")
	}
}

func showMenu() {
	fmt.Println("**********")
	fmt.Println("* 1.注册")
	fmt.Println("* 2.登陆")
	fmt.Println("* 3.忘记密码")
	fmt.Println("* 4.退出")
	fmt.Println("**********")
}

func initDB() (err error) {
	dsn := "root:root@tcp(localhost:3306)/lv2database"
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
	q := ""
	a := ""
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
			fmt.Println("set a Q&A in case of forgetting your password")
			fmt.Print("question:")
			fmt.Scan(&q)
			fmt.Print("answer:")
			fmt.Scan(&a)
			sqlStr2 := "insert into userdata values(?,?,?,?)"
			_, err := db.Exec(sqlStr2, u, p, q, a)
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
	fmt.Println(u, "login ok！")
	return

}
