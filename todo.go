package main
import (
	"fmt"
	"strings"
	"gopkg.in/alecthomas/kingpin.v2"
	"time"
)

type TaskAdder struct{
	Command *kingpin.CmdClause
	Task *[]string
	Due *time.Duration
	Priority *int
}

type TaskLister struct{
	Command *kingpin.CmdClause
	From *int
	To *int
}

var (
	taskAdder  = &TaskAdder{}
	taskLister = &TaskLister{}
)


func initialize(){
	taskAdder.Command = kingpin.Command("add", "Add Tasks")
	taskAdder.Task = taskAdder.Command.Arg("task", "The Task").Strings()
	taskAdder.Due = taskAdder.Command.Flag("due", "Completion Time").Default("24h").Short('d').Duration()
	taskAdder.Priority = taskAdder.Command.Flag("priority", "Task Priority").Short('p').Default("0").Int()

	taskLister.Command  = kingpin.Command("list", "List All Tasks").Default()
	taskLister.From      = taskLister.Command.Flag("from", "From Task ID").Default("0").Short('f').Int()
	taskLister.To      = taskLister.Command.Flag("to", "To Task ID").Short('t').Int()
}

func main(){
	initialize()
	kingpin.Version("0.0.1")
	switch kingpin.Parse() {
	case taskAdder.Command.FullCommand():
		fmt.Print("Your Task Is : ", strings.Join(*taskAdder.Task, " "),", Priority: ",*taskAdder.Priority,", Duration: ", taskAdder.Due.Seconds())
	case taskLister.Command.FullCommand():
		fmt.Print("Listing ")
	}
}