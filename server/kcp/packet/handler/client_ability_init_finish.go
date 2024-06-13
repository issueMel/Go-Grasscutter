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
	session.Router[base.ClientAbilityInitFinishNotify] = HandlerClientAbilityInitFinishNotify
}

func HandlerClientAbilityInitFinishNotify(sess *session.Session, header, payload []byte) {
	notif := pb.ClientAbilityInitFinishNotify{}
	err := proto.Unmarshal(payload, &notif)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
	// todo AbilityManager

	if len(notif.Invokes) > 0 {
		sess.Player.ClientAbility = append(sess.Player.ClientAbility, notif.Invokes...)
	}

	sess.Send(resp.PacketClientAbilityInitFinishNotify(sess.Player))
}
