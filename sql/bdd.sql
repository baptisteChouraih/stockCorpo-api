DROP TABLE IF EXISTS Users, Suggestion, Product, Logs;

CREATE TABLE Users
(
    idUsers SERIAL PRIMARY KEY,
    name    VARCHAR(20) NOT NULL,
    email   VARCHAR(50) NOT NULL UNIQUE,
    pwd     VARCHAR(20) NOT NULL,
    isAdmin BOOLEAN DEFAULT FALSE
);

CREATE TABLE Suggestion
(
    idSuggestion SERIAL PRIMARY KEY,
    idUsers      INT,
    suggestion   VARCHAR(255) NOT NULL,
    FOREIGN KEY (idUsers) REFERENCES Users (idUsers)
);

CREATE TABLE Product
(
    idProduct SERIAL PRIMARY KEY,
    types     VARCHAR(20)   NOT NULL,
    price     NUMERIC(4, 2) NOT NULL,
    stock     BOOLEAN       NOT NULL
);

CREATE TABLE Logs
(
    idLog     SERIAL PRIMARY KEY,
    idUser    INT,
    action    VARCHAR(20) NOT NULL,
    details   TEXT,
    timestamp TIMESTAMP,
    FOREIGN KEY (idUser) REFERENCES Users (idUsers)
);

INSERT INTO Users (name, email, pwd)
VALUES ('test', 'test@gmail.com', 'test');

INSERT INTO Suggestion (idUsers, suggestion)
VALUES (1, 'carl demission');

INSERT INTO Product (types, price, stock)
VALUES ('snack', 0.80, false);

INSERT INTO Logs (idUser, action, details, timestamp)
VALUES (1, 'ADD', 'AJOUT PRODUIT', NOW());