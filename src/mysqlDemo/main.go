package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
var(
	name string
	userid,age int
	Db *sql.DB
)

func init()  {
	db,err := sql.Open("mysql","root:root@(127.0.0.1:3306)/demo?charset=utf8")
	if err != nil{
		str := fmt.Sprintf("mysql open err:",err)
		panic(str)
	}

	err = db.Ping()
	if err != nil {
		str := fmt.Sprintf("mysql ping err:",err)
		panic(str)
	}
	Db = db

}

func main() {
    //add1()
    //add2()
    //read1()
    read2()
}

//单行数据
func read1()  {



    rows := Db.QueryRow("select * from user where userid = 5")
    defer Db.Close()

    rows.Scan(&userid,&name,&age)

    fmt.Println(userid,name,age)
}

//多行数据
func read2()  {
   rows,_ := Db.Query("select * from user")
   defer Db.Close()

	for rows.Next() {
		rows.Scan(&userid,&name,&age)
		fmt.Println(userid,name,age)
	}


}

func add1()  {
	result,err := Db.Exec("insert into user(name,age)values('wangwu',17)")
    defer Db.Close()
	if err != nil {
		fmt.Println("mysql insert into err:",err)
		return
	}
	fmt.Println(result.LastInsertId())    //返回添加的ID
	n,_ := result.RowsAffected()          //返回收影响的记录
	fmt.Println(n)
}

func add2()  {

	stmt,err := Db.Prepare("insert into user(name,age)values(?,?)")
	defer Db.Close()
	if err != nil {
		fmt.Println("mysql insert yuchuli err:",err)
		return
	}

	result,err := stmt.Exec("zhangsan","12")
	if err != nil {
		fmt.Println("mysql insert into err:",err)
		return
	}
	fmt.Println(result.LastInsertId())    //返回添加的ID
	n,_ := result.RowsAffected()          //返回收影响的记录
	fmt.Println(n)
}