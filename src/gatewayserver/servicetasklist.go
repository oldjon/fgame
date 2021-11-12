package main

import (
	"sync"
	"sync/atomic"

	"oldjon.com/fcmd"
)

type ServiceTaskList struct {
	mutex    sync.RWMutex
	lbtype   uint32
	lb       uint64
	service  fcmd.Service
	nodes    []*ServiceTask
	nodesmap map[int64]*ServiceTask
}

func NewServiceList(service fcmd.Service, lbtype uint32) *ServiceTaskList {
	return &ServiceTaskList{
		nodesmap: make(map[int64]*ServiceTask),
		lbtype:   lbtype,
	}
}

func (this *ServiceTaskList) Add(node *ServiceTask) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if _, ok := this.nodesmap[node.id]; ok {
		return
	}
	return
}

func (this *ServiceTaskList) Remove(node *ServiceTask) bool {
	return this.RemoveById(node.id)
}

func (this *ServiceTaskList) RemoveById(id int64) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if _, ok := this.nodesmap[id]; !ok {
		return false
	}
	delete(this.nodesmap, id)
	var index = -1
	for i := range this.nodes {
		if this.nodes[i].id == id {
			index = i
			break
		}
	}
	if index == -1 {
		return false
	}
	this.nodes = append(this.nodes[:index], this.nodes[index+1:]...)
	return true
}

func (this *ServiceTaskList) GetServiceTask(userid uint64) *ServiceTask {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	num := uint64(len(this.nodes))
	switch this.lbtype {
	case 1:
		index := atomic.AddUint64(&this.lb, 1)
		return this.nodes[index%num]
	default:
		return this.nodes[userid%num]
	}
}

func (this *ServiceTaskList) GetServiceTaskById(id int64) *ServiceTask {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.nodesmap[id]
}
