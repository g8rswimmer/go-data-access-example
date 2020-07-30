# go-data-access-example
This example takes a look at some common abstractions, [DAL](https://en.wikipedia.org/wiki/Data_access_layer) and [DAO](https://en.wikipedia.org/wiki/Data_access_object), and sees how it _could_ be applied in a go project.  This is an example and experiment with the goal of trying to make the project and code more readable.  In a project where information is easy to find and said information is easy to read, the maintainablity of the project is high.

## DAL
The data access layer is handled by a package `dal`.  This package would contain all of the entities, in this case `user`, and provide all of the functionality needed to interface with the database.

## DAO
The data access object are the interface use by packages that need to access the entity through the DAL.  In go, it is encouraged for the customer to define the interfaces.  This allows the package to define the data access functionality at the customer level.  Here the `user` api defines the DAO interface that it needs to use.  The DAL `user` is injected through that interface.

## Collections
Contained in this example is a postman collection under the `api/postman-collection` directory.
