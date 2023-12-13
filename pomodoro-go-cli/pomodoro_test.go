package main

import (
	"testing"
)

func TestAddTask(t *testing.T) {

	var taskName string = "Build Go CLI app"

	var currentTask Pomodoro = AddTask(taskName)

	if currentTask.TaskName != taskName {
		t.Errorf("AddTask failed - Expected: %v, Go: %v", taskName, currentTask.TaskName)
	}

}

func TestPrintTask_WithTask(t *testing.T) {
	taskName := "New Task"

	newTask := AddTask("New Task")

	printedOutPut := PrintTask(newTask)

	expectedPrintedOutput := "Task - " + taskName

	if printedOutPut != expectedPrintedOutput {
		t.Errorf("PrintTask failed - Expected: %v, Got: %v", expectedPrintedOutput, printedOutPut)
	}
}

func TestPrintTask_WithEmptyTask(t *testing.T) {
	printedOutPut := PrintTask(Pomodoro{})

	expectedPrintedOutput := "Current task is empty"

	if printedOutPut != expectedPrintedOutput {
		t.Errorf("PrintTask failed - Expected: %v, Got: %v", expectedPrintedOutput, printedOutPut)
	}
}
