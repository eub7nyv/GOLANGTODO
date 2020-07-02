# GOLANGTODO With MySql
Go Lang TODO App

Steps : 

1. Create SQL DB in your Local Machine
2. USerName : root , Password: *****
3. Create DB - goblog
4. select DB
5. Create Task Table with id, taskName, taskDscription
6. Id as Auto incremnent
7. create table task(id int(6) unsigned NOT NULL AUTO_INCREMENT,taskname varchar(30) NOT NULL,
taskdescription varchar(30) NOT NULL,PRIMARY KEY (id)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
8. Clone the application
9.  cd appname
10. run the driver command : go get -u github.com/go-sql-driver/mysql
11. Change DB Connection in main.go page
12. run the application by using command : go run main.go
13. Application URL # http://localhost:8980/
