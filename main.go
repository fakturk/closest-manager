package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	"github.com/fakturk/closest-manager/employee"
	"github.com/fatih/color"

)







func main()  {
	//adds mock organization called dragonlance (with dragonlance characters)
	employee.AddDragonlance()

	reader := bufio.NewReader(os.Stdin)
	color.Cyan("Welcome to Closest Manager Finder")
	fmt.Println("An Organization Chart for the Dragonlance Characters added to the system")
	fmt.Println("Please use", color.RedString("[exit]") ,"command for exiting from program")
	fmt.Println("For printing an employee (with it subortinates) please use", color.GreenString("[print Name]"), "or", color.GreenString("[p Name]"), "command")
	fmt.Println("For printing whole organization  please use", color.GreenString("[print Organization]"), "or", color.GreenString("[p Organization]"), "command")
	fmt.Println("For adding a new employee please use" ,color.GreenString("[newEmployee Name]"),",",color.GreenString("[new Name]"), "or" ,color.GreenString("[n Name]"), "command")
	fmt.Println("For adding a relation between a manager and an employee please use", color.GreenString("[addRelation ManagerName EmployeeName]"),",",color.GreenString("[add Manager Employee]"), "or", color.GreenString("[r Manager Employee]"), "command")
	fmt.Println("For finding common manager between two employees please use", color.GreenString("[findManager Employee1 Employee2]"),",",color.GreenString("[find Employee1 Employee2]"), "or", color.GreenString("[f Employee1 Employee2]"), "command")
	fmt.Println("Please use", color.YellowString("[help]") ,"command for help")
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
	case "help","h":
		fmt.Println("\n||\n  character meaning  or  and it means either commands can be used")
		fmt.Println("[print || p] employee - prints given employee with subordinates")
		fmt.Println("[newEmployee || new || n] employee - add  new employee (without relations)")
		fmt.Println("[addRelation || relation || r] manager employee - add a relation between two employees (manager-subortdinate)")
		fmt.Println("[findManager || find || f] employee1 employee2 - finds  closest common Manager ")

	
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
		fmt.Println("Closest Common Manager : ",color.CyanString(manager.Name))
		}
		// add another case here for custom commands.
	}
	
	return err
}


