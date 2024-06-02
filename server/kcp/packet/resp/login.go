package resp

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/utils/crypto"
	"google.golang.org/protobuf/proto"
)

var regionCache *pb.QueryCurrRegionHttpRsp

func PacketPlayerLoginRsp() *base.Packet {
	code := base.PlayerLoginRsp

	if regionCache == nil {
		gameInfo := config.Conf.Server.Game
		serverRegion := pb.RegionInfo{
			GateserverIp:   config.LrString(gameInfo.AccessAddress, gameInfo.BindAddress),
			GateserverPort: uint32(config.LrInt(gameInfo.AccessPort, gameInfo.BindPort)),
		}
		regionCache = &pb.QueryCurrRegionHttpRsp{
			RegionInfo:      &serverRegion,
			ClientSecretKey: crypto.DispatchKey,
		}
	}
	info := regionCache.GetRegionInfo()
	p := pb.PlayerLoginRsp{
		AbilityHashCode:            1844674,
		GameBiz:                    "hk4e_global",
		ClientDataVersion:          info.GetClientDataVersion(),
		ClientSilenceDataVersion:   info.GetClientSilenceDataVersion(),
		ClientMd5:                  info.GetClientDataMd5(),
		ClientSilenceMd5:           info.GetClientSilenceDataMd5(),
		ResVersionConfig:           info.GetResVersionConfig(),
		ClientVersionSuffix:        info.GetClientVersionSuffix(),
		ClientSilenceVersionSuffix: info.GetClientSilenceVersionSuffix(),
		CountryCode:                "US",
	}

	data, err := proto.Marshal(&p)
	if err != nil {
		log.SugaredLogger.Error(err)
	}

	packet := &base.Packet{
		Opcode: code,
		Data:   data,
	}
	packet.BuildHeader(1)
	return packet
}
