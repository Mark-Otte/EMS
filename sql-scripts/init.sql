CREATE TABLE Employee (
   ID INT IDENTITY(1,1) PRIMARY KEY,
   FirstName VARCHAR(50) NOT NULL,
   LastName VARCHAR(50) NOT NULL,
   Username VARCHAR(50) NOT NULL,
   Password VARCHAR(255) NOT NULL,
   Email VARCHAR(255) NOT NULL,
   DOB DATE NOT NULL,
   DepartmentID INT NOT NULL,
   Position VARCHAR(50) NOT NULL,
   CONSTRAINT FK_Employee_Department FOREIGN KEY (DepartmentID) REFERENCES Department (ID)
);

CREATE TABLE Department (
   ID INT IDENTITY(1,1) PRIMARY KEY,
   Name VARCHAR(50) NOT NULL
);