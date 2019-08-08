# 
Rest API for Creating, Finding, Deleting and Modifying a Collection a Movies Written in Golang with MongoDB Used as Database


2 Dependencies : "github.com/BurntSushi/toml

Make sure the structure of the directory remains intact when you clone it onto your computer.

1. model/model.go - Contains the structure for each of our document(MOVIE RECORD) inside the movie collection.

2. dao.moviesdao.go - Contains functions for CRUD operations happening on database-golang level which are inturn invoked by handler functions in myapp.go 

3. config/config.go - Contains the code for establishing a connection between GO and MongoDB and uses to .toml file given at the outermost level ( that contains database address and name of database to be used)


4. myapp.go is the code which contains handles routing on the first level using handler functions and will furthe interact with the frontend




A Blog shall be linked soon! Stay tuned. Write to narayansharma275@gmail.com in case of any queries.

