package main

import (
    "log"
    "net/http"
    "os"

    "fmt"
    "github.com/99designs/gqlgen/handler"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    go_orders_graphql_api "github.com/soberkoder/go-orders-graphql-api"
    "github.com/soberkoder/go-orders-graphql-api/models"
)

const defaultPort = "63932"

var db *gorm.DB

func initDB() {
    var err error
    dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
    db, err = gorm.Open("mysql", dataSourceName)

    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }

    db.LogMode(true)

    // Create the database. This is a one-time step.
    // Comment out if running multiple times - You may see an error otherwise
    db.Exec("CREATE DATABASE test_db")
    db.Exec("USE test_db")

    // Migration to create tables for Order and Item schema
    db.AutoMigrate(&models.Order{}, &models.Item{})
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    initDB()
    http.Handle("/", handler.Playground("GraphQL playground", "/query"))
    http.Handle("/query", handler.GraphQL(go_orders_graphql_api.NewExecutableSchema(go_orders_graphql_api.Config{Resolvers: &go_orders_graphql_api.Resolver{
        DB: db,
    }})))

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
