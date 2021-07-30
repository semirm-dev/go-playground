# go-playground
Just playing around with Go

## Layers and notes
* Domain layer, modeling, without any dependencies, unit testable -> can be merged with business logic layer
* Business logic layer, without any dependencies, unit testable, pure business logic
* Repository layer, interface between database logic layer and domain/business logic layer
* Database logic layer
* Controllers, usecases layer, use repository layer to apply business logic and against database layer, wire it all together

* Each layer must be changed without affecting other layers (not even single line of code should be changed in other layers)
* Each layer must communicate with another through interfaces, there must be no single concrete implementation dependency. 
Interfaces should be simple and their job is to transport and translate data between layers, plus they will make layers testable in isolation.
* Each layer should have its own rules:
> Business layer - not allowing customers to order unavailable items is a rule that applies to the web shop as well as to orders placed via a telephone hotline, orders must not exceed 250$... – such rules are business-wise, for everyone, not technical/application or specific use case

> Usecases - checking which users may add items to an order is use case specific, admin vs normaln user – plus, it contains an entity, User, that the domain layer must not be bothered with.

> Technical (Database layer) - the value of an item has to be saved to a database, and we must take care to only store floats to the value field within our database; however, this is a technical rule, not a business rule, and does not belong into our domain package

* Where is each layer used
* Where is its interface
* Where is its implementation
* Is it testable in isolation (no other layers should be touched/affected at all)


* io. - Reader, Writer (no buffers around - not always efficient, system calls on each Read(), Write())
* bufio. - wrap existing Reader, Writer with Buffer - makes less system calls, memory and cpu efficient
* ioutil - for small and simple task, not efficient at all for long-lived apps and big files
* bytes. - create new buffered Reader, Writer, ....