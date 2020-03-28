package tests

import (
	"net/http"
	"net/http/httptest"
	"pando/server/router"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gavv/httpexpect"
	"golang.org/x/crypto/bcrypt"
)

// Test endpoint login
func TestLogin(t *testing.T) {
	Setup()
	defer FinishTest()

	// Set mock
	clientsecret := "clienttestsecret"
	userpass := "usertestpass"
	encryptclientsecret, _ := bcrypt.GenerateFromPassword([]byte(clientsecret), bcrypt.DefaultCost)
	encryptuserpass, _ := bcrypt.GenerateFromPassword([]byte(userpass), bcrypt.DefaultCost)

	clientRows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "key", "secret"}).
		AddRow(1, time.Now(), time.Now(), nil, "clienttestkey", string(encryptclientsecret))
	userRows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "username", "password"}).
		AddRow(1, time.Now(), time.Now(), nil, "usertest", string(encryptuserpass))

	mock.ExpectQuery(`SELECT (.+) FROM \"clients\"`).
		WillReturnRows(clientRows)
	mock.ExpectQuery(`SELECT (.+) FROM \"users\"`).
		WillReturnRows(userRows)

	// Begin testing
	BeginTest()

	api := router.NewRouter()
	server := httptest.NewServer(api)
	defer server.Close()
	e := httpexpect.New(t, server.URL)
	htest := e.Builder(func(req *httpexpect.Request) {
		req.WithBasicAuth("clienttestkey", clientsecret)
	})

	htest.POST("/client/login").
		WithJSON(map[string]interface{}{
			"username": "usertest",
			"password": userpass,
		}).
		Expect().Status(http.StatusOK).JSON().Object().ValueEqual("code", "0000")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("error : %v", err)
	}
}
