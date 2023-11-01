CREATE TABLE students (
	id INT NOT NULL,
	first_name VARCHAR(100) NULL,
	last_name VARCHAR(100) NULL,
	age INT NOT NULL,
	grade FLOAT NULL,
	class_name VARCHAR(100) NULL,
	entrance_date TIME NULL,
	created_at TIME NOT NULL,
	updated_at TIME NOT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;