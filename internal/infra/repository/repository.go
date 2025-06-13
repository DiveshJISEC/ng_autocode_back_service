package repository

import (
	"context"
	"ng_autocode_back_service/pkg/logger"

	"go.uber.org/zap"
)

type DataObject struct {
	DBSet databaseSet
	Cache transientCache
}

func CreateNewAppRepo(c context.Context) (DataObject, error) {

	oRepo := DataObject{}

	readDB, err := ReadEbatestConnect()
	if err != nil {
		logger.Log(c).Error("Failed to connect to read database:", zap.Error(err))
		return oRepo, err
	}
	oRepo.DBSet.ReadDB.Db = readDB

	writeDB, err := ReadEbatestConnect()
	if err != nil {
		logger.Log(c).Error("Failed to connect to read database:", zap.Error(err))
		return oRepo, err
	}
	oRepo.DBSet.WriteDB.Db = writeDB

	return oRepo, nil
}
