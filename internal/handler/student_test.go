package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"srebootcamp/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type checkFunc func(t *testing.T, w *httptest.ResponseRecorder, id int)

func checkFoundStudent(t *testing.T, w *httptest.ResponseRecorder, id int) {
	var got model.Student
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &got))
	assert.Equal(t, "Laxman", got.FirstName)
	assert.Equal(t, "Sharma", got.LastName)
	assert.Equal(t, 10, got.Class)
	assert.Equal(t, "M", got.Gender)
	assert.Equal(t, id, got.ID)
}

func checkStudentInserted(t *testing.T, w *httptest.ResponseRecorder, _ int) {
	var resp map[string]any
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.EqualValues(t, 1, resp["student_insert"])
}

func checkClassUpdated(t *testing.T, w *httptest.ResponseRecorder, id int) {
	var resp map[string]any
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.EqualValues(t, 1, resp["rows_update"])

	got := sendRequest(t, http.MethodGet, fmt.Sprintf("/api/v1/students/%d", id), nil)
	var student model.Student
	require.NoError(t, json.Unmarshal(got.Body.Bytes(), &student))
	assert.Equal(t, 12, student.Class)
}

func checkStudentGone(t *testing.T, w *httptest.ResponseRecorder, id int) {
	notFound := sendRequest(t, http.MethodGet, fmt.Sprintf("/api/v1/students/%d", id), nil)
	assert.Equal(t, http.StatusNotFound, notFound.Code)
}

func TestGetAllStudents(t *testing.T) {
	resetTable(t)
	seedStudent(t, model.Student{FirstName: "Laxman", LastName: "Sharma", Class: 10, Gender: "M"})
	seedStudent(t, model.Student{FirstName: "Meena", LastName: "Kumari", Class: 8, Gender: "F"})

	w := sendRequest(t, http.MethodGet, "/api/v1/students", nil)
	assert.Equal(t, http.StatusOK, w.Code)

	var students []model.Student
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &students))
	assert.Len(t, students, 2)
}

func TestGetStudentByID(t *testing.T) {
	resetTable(t)
	id := seedStudent(t, model.Student{FirstName: "Laxman", LastName: "Sharma", Class: 10, Gender: "M"})

	cases := []struct {
		name       string
		path       string
		wantStatus int
		check      checkFunc
	}{
		{name: "existing id", path: fmt.Sprintf("/api/v1/students/%d", id), wantStatus: http.StatusOK, check: checkFoundStudent},
		{name: "nonexistent id", path: "/api/v1/students/999999", wantStatus: http.StatusNotFound},
		{name: "non-numeric id", path: "/api/v1/students/abc", wantStatus: http.StatusBadRequest},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := sendRequest(t, http.MethodGet, tc.path, nil)
			assert.Equal(t, tc.wantStatus, w.Code)
			if tc.check != nil {
				tc.check(t, w, id)
			}
		})
	}
}

func TestCreateStudent(t *testing.T) {
	resetTable(t)

	cases := []struct {
		name       string
		body       string
		wantStatus int
		check      checkFunc
	}{
		{name: "valid student", body: `{"first_name":"Laxman","last_name":"Sharma","class":10,"gender":"M"}`, wantStatus: http.StatusCreated, check: checkStudentInserted},
		{name: "malformed body", body: "{not valid json", wantStatus: http.StatusBadRequest},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/students", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			testRouter.ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatus, w.Code)
			if tc.check != nil {
				tc.check(t, w, 0)
			}
		})
	}
}

func TestUpdateStudent(t *testing.T) {
	resetTable(t)
	id := seedStudent(t, model.Student{FirstName: "Laxman", LastName: "Sharma", Class: 10, Gender: "M"})

	cases := []struct {
		name       string
		path       string
		body       map[string]any
		wantStatus int
		check      checkFunc
	}{
		{name: "existing id", path: fmt.Sprintf("/api/v1/students/%d", id), body: map[string]any{"class": 12}, wantStatus: http.StatusOK, check: checkClassUpdated},
		{name: "nonexistent id", path: "/api/v1/students/999999", body: map[string]any{"class": 5}, wantStatus: http.StatusNotFound},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := sendRequest(t, http.MethodPut, tc.path, tc.body)
			assert.Equal(t, tc.wantStatus, w.Code)
			if tc.check != nil {
				tc.check(t, w, id)
			}
		})
	}
}

func TestDeleteStudentByID(t *testing.T) {
	resetTable(t)
	id := seedStudent(t, model.Student{FirstName: "Laxman", LastName: "Sharma", Class: 10, Gender: "M"})

	cases := []struct {
		name       string
		path       string
		wantStatus int
		check      checkFunc
	}{
		{name: "existing id", path: fmt.Sprintf("/api/v1/students/%d", id), wantStatus: http.StatusOK, check: checkStudentGone},
		{name: "nonexistent id", path: "/api/v1/students/999999", wantStatus: http.StatusNotFound},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := sendRequest(t, http.MethodDelete, tc.path, nil)
			assert.Equal(t, tc.wantStatus, w.Code)
			if tc.check != nil {
				tc.check(t, w, id)
			}
		})
	}
}
