package internal

import (
	"dbkit/internal/common"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

type GlobalState struct {
	Config      *DBKitConfig
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
		log.Info("全局状态初始化成功")
	})
	return globalState
}

func makeGlobalState() *GlobalState {
	// 从配置文件读取配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件错误")
	}
	config := DBKitConfig{}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("解析配置文件错误")
	}

	state := GlobalState{
		Config:      &config,
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
		log.Fatalf("连接MySQL数据源失败，请确认DBKit基础数据库连接参数无误")
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
		log.Errorf("获取%s连接失败，连接参数未配置或连接失败", dbms)
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
		config := state.Config.MySQL
		if config.DBName == "" {
			return nil, 0
		}
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			config.Username, config.Password, config.Host, config.Port, config.DBName)
		db, err = sqlx.Open("mysql", connStr)
	case common.TIDB:
		config := state.Config.TiDB
		if config.DBName == "" {
			return nil, 0
		}
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			config.Username, config.Password, config.Host, config.Port, config.DBName)
		db, err = sqlx.Open("mysql", connStr)
	case common.MARIADB:
		config := state.Config.MariaDB
		if config.DBName == "" {
			return nil, 0
		}
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			config.Username, config.Password, config.Host, config.Port, config.DBName)
		db, err = sqlx.Open("mysql", connStr)
	case common.SQLITE:
		config := state.Config.SQLite
		if config.DBName == "" {
			return nil, 0
		}
		connStr := fmt.Sprintf("./db/%s.db", config.DBName)
		db, err = sqlx.Open("sqlite3", connStr)
	default:
		log.Fatalf("该类型数据库当前不支持:%s", dbms)
	}

	if err != nil {
		log.Warnf("%s连接失败:%s", dbms, err)
		return nil, -1
	}
	err = db.Ping()
	if err != nil {
		log.Warnf("%s连接失败:%s", dbms, err)
		return nil, -1
	}
	log.Infof("%s连接成功", dbms)
	return db, 1
}
