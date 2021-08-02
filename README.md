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



When a goroutine hangs without returning, the space in memory
remains used, contributing to the application's size in memory. The
goroutine and the variables it defines in the stack will get collected by
the GC only when the goroutine returns or panics.


#### SOLID
*S
Module -> source file
Only one reason to change
Module should be responsible to only one actor (service layer)
Separate code that different actors depend on
Database layer should change only when database changes! :O (no other layers/actors should be affected at all)
Database layer is responsible for database only, so it must have only one reason to change!
Separate those things that change for different reasons

Ex:
Sheet types: Tech sheet, Course sheet, Loan sheet
Changes in Tech sheets affect all other sheets (actors!) - wrong, unstable api
Tech sheet should be response for tech sheets only!

Solution: Each sheet type should be separate class with its own rules

*O
Adding new features should not force us to change existing features, at all!
Add new features by extending, not modifying! (wordpress plugins)
Printer -> print to console, print to web ui, print to file
Existing printers (console, web, file) should not be changed at all!
Just add new printer type and Printer should work

Ex:
SheetService
Adding new sheet types requires modifying sheet service and existing functions -> wrong
SheetService should be extended, not modified
No existing sheet types logic should be modified at all
New feature (sheet type), new class.

Solution: Sheet service should be wrapper around all sheet types,
so adding new sheet type (feature) should be as simple as adding new class that implements some sort of SheetInterface,
which would be mapped in sheet service

*L
Behavior of application should not depend on the type it uses
Types should be substitued with ease
Billing -> ILicense (Personal, Business)
Billing service behaviour should not depend on specific ILicense type
Never ever depend on what happens in ILicense (logging, data flow, absolutely nothing)

Ex:
SheetService's functions heavily depends on sheet types!

Solution: In conjuction with *O, sheet service should have SheetInterface reference only, and not care about each specific sheet type!

*I
Dont depend on something you dont need
Segragate interfaces
Interfaces should be small, specific to problem
Rather implement two, three smaller interfaces, than one big interace of 20 functions
That way our components can depend on as little as possible

*D
Depend only on abstractions, not concrete implementations
Simple as that

*CCP, common closure
Separate those things that change for different reasons
Gather into components those classes that change for the same reasons and at the same times.
Separate into different components those classes that change at different times and for different
reasons.
This is *S for components

*CRP, common reuse
Dont depend on components you dont need
Tells which classes not to keep together in component
This is *I for components


- start at business rules, problems
- leave database, ui, frameworks, all aside, focus on the problem being solved!
- the less you depend on database, ui, frameworks, the more decoupled system you have
- entity -> high level, business critical rules, data, knows nothing about usecases
- usecase -> low level, application rules, input, output, ref to entity, knows about entities, they know nothing about web!
- try to not modify existing functions, add new modified versions (v2) instead