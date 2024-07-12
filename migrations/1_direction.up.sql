CREATE TABLE direction (
    code VARCHAR(8) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    exams TEXT
);