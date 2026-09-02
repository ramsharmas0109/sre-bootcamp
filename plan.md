Build an API server.

API
- Add a new student.
- Get all students.
- Get a student with an ID.
- Update existing student information.
- Delete a student record.

/api/v1/students -> to manage students resources
/healthcheck -> response 200 ok 

HTTP verb 
- C (Create) - Add a new student - POST /api/v1/students + body (student data to be added)
- U (Update) - Update existing student information - PUT /api/v1/students/:id + body (field name and updated value)
- R (Read)
    - Get all students - GET /api/v1/students 
    - Get a student with an ID - GET /api/v1/students/:id
- D (Delete) - Delete a student record - DELETE /api/v1/students/:id

SERVER
- Host - localhost
- Port - any port (8000 to 65535)


start server -> handler (/healthcheck)
request sent on /healthcheck -> gin creates a context for this request -> gin passes this context to handler function -> handler function responds with ok.


rg.GET("/students", getAllStudents)
rg.GET("/students/:id", getStudentByID)
rg.POST("/students", createStudent)
rg.PUT("/students/:id", updateStudent)
rg.DELETE("/students/:id", deleteStudentByID)


Database 101
Here we will use postgresql which is a relational DB. In this we will create logical DB which has public schema and this public schema consists of tables and these tables will have rows where we will store the data.


curl -X GET localhost:8080/api/v1/students
GET "/students" -> getAllStudents -> SELECT * FROM students_data;

curl -X GET localhost:8080/api/v1/students/:id
GET "/students/:id" -> getStudentByID -> SELECT * FROM students_data WHERE student_id = id;


curl -X POST localhost:8080/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{"first_name":"Akshay","last_name":"Sharma","class":10,"gender":"m"}'
POST "/students" -> createStudent -> INSERT INTO students_data (first_name, last_name, class, gender) VALUES ('Akshay', 'Sharma', 10, 'm');


curl -X PUT localhost:8080/api/v1/students/:id \
  -H "Content-Type: application/json" \
  -d '{"last_name":"Kumar","class":12}'
PUT "/students/:id" -> updateStudent -> UPDATE students_data SET class = 12, last_name = 'Kumar' WHERE student_id = id;


curl -X DELETE localhost:8080/api/v1/students/:id
DELETE "/students/:id" -> deleteStudentByID -> DELETE FROM students_data WHERE student_id = id;


JSON -> student struct -> insert DB
GET DB -> student struct -> JSON

Actionable items for DB
1. Select and configure drivers for postgresql 
2. Connect to postgresql
3. Create a student struct
4. Implement getAllStudents 
5. Implement getStudentByID
6. Implement createStudent

we need a struct to get and put student data from DB and to DB. 

DB Driver using which connect to DB
And run some commands
When will these sql commands run? 
When we make request to certain endpoint

psql -h localhost -p 5432 -d students -U postgres



main.go -> server setup, db conn, handler define, handler Db querier. 
main -> handler.go, router.go, db.go, model.go



main -> router pkg (SetupRouter) -> handler pkg (diff handlers) -> db & models pkg with different funcs
