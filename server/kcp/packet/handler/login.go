package handler

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/packet/resp"
	"Go-Grasscutter/server/kcp/session"
	"google.golang.org/protobuf/proto"
)

func init() {
	session.Router[base.PlayerLoginReq] = HandlerPlayerLoginReq
}

func HandlerPlayerLoginReq(sess *session.Session, header, payload []byte) {
	// Check
	if sess.Account == nil {
		sess.Tunnel.Kcp.Close()
		return
	}

	// Parse request
	req := pb.PlayerLoginReq{}
	err := proto.Unmarshal(payload, &req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}

	// Authenticate session
	if req.Token != sess.Account.Token {
		sess.Tunnel.Kcp.Close()
		return
	}

	// Load character from db
	p := sess.Player

	// Show opening cutscene if player has no avatars
	if len(p.Avatars.Avatars) == 0 {
		// Pick character
		sess.State = session.PickingCharacter
		sess.Send(base.NewPacketWithCode(base.DoSetPlayerBornDataNotify))
	} else {
		// todo INCOMPLETE: onLogin()
		// Login done
		p.OnLogin()
		NotifyLogin(sess)
		NotifyManger(sess)
	}

	// Final packet to tell client logging in is done
	sess.Send(resp.PacketPlayerLoginRsp())
}

func NotifyManger(sess *session.Session) {
	// todo INCOMPLETE: other manager send packet
	sess.Send(resp.PacketOpenStateUpdateNotify(sess.Player))    // this.getProgressManager().onPlayerLogin();
	sess.Send(resp.PacketAchievementAllDataNotify(sess.Player)) // Achievements
	sess.Send(resp.PacketForgeDataNotify())                     // forgingManager
	sess.Send(resp.PacketResinChangeNotify(sess.Player))        // resinManager
	sess.Send(resp.PacketCookDataNotify())                      // cookingManager

	sess.Send(resp.PacketCompoundDataNotify()) // cookingCompoundManager
	// getTodayMoonCard PacketCardProductRewardNotify

	// todo getQuestManager().onLogin()

	sess.Send(resp.PacketBattlePassMissionUpdateNotify(sess.Player)) // triggerMission

	// furnitureManager
	sess.Send(resp.PacketUnlockedFurnitureFormulaDataNotify(sess.Player))
	sess.Send(resp.PacketUnlockedFurnitureSuiteDataNotify(sess.Player))

	// home.onOwnerLogin
	// PacketHomeBasicInfoNotify
	// PacketPlayerHomeCompInfoNotify
	// PacketHomeComfortInfoNotify
	// PacketFurnitureCurModuleArrangeCountNotify
	// PacketHomeMarkPointNotify
	// PacketHomeAvatarTalkFinishInfoNotify
	// PacketHomeAllUnlockedBgmIdListNotify
	// PacketHomeAvatarRewardEventNotify
	// PacketHomeAvatarAllFinishRewardNotify
	// PacketHomeResourceNotify
	sess.Send(resp.PacketActivityScheduleInfoNotify()) // ActivityManager
	temp(sess)
}

func NotifyLogin(sess *session.Session) {
	sess.Send(resp.PacketPlayerDataNotify(sess.Player))
	sess.Send(resp.PacketStoreWeightLimitNotify())
	sess.Send(resp.PacketPlayerStoreNotify(sess.Player))
	sess.Send(resp.PacketAvatarDataNotify(sess.Player))

	sess.Send(resp.PacketFinishedParentQuestNotify(sess.Player))
	sess.Send(resp.PacketBattlePassAllDataNotify(sess.Player))
	sess.Send(resp.PacketQuestListNotify(sess.Player))
	sess.Send(resp.PacketQuestGlobalVarNotify(sess.Player))
	sess.Send(resp.PacketCodexDataFullNotify(sess.Player))
	sess.Send(resp.PacketAllWidgetDataNotify(sess.Player))

	sess.Send(resp.PacketWidgetGadgetAllDataNotify())
	sess.Send(resp.PacketCombineDataNotify(sess.Player.UnlockedCombines))
	sess.Send(resp.PacketGetChatEmojiCollectionRsp(sess.Player.ChatEmojiIdList))

	sess.Send(resp.PacketPlayerEnterSceneNotify(sess.Player))
	sess.Send(resp.PacketPlayerLevelRewardUpdateNotify(sess.Player.RewardedLevels))

	// First notify packets sent
	sess.Player.HasSentLoginPackets = true

	sess.State = session.Active

	// game.Server.RegisterPlayer(sess.Player)
	// todo KickPlayer
}
