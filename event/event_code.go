package event

const (
	CLIENT_NICKNAME_SET              = "CLIENT_NICKNAME_SET"              // 设置昵称
	CLIENT_EXIT                      = "CLIENT_EXIT"                      // 客户端退出
	CODE_CLIENT_KICK                 = "CODE_CLIENT_KICK"                 // 客户端被踢出
	CLIENT_CONNECT                   = "CLIENT_CONNECT"                   // 客户端加入成功
	SHOW_OPTIONS                     = "SHOW_OPTIONS"                     // 全局选项列表
	CODE_SHOW_OPTIONS_SETTING        = "CODE_SHOW_OPTIONS_SETTING"        // 设置选项
	SHOW_OPTIONS_PVP                 = "SHOW_OPTIONS_PVP"                 // 玩家对战选项
	CODE_SHOW_OPTIONS_PVE            = "CODE_SHOW_OPTIONS_PVE"            // 人机对战选项
	SHOW_ROOMS                       = "SHOW_ROOMS"                       // 展示房间列表
	SHOW_POKERS                      = "SHOW_POKERS"                      // 展示Poker
	ROOM_CREATE_SUCCESS              = "ROOM_CREATE_SUCCESS"              // 创建房间成功
	ROOM_JOIN_SUCCESS                = "ROOM_JOIN_SUCCESS"                // 加入房间成功
	ROOM_JOIN_FAIL_BY_FULL           = "ROOM_JOIN_FAIL_BY_FULL"           // 房间人数已满
	ROOM_JOIN_FAIL_BY_INEXIST        = "ROOM_JOIN_FAIL_BY_INEXIST"        // 加入-房间不存在
	GAME_STARTING                    = "GAME_STARTING"                    // 开始游戏
	GAME_LANDLORD_ELECT              = "GAME_LANDLORD_ELECT"              // 抢地主
	GAME_LANDLORD_CONFIRM            = "GAME_LANDLORD_CONFIRM"            // 地主确认
	GAME_LANDLORD_CYCLE              = "GAME_LANDLORD_CYCLE"              // 地主一轮确认结束
	CODE_GAME_POKER_PLAY             = "CODE_GAME_POKER_PLAY"             // 出牌回合
	GAME_POKER_PLAY_REDIRECT         = "GAME_POKER_PLAY_REDIRECT"         // 出牌重定向
	GAME_POKER_PLAY_MISMATCH         = "GAME_POKER_PLAY_MISMATCH"         // 出牌不匹配
	CODE_GAME_POKER_PLAY_LESS        = "CODE_GAME_POKER_PLAY_LESS"        // 出牌太小
	GAME_POKER_PLAY_PASS             = "GAME_POKER_PLAY_PASS"             // 不出
	GAME_POKER_PLAY_CANT_PASS        = "GAME_POKER_PLAY_CANT_PASS"        // 不允许不出
	CODE_GAME_POKER_PLAY_INVALID     = "CODE_GAME_POKER_PLAY_INVALID"     // 无效
	CODE_GAME_POKER_PLAY_ORDER_ERROR = "CODE_GAME_POKER_PLAY_ORDER_ERROR" // 顺序错误
	GAME_OVER                        = "GAME_OVER"                        // 游戏结束
	CODE_PVE_DIFFICULTY_NOT_SUPPORT  = "CODE_PVE_DIFFICULTY_NOT_SUPPORT"  // 人机难度不支持
	ROOM_OWNER_SELECT                = "ROOM_OWNER_SELECT"                //重新开始游戏
)

const (
	SERVER_CODE_CLIENT_EXIT              = "CODE_CLIENT_EXIT"              // 玩家退出
	SERVER_CODE_CLIENT_OFFLINE           = "CODE_CLIENT_OFFLINE"           // 玩家离线
	SERVER_CODE_CLIENT_NICKNAME_SET      = "CODE_CLIENT_NICKNAME_SET"      // 设置昵称
	SERVER_CODE_CLIENT_HEAD_BEAT         = "CODE_CLIENT_HEAD_BEAT"         // 不出
	SERVER_CODE_ROOM_CREATE              = "CODE_ROOM_CREATE"              // 创建PVP房间
	SERVER_CODE_ROOM_CREATE_PVE          = "CODE_ROOM_CREATE_PVE"          // 创建PVE房间
	SERVER_CODE_GET_ROOMS                = "CODE_GET_ROOMS"                // 获取房间列表
	SERVER_CODE_ROOM_JOIN                = "CODE_ROOM_JOIN"                // 加入房间
	SERVER_CODE_ROOM_DISBAND             = "CODE_ROOM_DISBAND"             // 解散房间
	SERVER_CODE_GAME_STARTING            = "CODE_GAME_STARTING"            // 游戏开始
	SERVER_CODE_GAME_LANDLORD_ELECT      = "CODE_GAME_LANDLORD_ELECT"      // 抢地主
	SERVER_CODE_GAME_POKER_PLAY          = "CODE_GAME_POKER_PLAY"          // 出牌环节
	SERVER_CODE_GAME_POKER_PLAY_REDIRECT = "CODE_GAME_POKER_PLAY_REDIRECT" // 出牌重定向
	SERVER_CODE_GAME_POKER_PLAY_PASS     = "CODE_GAME_POKER_PLAY_PASS"     // 不出
)
