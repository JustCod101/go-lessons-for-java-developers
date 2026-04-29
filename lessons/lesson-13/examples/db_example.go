package main

import (
	"database/sql"
	"fmt"
	"log"
	// _ "github.com/mattn/go-sqlite3" // 导入驱动，但不直接使用其代码
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	// 1. 打开连接 (类似于获取 DataSource)
	// 注意：这里假设已经安装了 sqlite3 驱动
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. 执行 Exec (用于 INSERT, UPDATE, DELETE)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id INTEGER PRIMARY KEY, name TEXT, price REAL)")
	if err != nil {
		log.Fatal(err)
	}

	// 3. 插入数据 (使用占位符防止 SQL 注入)
	res, err := db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", "Golang Book", 49.9)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := res.LastInsertId()
	fmt.Printf("Inserted product with ID: %d\n", id)

	// 4. 查询单行 (QueryRow)
	var p Product
	err = db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found product: %+v\n", p)

	// 5. 查询多行 (Query)
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Product: %+v\n", p)
	}
}
