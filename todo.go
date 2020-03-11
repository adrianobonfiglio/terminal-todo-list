package main

	
import (
	"io/ioutil"
	"os"
	"encoding/json"
	"fmt"
	"strconv"
)

const CLEAR_ARG = "-clear"
const CREATE_ARG = "-c"
const FINISH_ARG = "-f"
const LIST_ARG = "-l"
const JSON_FILE_NAME = "todo.json"

type Todo struct {
	Description string
	IsClosed bool
}

func checkError(e error) {
	if(e != nil) {
		panic(e)
	}
}

func finishTodo(args []string, tds *[]Todo) {
	todoInt, _ := strconv.ParseInt(args[2], 10, 0)
	(*tds)[todoInt].IsClosed = true
}

func writeJson(tds []Todo) {
	todo, err := json.Marshal(tds)
	checkError(err)
	e := ioutil.WriteFile(JSON_FILE_NAME, todo, 0644)
	checkError(e)
}

func readTodos(tds *[]Todo) {
	jsonresult, notFound := ioutil.ReadFile(JSON_FILE_NAME)
	if(notFound != nil) {
		os.Create(JSON_FILE_NAME)
	}
	json.Unmarshal(jsonresult, tds)
}

func createTodo(args []string, tds []Todo) {
	readTodos(&tds)
	description := args[2]
	var t = Todo{description, false}

	tds = append(tds, t)
	writeJson(tds)
}

func main() {
	var tds []Todo
	var firstArg string

	if len(os.Args) > 1 {
		firstArg = os.Args[1]
	}else {
		firstArg = LIST_ARG
	}

	switch(firstArg) {
	case CLEAR_ARG:
		os.Remove(JSON_FILE_NAME)
	case CREATE_ARG:
		createTodo(os.Args, tds)
		readTodos(&tds)
	case FINISH_ARG:
		readTodos(&tds)
		finishTodo(os.Args, &tds)
		writeJson(tds)
	case LIST_ARG:
		readTodos(&tds)
	}
	
	for i := 0; i < len(tds); i++ {
		todo := tds[i]
		if todo.IsClosed {
			fmt.Println(i, " [X] -", todo.Description)			
		} else {
			fmt.Println(i, " [ ] -", todo.Description)
		} 
	}

}
