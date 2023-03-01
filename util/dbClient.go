/**
* @description :
* @author : Jarick
* @Date : 2022-11-25
* @Url : http://CloudWebOps
 */

// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.1.md

package util

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	_ "github.com/lib/pq"           // https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.4.md
// 	_ "github.com/mattn/go-sqlite3" // https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.3.md
// )

// func SQLite() {
// 	db, err := sql.Open("sqlite3", "./foo.db")
// 	checkErr(err)

// 	//插入数据
// 	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
// 	checkErr(err)

// 	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
// 	checkErr(err)

// 	id, err := res.LastInsertId()
// 	checkErr(err)

// 	fmt.Println(id)
// 	//更新数据
// 	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
// 	checkErr(err)

// 	res, err = stmt.Exec("astaxieupdate", id)
// 	checkErr(err)

// 	affect, err := res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println(affect)

// 	//查询数据
// 	rows, err := db.Query("SELECT * FROM userinfo")
// 	checkErr(err)

// 	for rows.Next() {
// 		var uid int
// 		var username string
// 		var department string
// 		var created time.Time
// 		err = rows.Scan(&uid, &username, &department, &created)
// 		checkErr(err)
// 		fmt.Println(uid)
// 		fmt.Println(username)
// 		fmt.Println(department)
// 		fmt.Println(created)
// 	}

// 	//删除数据
// 	stmt, err = db.Prepare("delete from userinfo where uid=?")
// 	checkErr(err)

// 	res, err = stmt.Exec(id)
// 	checkErr(err)

// 	affect, err = res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println(affect)

// 	db.Close()
// }

// func pgSQL() {
// 	db, err := sql.Open("postgres", "user=astaxie password=astaxie dbname=test sslmode=disable")
// 	checkErr(err)

// 	//插入数据
// 	stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
// 	checkErr(err)

// 	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
// 	checkErr(err)

// 	//pg不支持这个函数，因为他没有类似MySQL的自增ID
// 	// id, err := res.LastInsertId()
// 	// checkErr(err)
// 	// fmt.Println(id)

// 	var lastInsertId int
// 	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
// 	checkErr(err)
// 	fmt.Println("最后插入id =", lastInsertId)

// 	//更新数据
// 	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
// 	checkErr(err)

// 	res, err = stmt.Exec("astaxieupdate", 1)
// 	checkErr(err)

// 	affect, err := res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println(affect)

// 	//查询数据
// 	rows, err := db.Query("SELECT * FROM userinfo")
// 	checkErr(err)

// 	for rows.Next() {
// 		var uid int
// 		var username string
// 		var department string
// 		var created string
// 		err = rows.Scan(&uid, &username, &department, &created)
// 		checkErr(err)
// 		fmt.Println(uid)
// 		fmt.Println(username)
// 		fmt.Println(department)
// 		fmt.Println(created)
// 	}

// 	//删除数据
// 	stmt, err = db.Prepare("delete from userinfo where uid=$1")
// 	checkErr(err)

// 	res, err = stmt.Exec(1)
// 	checkErr(err)

// 	affect, err = res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println(affect)

// 	db.Close()
// }
