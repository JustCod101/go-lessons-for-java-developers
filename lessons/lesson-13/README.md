# Lesson 13: 数据库 CRUD (Database CRUD)

在 Java 中，我们经历了从 JDBC 到 MyBatis 再到 JPA/Hibernate 的演进。Go 的数据库操作哲学更接近于“增强版的 JDBC”，它提供了一个标准接口 `database/sql`，具体的数据库实现由驱动提供。

## 1. 学习目标
* 理解 `database/sql` 标准库的设计理念
* 掌握数据库驱动 (Driver) 的导入与使用
* 学会使用 `Exec`、`Query` 和 `QueryRow` 进行 CRUD 操作
* 掌握事务 (Transaction) 的处理方法
* 理解如何预防 SQL 注入
* 了解 Repository 模式在 Go 中的实现

## 2. 给 Java 开发者的类比
* **`database/sql`**: 相当于 Java 的 JDBC 接口规范。
* **`sql.DB`**: 相当于 `javax.sql.DataSource`。它内置了连接池，不是一个单一的数据库连接。
* **`rows.Scan`**: 相当于 `ResultSet.getXXX` 并手动映射到 POJO。
* **`db.Begin()`**: 相当于 `connection.setAutoCommit(false)`。
* **`?` 占位符**: 相当于 `PreparedStatement` 中的 `?`。

## 3. 核心概念

### 驱动 (Driver) 机制
Go 的 `database/sql` 包不包含任何数据库驱动。你需要通过 `import _ "github.com/go-sql-driver/mysql"` 这种方式匿名导入驱动。
* **Java 对比**: 类似于在 `pom.xml` 中引入 `mysql-connector-java`，然后在代码中 `Class.forName("com.mysql.jdbc.Driver")`。

### 连接池 (Connection Pool)
`sql.DB` 对象是并发安全的，并且内置了连接池。你不需要手动创建一个 `GenericObjectPool`。
* **Java 对比**: 相当于内置了 HikariCP 或 Druid。你可以通过 `db.SetMaxOpenConns()` 等方法配置连接池参数。

### Exec vs Query
* `Exec`: 用于不返回结果集的语句（INSERT, UPDATE, DELETE）。
* `Query`: 用于返回多行结果集的语句（SELECT）。
* `QueryRow`: 用于预期只返回一行的查询。
* **Java 对比**: 对应 `statement.executeUpdate()` 和 `statement.executeQuery()`。

### 事务处理
通过 `db.Begin()` 开启事务，返回一个 `Tx` 对象。必须显式调用 `tx.Commit()` 或 `tx.Rollback()`。
* **Java 对比**: 类似于手动管理 JDBC 事务，或者 Spring 的 `TransactionTemplate`。Go 没有 `@Transactional` 注解。

## 4. 代码示例

查看 `examples/db_example.go` 获取完整代码。

```go
// 插入数据示例
result, err := db.Exec("INSERT INTO users(name, age) VALUES(?, ?)", "Alice", 25)
if err != nil {
    log.Fatal(err)
}
lastId, _ := result.LastInsertId()
```

## 5. 常见误区
* **忘记关闭 Rows**: `rows, err := db.Query(...)` 之后必须 `defer rows.Close()`，否则会导致连接泄露。
* **错误地频繁打开/关闭 DB**: `sql.DB` 应该是一个长生命周期的对象（通常是全局单例或通过依赖注入传递），不要在每个请求中都 `sql.Open`。
* **SQL 注入**: 永远不要使用 `fmt.Sprintf` 来拼接 SQL 语句。务必使用 `?` 占位符。

## 6. 本节练习
1. 实现一个 `UpdateProductPrice` 函数，根据 ID 更新价格。
2. 实现一个带有事务的操作：同时插入一条订单记录和更新库存数量，如果其中一个失败则回滚。
3. 尝试使用 `sqlx`（一个流行的第三方库）来简化 `Scan` 过程。

## 7. 面试可能怎么问
* **Q: `sql.DB` 是线程安全的吗？**
  * A: 是的，它设计用于多个 Goroutine 并发使用，并内置了连接池。
* **Q: 如何在 Go 中处理数据库事务？**
  * A: 使用 `db.Begin()` 获取 `Tx` 对象，执行操作后根据是否有错误调用 `tx.Commit()` 或 `tx.Rollback()`。
* **Q: 为什么导入驱动时要使用 `_` 前缀？**
  * A: 这叫匿名导入。它会触发驱动包内的 `init()` 函数，将驱动注册到 `database/sql` 中，但我们不需要在代码中直接引用驱动包的任何变量或函数。

## 8. 本节总结
Go 的数据库操作非常“原始”但也非常透明。虽然没有 JPA 那样强大的 ORM 功能，但它让你对执行的每一条 SQL 都有完全的控制。在实际项目中，我们通常会配合 `sqlx` 或 `gorm` 来提高开发效率。
