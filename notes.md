#setup sql database

1. start mysql server
```bash
mysql.server start
```

2. create user 
```bash
mysql -u root -p <password>
```
ex. mysql -u root -p riddhi

3. create database
```bash
CREATE DATABASE stateStreet;
```
```bash
USE stateStreet;
```

4. create table to save ETFs
```bash
CREATE TABLE ETFs(Name VARCHAR(255) NOT NULL, 
Ticker VARCHAR(10) NOT NULL,
Identifier VARCHAR(30) NOT NULL, 
SEDOL VARCHAR(30) NOT NULL, 
Weight VARCHAR(30) NOT NULL, 
Sector VARCHAR(100) NOT NULL, 
Shares_Held VARCHAR(30) NOT NULL, 
Local_Currency VARCHAR(10) NOT NULL, PRIMARY KEY(Ticker));
```

5. verify the creation of table 
```bash
SHOW FIELDS FROM ETFs;
```

6. create table for users
```bash
CREATE TABLE USER(username VARCHAR(30) NOT NULL, password VARCHAR(20) NOT NULL, PRIMARY KEY(username));

```
7. insert user details
```bash
INSERT INTO USER VALUES("username1", "password1");
```