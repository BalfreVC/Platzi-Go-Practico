package main

import "fmt"

type taskList struct {
	// Crea un SLICE de tipo TASK
	tasks []task
}

func (l *taskList) agregarTarea(t task) {
	l.tasks = append(l.tasks, t)
}
func (l *taskList) eliminarTarea(indice int) {
	l.tasks = append(l.tasks[:indice], l.tasks[indice+1:]...)
}
func (l taskList) muestraTareas() {
	// for i := 0; i < len(l.tasks); i++ {
	// 	fmt.Printf("%v", l.tasks[i].descripcion)
	// 	if l.tasks[i].completado {
	// 		fmt.Println(" FINALIZADA")
	// 	} else {
	// 		fmt.Println(" PENDIENTE")
	// 	}
	// }

	for indice, tarea := range l.tasks {
		fmt.Printf("%d .- %v", indice, tarea.descripcion)
		if tarea.completado {
			fmt.Println(" FINALIZADA")
		} else {
			fmt.Println(" PENDIENTE")
		}
	}

	// for _, tarea := range l.tasks {
	// 	fmt.Printf("%v", tarea.descripcion)
	// 	if tarea.completado {
	// 		fmt.Println(" FINALIZADA")
	// 	} else {
	// 		fmt.Println(" PENDIENTE")
	// 	}
	// }
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}
func (t *task) actDescripcion(descripcion string) {
	t.descripcion = descripcion
}
func (t *task) actNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := task{"t1", "Primera", false}
	fmt.Printf("%v\n", t1)
	t2 := task{"t2", "Segunda", false}
	fmt.Printf("%v\n", t2)
	t3 := task{"t3", "Tercera", false}
	fmt.Printf("%v\n", t3)

	lista := taskList{
		[]task{
			t1, t2,
		},
	}
	fmt.Println(lista)
	fmt.Println(len(lista.tasks))
	lista.agregarTarea(t3)
	lista.muestraTareas()
	// fmt.Println(len(lista.tasks))
	// fmt.Println(lista)
	// lista.eliminarTarea(1)
	// fmt.Println(len(lista.tasks))
	// fmt.Println(lista)

}
