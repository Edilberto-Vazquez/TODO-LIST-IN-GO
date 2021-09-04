package main

import "fmt"

type task struct {
	name        string
	description string
	complete    bool
}

type taskList struct {
	tasks []*task
}

func (t *task) updateName(name string) {
	t.name = name
}
func (t *task) updateDescription(description string) {
	t.description = description
}
func (t *task) checkComplete() {
	t.complete = true
}

func (t *taskList) deleteFromList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (tl *taskList) addToList(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (t *taskList) printList() {
	for _, task := range t.tasks {
		fmt.Println("Nombre", task.name)
		fmt.Println("Description", task.description)
		fmt.Println("Complete", task.complete)
	}
}

func (t *taskList) printListComplete() {
	for _, task := range t.tasks {
		if task.complete == true {
			fmt.Println("Nombre", task.name)
			fmt.Println("Description", task.description)
			fmt.Println("Complete", task.complete)
		}
	}
}

func main() {
	t1 := &task{
		name:        "Completar curso Go",
		description: "Completar curso Go en esta semana",
		complete:    false,
	}
	t2 := &task{
		name:        "Completar curso Python",
		description: "Completar curso Python en esta semana",
		complete:    false,
	}
	t3 := &task{
		name:        "Completar curso JavaScript",
		description: "Completar curso JavaScript en esta semana",
		complete:    false,
	}
	list := taskList{tasks: []*task{t1, t2}}
	list.addToList(t3)
	list.tasks[0].checkComplete()
	fmt.Println("Tareas completadas")
	list.printListComplete()
}
