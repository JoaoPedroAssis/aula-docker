package main

import (
    // "database/sql"
    // "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "postgres"
)

func main() {
    // Configuração de conexão com o banco de dados PostgreSQL
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    // Crie uma instância do framework Gin
    r := gin.Default()

    // Rota "Hello World"
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    // Rota para testar a conexão com o banco de dados
    // r.GET("/test-db-connection", func(c *gin.Context) {
    //     c.JSON(http.StatusOK, gin.H{"message": "Conexão com o banco de dados estabelecida com sucesso!"})
    // })

    // Inicie o servidor na porta 8080
    r.Run(":8080")
}
