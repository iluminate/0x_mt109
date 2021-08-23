package container

import (
	"0x_mt109/application/configs"
	"0x_mt109/application/controllers"
	"0x_mt109/application/repositories"
	"0x_mt109/application/services"
	"0x_mt109/helpers/database"
	"0x_mt109/helpers/loader"
	"sync"
)

var syncDbHelper sync.Once
var mysqlDbHelper *database.MysqldbHelper
func dbConfig() *database.MysqlConfig {
	dbConf := &configs.MysqlConfig{}
	loader.ReadConf(dbConf)
	return &database.MysqlConfig{
		Host:            dbConf.Mysql.Host,
		Port:            dbConf.Mysql.Port,
		User:            dbConf.Mysql.User,
		Password:        dbConf.Mysql.Password,
		Database:        dbConf.Mysql.Database,
		ConnMaxLifetime: dbConf.Mysql.ConnMaxLifetime,
		MaxIdleConns:    dbConf.Mysql.MaxIdleConns,
		MaxOpenConns:    dbConf.Mysql.MaxOpenConns,
	}
}
func MysqldbHelper() *database.MysqldbHelper {
	syncDbHelper.Do(func() {
		mysqldbHelper := database.NewMysqldbHelper(dbConfig())
		go mysqldbHelper.OpenConnection()
		mysqlDbHelper = mysqldbHelper
	})
	return mysqlDbHelper
}
func ActorRepository() repositories.IActorRepository {
	return repositories.NewActorRepository(MysqldbHelper())
}
func ActorService() services.IActorService {
	return services.NewActorService(ActorRepository())
}
func ActorHandler() *controllers.ActorHandler {
	return controllers.NewActorHandler(ActorService())
}

