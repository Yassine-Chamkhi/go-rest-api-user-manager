package repository

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"target/onboarding-assignment/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(fmt.Sprintf("an error '%s' was not expected when opening a stub database connection", err))
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Id", "Name", "Age"})

	for i := 1; i < 11; i++ {
		rand.Seed(time.Now().UnixNano())
		rows.AddRow(i, fmt.Sprintf("Username%d", i), rand.Intn(90-18+1)+18)
	}
	query := `SELECT \* FROM people;`

	mock.ExpectQuery(query).WillReturnRows(rows)

	repo := UserRepository{Db: db}
	users, err := repo.GetAllUsers()

	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 10)
}

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(fmt.Sprintf("an error '%s' was not expected when opening a stub database connection", err))
	}
	defer db.Close()

	insertedUser := models.User{
		Id:   2,
		Name: "Username2",
		Age:  23,
	}

	rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(insertedUser.Id, insertedUser.Name, insertedUser.Age)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id,name,age FROM people WHERE id=$1;`)).WithArgs(2).WillReturnRows(rows)

	repo := UserRepository{Db: db}
	user, err := repo.GetUserById(2)

	assert.NotEmpty(t, user)
	assert.NoError(t, err)
	assert.Equal(t, insertedUser, user)
}

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(fmt.Sprintf("an error '%s' was not expected when opening a stub database connection", err))
	}
	defer db.Close()

	insertedUser := models.User{
		Id:   1,
		Name: "Username1",
		Age:  23,
	}

	rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(insertedUser.Id, insertedUser.Name, insertedUser.Age)
	query := regexp.QuoteMeta(`INSERT INTO people (name, age) VALUES ($1, $2) RETURNING id, name, age;`)
	mock.ExpectQuery(query).WithArgs(insertedUser.Name, insertedUser.Age).WillReturnRows(rows)

	repo := UserRepository{Db: db}

	user, err := repo.AddUser(insertedUser)
	assert.Equal(t, insertedUser, user)
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(fmt.Sprintf("an error '%s' was not expected when opening a stub database connection", err))
	}
	defer db.Close()

	insertedUser := models.User{
		Id:   1,
		Name: "Username1",
		Age:  23,
	}

	rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(insertedUser.Id, insertedUser.Name, insertedUser.Age)

	repo := UserRepository{Db: db}
	query1 := regexp.QuoteMeta(`SELECT id,name,age FROM people WHERE id=$1;`)
	query2 := regexp.QuoteMeta(`DELETE FROM people WHERE id=$1;`)

	mock.ExpectQuery(query1).WithArgs(1).WillReturnRows(rows)
	mock.ExpectExec(query2).WithArgs(1).WillReturnResult(driver.ResultNoRows).WillReturnError(nil)

	err = repo.DeleteUserById(1)
	assert.NoError(t, err)

	mock.ExpectQuery(query1).WithArgs(0).WillReturnError(errors.New("Id must be an integer value superior to 0"))
	err = repo.DeleteUserById(0)
	assert.EqualError(t, err, "Id must be an integer value superior to 0")

	mock.ExpectQuery(query1).WithArgs(15).WillReturnError(errors.New("sql: no rows in result set"))

	err = repo.DeleteUserById(15)
	fmt.Println(err)
	assert.EqualError(t, err, "sql: no rows in result set")
}
