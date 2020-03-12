package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	"github.com/fakturk/closest-manager/employee"

)







func main()  {
	employee.AddDragonlance()
	// fmt.Printf("%+v\n",Raistlin.Subordinates[0])
	// fmt.Println(Caramon)
	// Raistlin.Print(0)
	// Caramon.Print(0)
	// Raistlin.Print(0)
	employee.Employees["Tanis"].Print(0)
	// Employees["Raistlin"].Print(0)
	// fmt.Println(Employees["Tanis"])
	// fmt.Println("flint under tanis: ",FindByNameDFS(Employees["Tanis"],"Flint"))
	// path:=pathToCEO(Employees["Tanis"],"Flint",nil)
	// fmt.Printf("flint under tanis by ceo: %+v\n",PrintEmployees(pathToCEO(Employees["Tanis"],"Flint",nil)))
	// fmt.Printf("caramon under tanis by ceo: %+v\n",PrintEmployees(pathToCEO(Employees["Tanis"],"Caramon",nil)))
	// fmt.Printf("tas under tanis by ceo: %+v\n",PrintEmployees(pathToCEO(Employees["Tanis"],"Tasslehoff",nil)))
	// PrintEmployees(pathToCEO(Employees["Tanis"],"Flint",nil))
	// PrintEmployees(pathToCEO(Employees["Tanis"],"Caramon",nil))
	// PrintEmployees(pathToCEO(Employees["Tanis"],"Tasslehoff",nil))
	// fmt.Println(FindCommonManager(Employees["Tasslehoff"],Employees["Sturm"],Employees["Tanis"]))
	// fmt.Println(FindCommonManager(Employees["Caramon"],Employees["Sturm"],Employees["Tanis"]))
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Closest Manager Finder")
	fmt.Println("An Organization Chart for the Dragonlance Characters added to the system")
	fmt.Println("For printing an employee (with it subortinates) please use [print Name] or [p Name] command")
	fmt.Println("For printing whole organization  please use [print Organization] or [p Organization] command")
	fmt.Println("For adding a new employee please use [newEmployee Name],[new Name] or [n Name] command")
	fmt.Println("For adding a relation between a manager and an employee please use [addRelation ManagerName EmployeeName],[add Manager Employee] or [r Manager Employee] command")
	fmt.Println("For finding common manager between two employees please use [findManager Employee1 Employee2],[find Employee1 Employee2] or [f Employee1 Employee2] command")
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
	
	case "print","p":
		if len(arrCommandStr)<2{
			err=errors.New("name not given for  print")
			break
		}else{
			name:=arrCommandStr[1]
			if name =="Organization"{
				name = employee.GetCEO().Name
				fmt.Println("CEO :",name)
			} else if name =="All"{
				fmt.Println(employee.Employees)
				break
			} else if (employee.Employees[name]==nil){
				err=errors.New("name not found on employee list")
				break
			}
	
			employee.Employees[name].Print(0)
		}
	
	case "newEmployee","new","n":
		if len(arrCommandStr)<2{
			err=errors.New("name not given for Employee")
			break
		}else{
			name:=arrCommandStr[1]
			employee.AddEmployee(name)
		}
	
	case "addRelation","relation","r":
		if len(arrCommandStr)<3{
			err=errors.New("usage: addRelation manager employee")
			break
		}else{
			manager:=arrCommandStr[1]
			e:=arrCommandStr[2]
			if (employee.Employees[manager]==nil){
				err=errors.New("manager not found on employee list")
				break
			}else if (employee.Employees[e]==nil){
				err=errors.New("employee not found on employee list")
				break
			}
			employee.AddRelation(manager,e)
		}
	case "findManager","find","f":
		if len(arrCommandStr)<3{
			err=errors.New("usage: findManager employee1 employee2")
			break
		}else{
			e1:=arrCommandStr[1]
			e2:=arrCommandStr[2]
			if (employee.FindEmployee(e1)==nil){
				err=errors.New("first employee not found on employee list")
				break
			}else if (employee.FindEmployee(e2)==nil){
				err=errors.New("second employee not found on employee list")
				break
			}
		manager:=employee.FindManager(e1,e2)
		fmt.Println("common manager : ",manager.Name)
		}
		// add another case here for custom commands.
	}
	// cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	return err
}


