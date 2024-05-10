package handler

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/src/config"
	"Go-Grasscutter/src/consts"
	"Go-Grasscutter/src/server/event/dispatch"
	"Go-Grasscutter/src/server/http/object"
	"Go-Grasscutter/src/utils"
	"Go-Grasscutter/src/utils/crypto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/golang/protobuf/proto"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	regions              sync.Map // map[string]RegionData
	regionListResponse   string
	regionListResponseCN string
	doOnce               sync.Once
)

// Initialize Configures region data according to configuration.
func Initialize() {
	var head string
	c := config.Conf
	httpInfo := c.Server.Http
	if httpInfo.Encryption.UseInRouting {
		head = "https://"
	} else {
		head = "http://"
	}
	// todo use https
	head = "http://"

	dispatchDomain :=
		head +
			config.LrString(httpInfo.AccessAddress, httpInfo.BindAddress) +
			":" +
			strconv.Itoa(config.LrInt(httpInfo.AccessPort, httpInfo.BindPort))
	var (
		servers           = make([]*pb.RegionSimpleInfo, 0)
		usedNames         = make(map[string]struct{})
		configuredRegions = make([]*config.Region, 0) // todo raed config data
	)

	// todo getRunMode() != ServerRunMode.HYBRID && set const "HYBRID"
	if c.Server.RunMode != "HYBRID" && len(configuredRegions) == 0 {
		log.Fatal("[Dispatch] There are no game servers available. Exiting due to unplayable state.")
	} else if len(configuredRegions) == 0 {
		gameInfo := c.Server.Game
		configuredRegions = append(configuredRegions, &config.Region{
			Name:  "os_usa",
			Title: c.Server.Dispatch.DefaultName,
			Ip:    config.LrString(gameInfo.AccessAddress, gameInfo.BindAddress),
			Port:  config.LrInt(gameInfo.AccessPort, gameInfo.BindPort),
		})
	}

	for _, region := range configuredRegions {
		if _, ok := usedNames[region.Name]; ok {
			log.Println("Region name already in use.")
			continue
		}
		// Create a region identifier.
		identifier := &pb.RegionSimpleInfo{
			Name:        region.Name,
			Title:       region.Title,
			Type:        "DEV_PUBLIC",
			DispatchUrl: dispatchDomain + "/query_cur_region/" + region.Name,
		}
		usedNames[region.Name] = struct{}{}
		servers = append(servers, identifier)

		// Create a region info object.
		// todo load config regions
		updatedQuery := &pb.QueryCurrRegionHttpRsp{
			RegionInfo: &pb.RegionInfo{
				GateserverIp:   region.Ip,
				GateserverPort: uint32(region.Port),
			},
			ClientSecretKey: crypto.DispatchSeed,
		}
		data, err := proto.Marshal(updatedQuery)
		if err != nil {
			panic(err)
		}
		regions.Store(region.Name, &RegionData{
			RegionQuery: updatedQuery,
			Base64:      utils.Base64Encode(data),
		})
	}

	// Determine config settings.
	// Create a config object.
	customConfig := struct {
		Sdkenv        string `json:"sdkenv"`
		CheckDevice   string `json:"checkdevice"`
		LoadPatch     string `json:"loadPatch"`
		ShowException string `json:"showexception"`
		RegionConfig  string `json:"regionConfig"`
		DownloadMode  string `json:"downloadMode"`
		CodeSwitch    []int  `json:"codeSwitch"`
		CoverSwitch   []int  `json:"coverSwitch"`
	}{
		Sdkenv:        "2",
		CheckDevice:   "false",
		LoadPatch:     "false",
		ShowException: "false", // todo GameConstants.DEBUG
		RegionConfig:  "pm|fk|add",
		DownloadMode:  "0",
		CodeSwitch:    []int{3628}, // codeSwitch
		CoverSwitch:   []int{40},   // hiddenIcons
	}

	// XOR the config with the key.

	encodedConfig, err := json.Marshal(customConfig)

	encodedConfig = crypto.Xor(encodedConfig, crypto.DispatchKey)

	// Create an updated region list.
	updatedRegionList := pb.QueryRegionListHttpRsp{
		RegionList:                  servers,
		ClientSecretKey:             crypto.DispatchSeed,
		ClientCustomConfigEncrypted: encodedConfig,
		EnableLoginPc:               true,
	}
	// Set the region list response.
	regionListResponse = utils.Base64Encode([]byte(updatedRegionList.String()))

	// CN
	// Modify the existing config option.
	customConfig.Sdkenv = "0"
	// XOR the config with the key.
	encodedConfig, err = json.Marshal(customConfig)
	if err != nil {
		log.Println(err)
	}

	encodedConfig = crypto.Xor(encodedConfig, crypto.DispatchKey)
	// Create an updated region list.
	updatedRegionListCN := pb.QueryRegionListHttpRsp{
		RegionList:                  servers,
		ClientSecretKey:             crypto.DispatchSeed,
		ClientCustomConfigEncrypted: encodedConfig,
		EnableLoginPc:               true,
	}

	// todo fix it
	// Set the region list response.
	marshal, err := proto.Marshal(&updatedRegionListCN)
	if err != nil {
		log.Println("updatedRegionListCN Marshal error!")
		return
	}
	regionListResponseCN = utils.Base64Encode(marshal)
	// todo remove temp way
	regionListResponseCN = "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"

}

type RegionData struct {
	RegionQuery *pb.QueryCurrRegionHttpRsp
	Base64      string
}

func ApplyRegionHandler(r *server.Hertz) {
	r.GET("/query_region_list", queryRegionList)
	// login/v2 -> /query_cur_region
	r.GET("/query_cur_region", queryCurrentRegion)
	r.GET("/query_cur_region/:region", queryCurrentRegion)
}

// todo fix it
func queryRegionList(c context.Context, ctx *app.RequestContext) {
	doOnce.Do(func() {
		Initialize()
	})
	if ctx.QueryArgs().Has("version") && ctx.QueryArgs().Has("platform") {
		versionName := ctx.Query("version")
		versionCode := versionName[0:8]
		//platformName := ctx.Query("platform")
		switch versionCode {
		case "CNRELiOS", "CNRELWin", "CNRELAnd":
			// Use the CN region list.
			// todo QueryAllRegionsEvent

			// Respond with the event result.
			ctx.String(200, regionListResponseCN)
		case "OSRELiOS", "OSRELWin", "OSRELAnd":
			// Use the OS region list.

			// Respond with the event result.
			ctx.String(200, regionListResponse)
		default:
			ctx.String(200, regionListResponse)
		}
	} else {
		// Use the default region list.

		// Respond with the event result.
		ctx.String(200, regionListResponse)
	}
	// Log the request to the console.

}

// GET /query_cur_region/
// os_usa?
// version=CNRELWin4.0.0
// &lang=2
// &platform=3
// &binary=1
// &time=778
// &channel_id=1
// &sub_channel_id=1
// &account_type=1
// &dispatchSeed=f586f0569733d33a
// &key_id=4
// &aid=10001
// todo check make game can connect to game server
func queryCurrentRegion(c context.Context, ctx *app.RequestContext) {
	doOnce.Do(func() {
		Initialize()
	})

	// Get region to query.
	regionName := ctx.Param("region")
	versionName := ctx.Query("version")

	// Get region data.
	regionData := "CAESGE5vdCBGb3VuZCB2ZXJzaW9uIGNvbmZpZw=="
	val, ok := regions.Load(regionName)
	if ctx.QueryArgs().Len() > 0 {
		if ok {
			regionData = val.(*RegionData).Base64
		}
	}

	re := regexp.MustCompile("[a-zA-Z]")
	clientVersion := re.ReplaceAllString(versionName, "")
	versionCode := strings.Split(clientVersion, ".")

	versionMajor, _ := strconv.Atoi(versionCode[0])
	versionMinor, _ := strconv.Atoi(versionCode[1])
	versionFix, _ := strconv.Atoi(versionCode[2])
	if versionMajor >= 3 ||
		(versionMajor == 2 && versionMinor == 7 && versionFix >= 50) ||
		(versionMajor == 2 && versionMinor == 8) {
		event := dispatch.NewQueryCurrentRegionEvent(regionData)
		// event.call();
		keyId := ctx.Query("key_id")
		// The 'fix' or 'patch' version is not checked because it is only used
		// when miHoYo is desperate and fucks up big time.
		if versionMajor != consts.VersionParts[0] ||
			versionMinor != consts.VersionParts[1] {
			// Reject clients when there is a version mismatch
			updateClient := strings.Compare(consts.Version, clientVersion) > 0
			var contentMsg string
			if updateClient {
				contentMsg = fmt.Sprintf("\nVersion mismatch outdated client! \n\nServer version: %s\nClient version: %s",
					consts.Version, clientVersion)
			} else {
				contentMsg = fmt.Sprintf("\nVersion mismatch outdated server! \n\nServer version: %s\nClient version: %s",
					consts.Version, clientVersion)
			}
			rsp := pb.QueryCurrRegionHttpRsp{
				Retcode:    int32(pb.Retcode_RET_STOP_SERVER),
				Msg:        "Connection Failed!",
				RegionInfo: &pb.RegionInfo{},
				Detail: &pb.QueryCurrRegionHttpRsp_StopServer{
					StopServer: &pb.StopServerInfo{
						StopBeginTime: uint32(time.Now().Unix()),
						StopEndTime:   uint32(time.Now().Unix() + 1),
						Url:           "https://discord.gg/T5vZU6UyeG",
						ContentMsg:    contentMsg,
					},
				},
			}
			log.Printf("Connection denied for %s due to %t.\n", utils.Address(ctx), updateClient)
			ctx.JSON(200, &rsp)
			return
		}
		if len(ctx.Query("dispatchSeed")) == 0 {
			rsp := object.QueryCurRegionRspJson{
				Content: event.RegionInfo,
				Sign:    "TW9yZSBsb3ZlIGZvciBVQSBQYXRjaCBwbGF5ZXJz",
			}
			ctx.JSON(200, rsp)
			return
		}
		regionInfo := utils.Base64Decode(event.RegionInfo)
		rsp, err := crypto.EncryptAndSignRegionData(regionInfo, keyId)
		if err != nil {
			log.Println("An error occurred while handling query_cur_region.", rsp)
			return
		}
		ctx.JSON(200, rsp)
	} else {
		// todo Invoke event.
		event := dispatch.NewQueryCurrentRegionEvent(regionData)
		// Respond with event result.
		ctx.String(200, event.RegionInfo)
	}
}
