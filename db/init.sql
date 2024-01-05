CREATE TABLE tweets (
    id int PRIMARY KEY,
    user_id varchar(30),
    content varchar(100)
);

CREATE TABLE users (
    username varchar(100)
);

INSERT INTO users (username)
VALUES
('b');

INSERT INTO tweets (id, user_id, content)
VALUES
(1, 'b', 'First Tweet'),
(2, 'b', 'Hello World!');

