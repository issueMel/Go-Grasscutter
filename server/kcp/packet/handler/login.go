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
		sess.Tunnel.Close()
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
		sess.Tunnel.Close()
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
		// todo onLogin()
		p.OnLogin()
		NotifyLogin(sess)
	}

	// Final packet to tell client logging in is done
	sess.Send(resp.PacketPlayerLoginRsp())
}

func NotifyLogin(sess *session.Session) {
	sess.Send(resp.PacketPlayerDataNotify(sess.Player))
	sess.Send(resp.PacketStoreWeightLimitNotify())
	sess.Send(resp.PacketPlayerStoreNotify(sess.Player))
	sess.Send(resp.PacketAvatarDataNotify(sess.Player))

	//sess.Send() // this.getProgressManager().onPlayerLogin();

	sess.Send(resp.PacketFinishedParentQuestNotify(sess.Player))
	sess.Send(resp.PacketBattlePassAllDataNotify(sess.Player))
	sess.Send(resp.PacketQuestListNotify(sess.Player))
	sess.Send(resp.PacketQuestGlobalVarNotify(sess.Player))
	sess.Send(resp.PacketCodexDataFullNotify(sess.Player))
	sess.Send(resp.PacketAllWidgetDataNotify(sess.Player))
	//
	////Achievements
	//sess.Send()
	//sess.Send()
	//sess.Send()
	//sess.Send()
	//
	//sess.Send()
	//sess.Send()

	sess.State = session.Active
}
