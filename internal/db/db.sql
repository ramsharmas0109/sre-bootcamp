CREATE DATABASE students;

CREATE TABLE students_data (
    student_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    class INT NOT NULL,
    gender CHAR NOT NULL
);

INSERT INTO students_data (first_name, last_name, class, gender) VALUES ('Ram', 'Sharma', 10, 'M');
INSERT INTO students_data (first_name, last_name, class, gender) VALUES ('AbRam', 'Khan', 5, 'M');
INSERT INTO students_data (first_name, last_name, class, gender) VALUES ('Nagaraju Segu', 'Kumar', 12, 'M');

SELECT * FROM students_data;
SELECT last_name, first_name FROM students_data;
SELECT student_id, first_name FROM students_data WHERE student_id = 2;
SELECT * FROM students_data WHERE student_id BETWEEN 1 AND 10;

UPDATE students_data SET class = 10 WHERE student_id = 2;
UPDATE students_data SET class = 6, last_name = 'Sharma' WHERE student_id = 2;


DELETE FROM students_data WHERE student_id = 2;
======================================================

CREATE DATABASE teachers;

CREATE TABLE teachers_data (
    teacher_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    class_teacher INT NOT NULL,
    gender CHAR NOT NULL
);

INSERT INTO teachers_data (first_name, last_name, class_teacher, gender) VALUES ('Vinod', 'Yadav', 5, 'M');
INSERT INTO teachers_data (first_name, last_name, class_teacher, gender) VALUES ('Sonia', 'Bansal', 10, 'M');
INSERT INTO teachers_data (first_name, last_name, class_teacher, gender) VALUES ('Arjun', 'Sardana', 12, 'M');

SELECT * FROM teachers_data;
SELECT first_name FROM teachers_data WHERE teacher_id = 1;
SELECT first_name FROM teachers_data WHERE teacher_id  BETWEEN 1 and 3;

UPDATE teachers_data SET last_name = 'Kumar' WHERE teacher_id = 1;
UPDATE teachers_data SET class_teacher = 4 WHERE teacher_id = 3;


DELETE FROM teachers_data WHERE teacher_id = 1;
DELETE FROM teachers_data WHERE teacher_id BETWEEN 2 and 3;
