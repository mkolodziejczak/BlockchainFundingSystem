package db

import (
	"crypto/md5"
	"database/sql"
	"fundingSystem/common"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

type Database struct {
	mu sync.Mutex
	db *sql.DB
}

const file string = "resources\\db\\projects.db"

var connection *Database

func GetConnection() (*Database, error) {

	if connection != nil {
		return connection, nil
	}

	db, err := sql.Open("sqlite3", file)

	if err != nil {
		return nil, err
	}

	connection = &Database{
		db: db,
	}
	return connection, nil
}

func CreateUser(dto common.UserDto) (int, error) {
	db, err := GetConnection()
	pass := []byte(dto.Password)
	res, err := db.db.Exec("INSERT INTO users VALUES(NULL,?,?,?,?,?,?);",
		dto.Login, md5.Sum(pass), dto.FirstName, dto.LastName, dto.Email, dto.WalletAddress)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func LoginUser(dto common.LoginDto) (int, error) {
	db, err := GetConnection()
	pass := []byte(dto.Password)
	row := db.db.QueryRow("SELECT id FROM users WHERE login=? AND  password=?", dto.User, md5.Sum(pass))

	if err != nil {
		return 0, err
	}
	user := common.User{}
	if err = row.Scan(&user.Id); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return 0, sql.ErrNoRows
	}
	return user.Id, err

}

func GetUserWalletId(userId int) (string, error) {
	db, err := GetConnection()

	row := db.db.QueryRow("SELECT wallet_id FROM users WHERE id=?", userId)

	if err != nil {
		return "", err
	}
	user := common.User{}
	if err = row.Scan(&user.WalletAddress); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return "", sql.ErrNoRows
	}
	return user.WalletAddress, err

}

func CreateProject(dto common.ProjectDto, userId int, contractAddress string) (int, error) {
	db, err := GetConnection()
	res, err := db.db.Exec("INSERT INTO projects VALUES(NULL,?,?,?,?,?,?,?,?);",
		userId, dto.Title, dto.Description, dto.GoalBacking, dto.Milestone1Date, dto.Milestone2Date, dto.Milestone3Date, contractAddress)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}
