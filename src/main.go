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
	list1 := &taskList{tasks: []*task{t1, t2, t3}}
	// list.addToList(t3)
	// list.tasks[0].checkComplete()
	// fmt.Println("Tareas completadas")
	// list.printListComplete()

	t4 := &task{
		name:        "Completar curso C#",
		description: "Completar curso C# en esta semana",
		complete:    false,
	}
	t5 := &task{
		name:        "Completar curso Rust",
		description: "Completar curso Rust en esta semana",
		complete:    false,
	}

	list2 := &taskList{tasks: []*task{t4, t5}}

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Edilberto"] = list1

	mapaTareas["Maria"] = list2

	fmt.Println("Tareas de Edilberto")
	mapaTareas["Edilberto"].printList()

	fmt.Println("<---------->")
	fmt.Println("Tareas de Maria")
	mapaTareas["Maria"].printList()
}
