package main

import (
	"time"

	"oldjon.com/base/env"
	"oldjon.com/base/glog"
	"oldjon.com/com"
)

const (
	_ = iota
	OPTP_AddPlayer
	OPTP_Move
)

type Scene struct {
	id          uint32  //场景id
	zero        *Vector //坐标零点
	scenewidth  float64
	sceneheight float64
	now         time.Time
	timeloop    uint32
	playermgr   *PlayerMgr
	playerop    chan *PlayerOp
}

var scene *Scene

func GetScene() *Scene {
	if scene == nil {
		scene = &Scene{}
		scene.Init(com.StringToUint32(env.Get("sceneserver", "sceneid")), Vector{0, 0}, 10, 10)
		//TODO 根据sceneid从配置中读取相应配置初始化scene
		go scene.Loop()
	}
	return scene
}

func (this *Scene) Init(id uint32, v Vector, width, height float64) *Scene {
	return &Scene{id: id,
		zero:        &v,
		scenewidth:  width,
		sceneheight: height,
		playerop:    make(chan *PlayerOp),
	}
}

func (this *Scene) RandPosInScene() *Position {
	return &Position{
		centre: &Vector{scene.zero.x + com.RandBetweenFloat(0, scene.scenewidth),
			scene.zero.y + com.RandBetweenFloat(0, scene.sceneheight)},
		radius: 1,
	}
}

func (this *Scene) Stop() { //关闭场景处理
}

func (this *Scene) Loop() {
	var timeTicker = time.NewTicker(time.Millisecond)
	defer func() {
		this.Stop()
		timeTicker.Stop()
		if err := recover(); err != nil {
			glog.Error("[异常] 场景线程出错")
		}
	}()
	for {
		this.now = time.Now()
		select {
		case <-timeTicker.C:
			{ //处理定时计算任务
				//0.05s  //逻辑处理帧时间可以放大
				if this.timeloop%50 == 0 {
					//更新玩家位置
					this.playermgr.UpdatePos()
				}

				if this.timeloop%100 == 0 {
					//广播场景边缘玩家
					this.BroadcastBorderPlayers()
					//广播玩家信息（位置状态）
					this.BroadcastPlayerInfo()
				}

				this.timeloop++
				if this.timeloop == 1000000 {
					this.timeloop = 0
				}
			}
		case playerop := <-this.playerop:
			{
				switch playerop.optype {
				case OPTP_AddPlayer:
					player, ok := playerop.value.(*Player)
					if !ok || player == nil {
						break
					}
					this.playermgr.PlayerList[player.id] = player
				case OPTP_Move:
					move, ok := playerop.value.(*PlayerMove)
					if !ok || move == nil {
						break
					}
					player, ok := this.playermgr.PlayerList[move.UserId]
					if !ok || player == nil {
						break
					}
					player.power = move.Power
					player.way = move.Way
					player.movetime = this.now.UnixNano() / 1000000 //单位毫秒
				}
			}
		}
	}
}

func (this *Scene) BroadcastPlayerInfo() {

}

func (this *Scene) BroadcastBorderPlayers() {

}
