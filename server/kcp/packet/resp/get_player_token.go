package resp

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/session"
	"Go-Grasscutter/utils/crypto"
	"google.golang.org/protobuf/proto"
)

func PacketGetPlayerTokenRsp(sess *session.Session, encryptedSeed, encryptedSeedSign string) *base.Packet {
	code := base.GetPlayerTokenRsp

	p := &pb.GetPlayerTokenRsp{
		Uid:                    uint32(sess.Player.ID),
		Token:                  sess.Account.Token,
		AccountType:            1,
		IsProficientPlayer:     true, // Not sure where this goes
		SecretKeySeed:          sess.EncryptSeed,
		SecurityCmdBuffer:      crypto.EncryptSeedBuffer,
		PlatformType:           3,
		ChannelId:              1,
		CountryCode:            "US",
		ClientVersionRandomKey: "c25-314dd05b0b5f",
		RegPlatform:            3,
		ClientIpStr:            "127.0.0.1", // todo INCOMPLETE: addr, modify below
		ServerRandKey:          encryptedSeed,
		Sign:                   encryptedSeedSign,
	}

	data, err := proto.Marshal(p)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}
	return &base.Packet{
		Opcode:            code,
		ShouldBuildHeader: true,
		Data:              data,
		UseDispatchKey:    true,
	}
}

func PacketGetPlayerTokenNormalRsp(sess *session.Session) *base.Packet {
	code := base.GetPlayerTokenRsp
	p := &pb.GetPlayerTokenRsp{
		Uid:                    uint32(sess.Player.ID),
		Token:                  sess.Account.Token,
		AccountType:            1,
		IsProficientPlayer:     true,
		SecretKeySeed:          sess.EncryptSeed,
		SecurityCmdBuffer:      crypto.EncryptSeedBuffer,
		PlatformType:           3,
		ChannelId:              1,
		CountryCode:            "US",
		ClientVersionRandomKey: "c25-314dd05b0b5f",
		RegPlatform:            3,
		ClientIpStr:            "127.0.0.1",
	}

	data, err := proto.Marshal(p)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	return &base.Packet{
		Opcode:            code,
		ShouldBuildHeader: true,
		Data:              data,
		UseDispatchKey:    true,
	}
}

func PacketGetPlayerTokenBannedRsp(sess *session.Session, retcode int, msg string, blackEndTime int) *base.Packet {
	code := base.GetPlayerTokenRsp
	p := pb.GetPlayerTokenRsp{
		Uid:                uint32(sess.Player.ID),
		IsProficientPlayer: true, // Not sure where this goes
		Retcode:            int32(retcode),
		Msg:                msg,
		BlackUidEndTime:    uint32(blackEndTime),
		RegPlatform:        3,
		CountryCode:        "US",
		ClientIpStr:        "127.0.0.1",
	}
	data, err := proto.Marshal(&p)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}
	return &base.Packet{
		Opcode:            code,
		ShouldBuildHeader: true,
		Data:              data,
		UseDispatchKey:    true,
	}
}
