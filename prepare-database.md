# Commands to run

1. Connect to MySQL server on command line

> mysql -h 127.0.0.1 -u golang -p

2. Select devbook database

> use devbook;

3. Create table users

> CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    email varchar(50) not null
) ENGINE=INNODB;
