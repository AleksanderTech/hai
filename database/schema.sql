CREATE TABLE IF NOT EXISTS messages (
    id bigserial PRIMARY KEY,
    code varchar(100) NOT NULL,
    title varchar(255) NOT NULL,
    content text NOT NULL,
    email varchar(100) NOT NULL
);