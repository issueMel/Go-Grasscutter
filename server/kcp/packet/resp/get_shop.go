package resp

import (
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketGetShopRsp(inv *player.Player, shopType uint32) *base.Packet {
	code := base.GetShopRsp
	shop := pb.Shop{
		ShopType:            shopType,
		CityId:              1,
		CityReputationLevel: 10,
	}

	// todo ShopSystem

	msg := pb.GetShopRsp{
		Shop: &shop,
	}

	d, err := proto.Marshal(&msg)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	return &base.Packet{
		Opcode: code,
		Data:   d,
	}
}
