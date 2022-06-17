package internal

import (
	"dbkit/config"
	"dbkit/internal/common"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type GlobalState struct {
	Config      *config.DBKitConfig
	Tests       []*TestContext
	ConnStates  map[common.DBMS]int // -1连接失败 0未配置 1成功
	Connections map[common.DBMS]*sqlx.DB
	DataSource  *sqlx.DB
	TableCount  int
}

var globalState *GlobalState

var once sync.Once

func GetState() *GlobalState {
	once.Do(func() {
		globalState = makeGlobalState()
		log.Info("Initialize global state")
	})
	return globalState
}

func makeGlobalState() *GlobalState {
	// 从配置文件读取配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fail to read config")
	}
	dbKitConfig := config.DBKitConfig{}
	err = viper.Unmarshal(&dbKitConfig)
	if err != nil {
		log.Fatalf("Fail to parse config")
	}

	state := GlobalState{
		Config:      &dbKitConfig,
		Tests:       make([]*TestContext, 0),
		ConnStates:  make(map[common.DBMS]int),
		Connections: make(map[common.DBMS]*sqlx.DB),
		DataSource:  nil,
	}

	// 根据配置文件建立连接
	dataSource := state.Config.DataSource
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/dbkit",
		dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port)
	state.DataSource, err = sqlx.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Fail to connect database")
	}
	for _, dbms := range common.DBMSSet {
		db, res := state.initConnPool(dbms)
		state.ConnStates[dbms] = res
		state.Connections[dbms] = db
	}

	return &state
}

func (state *GlobalState) GetDataSourceConn() *sqlx.DB {
	return state.DataSource
}

func (state *GlobalState) GetConnPool(dbms common.DBMS) *sqlx.DB {
	if state.ConnStates[dbms] != 0 {
		log.Errorf("Fail to connect %s", dbms)
	}
	return state.Connections[dbms]
}

// -1连接失败 0未配置 1成功
func (state *GlobalState) initConnPool(dbms common.DBMS) (*sqlx.DB, int) {
	var (
		db  *sqlx.DB
		err error
	)

	switch dbms {
	case common.MYSQL:
		mySQLConfig := state.Config.MySQL
		if mySQLConfig.DBName == "" {
			return nil, 0
		}
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			mySQLConfig.Username, mySQLConfig.Password, mySQLConfig.Host, mySQLConfig.Port, mySQLConfig.DBName)
		db, err = sqlx.Open("mysql", connStr)
	case common.TIDB:
		tiDBConfig := state.Config.TiDB
		if tiDBConfig.DBName == "" {
			return nil, 0
		}
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			tiDBConfig.Username, tiDBConfig.Password, tiDBConfig.Host, tiDBConfig.Port, tiDBConfig.DBName)
		db, err = sqlx.Open("mysql", connStr)
	case common.MARIADB:
		mariaDBConfig := state.Config.MariaDB
		if mariaDBConfig.DBName == "" {
			return nil, 0
		}
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			mariaDBConfig.Username, mariaDBConfig.Password, mariaDBConfig.Host, mariaDBConfig.Port, mariaDBConfig.DBName)
		db, err = sqlx.Open("mysql", connStr)
	case common.SQLITE:
		sqLiteConfig := state.Config.SQLite
		if sqLiteConfig.DBName == "" {
			return nil, 0
		}
		connStr := fmt.Sprintf("./db/%s.db", sqLiteConfig.DBName)
		db, err = sqlx.Open("sqlite3", connStr)
	default:
		log.Fatalf("Do not support %s", dbms)
	}

	if err != nil {
		log.Warnf("Fail to connect %s [%s]", dbms, err)
		return nil, -1
	}
	err = db.Ping()
	if err != nil {
		log.Warnf("Fail to connect %s [%s]", dbms, err)
		return nil, -1
	}
	log.Infof("Connect %s", dbms)
	return db, 1
}
