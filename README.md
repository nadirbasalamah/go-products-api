# go-products-api

REST API implementation to manage products data. Implemeted with Go using fiber.

## How To Use

1. There are two branches in this project.

- choose `master` for implementation with MySQL database.
- choose `local-storage` for implementation using local storage.

2. Clone the repository into your local machine.

3. If `master` is cloned. Create a new database called `goproducts` with this command.

```sql
CREATE DATABASE goproducts;
```

4. Make sure the MySQL server is online then run the application with this command.

```
go run server.go
```

5. If `local-storage` is cloned. Run the application with this command.

```
go run server.go
```
