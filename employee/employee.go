package employee

import(
	"fmt"
	"strings"
	"github.com/fatih/color"
)

type Employee struct{
	Name string
	Subordinates []*Employee
}

// adds subordinate to and employee
func (e *Employee) AddSubordinate(sub *Employee)  {
	e.Subordinates=append(e.Subordinates,sub)
	Employees[e.Name]=e
}
var Employees = map[string]*Employee{}
var CEO =Employee{
	Name:"",
}

// prints an employee with its subordinates
func (e *Employee)  Print(level int){
	space := strings.Repeat(" ", level)
	fmt.Print(space,"Name: ",color.CyanString(e.Name))

	if len(e.Subordinates)>0 {
		level++
		fmt.Println(" ,Subordinates ↓")
		for _,v:=range e.Subordinates{
			v.Print(level)
		}
		fmt.Println()
	}
	if level==0 {
		fmt.Println()
	}
	

}

// prints all the employees even if they are not belong to the organization yet
func PrintEmployees(list []*Employee){
	// fmt.Println("inside Print Employees")
	// fmt.Println("len: ",len(list))
	for _,v:=range list{
		fmt.Println(v.Name)
	}
}

//finds given employee on the hierachy tree, first parameter is the head of the tree (ceo) and second paramter is the name of the employee that we search
//finds with Depth First Search
func FindByNameDFS(e *Employee, name string) *Employee {
	// fmt.Println(e, e.Name,name)
	if e.Name == name {
		// fmt.Println("equal")
			return e
	} else if len(e.Subordinates) > 0 {
			for _, child := range e.Subordinates {
				if result:= FindByNameDFS(child, name); result!=nil{
					return result
				} 
			}
	}
	return nil
}

// finds the employee on the organization with DFS
func FindEmployee(name string) *Employee{
	return FindByNameDFS(GetCEO(),name)
}

// gets CEO of the orgnaization
func GetCEO()*Employee{
	return &CEO
}

// checks if the employee list contains given employee
func contains(s []*Employee, e *Employee) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

// finds the path between CEO and given empmployee
// path means an emmployee list that have all the middle managers include CEO and employee
func pathToCEO(e *Employee,name string,path []*Employee)[]*Employee{
	// fmt.Println(e, e.Name,name)
	if e.Name == name {
		// fmt.Println("equal")
			return append(path,e)
	}else if len(e.Subordinates) > 0 {
		for _, child := range e.Subordinates {
			if !contains(path,e) {
				path=append(path,e)
			}
			
			if result:= pathToCEO(child, name, path); result!=nil{
				// fmt.Println("result: ",result)
				return result
			} 
		}
	}
	return nil
}

// finds the common manager between two employees
// this method only take employees names
func FindManager(e1,e2 string) *Employee{
	return FindCommonManager(FindEmployee(e1),FindEmployee(e2),GetCEO())
}

// finds the common manager between two employees
func FindCommonManager(e1,e2,ceo *Employee) *Employee{
	firstPath:=pathToCEO(ceo,e1.Name,nil)
	secondPath:=pathToCEO(ceo,e2.Name,nil)
	// fmt.Println("first path")
	// PrintEmployees(firstPath)
	// fmt.Println("second path")
	// PrintEmployees(secondPath)
	return FindManagerByPaths(firstPath,secondPath)
}

// compare two paths (ceo to employee1 and ceo to employee2) and finds the closest common manager
func FindManagerByPaths(p1,p2 []*Employee) *Employee{
	lenp1:=len(p1)
	lenp2:=len(p2)
	smallestLen:=lenp2
	if lenp1<lenp2{
		smallestLen=lenp1
	}
	manager:=&Employee{}
	for i := 0; i < smallestLen; i++ {
		if p1[i]==p2[i] {
			manager=p1[i]
			// fmt.Println("Closest Common Manager: ",manager)
		}else{
			break
		}
	}
	return manager
}

// adds a new employee (without relations)
// if employee already exist this method does nothing
func AddEmployee(name string){
	// fmt.Println("inside add employee")
	e:=Employee{
		Name:name,
	}
	if Employees[name]==nil{
		Employees[name]=&e
	}
	
	// fmt.Println(Employees)
	// fmt.Println(Employees[name])
}

// adds a relation between two employees and one employee become manager and other one become subordinate
func AddRelation(manager, employee string){
	m:=Employees[manager]
	e:=Employees[employee]
	m.AddSubordinate(e)
	UpdateCEO(m,e)
	// fmt.Println(m)
	// fmt.Println(e)
	// fmt.Println("inside add relation")
}

// everytime when we are adding a relation between employees we check the relation  between them and update the CEO of the organization
func UpdateCEO(manager, employee *Employee)  {
	// fmt.Println("inside updateCEO, ",manager,employee)
	if (CEO.Name=="") || employee.Name==CEO.Name{
		CEO=*manager
		// fmt.Println("CEO  updated, ",CEO)
	}
}

// adds a mock organization that named Dragonlance and have Dragonlance characters
func AddDragonlance()  {
	// Drangonlance characters added to employee list (without relations)
	AddEmployee("Raistlin")
	AddEmployee("Caramon")
	AddEmployee("Tanis")
	AddEmployee("Flint")
	AddEmployee("Goldmoon")
	AddEmployee("Riverwind")
	AddEmployee("Mira")
	AddEmployee("Tasslehoff")
	AddEmployee("Sturm")
	AddEmployee("Dalamar")
	AddEmployee("Valin")
	AddEmployee("Parsalian")
	AddEmployee("Crysina")
	AddEmployee("Bupu")
	AddEmployee("Laurana")
	AddEmployee("Kitiara")
	AddEmployee("Huma")
	AddEmployee("Elistan")
	AddEmployee("Fistandantilus")
	AddEmployee("Tika")
	
	// Add relations between dragonlance hierachy
	AddRelation("Tanis","Raistlin")
	AddRelation("Tanis","Flint")
	AddRelation("Tanis","Goldmoon")

	AddRelation("Raistlin","Caramon")
	AddRelation("Raistlin","Dalamar")
	AddRelation("Raistlin","Bupu")


	AddRelation("Flint","Tasslehoff")
	AddRelation("Flint","Sturm")

	AddRelation("Goldmoon","Riverwind")
	AddRelation("Goldmoon","Mira")

	AddRelation("Dalamar","Valin")
	AddRelation("Dalamar","Parsalian")
	AddRelation("Dalamar","Crysina")

	AddRelation("Caramon","Laurana")
	AddRelation("Caramon","Kitiara")
	AddRelation("Caramon","Huma")

	AddRelation("Tasslehoff","Elistan")
	AddRelation("Tasslehoff","Fistandantilus")
	AddRelation("Tasslehoff","Tika")

}