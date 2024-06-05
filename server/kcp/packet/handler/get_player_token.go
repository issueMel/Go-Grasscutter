package handler

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/game"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/packet/resp"
	"Go-Grasscutter/server/kcp/session"
	"Go-Grasscutter/utils"
	"Go-Grasscutter/utils/crypto"
	crypto2 "crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"google.golang.org/protobuf/proto"
)

func init() {
	session.Router[base.GetPlayerTokenReq] = HandlerGetPlayerTokenReq
}

func HandlerGetPlayerTokenReq(sess *session.Session, header, payload []byte) {
	req := &pb.GetPlayerTokenReq{}
	err := proto.Unmarshal(payload, req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}

	// Fetch the account from the ID and token.
	account := DispatchAuthenticate(req.AccountUid, req.AccountToken)

	// Check the account.
	if account == nil {
		sess.Tunnel.Close()
		return
	}

	// Set account
	sess.Account = account

	// Check if player object exists in server
	// NOTE: CHECKING MUST SITUATE HERE (BEFORE getPlayerByUid)! because to save firstly ,to load
	// secondly !!!
	kicked := false
	val, exist := game.Server.Players.Load(account.ID)
	if exist {
		p := val.(player.Player)
		if p.Session != sess.Tunnel { // No self-kicking
			// must save immediately , or the below will load old data
			p.Save()
			p.Session.Close()
			log.SugaredLogger.Infof("Player %s was kicked due to duplicated login\n", p.Account.Username)
			kicked = true
		}
	}

	// NOTE: If there are 5 online players, max count of player is 5,
	// a new client want to login by kicking one of them ,
	// I think it should be allowed
	if !kicked {
		// Max players limit
		maxPlayer := config.Conf.Account.MaxPlayer
		if maxPlayer > -1 && game.Server.PlayerNum.Load() > int32(maxPlayer) {
			sess.Tunnel.Close()
			log.SugaredLogger.Info("max players limit")
			return
		}
	}

	// Checks if the player is banned
	if account.IsBanned {
		sess.State = session.AccountBanned
		rsp := resp.PacketGetPlayerTokenBannedRsp(sess, 21, "FORBID_CHEATING_PLUGINS", sess.Account.BanEndTime)
		sess.Send(rsp)
		return
	}

	// Get player.
	p := player.GetPlayerByAccount(account)
	if p == nil {
		nextPlayerUid := player.GetNextPlayerId(account.ReservedPlayerId)
		p = player.CreatePlayer(account, sess.Tunnel)
		// Save to db
		p.GeneratePlayerUid(nextPlayerUid)
	}
	// Set player object for session
	sess.Player = p

	// todo INCOMPLETE: Load player from database
	sess.Player.LoadFromDatabase()

	// Set session state
	sess.UseSecretKey = true
	sess.State = session.WaitingForLogin

	if req.KeyId > 0 {
		clientSeedEncrypted, err := utils.Base64Decode(req.ClientRandKey)
		if err != nil {
			log.SugaredLogger.Error(err)
			return
		}
		clientSeedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, crypto.CurSigningKey, clientSeedEncrypted)
		if err != nil {
			log.SugaredLogger.Error(err)
			return
		}
		clientSeed := binary.BigEndian.Uint64(clientSeedBytes)

		encryptSeed := sess.EncryptSeed
		seedBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(seedBytes, encryptSeed^clientSeed)

		publicKey := crypto.EncryptionKeys[int(req.KeyId)]
		seedEncrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, seedBytes)
		if err != nil {
			log.SugaredLogger.Error(err)
			return
		}

		hashed := sha256.Sum256(seedBytes)
		signature, err := rsa.SignPKCS1v15(rand.Reader, crypto.CurSigningKey, crypto2.SHA256, hashed[:])
		if err != nil {
			log.SugaredLogger.Error(err)
			return
		}
		rsp := resp.PacketGetPlayerTokenRsp(sess, utils.Base64Encode(seedEncrypted), utils.Base64Encode(signature))
		sess.Send(rsp)
	} else {
		// todo INCOMPLETE: Send packet
	}
}

func DispatchAuthenticate(accountId, token string) *player.Account {
	account := player.GetAccountById(accountId)
	if account == nil {
		return nil
	}
	if account.Token != token {
		return nil
	}
	return account
}
