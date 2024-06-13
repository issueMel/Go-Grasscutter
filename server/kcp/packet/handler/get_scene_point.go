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
	session.Router[base.GetScenePointReq] = HandlerGetScenePointReq
}

func HandlerGetScenePointReq(sess *session.Session, header, payload []byte) {
	req := pb.GetSceneAreaReq{}
	err := proto.Unmarshal(payload, &req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}

	sess.Send(resp.PacketGetScenePointRsp(sess.Player, req.SceneId))
}
