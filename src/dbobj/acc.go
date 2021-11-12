package db

type Account struct {
	Id       uint64 `redis:"id"`       //玩家id
	Password string `redis:"password"` //密码
	IsBind   uint32 `redis:"isbind"`   //是否绑定
}

type GameData struct {
	Account   string `redis:"account"`   //账号
	Device    string `redis:"device"`    //设备号
	Age       uint32 `redis:"age"`       //年龄
	Sex       uint32 `redis:"sex"`       // 0男1女
	HeadIcon  uint32 `redis:"headicon"`  //默认头像
	HeadUrl   string `redis:"headurl"`   //头像url
	RegTime   int64  `redis:"regtime"`   //注册时间
	FollowNum uint32 `redis:"follownum"` //关注数
	FansNum   uint32 `redis:"fansnum"`   //粉丝数
	FriendNum uint32 `redis:"friendnum"` //好友数
	Platform  uint32 `redis:"platform"`  //平台
	Tel       string `redis:"tel"`       //电话
	OfflineTS int64  `redis:"offlinets"` //最后线下时间
}
