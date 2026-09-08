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






--- Notes

REST API server that manages student data, the API server should support following operations. 

1) Create -> Add new students in DB. -> {first name, last name, class, gender} -> id
2) Update -> Update existing students data. -> we need to pass an id and some data that needs to be changed
3) Read   -> Get all students, Get one student by id. -> we only need to pass id, if we dont pass it, every row will be returned.
4) Delete -> Remove data from DB for that student -> we need to pass an id of target student that needs to be removed.

server -> port 
       -> router -> endpoints that server supports.
            GET /healthcheck
            GET /api/v1/students
            GET /api/v1/students/:id
            DELETE /api/v1/students/:id
            PUT /api/v1/students/:id
            POST /api/v1/students
       -> handler func that implements these endpoints logic

config -> db details, and server port
     
cmd/server/main.go
internal/
    handler/
    router/
    config/

main.go -> config -> setup server -> router -> handler

config
-> .env values -> load -> env for program
 (input)                 (output)
 
-> error handling for missing values

1. try keeping less env value
2. try using default values for not important envs.
3. error handling and validation is important!!!

logical and physical

physical -> localhost 5432

logical db for students data
inside students data we need table, we need a row and we need a col

// {
// "first_name": "Sindhi",
// "id": "1",
// "last_name": "Sharma",
// "age": "10",
// "class": "12",
// "gender": "m"
// },

// json -> first_name
// student.Firstname = json's firstname
// DB write student struct

// user -> {first name, last name, class, gender, age} -> DB
// json -> decode -> DB

// json
// struct
// db


// we need to establish connection with DB
// for this we need what the conn URl, that can be generated using the env vars. *cfg
// open a db connection
// *DB Open(driver, dbdetails)
// DB.Ping() -> Ping
// DB.Query("Insert")

// get all sutdents from db

endpoints (router)
func of endpoints(handler)
model (that bridges the gap between json and db col)
config (that maps env vars with config in code)
db (that establish conn with db)
handler -> db query

// this should actually get all rows from my students DB
// but for this to happen there should be a DB connection should be estd via my code.
// assumption is db has some data.

// we need to establish connection with DB
// for this we need what the conn URl, that can be generated using the env vars. cfg
// open a db connection
// get all sutdents from db

// *DB Open(driver, dbdetails)

// DB.Ping() -> Ping
// DB.Query("Insert")

DB -> GET -> handler -> null
curl -X POST {} -> DB

student struct, 

curl -X POST localhost:8080/api/v1/students \
  -H "Content-Type: application/json" \
  -d '{"first_name":"Ram","last_name":"Kumar","age":22,"class":10,"gender":"M"}' -> inside my DB

AWS -> Host, password.. 
