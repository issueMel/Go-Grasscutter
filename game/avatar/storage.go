package avatar

import (
	"Go-Grasscutter/db"
	"Go-Grasscutter/log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

const collName = "avatars"

type Storage struct {
	Uid         int // Get from player
	Loaded      bool
	Avatars     map[int]*Avatar
	AvatarsGuid map[int64]*Avatar
}

func (s *Storage) LoadFromDatabase() {
	if s.Loaded {
		return
	}
	avatars := getAvatars(s.Uid)

	for _, avatar := range avatars {
		// todo INCOMPLETE: GameData

		// todo INCOMPLETE: Set ownerships

		// todo INCOMPLETE: Force recalc of const boosted skills

		// Add to avatar storage
		s.Avatars[avatar.AvatarId] = avatar
		s.AvatarsGuid[avatar.Guid] = avatar
	}
	s.Loaded = true
}

func getAvatars(uid int) []*Avatar {
	find, err := db.DB.Collection(collName).Find(context.Background(), bson.D{{"ownerId", uid}})
	if err != nil {
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
