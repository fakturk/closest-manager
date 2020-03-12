package employee

import(
	"fmt"
)

type Employee struct{
	Name string
	Subordinates []*Employee
}
func (e *Employee) AddSubordinate(sub *Employee)  {
	e.Subordinates=append(e.Subordinates,sub)
	Employees[e.Name]=e
}
var Employees = map[string]*Employee{}

func (e *Employee)  Print(level int){
	fmt.Print("Name: ",e.Name,", ")
	

	if len(e.Subordinates)>0 {
		fmt.Print("Subordinates: [")
		for _,v:=range e.Subordinates{
			if v!=nil {
				
			}
			level++
			v.Print(level)
		}
		fmt.Println("]")
	}
	if level==0 {
		fmt.Println()
	}
	

}

func PrintEmployees(list []*Employee){
	fmt.Println("inside Print Employees")
	fmt.Println("len: ",len(list))
	for _,v:=range list{
		fmt.Println(v.Name)
	}
}

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

func FindEmployee(name string) *Employee{
	return FindByNameDFS(GetCEO(),name)
}

func GetCEO()*Employee{
	return Employees["Tanis"]
}

func contains(s []*Employee, e *Employee) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func pathToCEO(e *Employee,name string,path []*Employee)[]*Employee{
	fmt.Println(e, e.Name,name)
	if e.Name == name {
		fmt.Println("equal")
			return append(path,e)
	}else if len(e.Subordinates) > 0 {
		for _, child := range e.Subordinates {
			if !contains(path,e) {
				path=append(path,e)
			}
			
			if result:= pathToCEO(child, name, path); result!=nil{
				fmt.Println("result: ",result)
				// return append(path,result...)
				return result
			} 
		}
	}
	return nil
}

func FindManager(e1,e2 string) *Employee{
	return FindCommonManager(FindEmployee(e1),FindEmployee(e2),GetCEO())
}

func FindCommonManager(e1,e2,ceo *Employee) *Employee{
	firstPath:=pathToCEO(ceo,e1.Name,nil)
	secondPath:=pathToCEO(ceo,e2.Name,nil)
	fmt.Println("first path")
	PrintEmployees(firstPath)
	fmt.Println("second path")
	PrintEmployees(secondPath)
	return FindManagerByPaths(firstPath,secondPath)
}

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
			fmt.Println("manager: ",manager)
		}else{
			break
		}
	}
	return manager
}


func AddEmployee(name string){
	fmt.Println("inside add employee")
	e:=Employee{
		Name:name,
	}
	Employees[name]=&e
	fmt.Println(Employees)
	fmt.Println(Employees[name])
}

func AddRelation(manager, employee string){
	m:=FindEmployee(manager)
	e:=Employees[employee]
	m.AddSubordinate(e)
	fmt.Println(m)
	fmt.Println(e)
	fmt.Println("inside add relation")
}

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