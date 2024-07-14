CREATE TABLE profiles (
    id INTEGER PRIMARY KEY,
    userId INT NOT NULL UNIQUE,
    firstName VARCHAR(255),
    middleName VARCHAR(255),
    lastName VARCHAR(255)
);