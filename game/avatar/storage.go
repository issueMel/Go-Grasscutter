package avatar

import (
	"Go-Grasscutter/data"
	"Go-Grasscutter/db"
	"Go-Grasscutter/log"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync/atomic"
)

const collName = "avatars"

type Storage struct {
	Uid         int         // Get from player
	Loaded      atomic.Bool // Inventory need atomic
	Avatars     map[int]*Avatar
	AvatarsGuid map[int64]*Avatar
}

func (s *Storage) LoadFromDatabase() {
	if s.Loaded.Load() {
		return
	}
	avatars := getAvatars(s.Uid)
	if avatars == nil {
		return
	}

	for _, avatar := range avatars {
		if avatar.Id.IsZero() {
			// todo CHECK: right?
			continue
		}

		avatarData, ok1 := data.GameData.AvatarDataMap[avatar.AvatarId]
		skillDepot, ok2 := data.GameData.AvatarSkillDepotDataMap[avatar.SkillDepotId]
		if !ok1 || !ok2 {
			continue
		}

		// todo INCOMPLETE: Set ownerships
		avatar.AvatarData = avatarData
		avatar.SkillDepot = skillDepot
		// avatar.Owner = player

		// todo INCOMPLETE: Force recalc of const boosted skills

		// Add to avatar storage
		s.Avatars[avatar.AvatarId] = avatar
		s.AvatarsGuid[avatar.Guid] = avatar

		// Set main character skill depot data, fixes loading with no element every login
		if avatar.AvatarId == 10000007 || avatar.AvatarId == 10000005 {
			// avatar.setSkillDepotData()
			// avatar.save();
		}
	}
	s.Loaded.Store(true)
}

func getAvatars(uid int) []*Avatar {
	find, err := db.DB.Collection(collName).Find(context.Background(), bson.D{{"ownerId", uid}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}
		log.SugaredLogger.Error(err)
		return nil
	}
	var avatars []*Avatar
	err = find.All(context.Background(), &avatars)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}
	return avatars
}
