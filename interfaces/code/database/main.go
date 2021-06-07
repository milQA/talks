package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"code/database/repository"
	"code/database/structs"
)

const connDSN = "postgres://postgres:postgres@localhost/postgres?sslmode=disable"

func main() {
	conn, err := sql.Open("postgres", connDSN)
	if err != nil {
		log.Fatalf("cannot connection to database: %v", err)
	}
	defer conn.Close()

	if err = conn.Ping(); err != nil {
		log.Fatalf("cannot ping database: %v", err)
	}

	repo := repository.NewUserRepository(conn)

	user := &structs.User{
		Name:  "nur",
		Email: "nurke.ru@gmail.com",
		Address: &structs.UserAddress{
			Valid:  true,
			City:   "foo",
			Street: "boo",
			Home:   42,
			Flat:   42,
		},
	}

	if err = repo.SaveUser(user); err != nil {
		log.Fatalf("cannot save user: %s", err)
	}

	log.Printf("saved userID = %d\n", user.ID)

	if user, err = repo.GetUserByID(user.ID); err != nil {
		log.Fatalf("cannot get user: %s", err)
	}

	log.Printf("geted user address: %+v\n", user.Address)
}
