// Command dbtool: admin-zero 开发辅助工具。
//   dbtool truncate            重建 admin_zero 库（DROP + CREATE）
//   dbtool captcha <id>        读取 Redis 中的图形验证码（CAPTCHA:<id>）
//   dbtool captcha <email/phone> 同样可读邮箱/短信验证码（同前缀）
//   dbtool del <key...>        删除 Redis 键（如 initDatabase:member 守卫）
package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

const (
	dsnBase = "root:@tcp(127.0.0.1:3306)/"
	dbName  = "admin_zero"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: dbtool truncate | captcha <captchaId>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "truncate":
		db, err := sql.Open("mysql", dsnBase+"mysql?charset=utf8mb4&multiStatements=true")
		if err != nil {
			fmt.Println("open:", err)
			os.Exit(1)
		}
		defer db.Close()
		if _, err = db.Exec("DROP DATABASE IF EXISTS " + dbName); err != nil {
			fmt.Println("drop:", err)
			os.Exit(1)
		}
		if _, err = db.Exec("CREATE DATABASE " + dbName + " DEFAULT CHARACTER SET utf8mb4"); err != nil {
			fmt.Println("create:", err)
			os.Exit(1)
		}
		fmt.Println("admin_zero recreated")

	case "captcha":
		if len(os.Args) < 3 {
			fmt.Println("usage: dbtool captcha <captchaId>")
			os.Exit(1)
		}
		rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
		defer rdb.Close()
		v, err := rdb.Get(context.Background(), "CAPTCHA:"+os.Args[2]).Result()
		if err != nil {
			fmt.Println("get:", err)
			os.Exit(1)
		}
		fmt.Println(v)

	case "del":
		if len(os.Args) < 3 {
			fmt.Println("usage: dbtool del <key...>")
			os.Exit(1)
		}
		rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
		defer rdb.Close()
		n, err := rdb.Del(context.Background(), os.Args[2:]...).Result()
		if err != nil {
			fmt.Println("del:", err)
			os.Exit(1)
		}
		fmt.Println("deleted", n, "key(s)")
	}
}
