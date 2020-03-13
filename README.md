# closest-manager
Find closest parents of two children on a tree with golang

#Instructions
Bureaucr.at is a typical hierarchical organisation. Claire, its CEO, has a hierarchy of employees reporting to her and each employee can have a list of other employees reporting to him/her. An employee with at least one report is called a Manager.

Your task is to implement a corporate directory for Bureaucr.at with an interface to find the closest common Manager (i.e. farthest from the CEO) between two employees. You may assume that all employees eventually report up to the CEO.

Here are some guidelines:
- Resolve ambiguity with assumptions.
- The directory should be an in-memory structure.
- A Manager should link to Employees and not the other way around.
- We prefer that you to use Go, but accept other languages too.
- How the program takes its input and produces its output is up to you.

____

# Assumptions

- Hierarchical organisation is structured as a tree (acyclic graph)
- Each employee only has one manager
- All employee names are uniqe
- Names are case sensitive
- Remove Employee not implemented (it could be overengineering because goal of this program is finding common manager between two employees)
- Relation can not be changed after first add (it could be overengineering because goal of this program is finding common manager between two employees)

# Hierarchical Organisation of Dragonlance

![Dragonlance](/images/dragonlance.png)

# Algorithm
- Finds the road between the given employee and ceo (all the employees between ceo and given employee) with Depth First Search for both employees.
- We have two list of paths for "ceo to employee1" and "ceo to employee2"
- Start comparing two list from the beginning and find first different element (manager) between the list
- Return the manager we found

# Usage
- go run main.go
- after running program following commands can be used inside
- For printing an employee (with it subortinates) please use [print Name] or [p Name] command
- For adding a new employee please use [newEmployee Name],[new Name] or [n Name] command
- For adding a relation between a manager and an employee please use [addRelation ManagerName EmployeeName],[add Manager Employee] or [r Manager Employee] command
- For finding common manager between two employees please use [findManager Employee1 Employee2],[find Employee1 Employee2] or [f Employee1 Employee2] command