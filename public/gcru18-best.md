# Code Like the Go Team

---
## About Me

- Microsoft Azure - Cloud Developer Advocate
- Twitter: @ashleymcnamara
- Github: ashleymcnamara
- GOPHERS!: https://github.com/ashleymcnamara/gophers

---
## This Talk

[https://cda.ms/jP](https://cda.ms/jP)
[video](https://youtu.be/MzTcsI6tn-0)

---

## My Hero
<!-- .slide: data-background-image="/images/brian.jpeg" -->


---

#### Write Code Like the Go Team

- how to organize your code into packages, and what those packages should contain  
- code patterns and conventions that are prevalent in the standard library  
- how to write your code to be more clear and understandable  
- unwritten Go conventions that go beyond “go fmt” and make you look like a veteran Go contributor wrote it

---

#### Outline

- Packages
- Naming Conventions
- Source Code Conventions

---

## Packages

---

#### Code Organization

There are two key areas of code organization in Go that will make a huge impact on the usability, testability, and functionality of your code:

---

#### Code Organization

- Package Naming
- Package Organization

---

#### Code Organization

We can't talk about them separately because one influences the other.

---

#### Package Organization - Libraries

- Packages contain code that has a single purpose

```
	archive	cmd	crypto	errors	go	index	math	path	sort	syscall	unicode bufio	compress  database  expvar	hash	internal  mime      reflect	strconv	testing	unsafe builtin	container debug	flag  
```

---

#### Package Organization - Libraries

When a group of packages provides a common set of functionalities with different implementations, they're organized under a parent.  Look at the contents of package *encoding*:

```
	ascii85     base32      binary      encoding.go hex         pem
	asn1        base64      csv         gob         json        xml
```

---

#### Package Organization - Libraries

Some commonalities:

- Packages names describe their purpose
- It's very easy to see what a package does by looking at the name
- Names are generally short
- When necessary, use a descriptive parent package and several children implementing the functionality -- like the *encoding* package

---


#### Package Organization - Applications

The packages we've seen are all libraries.  They're intended to be imported and used by some executable program like a service or command line tool.

What should the organization of your executable applications look like?

---

#### Package Organization - Applications

When you have an application, the package organization is subtly different.  The difference is the command, the executable that ties all of those packages together.

---

#### Package Organization - Applications

Application package organization has a huge impact on the testability and functionality of your system.

---

#### Package Organization - Applications

When writing an application your goal should be to write code that is easy to understand, easy to refactor, and easy for someone else to maintain.

---


#### Package Organization - Applications

Most libraries focus on providing a singularly scoped function; logging, encoding, network access.

Your application will tie all of those libraries together to create a tool or service.  That tool or service will be much larger in scope.  

---

#### Package Organization - Applications

When you're building an application, you should organize your code into packages, but those packages should be centered on two categories:

- Domain Types
- Services

---

#### Package Organization - Applications

*Domain Types* are types that model your business functionality and objects. 

*Services* are packages that operate on or with the domain types.

https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1

---

#### Package Organization - Applications

A domain type is the substance of your application.  If you have an inventory application, your domain types might include *Product* and *Supplier*.  If you have an HR administration system, your domain types might include *Employee*, *Department*, and *Business Unit*.

---

#### Package Organization - Applications

The package containing your domain types should also define the interfaces between your domain types and the rest of the world.  These interfaces define the things you want to do with your domain types.

- *ProductService*  
- *SupplierService*
- *AuthenticationService*
- *EmployeeStorage*
- *RoleStorage*

---

#### Package Organization - Applications

Your domain type package should be the root of your application repository.  This makes it clear to anyone opening the codebase what types are being used, and what operations will be performed on those types.

---

#### Package Organization - Applications

The domain type package, or *root* package of your application should not have any external dependencies.  
> It exists for the sole purpose of describing your types and their behaviors.

---

#### Package Organization - Applications

The implementations of your domain interfaces should be in separate packages, organized by dependency.

---

#### Package Organization - Applications

Dependencies include:

- External data sources
- Transport logic (http, RPC)

You should have one package per dependency.

---

#### Package Organization - Applications

Why one package per dependency?

- Simple testing
- Easy substitution/replacement
- No circular dependencies


---

## Conventions

---

#### Naming Conventions

> there are two hard things in computer science: 
> cache invalidation, naming things, and off-by-one errors
> - Every developer on Twitter


Naming things *is* hard, but putting some thought into your type, function, and package names will make your code more readable.

---

#### Naming Conventions - Packages

A package name should have the following characteristics:

- short
	- Prefer "transport" over "transportmechanisms"

- clear 
	- Name for clarity based on function: "bytes" 
	- Name to describe implementation of external dependency: "postgres" 

---

#### Naming Conventions - Packages

Packages should provide functionality for one and only one purpose.  Avoid *catchall* packages:

-	util
-	helpers
-	etc.

Frequently they're a sign that you're missing an interface somewhere.

---

#### Naming Conventions - Packages

 `util.ConvertOtherToThing()` should probably be refactored into a Thinger interface

*catchall* packages are always the first place you'll run into problems with testing and circular dependencies, too.

---

#### Naming Conventions - Variables

Some common conventions for variable names:

- use camelCase not snake_case
- use single letter variables to represent indexes
	- `for i:=0; i < 10; i++ {}`
- use short but descriptive variable names for other things
	- `var count int`
	- `var cust Customer`

---

#### Naming Conventions - Variables

There are no bonus points in Go for obfuscating your code by using unnecessarily short variables.  Use the scope of the variable as your guide.  The farther away from declaration you use it, the longer the name should be.

---

#### Naming Conventions - Variables

- use repeated letters to represent a collection/slice/array
	- `var tt []*Thing`
- inside a loop/range, use the single letter
	- `for i, t := range tt {}`

These conventions are common inside Go's own source code.

---

#### Naming Conventions - Functions and Methods

Avoid a package-level function name that repeats the package name.  

	GOOD:  log.Info()
	BAD:   log.LogInfo()

The package name already declares the purpose of the package, so there's no need to repeat it.

---

#### Naming Conventions - Functions and Methods

Go code doesn't have setters and getters.

	GOOD:  custSvc.Customer()
	BAD:   custSvc.GetCustomer()

---

#### Naming Conventions - Interfaces

If your interface has only one function, append "-er" to the function name:
```go
	type Stringer interface{
		String() string
	}
```

---

#### Naming Conventions - Interfaces

If your interface has more than one function, use a name to represent its functionality:

```go
	type CustomerStorage interface {
		Customer(id int) (*Customer, error)
		Save(c *Customer)  error
		Delete(id int) error
	}
```

---

#### Naming Conventions - Interfaces

Some purists think that all interfaces should end in `-er`.  I think interfaces should be descriptive and readable.  

```go
	type CustomerStorer interface {} // MEH
	type CustomerStorage interface {} // Better
```
---

#### Naming Conventions - Source Code

Inside a package separate code into logical concerns.

If the package deals with multiple types, keep the logic for each type in its own source file:
```
	package: postgres

	orders.go
	suppliers.go
	products.go
```
---

#### Naming Conventions - Source Code

In the package that defines your domain objects, define the types and interfaces for each object in the same source file:

```
	package: inventory

	orders.go 
	-- contains Orders type 
	   and OrderStorage interface
```
---

#### Smaller Tips

- Make comments in full sentences, always.
```go
	// An Order represents an order from a customer.
	type Order struct {}
```
---

#### Smaller Tips

Use `goimports` to manage your imports, and they'll always be in canonical order. Standard lib first, external next. 

---

#### Smaller Tips

Avoid the `else` clause.  Especially in error handling.  
```go
	if err != nil {
		// error handling
		return // or continue, etc.
	} 
	// normal code
```

---

## Coding Like the Go Team

Following these conventions will make your source code easier to read, easier to maintain, and easier for someone else to understand.

---

### Questions?

<br>

@fa[twitter gp-contact](@ashleymcnamara)



---
<!-- .slide: data-background-image="/images/gitpitch-audience.jpg" -->

## Thank You