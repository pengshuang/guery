package Agent

import (
	"fmt"
	"sync"

	"github.com/xitongsys/guery/pb"
)

type TaskMap struct {
	sync.Mutex
	Tasks map[int64]*pb.Task
}

func NewTaskMap() *TaskMap {
	return &TaskMap{
		Tasks: make(map[int64]*pb.Task),
	}
}

func (self *TaskMap) HasTask(id int64) bool {
	self.Lock()
	defer self.Unlock()
	_, ok := self.Tasks[id]
	return ok
}

func (self *TaskMap) GetTask(id int64) *pb.Task {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Tasks[id]; ok {
		res := self.Tasks[id]
		return res
	} else {
		return nil
	}
}

func (self *TaskMap) GetTaskInfos() []*pb.TaskInfo {
	self.Lock()
	defer self.Unlock()
	res := make([]*pb.TaskInfo, 0)
	for _, task := range self.Tasks { //should copy?
		res = append(res, task.Info)
	}
	return res
}

func (self *TaskMap) GetTaskNumber() int32 {
	self.Lock()
	defer self.Unlock()
	return int32(len(self.Tasks))
}

func (self *TaskMap) PopTask(id int64) *pb.Task {
	self.Lock()
	defer self.Unlock()
	if task, ok := self.Tasks[id]; ok {
		delete(self.Tasks, id)
		return task

	} else {
		return nil
	}
}

func (self *TaskMap) AddTask(task *pb.Task) error {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Tasks[task.TaskId]; ok {
		return fmt.Errorf("task already exists")
	}
	self.Tasks[task.TaskId] = task
	return nil
}

func (self *TaskMap) DeleteTask(task *pb.Task) error {
	self.Lock()
	defer self.Unlock()
	if _, ok := self.Tasks[task.TaskId]; !ok {
		return nil
	}
	delete(self.Tasks, task.TaskId)
	return nil
}