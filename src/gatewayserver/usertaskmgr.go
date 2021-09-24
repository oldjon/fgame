package main

import (
	"runtime/debug"
	"sync"
	"time"

	"oldjon.com/glog"
)

type UserTaskMgr struct {
	mutex sync.RWMutex
	users map[uint64]*UserTask
}

var userTaskMgr *UserTaskMgr

func UserTaskMgr_GetMe() *UserTaskMgr {
	if userTaskMgr == nil {
		userTaskMgr = &UserTaskMgr{
			users: make(map[uint64]*UserTask),
		}
	}
	return userTaskMgr
}

func (this *UserTaskMgr) Init() {
	go this.TimeLoop()
	return
}

func (this *UserTaskMgr) TimeLoop() {
	tick := time.NewTimer(time.Second * 5)
	defer func() {
		tick.Stop()
		if err := recover(); err != nil {
			glog.Error("[异常] ", err, "\n", string(debug.Stack()))
		}
	}()
	for {
		select {
		case <-tick.C:
			this.Tick()
		}
	}
	return
}

func (this *UserTaskMgr) Tick() {
	var (
		deltasks []*UserTask
		nowunix  = time.Now().Unix()
	)
	this.mutex.RLock()
	for _, t := range this.users {
		if t.activeTime == 0 {
			continue
		}
		if t.activeTime+Task_Max_TimeOut < nowunix {
			deltasks = append(deltasks, t)
			if len(deltasks) >= 200 {
				break
			}
		}
	}
	this.mutex.RUnlock()
	for _, t := range deltasks {
		if !t.Stop() {
			this.Remove(t)
		}
		glog.Info("[玩家] 连接超时 ", t.Id)
	}
	return
}

func (this *UserTaskMgr) Add(task *UserTask) {
	this.mutex.Lock()
	this.users[task.Id] = task
	this.mutex.Unlock()
	return
}

func (this *UserTaskMgr) GetNum() uint32 {
	this.mutex.RLock()
	l := uint32(len(this.users))
	this.mutex.RUnlock()
	return l
}

func (this *UserTaskMgr) Replace(task *UserTask) *UserTask {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	t := this.users[task.Id]
	this.users[task.Id] = task
	return t
}

func (this *UserTaskMgr) Remove(task *UserTask) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	_, ok := this.users[task.Id]
	if !ok {
		return false
	}
	delete(this.users, task.Id)
	return true
}

func (this *UserTaskMgr) Get(id uint64) *UserTask {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	t := this.users[id]
	return t
}

func (this *UserTaskMgr) GetAll() (users []*UserTask) {

	this.mutex.RLock()
	defer this.mutex.RUnlock()
	users = make([]*UserTask, len(this.users))
	i := 0
	for _, task := range this.users {
		users[i] = task
		i++
	}
	return users
}

func (this *UserTaskMgr) CloseAll() {
	users := this.GetAll()
	for _, user := range users {
		user.Close()
	}
	return
}

func (this *UserTaskMgr) SendBytes(id uint64, data []byte) bool {
	task := this.Get(id)
	if task == nil {
		glog.Error("[发送] 找不到玩家 ", id, ",", len(data))
		return false
	}
	task.SendBytes(data)
	return true
}
