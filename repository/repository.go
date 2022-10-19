package repository

import (
	"database/sql"
	"fmt"
	"go-rest-api/models"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

type UserRepositoryInterface interface {
	GetUserById(id int) (user models.User, err error)
	GetAllUsers() (users []models.User, err error)
	AddUser(user models.User) (addedUser models.User, err error)
	DeleteUserById(id int) (err error)
}

type UserRepository struct {
	Db *sql.DB
}

func ConnectToDatabase() (*sql.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}

func (repo *UserRepository) GetUserById(id int) (user models.User, err error) {

	err = repo.Db.QueryRow("SELECT id,name,age FROM people WHERE id=$1;", id).Scan(&user.Id, &user.Name, &user.Age)
	return
}

func (repo *UserRepository) GetAllUsers() (users []models.User, err error) {

	userArray, err := repo.Db.Query("SELECT * FROM people;")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer userArray.Close()

	for userArray.Next() {
		var user models.User
		err = userArray.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Println(err)
			return
		}
		users = append(users, user)
	}
	return
}

func (repo *UserRepository) AddUser(user models.User) (returnedUser models.User, err error) {

	err = repo.Db.QueryRow("INSERT INTO people (name, age) VALUES ($1, $2) RETURNING id, name, age;", user.Name, user.Age).Scan(&returnedUser.Id, &returnedUser.Name, &returnedUser.Age)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (repo *UserRepository) DeleteUserById(id int) (err error) {

	_, err = repo.GetUserById(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = repo.Db.Exec("DELETE FROM people WHERE id=$1;", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
