package main

import(
	"fmt"
)

type Employee struct{
	ID int
	Name string
	Subordinates []*Employee
}



func (e *Employee) AddSubordinate(sub *Employee)  {
	e.Subordinates=append(e.Subordinates,sub)
}
var Employees = map[string]*Employee{}

func (e *Employee)  Print(level int){
	fmt.Print("Name: ",e.Name)
	fmt.Print(",ID: ",e.ID,", ")
	

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

func main()  {
	AddDragonlance()
	// fmt.Printf("%+v\n",Raistlin.Subordinates[0])
	// fmt.Println(Caramon)
	// Raistlin.Print(0)
	// Caramon.Print(0)
	// Raistlin.Print(0)
	Employees["Tanis"].Print(0)
	// fmt.Println(Employees["Tanis"])
	// fmt.Println("flint under tanis: ",findByNameDFS(Employees["Tanis"],"Flint"))
	// path:=pathToCEO(Employees["Tanis"],"Flint",nil)
	// fmt.Printf("flint under tanis by ceo: %+v\n",PrintEmployees(pathToCEO(Employees["Tanis"],"Flint",nil)))
	// fmt.Printf("caramon under tanis by ceo: %+v\n",PrintEmployees(pathToCEO(Employees["Tanis"],"Caramon",nil)))
	// fmt.Printf("tas under tanis by ceo: %+v\n",PrintEmployees(pathToCEO(Employees["Tanis"],"Tasslehoff",nil)))
	// PrintEmployees(pathToCEO(Employees["Tanis"],"Flint",nil))
	// PrintEmployees(pathToCEO(Employees["Tanis"],"Caramon",nil))
	// PrintEmployees(pathToCEO(Employees["Tanis"],"Tasslehoff",nil))
	fmt.Println(findCommonManager(Employees["Tasslehoff"],Employees["Sturm"],Employees["Tanis"]))
	fmt.Println(findCommonManager(Employees["Caramon"],Employees["Sturm"],Employees["Tanis"]))

	
}
func findByNameDFS(e *Employee, name string) *Employee {
	fmt.Println(e, e.Name,name)
	if e.Name == name {
		fmt.Println("equal")
			return e
	} else if len(e.Subordinates) > 0 {
			for _, child := range e.Subordinates {
				if result:= findByNameDFS(child, name); result!=nil{
					return result
				} 
			}
	}
	return nil
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

func findCommonManager(e1,e2,ceo *Employee) *Employee{
	firstPath:=pathToCEO(ceo,e1.Name,nil)
	secondPath:=pathToCEO(ceo,e2.Name,nil)
	fmt.Println("first path")
	PrintEmployees(firstPath)
	fmt.Println("second path")
	PrintEmployees(secondPath)
	return findManagerByPaths(firstPath,secondPath)
}

func findManagerByPaths(p1,p2 []*Employee) *Employee{
	lenp1:=len(p1)
	// lenp2:=len(p2)
	manager:=&Employee{}
	for i := 0; i <= lenp1; i++ {
		if p1[i]==p2[i] {
			manager=p1[i]
			fmt.Println("manager: ",manager)
		}else{
			break
		}
	}
	return manager
}

func AddDragonlance()  {
	Raistlin:=Employee{
		ID:1,
		Name:"Raistlin",
	}
	Caramon:=Employee{
		ID:2,
		Name:"Caramon",
	}
	Tanis:=Employee{
		ID:3,
		Name:"Tanis",
	}
	Flint:=Employee{
		ID:4,
		Name:"Flint",
	}
	Goldmoon:=Employee{
		ID:5,
		Name:"Goldmoon",
	}
	Riverwind:=Employee{
		ID:6,
		Name:"Riverwind",
	}
	Tasslehoff:=Employee{
		ID:7,
		Name:"Tasslehoff",
	}
	Sturm:=Employee{
		ID:8,
		Name:"Sturm",
	}
	Tanis.AddSubordinate(&Raistlin)
	Tanis.AddSubordinate(&Flint)
	Tanis.AddSubordinate(&Goldmoon)
	Raistlin.AddSubordinate(&Caramon)
	Flint.AddSubordinate(&Tasslehoff)
	Flint.AddSubordinate(&Sturm)
	Goldmoon.AddSubordinate(&Riverwind)

	Employees["Raistlin"]=&Raistlin
	Employees["Caramon"]=&Caramon
	Employees["Tanis"]=&Tanis
	Employees["Flint"]=&Flint
	Employees["Goldmoon"]=&Goldmoon
	Employees["Riwervind"]=&Riverwind
	Employees["Tasslehoff"]=&Tasslehoff
	Employees["Sturm"]=&Sturm

}

