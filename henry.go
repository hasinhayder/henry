package main

import (
	"fmt"
	"strings"
	"gopkg.in/alecthomas/kingpin.v2"
	"time"
	"github.com/nanobox-io/golang-scribble"
	"./locker"
	"./utility"
	"io/ioutil"
)

type (
	TaskAdder struct {
		Command  *kingpin.CmdClause
		Task     *[]string
		Due      *time.Duration
		Priority *int
	}

	TaskLister struct {
		Command *kingpin.CmdClause
		From    *int
		To      *int
	}

	Task struct {
		ID       int
		Title    string
		Timeout  float64
		Priority int
	}
)

var (
	taskAdder   = &TaskAdder{}
	taskLister  = &TaskLister{}
	db          = &scribble.Driver{}
	counter     = 1
	key         = []byte("Fuck The World-**-Duck The World")
	dbpath      = "/tmp/henry"
	pwdpath     = "/tmp/henry/henry-password.txt"
	counterpath = "/tmp/henry/henry-counter.txt"
)

func initialize() {
	taskAdder.Command = kingpin.Command("add", "Add Tasks")
	taskAdder.Task = taskAdder.Command.Arg("task", "The Task").Strings()
	taskAdder.Due = taskAdder.Command.Flag("due", "Completion Time").Default("24h").Short('d').Duration()
	taskAdder.Priority = taskAdder.Command.Flag("priority", "Task Priority").Short('p').Default("0").Int()

	taskLister.Command = kingpin.Command("list", "List All Tasks").Default()
	taskLister.From = taskLister.Command.Flag("from", "From Task ID").Default("0").Short('f').Int()
	taskLister.To = taskLister.Command.Flag("to", "To Task ID").Short('t').Int()

	if (!utility.DoesFileExist(pwdpath)) {
		ioutil.WriteFile(pwdpath, []byte("He Who Must Not Be Named"), 0644)
	}

	if (!utility.DoesFileExist(counterpath)) {
		ioutil.WriteFile(counterpath, []byte("1"), 0644)
	} else {
		counter = utility.GetCounter(counterpath)
	}
}

func main() {
	initialize()
	db, _ := scribble.New(dbpath, nil)

	kingpin.Version("0.0.1")
	switch kingpin.Parse() {
	case taskAdder.Command.FullCommand():
		fmt.Print("Your Task Is : ", strings.Join(*taskAdder.Task, " "), ", Priority: ", *taskAdder.Priority, ", Duration: ", taskAdder.Due.Seconds())
		var task = &Task{}
		task.Title = locker.Encrypt(key, strings.Join(*taskAdder.Task, " "))
		task.Priority = *taskAdder.Priority
		task.Timeout = taskAdder.Due.Seconds()
		task.ID = 1

		db.Write("tasks", "task-"+utility.IntToString(counter), task)
		utility.UpdateCounter(counterpath,counter)

	case taskLister.Command.FullCommand():
		fmt.Print("Listing ")
	}
}
