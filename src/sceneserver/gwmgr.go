package main

import (
	"sync"
)

type GWMgr struct {
	runmutex sync.RWMutex
	GWList   map[uint32]*GWTask
}

var gwmgr *GWMgr

func GWMgr_GetMe() *GWMgr {
	if gwmgr == nil {
		gwmgr = &GWMgr{
			GWList: make(map[uint32]*GWTask),
		}
		gwmgr.Init()
	}
	return gwmgr
}

func (this *GWMgr) Init() {

}

func (this *GWMgr) AddGWTask(gwtask *GWTask) {
	this.runmutex.Lock()
	defer this.runmutex.Unlock()
	oldgwtask, ok := this.GWList[gwtask.GWId]
	if ok {
		go oldgwtask.Close()
	}
	this.GWList[gwtask.GWId] = gwtask
}

func (this *GWMgr) RemoveGWTask(gwtask *GWTask) bool {
	this.runmutex.Lock()
	defer this.runmutex.Unlock()
	delete(this.GWList, gwtask.GWId)
	go gwtask.Close()
	return true
}
