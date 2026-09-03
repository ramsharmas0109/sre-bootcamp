package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"srebootcamp/internal/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetAllStudents(c *gin.Context) {
	var students []model.Student

	rows, err := h.DB.Query("SELECT * FROM students_data;")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		var student model.Student
		// Pass pointers to rows.Scan matching the SELECT column order
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Class, &student.Gender)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	c.JSON(http.StatusOK, students)

}

func (h *Handler) GetStudentByID(c *gin.Context) {
	var students []model.Student

	id := c.Param("id")
	tempID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := h.DB.Query("SELECT * FROM students_data WHERE student_id = $1;", tempID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Class, &student.Gender)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()
	if len(students) > 0 {
		c.JSON(http.StatusOK, students)
	} else {
		c.JSON(http.StatusNotFound, nil)
	}
}

func (h *Handler) CreateStudent(c *gin.Context) {
	var student model.Student

	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := h.DB.Query(
		`INSERT INTO students_data (first_name, last_name, class, gender)
    	 VALUES ($1, $2, $3, $4)`, student.FirstName, student.LastName, student.Class, student.Gender,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()
	c.JSON(http.StatusCreated, gin.H{"insert": true})
}

func (h *Handler) UpdateStudent(c *gin.Context) {
	var student, studentFromDB model.Student

	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	tempID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := h.DB.Query("SELECT * FROM students_data WHERE student_id = $1;", tempID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		err = rows.Scan(&studentFromDB.ID, &studentFromDB.FirstName, &studentFromDB.LastName, &studentFromDB.Class, &studentFromDB.Gender)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if student.Class != 0 {
		studentFromDB.Class = student.Class
	}

	if student.FirstName != "" {
		studentFromDB.FirstName = student.FirstName
	}

	if student.LastName != "" {
		studentFromDB.LastName = student.LastName
	}

	if student.Gender != "" {
		studentFromDB.Gender = student.Gender
	}

	rows, err = h.DB.Query("UPDATE students_data SET class = $4, gender = $5, last_name = $3, first_name = $2 WHERE student_id = $1;", tempID, studentFromDB.FirstName, studentFromDB.LastName, studentFromDB.Class, studentFromDB.Gender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{"update": true})

}

func (h *Handler) DeleteStudentByID(c *gin.Context) {
	id := c.Param("id")
	tempID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := h.DB.Query("DELETE FROM students_data WHERE student_id = $1;", tempID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{"delete": true})
}
