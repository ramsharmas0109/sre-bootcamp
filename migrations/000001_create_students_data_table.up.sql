CREATE TABLE students_data (
    student_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    class INT NOT NULL,
    gender CHAR NOT NULL
);
