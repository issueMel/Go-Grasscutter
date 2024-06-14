package pros

type PlayerProperty struct {
	ID           int
	Min          int
	Max          int
	DynamicRange bool
}

const (
	PROP_NONE                            = 0
	PROP_EXP                             = 1001
	PROP_BREAK_LEVEL                     = 1002
	PROP_SATIATION_VAL                   = 1003
	PROP_SATIATION_PENALTY_TIME          = 1004
	PROP_LEVEL                           = 4001
	PROP_LAST_CHANGE_AVATAR_TIME         = 10001
	PROP_MAX_SPRING_VOLUME               = 10002
	PROP_CUR_SPRING_VOLUME               = 10003
	PROP_IS_SPRING_AUTO_USE              = 10004
	PROP_SPRING_AUTO_USE_PERCENT         = 10005
	PROP_IS_FLYABLE                      = 10006
	PROP_IS_WEATHER_LOCKED               = 10007
	PROP_IS_GAME_TIME_LOCKED             = 10008
	PROP_IS_TRANSFERABLE                 = 10009
	PROP_MAX_STAMINA                     = 10010
	PROP_CUR_PERSIST_STAMINA             = 10011
	PROP_CUR_TEMPORARY_STAMINA           = 10012
	PROP_PLAYER_LEVEL                    = 10013
	PROP_PLAYER_EXP                      = 10014
	PROP_PLAYER_HCOIN                    = 10015
	PROP_PLAYER_SCOIN                    = 10016
	PROP_PLAYER_MP_SETTING_TYPE          = 10017
	PROP_IS_MP_MODE_AVAILABLE            = 10018
	PROP_PLAYER_WORLD_LEVEL              = 10019
	PROP_PLAYER_RESIN                    = 10020
	PROP_PLAYER_WAIT_SUB_HCOIN           = 10022
	PROP_PLAYER_WAIT_SUB_SCOIN           = 10023
	PROP_IS_ONLY_MP_WITH_PS_PLAYER       = 10024
	PROP_PLAYER_MCOIN                    = 10025
	PROP_PLAYER_WAIT_SUB_MCOIN           = 10026
	PROP_PLAYER_LEGENDARY_KEY            = 10027
	PROP_IS_HAS_FIRST_SHARE              = 10028
	PROP_PLAYER_FORGE_POINT              = 10029
	PROP_CUR_CLIMATE_METER               = 10035
	PROP_CUR_CLIMATE_TYPE                = 10036
	PROP_CUR_CLIMATE_AREA_ID             = 10037
	PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE   = 10038
	PROP_PLAYER_WORLD_LEVEL_LIMIT        = 10039
	PROP_PLAYER_WORLD_LEVEL_ADJUST_CD    = 10040
	PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM = 10041
	PROP_PLAYER_HOME_COIN                = 10042
	PROP_PLAYER_WAIT_SUB_HOME_COIN       = 10043
	PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP   = 10044
	PROP_PLAYER_GCG_COIN                 = 10045
	PROP_PLAYER_WAIT_SUB_GCG_COIN        = 10046
	PROP_PLAYER_ONLINE_TIME              = 10047
	PROP_PLAYER_CAN_DIVE                 = 10048
	PROP_DIVE_MAX_STAMINA                = 10049
	PROP_DIVE_CUR_STAMINA                = 10050
)

var (
	propMap = map[int]*PlayerProperty{}
)

func init() {
	propMap[0] = &PlayerProperty{0, -1 << 31, 1<<31 - 1, false}
	propMap[1001] = &PlayerProperty{1001, 0, 1<<31 - 1, false}
	propMap[1002] = &PlayerProperty{1002, -1 << 31, 1<<31 - 1, false}
	propMap[1003] = &PlayerProperty{1003, -1 << 31, 1<<31 - 1, false}
	propMap[1004] = &PlayerProperty{1004, -1 << 31, 1<<31 - 1, false}
	propMap[4001] = &PlayerProperty{4001, 0, 90, false}
	propMap[10001] = &PlayerProperty{10001, -1 << 31, 1<<31 - 1, false}
	propMap[10002] = &PlayerProperty{10002, 0, 8500000, false}
	propMap[10003] = &PlayerProperty{10003, -1 << 31, 1<<31 - 1, true}
	propMap[10004] = &PlayerProperty{10004, 0, 1, false}
	propMap[10005] = &PlayerProperty{10005, 0, 100, false}
	propMap[10006] = &PlayerProperty{10006, 0, 1, false}
	propMap[10007] = &PlayerProperty{10007, 0, 1, false}
	propMap[10008] = &PlayerProperty{10008, 0, 1, false}
	propMap[10009] = &PlayerProperty{10009, 0, 1, false}
	propMap[10010] = &PlayerProperty{10010, 0, 24000, false}
	propMap[10011] = &PlayerProperty{10011, -1 << 31, 1<<31 - 1, true}
	propMap[10012] = &PlayerProperty{10012, -1 << 31, 1<<31 - 1, false}
	propMap[10013] = &PlayerProperty{10013, 1, 60, false}
	propMap[10014] = &PlayerProperty{10014, 0, 1<<31 - 1, false}
	propMap[10015] = &PlayerProperty{10015, -1 << 31, 1<<31 - 1, false}
	propMap[10016] = &PlayerProperty{10016, 0, 1<<31 - 1, false}
	propMap[10017] = &PlayerProperty{10017, 0, 2, false}
	propMap[10018] = &PlayerProperty{10018, 0, 1, false}
	propMap[10019] = &PlayerProperty{10019, 0, 8, false}
	propMap[10020] = &PlayerProperty{10020, 0, 2000, false}
	propMap[10022] = &PlayerProperty{10022, -1 << 31, 1<<31 - 1, false}
	propMap[10023] = &PlayerProperty{10023, -1 << 31, 1<<31 - 1, false}
	propMap[10024] = &PlayerProperty{10024, 0, 1, false}
	propMap[10025] = &PlayerProperty{10025, -1 << 31, 1<<31 - 1, false}
	propMap[10026] = &PlayerProperty{10026, -1 << 31, 1<<31 - 1, false}
	propMap[10027] = &PlayerProperty{10027, 0, 3, false}
	propMap[10028] = &PlayerProperty{10028, -1 << 31, 1<<31 - 1, false}
	propMap[10029] = &PlayerProperty{10029, 0, 300000, false}
	propMap[10035] = &PlayerProperty{10035, -1 << 31, 1<<31 - 1, false}
	propMap[10036] = &PlayerProperty{10036, -1 << 31, 1<<31 - 1, false}
	propMap[10037] = &PlayerProperty{10037, -1 << 31, 1<<31 - 1, false}
	propMap[10038] = &PlayerProperty{10038, -1 << 31, 1<<31 - 1, false}
	propMap[10039] = &PlayerProperty{10039, 0, 8, false}
	propMap[10040] = &PlayerProperty{10040, -1 << 31, 1<<31 - 1, false}
	propMap[10041] = &PlayerProperty{10041, -1 << 31, 1<<31 - 1, false}
	propMap[10042] = &PlayerProperty{10042, 0, 1<<31 - 1, false}
	propMap[10043] = &PlayerProperty{10043, -1 << 31, 1<<31 - 1, false}
	propMap[10044] = &PlayerProperty{10044, -1 << 31, 1<<31 - 1, false}
	propMap[10045] = &PlayerProperty{10045, -1 << 31, 1<<31 - 1, false}
	propMap[10046] = &PlayerProperty{10046, -1 << 31, 1<<31 - 1, false}
	propMap[10047] = &PlayerProperty{10047, -1 << 31, 1<<31 - 1, false}
	propMap[10048] = &PlayerProperty{10048, 0, 1, false}
	propMap[10049] = &PlayerProperty{10049, 0, 10000, false}
	propMap[10050] = &PlayerProperty{10050, 0, 10000, false}
}

func GetPropById(id int) *PlayerProperty {
	if prop, exists := propMap[id]; exists {
		return prop
	}
	return nil
}
