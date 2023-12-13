package main

import "time"

//:TODO: schedule task

type Pomodoro struct {

	//name of task created by user
	TaskName string

	//start time of the task created by the user

	StartTime time.Time
}

var currentTask Pomodoro = Pomodoro{}

func AddTask(taskName string) Pomodoro {
	return Pomodoro{
		TaskName:  taskName,
		StartTime: time.Now(),
	}
}

func PrintTask(task Pomodoro) string {
	if (Pomodoro{} == task) {
		return "Current task is empty"
	} else {
		return "Task - " + task.TaskName
	}

}
