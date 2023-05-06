CREATE DATABASE state_management;
USE state_management;

CREATE TABLE test_attempt (
    attempt_id INT NOT NULL AUTO_INCREMENT,
    user_id INT DEFAULT NULL, 
    test_id INT DEFAULT NULL,

    state INT DEFAULT NULL,
    started_at DATETIME,
    ended_at DATETIME,

    PRIMARY KEY(attempt_id)
);