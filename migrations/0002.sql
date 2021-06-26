CREATE TABLE IF NOT EXISTS users (
    id INT auto_increment,
    admin BOOLEAN,
    username VARCHAR(255),
    password VARCHAR(255),

    PRIMARY KEY (id)
);
