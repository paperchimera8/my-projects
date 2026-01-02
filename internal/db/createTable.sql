CREATE TABLE Shops (
    ID INT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Time VARCHAR(50)  NOT NULL
);

INSERT INTO Shops(ID, Name, Time)
VALUES (1, 'Пятерочка', '8.30 - 23.00'),
       (2, 'Магнит', '8.00 - 23.30');