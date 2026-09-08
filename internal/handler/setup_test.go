package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"srebootcamp/internal/db"
	"srebootcamp/internal/handler"
	"srebootcamp/internal/model"
	"srebootcamp/internal/router"

	"github.com/stretchr/testify/require"
)

var (
	testHandler *handler.Handler
	testRouter  http.Handler
)

func TestMain(m *testing.M) {
	c := model.Config{
		DBHost:  "localhost",
		DBPort:  5432,
		DBUser:  "postgres",
		DBPass:  "dummy",
		DBName:  "postgres",
		SSLMode: "disable",
	}

	testDB := db.InitDB(c)

	testHandler = &handler.Handler{DB: testDB}
	testRouter = router.SetupRouter(testHandler)

	code := m.Run()
	testDB.Close()
	os.Exit(code)
}

func resetTable(t *testing.T) {
	t.Helper()
	_, err := testHandler.DB.Exec("TRUNCATE students_data RESTART IDENTITY;")
	require.NoError(t, err)
}

func seedStudent(t *testing.T, s model.Student) int {
	t.Helper()
	var id int
	err := testHandler.DB.QueryRow(
		"INSERT INTO students_data (first_name, last_name, class, gender) VALUES ($1, $2, $3, $4) RETURNING student_id;",
		s.FirstName, s.LastName, s.Class, s.Gender,
	).Scan(&id)
	require.NoError(t, err)
	return id
}

func sendRequest(t *testing.T, method, path string, body any) *httptest.ResponseRecorder {
	t.Helper()
	var raw string
	if body != nil {
		b, err := json.Marshal(body)
		require.NoError(t, err)
		raw = string(b)
	}

	req := httptest.NewRequest(method, path, strings.NewReader(raw))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	return w
}
