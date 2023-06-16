CREATE TABLE IF NOT EXISTS admin (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(80),
    last_name VARCHAR(80),
    email VARCHAR(120),
    password varchar(150),
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS client (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(80),
    last_name VARCHAR(80),
    email VARCHAR(120),
    password VARCHAR(150),
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    client_id INTEGER REFERENCES client(id) ON DELETE CASCADE,
    number UUID,
    balance INTEGER, 
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transfer_demand (
    id SERIAL PRIMARY KEY,
    closed BOOLEAN,
    from_account UUID,
    to_account UUID,
    message VARCHAR(250),
    amount INTEGER,
    accepted BOOLEAN,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tansfer (
    id SERIAL PRIMARY KEY,
    demand_id INTEGER REFERENCES transfer_demand(id),
    done BOOLEAN,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS closed_account (
    id SERIAL PRIMARY KEY,
    number UUID,
    created_at TIMESTAMP
);
