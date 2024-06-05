package consts

import "Go-Grasscutter/utils"

var (
	Version                 = "4.0.0"
	VersionParts            = []int{4, 0, 0}
	DEBUG                   = false
	DefaultTeams            = 4
	MaxTeams                = 10
	MainCharacterMale       = 10000005
	MainCharacterFemale     = 10000007
	StartPosition           = Position{X: 2747, Y: 194, Z: -1719}
	MaxFriends              = 60
	MaxFriendRequests       = 50
	ServerConsoleUID        = 99
	BattlePassMaxLevel      = 50
	BattlePassLevelPrice    = 150
	BattlePassCurrentIndex  = 2
	BattlePassPointPerLevel = 1000
	BattlePassPointPerWeek  = 10000
	DefaultAbilityStrings   = []string{
		"Avatar_DefaultAbility_VisionReplaceDieInvincible",
		"Avatar_DefaultAbility_AvartarInShaderChange",
		"Avatar_SprintBS_Invincible",
		"Avatar_Freeze_Duration_Reducer",
		"Avatar_Attack_ReviveEnergy",
		"Avatar_Component_Initializer",
		"Avatar_FallAnthem_Achievement_Listener",
		"GrapplingHookSkill_Ability",
		"Avatar_PlayerBoy_DiveStamina_Reduction",
		"Ability_Avatar_Dive_SealEcho",
		"Absorb_SealEcho_Bullet_01",
		"Absorb_SealEcho_Bullet_02",
		"Ability_Avatar_Dive_CrabShield",
		"ActivityAbility_Absorb_Shoot",
		"SceneAbility_DiveVolume",
	}
	DefaultTeamAbilityStrings = []string{
		"Ability_Avatar_Dive_Team",
	}
	// todo INCOMPLETE: SparseSet
	IllegalWeapons = []string{
		"10000-10008", "11411", "11506-11508", "12505", "12506", "12508", "12509",
		"13503", "13506", "14411", "14503", "14505", "14508", "15504-15506",
	}
	// todo INCOMPLETE: SparseSet
	IllegalRelics = []string{
		"20001", "23300-23340", "23383-23385", "78310-78554", "99310-99554",
	}
	// todo INCOMPLETE: SparseSet
	IllegalItems = []string{
		"100086", "100087", "100100-101000", "101106-101110", "101306", "101500-104000",
		"105001", "105004", "106000-107000", "107011", "108000", "109000-110000",
		"115000-130000", "200200-200899", "220050", "220054",
	}
	DefaultAbilityHashes = []int{} // todo INCOMPLETE: Arrays.stream(DEFAULT_ABILITY_STRINGS).mapToInt(Utils::abilityHash).toArray()
	DefaultAbilityName   = utils.AbilityHash("Default")
)

type Position struct {
	X, Y, Z int
}
