package main

import (
	"oldjon.com/com"
	"oldjon.com/fcmd"
)

type PlayerMgr struct {
	//runmutex   sync.RWMutex
	PlayerList map[uint64]*Player //场景采用单协程处理
}

func PlayerMgr_GetMe() *PlayerMgr {
	scene := GetScene()
	if scene.playermgr == nil {
		scene.playermgr = &PlayerMgr{
			PlayerList: make(map[uint64]*Player),
		}
		scene.playermgr.Init()
	}
	return scene.playermgr
}

func (this *PlayerMgr) Init() {

}

func (this *PlayerMgr) ParseMsg(data []byte, flag byte) bool {
	cmd := fcmd.UCmd(com.GetCmd(data))

	switch cmd {
	case fcmd.UCmd_Login:
		reqcmd, ok := com.DecodeCmd(data, flag, &fcmd.ReqLogin{}).(*fcmd.ReqLogin)
		if !ok {
			return false
		}

		player := this.NewPlayer(reqcmd.UserId)
		GetScene().playerop <- &PlayerOp{optype: OPTP_AddPlayer, value: player}
	case fcmd.UCmd_Move:
		reqcmd, ok := com.DecodeCmd(data, flag, &fcmd.ReqMove{}).(*fcmd.ReqMove)
		if !ok {
			return false
		}
		playerop := &PlayerMove{UserId: reqcmd.UserId, Power: reqcmd.Power}

		GetScene().playerop <- &PlayerOp{optype: OPTP_AddPlayer, value: playerop}

	}
	return true
}

func (this *PlayerMgr) NewPlayer(userid uint64) *Player {
	return &Player{
		id:   userid,
		skin: com.RandBetweenUint32(1, 255)<<16 | com.RandBetweenUint32(1, 255)<<8 | com.RandBetweenUint32(1, 255), //随机一个颜色
		pos:  GetScene().RandPosInScene(),
	}
}

func (this *PlayerMgr) AddPlayer(player *Player) {
	this.PlayerList[player.id] = player
}

func (this *PlayerMgr) RemovePlayer(userid uint64) {
	delete(this.PlayerList, userid)
}

func (this *PlayerMgr) UpdatePos() {

}
