package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
)

type Employee struct{
	ID int
	Name string
	Subordinates []*Employee
}



func (e *Employee) AddSubordinate(sub *Employee)  {
	e.Subordinates=append(e.Subordinates,sub)
	Employees[e.Name]=e
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
	// Employees["Raistlin"].Print(0)
	// fmt.Println(Employees["Tanis"])
	// fmt.Println("flint under tanis: ",findByNameDFS(Employees["Tanis"],"Flint"))
	// path:=pathToCEO(Employees["Tanis"],"Flint",nil)
	// fmt.Printf("flint under tanis by ceo: %+v\n",PrintEmployees(pathToCEO(Employees["Tanis"],"Flint",nil)))
	// fmt.Printf("caramon under tanis by ceo: %+v\n",PrintEmployees(pathToCEO(Employees["Tanis"],"Caramon",nil)))
	// fmt.Printf("tas under tanis by ceo: %+v\n",PrintEmployees(pathToCEO(Employees["Tanis"],"Tasslehoff",nil)))
	// PrintEmployees(pathToCEO(Employees["Tanis"],"Flint",nil))
	// PrintEmployees(pathToCEO(Employees["Tanis"],"Caramon",nil))
	// PrintEmployees(pathToCEO(Employees["Tanis"],"Tasslehoff",nil))
	// fmt.Println(findCommonManager(Employees["Tasslehoff"],Employees["Sturm"],Employees["Tanis"]))
	// fmt.Println(findCommonManager(Employees["Caramon"],Employees["Sturm"],Employees["Tanis"]))
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Closest Manager Finder")
	fmt.Println("An Organization Chart for the Dragonlance Characters added to the system")
	fmt.Println("For printing an employee (with it subortinates) please use [print Name] command")
	fmt.Println("For printing whole organization  please use [print Organization] command")
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	
}
func runCommand(commandStr string) error {
	
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	fmt.Println(arrCommandStr)
	var err error
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
		// add another case here for custom commands.
	
	case "print":
		if len(arrCommandStr)<2{
			err=errors.New("name not given for  print")
			break
		}else{
			name:=arrCommandStr[1]
			if name =="Organization"{
				name = getCEO().Name
			} else if name =="All"{
				fmt.Println(Employees)
				break
			} else if (Employees[name]==nil){
				err=errors.New("name not found on employee list")
				break
			}
	
			Employees[name].Print(0)
		}
	
	case "newEmployee":
		if len(arrCommandStr)<2{
			err=errors.New("name not given for Employee")
			break
		}else{
			name:=arrCommandStr[1]
			AddEmployee(name)
		}
	
	case "addRelation":
		if len(arrCommandStr)<3{
			err=errors.New("usage: addRelation manager employee")
			break
		}else{
			manager:=arrCommandStr[1]
			employee:=arrCommandStr[2]
			if (findEmployee(manager)==nil){
				err=errors.New("manager not found on employee list")
				break
			}else if (Employees[employee]==nil){
				err=errors.New("employee not found on employee list")
				break
			}
			AddRelation(manager,employee)
		}
		// add another case here for custom commands.
	}
	// cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	return err
}
func findByNameDFS(e *Employee, name string) *Employee {
	// fmt.Println(e, e.Name,name)
	if e.Name == name {
		// fmt.Println("equal")
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

func findEmployee(name string) *Employee{
	return findByNameDFS(getCEO(),name)
}

func getCEO()*Employee{
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
// type Counter struct {
// 	count:=1
// }

// func (self Counter) currentValue() int {
// 	return self.count
// }
// func (self *Counter) increment() {
// 	self.count++
// }

func AddEmployee(name string){
	fmt.Println("inside add employee")
	e:=Employee{
		ID:9,
		Name:name,
	}
	Employees[name]=&e
	fmt.Println(Employees)
	fmt.Println(Employees[name])
}

func AddRelation(manager, employee string){
	m:=findEmployee(manager)
	e:=Employees[employee]
	m.AddSubordinate(e)
	fmt.Println(m)
	fmt.Println(e)
	fmt.Println("inside add relation")
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
	Employees["Raistlin"]=&Raistlin
	Employees["Caramon"]=&Caramon
	Employees["Tanis"]=&Tanis
	Employees["Flint"]=&Flint
	Employees["Goldmoon"]=&Goldmoon
	Employees["Riwervind"]=&Riverwind
	Employees["Tasslehoff"]=&Tasslehoff
	Employees["Sturm"]=&Sturm
	
	Tanis.AddSubordinate(&Raistlin)
	Tanis.AddSubordinate(&Flint)
	Tanis.AddSubordinate(&Goldmoon)
	Raistlin.AddSubordinate(&Caramon)
	Flint.AddSubordinate(&Tasslehoff)
	Flint.AddSubordinate(&Sturm)
	Goldmoon.AddSubordinate(&Riverwind)

	

	AddEmployee("Dalamar")
	AddRelation("Raistlin","Dalamar")


}

