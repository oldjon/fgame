package main

import (
	"strings"
	"sync"

	"oldjon.com/com"
	"oldjon.com/env"
	"oldjon.com/fcmd"
	"oldjon.com/glog"
	pb "oldjon.com/protobuf/proto"
)

type ServiceMgr struct {
	mutex          sync.RWMutex
	Services       map[fcmd.Service]*ServiceTaskList //可用的服务
	GrayServices   map[fcmd.Service]*ServiceTaskList //不可用的服务
	onlineservice  []fcmd.Service                    //玩家上线需要处理的服务
	offlineservice []fcmd.Service                    //玩家下线需要处理的服务
	lbtypes        map[fcmd.Service]uint32           //负载均衡类型
}

var servicemgr *ServiceMgr

func ServiceMgr_GetMe() *ServiceMgr {
	if servicemgr == nil {
		servicemgr = &ServiceMgr{
			Services:     make(map[fcmd.Service]*ServiceTaskList),
			GrayServices: make(map[fcmd.Service]*ServiceTaskList),
			lbtypes:      make(map[fcmd.Service]uint32),
		}
	}
	return servicemgr
}

func (this *ServiceMgr) Init() {
	return
}

// GetService
/*
获取服务节点
模式默认:根据id一致性hash
模式1:根据负载请求数
*/
func (this *ServiceMgr) GetService(service fcmd.Service, userid uint64) *ServiceTask {
	this.mutex.RLock()
	ss, ok := this.Services[service]
	this.mutex.RUnlock()
	if !ok {
		return nil
	}
	num := uint64(len(ss.nodes))
	if num == 0 {
		return nil
	}
	return ss.GetServiceTask(userid)
}

func (this *ServiceMgr) UserOnline(userid uint64) {
	msg := &fcmd.ReqServiceOnline{}
	data, err := pb.Marshal(msg)
	if err != nil {
		glog.Error("[协议] 上线协议序列化失败 ", userid)
		return
	}

	for _, service := range this.onlineservice {
		servicenode := this.GetService(service, userid)
		if servicenode == nil {
			continue
		}
		buff := com.CreateServiceMsg(service, fcmd.UCmd_ServiceOnline, 0, data)
		servicenode.Send(userid, buff)
	}
	return
}

func (this *ServiceMgr) UserOffline(userid uint64) {
	msg := &fcmd.ReqServiceOffline{}
	data, err := pb.Marshal(msg)
	if err != nil {
		glog.Error("[协议] 下线协议序列化失败 ", userid)
		return
	}

	for _, service := range this.offlineservice {
		servicenode := this.GetService(service, userid)
		if servicenode == nil {
			continue
		}
		buff := com.CreateServiceMsg(service, fcmd.UCmd_ServiceOffline, 0, data)
		servicenode.Send(userid, buff)
	}
	return
}

func (this *ServiceMgr) AddServiceNode(node *ServiceTask) bool {
	this.mutex.Lock()
	sl, ok := this.Services[node.service]
	if !ok {
		sl = NewServiceList(node.service, this.lbtypes[node.service])
		this.Services[node.service] = sl
	}
	this.mutex.Unlock()
	sl.Add(node)
	return true
}

func (this *ServiceMgr) RemoveServiceNode(node *ServiceTask) {
	this.mutex.Lock()
	sl, ok := this.Services[node.service]
	if !ok {
		sl = NewServiceList(node.service, this.lbtypes[node.service])
		this.Services[node.service] = sl
	}
	this.mutex.Unlock()
	sl.Add(node)
	return
}

//初始化代理的服务，启动时调用，无需加锁
func (this *ServiceMgr) LoadServices() bool {
	//连接中心服 TODO

	servicesinfo := &fcmd.RetLoadlServiceInfo{}

	//暂时从配置获取，后续由中心服分发 TODO
	servicestr := env.Get("gatewayserver", "services")
	if servicestr == "" {
		glog.Error("[启动] 未配置服务 ", servicestr)
		return true
	}
	for _, servicegroup := range strings.Split(servicestr, "|") {
		strs := strings.Split(servicegroup, ",")
		if len(strs) < 3 {
			glog.Error("[启动] 服务配置错误 ", servicestr)
			return false
		}
		sg := &fcmd.ServiceGroupInfo{
			Service: fcmd.Service(com.StrToUint32(strs[0])),
			LBType:  com.StrToUint32(strs[1]),
		}

		nodestrs := strings.Split(strs[2], "$")
		if len(nodestrs) != 2 {
			glog.Error("[启动] 服务配置错误 ", servicestr)
			return false
		}
		sn := &fcmd.ServiceNodeInfo{
			Id:   com.StrToInt64(nodestrs[0]),
			Addr: nodestrs[1],
		}
		sg.Nodes = append(sg.Nodes, sn)
		servicesinfo.ServiceGroup = append(servicesinfo.ServiceGroup, sg)
	}
	//暂时从配置获取，后续由中心服分发 TODO

	for _, group := range servicesinfo.ServiceGroup {
		this.lbtypes[group.Service] = group.LBType
		for _, service := range group.Nodes {
			s := NewServiceTask(service.Id, group.Service, service.Addr)
			if !s.Connect() {
				return false
			}
		}
	}
	return true
}
