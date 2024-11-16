package main

import (
	"fmt"
	"time"
)

type Task struct {
	TaskType     string
	ScheduleTime int64
	TaskData     interface{}
}

type TaskScheduler struct {
	Tasks []Task
}

func (scheduler *TaskScheduler) AddTask(task Task) {
	scheduler.Tasks = append(scheduler.Tasks, task)
}

func (scheduler *TaskScheduler) DeleteTask(task Task) {
	for i, t := range scheduler.Tasks {
		if t == task {
			scheduler.Tasks = append(scheduler.Tasks[:i], scheduler.Tasks[i+1:]...)
			break
		}
	}
}

func (scheduler *TaskScheduler) ExecuteTasks() {
	for {
		currentTime := time.Now().UnixNano() / int64(time.Millisecond)

		for _, task := range scheduler.Tasks {
			if task.TaskType == "one-time" {
				if task.ScheduleTime <= currentTime {
					scheduler.ExecuteTask(task)
					scheduler.DeleteTask(task)
				}
			} else if task.TaskType == "recurring" {
				if currentTime%task.ScheduleTime == 0 {
					scheduler.ExecuteTask(task)
				}
			}
		}

		time.Sleep(time.Millisecond)
	}
}

func (scheduler *TaskScheduler) ExecuteTask(task Task) {
	// Perform task execution here, e.g. printing 'Hello World!'
	fmt.Println("Hello World!")
}

func main() {
	scheduler := TaskScheduler{}

	// Add a task to output 'Hello World!' every 10 milliseconds
	task1 := Task{
		TaskType:     "recurring",
		ScheduleTime: 10,
	}
	scheduler.AddTask(task1)

	// Add a task to output 'Hello World!' on the 15th day of every month at 19:55:55.270
	task2 := Task{
		TaskType:     "recurring",
		ScheduleTime: 50155755270, // milliseconds equivalent of the desired time
	}
	scheduler.AddTask(task2)

	// Start executing tasks
	scheduler.ExecuteTasks()
}
