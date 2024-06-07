package handler

import (
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/packet/resp"
	"Go-Grasscutter/server/kcp/session"
)

func init() {
	session.Router[base.GetChatEmojiCollectionReq] = HandlerGetChatEmojiCollectionReq
}

func HandlerGetChatEmojiCollectionReq(sess *session.Session, header, payload []byte) {
	sess.Send(resp.PacketGetChatEmojiCollectionRsp(sess.Player.ChatEmojiIdList))
}
