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
func (l taskList) muestraTareas(tipo int) { // 0 = Todas, 1 = Completas, 2 = Pendintes
	// for i := 0; i < len(l.tasks); i++ {
	// 	fmt.Printf("%v", l.tasks[i].descripcion)
	// 	if l.tasks[i].completado {
	// 		fmt.Println(" FINALIZADA")
	// 	} else {
	// 		fmt.Println(" PENDIENTE")
	// 	}
	// }

	// for indice, tarea := range l.tasks {
	// 	fmt.Printf("%d .- %v", indice, tarea.descripcion)
	// 	if tarea.completado {
	// 		fmt.Println(" FINALIZADA")
	// 	} else {
	// 		fmt.Println(" PENDIENTE")
	// 	}
	// }

	for _, tarea := range l.tasks {
		if tipo == 1 && tarea.completado == false {
			continue
		}
		if tipo == 2 && tarea.completado == true {
			continue
		}
		fmt.Printf("%v", tarea.descripcion)
		if tarea.completado {
			fmt.Println(" FINALIZADA")
		} else {
			fmt.Println(" PENDIENTE")
		}
	}
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
	t2 := task{"t2", "Segunda", true}
	fmt.Printf("%v\n", t2)
	t3 := task{"t3", "Tercera", false}
	fmt.Printf("%v\n", t3)
	t4 := task{"t4", "Cuarta", false}

	lista := taskList{
		[]task{
			t1, t2,
		},
	}
	fmt.Println(lista)
	fmt.Println(len(lista.tasks))
	lista.agregarTarea(t3)
	lista.muestraTareas(0)
	// fmt.Println(len(lista.tasks))
	// fmt.Println(lista)
	// lista.eliminarTarea(1)
	// fmt.Println(len(lista.tasks))
	// fmt.Println(lista)

	// Crear una lista de tareas por PERSONA
	tareasPorPersona := make(map[string]taskList)
	tareasPorPersona["balfre"] = lista

	lista2 := taskList{
		[]task{
			t3, t4,
		},
	}

	tareasPorPersona["ivana"] = lista2
	fmt.Println(tareasPorPersona)
	fmt.Println("Tareas de Balfre")
	tareasPorPersona["balfre"].muestraTareas(0)
	fmt.Println("Tareas de Ivana")
	tareasPorPersona["ivana"].muestraTareas(0)

}
