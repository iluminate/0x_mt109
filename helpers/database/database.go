package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type MysqlConfig struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	Database        string `yaml:"database"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
	MaxIdleConns    int `yaml:"maxIdleConns"`
	MaxOpenConns    int `yaml:"maxOpenConns"`
}

type MysqldbHelper struct {
	conn   *sql.Conn
	conf   *MysqlConfig
	db     *sql.DB
}

func NewMysqldbHelper(conf *MysqlConfig) *MysqldbHelper {
	return &MysqldbHelper{conf: conf}
}

func (helper *MysqldbHelper) DatabaseName() string {
	return helper.conf.Database
}

func (helper *MysqldbHelper) GetConnection() *sql.DB {
	return helper.db
}

func (helper *MysqldbHelper) OpenConnection() error {
	if helper.db != nil {
		return nil
	}
	cfg := mysql.Config{
		User:   helper.conf.User,
		Passwd: helper.conf.Password,
		Net:    "tcp",
		Addr:   helper.conf.Host + ":" + helper.conf.Port,
		DBName: helper.conf.Database,
		Collation: "utf8mb4_unicode_ci",
		ParseTime: true,
		Loc: time.Local,
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		log.Printf("%v",err)
		return err
	}
	helper.db = sql.OpenDB(connector)
	if err := helper.db.Ping(); err != nil {
		defer helper.closeConnection()
		log.Printf("%v",err)
		return err
	}
	helper.db.SetConnMaxLifetime(helper.conf.ConnMaxLifetime)
	helper.db.SetMaxIdleConns(helper.conf.MaxIdleConns)
	helper.db.SetMaxOpenConns(helper.conf.MaxOpenConns)

	log.Println("Connected to MysqlDB!")
	return nil
}

func (helper *MysqldbHelper) closeConnection() error {
	err := helper.db.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nConnection to MysqlDB closed.")
	return err
}
