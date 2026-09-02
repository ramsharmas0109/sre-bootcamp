package model

type Student struct {
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	ID        int    `json:"id" db:"student_id"`
	Class     int    `json:"class" db:"class"`
	Gender    string `json:"gender" db:"gender"`
}

type Config struct {
	DBHost  string
	DBPass  string
	DBPort  int
	DBName  string
	DBUser  string
	SSLMode string
	Port    string
}
