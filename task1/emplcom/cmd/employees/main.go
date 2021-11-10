package main

import (
	"database/sql"
	"emplcom/internal/api/handler"
	"emplcom/internal/app/employees"
	"emplcom/internal/app/environment"
	"emplcom/internal/repositories/pg"
	"fmt"
	"log"
)

func main() {

	con, err := sql.Open("", "")
	if err != nil {
		log.Fatal(err)
	}

	pg := pg.NewPG(con)
	emp := employees.NewEmployees(pg)
	env := environment.NewEnvironments(pg)

	h := handler.NewHandlers(emp, env)

	fmt.Println(h)
	//Ну и дальше h передается http серверу или реализации запуска http сервера

}