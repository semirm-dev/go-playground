#### Concurrency vs Parallelism: 
* Concurrency can be implemented with parallelism (execution at same time) or context switching (not at same time)
* Concurrent, many actions in progress at same time, execution does not have to happen at same time (but it can!)
* Parallelism, many actions executing at the same time, thanks to multi-core processors

* race-condition, concurrent execution of a function can lead to undesired behavior, different outputs can happen each time
* data-race, concurrent read + write, will crash application
* deadlock, two threads wait for each other to unlock, infinitely waiting for each other

#### goroutine leaking
When a goroutine hangs without returning, the space in memory
remains used, contributing to the application's size in memory. The
goroutine and the variables it defines in the stack will get collected by
the GC only when the goroutine returns or panics.

#### io
* io. - Reader, Writer (no buffers around - not always efficient, system calls on each Read(), Write())
* ioutil. - for small and simple task, not efficient at all for long-lived apps and big files
* bufio. - wrap existing Reader, Writer with Buffer - makes less system calls, memory and cpu efficient
* bytes. - create new buffered Reader, Writer, really/most efficient!!

#### mem
- stack, local variables, not referenced after function return, short lifetime, lifo cleaned
- heap, variables referenced after function return, too big values, unknown size, longer lifetime, gc collected
- stack is limited to 1gb on 64bits, and 250mb on 32bits
- heap is limited to the all available memory
- copied variable (struct{}) will usually remain on stack
- pointer variable (*struct{}) will usually go to heap
- goroutine stack size starts at 2kb and grows/shrinks, after func return stack is cleaned and returned value is copied to caller
- both storages (stack and heap) will clean a variable when it is no longer needed

#### architecture
* most inner circle, business logic, interfaces, knows absolutely nothing about any other component
* middle circle, use case, uses most inner circle interfaces to apply some business logic
* adapters, most outer circle, implement most inner circle, knows about other components, replace-able
* dependency always points inwards, inner circle must never know anything about outer circle (world)

#### SOLID
*S
Module -> source file
Only one reason to change
Module should be responsible to only one actor
Separate code that different actors depend on
Separate those things that change for different reasons

*O
Adding new features should not force us to change existing features, at all!
Add new features by extending, not modifying! (wordpress plugins)
Printer -> print to console, print to web ui, print to file
Existing printers (console, web, file) should not be changed at all!
Just add new printer type and Printer should work

*L
Behavior of application should not depend on the type it uses
Types should be substitued with ease
Billing -> ILicense (Personal, Business)
Billing service behaviour should not depend on specific ILicense type
Never ever depend on what happens in ILicense (logging, data flow, absolutely nothing)

*I
Dont depend on something you dont need
Segregate interfaces
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
- entity -> high severity, business critical rules, data, knows nothing about usecases
- usecase -> low severity, application rules, input, output, ref to entity, knows about entities, they know nothing about web!
- try to not modify existing functions, add new modified versions (v2) instead