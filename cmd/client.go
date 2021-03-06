package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	todo "github.com/kolodach/golang-todo"
)

const (
	helpCmd   = "help"
	listCmd   = "list"
	addCmd    = "add"
	statusCmd = "status"
)

// Client represents command line client.
type Client struct {
	serv todo.TodoService
}

// CreateClient creates new instanse of command line client and returns pointer on it.
func CreateClient(s todo.TodoService) *Client {
	return &Client{
		serv: s,
	}
}

// Process parses user command and retreives result.
func (c *Client) Process() {
	if len(os.Args) < 2 {
		fmt.Printf("Command is required. Type todo help to see available commands.")
		os.Exit(1)
	}

	helpFS := flag.NewFlagSet(helpCmd, flag.ExitOnError)
	listFS := flag.NewFlagSet(listCmd, flag.ExitOnError)
	addFS := flag.NewFlagSet(addCmd, flag.ExitOnError)
	statusFS := flag.NewFlagSet(statusCmd, flag.ExitOnError)

	switch os.Args[1] {
	case helpCmd:
		helpFS.Parse(os.Args[2:])
	case listCmd:
		listFS.Parse(os.Args[2:])
	case addCmd:
		addFS.Parse(os.Args[2:])
	case statusCmd:
		statusFS.Parse(os.Args[2:])
	}

	switch {
	case helpFS.Parsed():
		printHelp()
		os.Exit(0)
	case listFS.Parsed():
		printList(c)
		os.Exit(0)
	case addFS.Parsed():
		addTodo(c)
		os.Exit(0)
	case statusFS.Parsed():
		setStatus(c)
		os.Exit(0)
	}
}

// Help
func printHelp() {
	if len(os.Args) > 3 {
		fmt.Printf("Too many arguments provided.")
		os.Exit(1)
	}

	if len(os.Args) == 2 {
		printMainHelp()
		os.Exit(0)
	}

	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "list":
			printListHelp()
			os.Exit(0)
		case "add":
			printAddHelp()
			os.Exit(0)
		case "status":
			printStatusHelp()
			os.Exit(0)
		default:
			fmt.Printf("\nThere is no command such as %s. Type todo help to see list of available commands.\n", os.Args[2])
			os.Exit(0)
		}
	}
}

func printMainHelp() {
	sb := &strings.Builder{}
	sb.WriteString("Todo is a tool for managing your task list.")
	sb.WriteString("\n\nUsage:\n\ttodo <command> [arguments]")
	sb.WriteString("\n\nAvailable commands are:")
	sb.WriteString("\n\tlist\tView your agenda")
	sb.WriteString("\n\tadd\tCreate new todo item")
	sb.WriteString("\n\tstatus\tSwitch todo status")
	sb.WriteString("\n\nUse todo help <command> for more information about the command.")
	sb.WriteString("\n\n")
	fmt.Printf(sb.String())
}

func printListHelp() {
	sb := &strings.Builder{}
	sb.WriteString("Usage: todo list [arguments...]")
	sb.WriteString("\n\nList displays todo items.")
	sb.WriteString("\n\nBy default list displays all todo items, but it can be filtered using following args:")
	sb.WriteString("\n\t-p\tdisplays pending only")
	sb.WriteString("\n\t-d\tdisplays done only")
	sb.WriteString("\n\t-i\tdisplays in progress only")
	sb.WriteString("\n\nNote that flags can be combined.")
	sb.WriteString("\n\n")
	fmt.Printf(sb.String())
}

func printAddHelp() {
	sb := &strings.Builder{}
	sb.WriteString("Usage: todo add \"<name>\"")
	sb.WriteString("\n\nInserts specific todo item into agenda. Default state for new item is Pending.")
	sb.WriteString("\n<name> parameter is required and should be unique, otherwise it wount be added.")
	sb.WriteString("\n\n")
	fmt.Printf(sb.String())
}

func printStatusHelp() {
	sb := &strings.Builder{}
	sb.WriteString("Usage: todo status <id> <status-flag>.")
	sb.WriteString("\nExample:")
	sb.WriteString("\n\n\ttodo status 23 -d")
	sb.WriteString("\n\nSwitches todo item state.")
	sb.WriteString("\nAvailable statuses are:")
	sb.WriteString("\n\t-p\tpending")
	sb.WriteString("\n\t-d\tdone")
	sb.WriteString("\n\t-i\tin progress")
	sb.WriteString("\n\n<id> and <status-flag> are both required.")
	sb.WriteString("\n\n")
	fmt.Printf(sb.String())
}

// List
func printList(c *Client) {
	todos := c.serv.All()
	if todos == nil || len(todos) == 0 {
		fmt.Printf("\nNothing to show\n")
		return
	}
	sb := &strings.Builder{}
	sb.WriteString("\nAgenda:\n")
	sb.WriteString("\tid\tname\tstatus")
	sb.WriteString("\n\t--\t----\t------")
	for _, e := range todos {
		sb.WriteString(fmt.Sprintf("\n\t%s\t%s\t%s", e.ID, e.Name, e.Status.String()))
	}
	sb.WriteString("\n\n")
	fmt.Printf(sb.String())
}

// Add
func addTodo(c *Client) {
	if len(os.Args) != 3 {
		fmt.Println("Invalid arguments count. Expected todo add \"<name>\". Type todo help add to read more.")
		return
	}
	name := os.Args[2]
	fmt.Println(name)
	todo := &todo.Todo{
		ID:     "0",
		Name:   name,
		Status: todo.Status(todo.Pending),
	}
	err := c.serv.Create(todo)
	if err != nil {
		fmt.Printf("Unable to create todo item. Reason: %s\n", err.Error())
		return
	}
	printList(c)
}

// Status
func setStatus(c *Client) {
	if len(os.Args) < 4 {
		fmt.Println("Not enought parms provided. Type todo help status to view docs.")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("Too many params provided. Type todo help status to view docs.")
		return
	}
	statePrefMap := map[string]todo.Status{
		"-p": todo.Status(todo.Pending),
		"-d": todo.Status(todo.Done),
		"-i": todo.Status(todo.InProgress),
	}
	id := os.Args[2]
	todo, err := c.serv.ByID(id)
	if err != nil {
		fmt.Printf("Can not find todo with id:%s", id)
		return
	}
	newState, exists := statePrefMap[os.Args[3]]
	if !exists {
		fmt.Printf("Invalid state %s. Type \n\ttodo help state\n to see available statuses.", os.Args[3])
		return
	}
	if todo.Status != newState {
		c.serv.Status(id, newState)
	}
	fmt.Printf("\n\t%s\t%s\t%s\n", todo.ID, todo.Name, todo.Status.String())
}
	