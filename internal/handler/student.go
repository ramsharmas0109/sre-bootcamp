package handler

import (
	"log"
	"net/http"
	"strconv"

	"srebootcamp/internal/db"
	"srebootcamp/internal/model"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var students []model.Student

	db := db.InitDB()

	rows, err := db.Query("SELECT * FROM students_data;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var student model.Student
		// Pass pointers to rows.Scan matching the SELECT column order
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Class, &student.Gender)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}

	defer rows.Close()

	c.JSON(http.StatusOK, students)

}

func GetStudentByID(c *gin.Context) {
	var students []model.Student

	db := db.InitDB()
	id := c.Param("id")
	temp, _ := strconv.Atoi(id)

	rows, err := db.Query("SELECT * FROM students_data WHERE student_id = $1;", temp)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var student model.Student
		// Pass pointers to rows.Scan matching the SELECT column order
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Class, &student.Gender)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}

	defer rows.Close()

	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student model.Student

	db := db.InitDB()

	err := c.ShouldBindJSON(&student)
	if err != nil {
		//log.Fatal(err)
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := db.Query(
		`INSERT INTO students_data (first_name, last_name, class, gender)
    	 VALUES ($1, $2, $3, $4)`, student.FirstName, student.LastName, student.Class, student.Gender,
	)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{"insert": true})
}

func UpdateStudent(c *gin.Context) {
	var student, studentFromDB model.Student
	db := db.InitDB()

	err := c.ShouldBindJSON(&student)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	temp, _ := strconv.Atoi(id)

	rows, err := db.Query("SELECT * FROM students_data WHERE student_id = $1;", temp)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&studentFromDB.ID, &studentFromDB.FirstName, &studentFromDB.LastName, &studentFromDB.Class, &studentFromDB.Gender)
		if err != nil {
			log.Fatal(err)
		}
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

	rows, err = db.Query("UPDATE students_data SET class = $4, gender = $5, last_name = $3, first_name = $2 WHERE student_id = $1;", temp, studentFromDB.FirstName, studentFromDB.LastName, studentFromDB.Class, studentFromDB.Gender)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{"update": true})

}

func DeleteStudentByID(c *gin.Context) {
	db := db.InitDB()
	id := c.Param("id")
	temp, _ := strconv.Atoi(id)
	rows, err := db.Query("DELETE FROM students_data WHERE student_id = $1;", temp)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{"delete": true})
}
