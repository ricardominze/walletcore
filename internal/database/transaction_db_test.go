package database

import (
	"database/sql"
	"testing"

	"github.com/ricardominze/ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client1       *entity.Client
	client2       *entity.Client
	account1      *entity.Account
	account2      *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")
	s.transactionDB = NewTransactionDB(db)
	s.client1, err = entity.NewClient("John1", "j1@j.com")
	s.Nil(err)
	s.client2, err = entity.NewClient("John2", "j2@j.com")
	s.Nil(err)
	//creating accounts
	s.account1 = entity.NewAccount(s.client1)
	s.account1.Balance = 1000
	s.account2 = entity.NewAccount(s.client2)
	s.account2.Balance = 1000
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TALE clients")
	s.db.Exec("DROP TALE accounts")
	s.db.Exec("DROP TALE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, &TransactionDBTestSuite{})
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.account1, s.account2, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
