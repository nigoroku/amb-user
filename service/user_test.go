package service

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"local.packages/models"
)

var (
	userDBTypes            = map[string]string{`UserID`: `int`, `AccountName`: `varchar`, `Email`: `varchar`, `Password`: `varchar`, `Role`: `tinyint`, `LastLoginAt`: `timestamp`, `CreatedBy`: `int`, `CreatedAt`: `timestamp`, `ModifiedBy`: `int`, `ModifiedAt`: `timestamp`, `AccountImg`: `blob`}
	_                      = bytes.MinRead
	userColumnsWithDefault = []string{"user_id"}
)

func Init() {
	// DB接続
	db, err := sql.Open("mysql", "moizumi:base0210@tcp(localhost:3306)/ambitious_test?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)
}

func MustTx(transactor boil.ContextTransactor, err error) boil.ContextTransactor {
	if err != nil {
		panic(fmt.Sprintf("Cannot create a transactor: %s", err))
	}
	return transactor
}

func TestAddUser(t *testing.T) {
	// DB接続
	Init()

	t.Parallel()
	seed := randomize.NewSeed()
	var err error
	o := &models.User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	o.Email = "test@test.test"

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	userService := &UserService{ctx, tx}
	got, err2 := userService.AddUser(o)

	if err2 != nil {
		t.Error(err2)
		errorfHelper(t, o.Email, got.Email)
	}

	tx.Commit()

	got2, err3 := userService.FindUserById(got.UserID)

	if err3 != nil {
		t.Error(err3)
		errorfHelper(t, o.Email, got2.Email)
	}

	assert.Equal(t, o.Email, got2.Email)
}

func TestFindRegisteredUser(t *testing.T) {
	// DB接続
	Init()

	t.Parallel()
	seed := randomize.NewSeed()
	var err error
	o := &models.User{}
	if err = randomize.Struct(seed, o, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	email := "test2@test.com"
	password := "password"

	o.Email = email
	o.Password = password

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	userService := &UserService{ctx, tx}
	got, err2 := userService.AddUser(o)

	if err2 != nil {
		t.Error(err2)
		errorfHelper(t, o.Email, got.Email)
	}

	tx.Commit()

	got2, err3 := userService.FindRegisteredUser(email, password)

	if err3 != nil {
		t.Error(err3)
		errorfHelper(t, o.Email, got2.Email)
	}

	assert.NotNil(t, got2)
}

func errorfHelper(tb testing.TB, want string, got string) {
	tb.Helper()
	tb.Errorf("want = %v, got = %v", want, got)
}
