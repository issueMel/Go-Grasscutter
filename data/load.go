package data

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/data/common"
	"Go-Grasscutter/log"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type ScenePointEntry struct {
	SceneID   int
	PointData *common.PointData
}

type ScenePointConfig struct {
	Points map[int]*common.PointData `json:"points"`
}

func loadScenePoints() {
	fp := filepath.Join(config.Conf.FolderStructure.Resources, "BinOutput/Scene/Point")
	pattern := regexp.MustCompile(`scene([0-9]+)_point\.json`)

	files, err := os.ReadDir(fp)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		matcher := pattern.FindStringSubmatch(file.Name())
		if matcher == nil {
			continue
		}

		sceneId, err := strconv.Atoi(matcher[1])
		if err != nil {
			log.SugaredLogger.Error(err)
			continue
		}

		var config ScenePointConfig
		data, err := os.ReadFile(filepath.Join(fp, file.Name()))
		if err != nil {
			log.SugaredLogger.Error(err)
			continue
		}

		err = json.Unmarshal(data, &config)
		if err != nil {
			log.SugaredLogger.Error(err)
			continue
		}
		if err != nil {
			log.SugaredLogger.Error(err)
			continue
		}

		if config.Points == nil {
			continue
		}

		var scenePoints []int
		for pointId, pointData := range config.Points {
			scenePoint := &ScenePointEntry{SceneID: sceneId, PointData: pointData}
			scenePoints = append(scenePoints, pointId)
			pointData.Id = pointId

			GameData.ScenePointIdList = append(GameData.ScenePointIdList, pointId)
			GameData.ScenePointEntryMap[(sceneId<<16)+pointId] = scenePoint

			pointData.UpdateDailyDungeon() // todo UpdateDailyDungeon()
		}
		GameData.ScenePointsPerScene[sceneId] = scenePoints
	}
}
