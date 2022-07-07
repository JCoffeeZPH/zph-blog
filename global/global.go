package global

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reflect"
	"strconv"
	"strings"
	"time"
	"zph/config"
	"zph/logger"
)

var (
	MysqlClient *gorm.DB
	RedisClient *redis.Client
	log = logger.Log
)

func init() {
	mysqlErr := initMysql()
	if mysqlErr != nil {
		log.Error(mysqlErr.Error())
	}
	redisErr := initRedis()
	if redisErr != nil {
		log.Error(redisErr.Error())
	}
}

func initRedis() (err error) {
	redisConf := config.GetRedisConfig()
	addr := redisConf.Host + ":" + strconv.Itoa(int(redisConf.Port))
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     redisConf.Password,
		DB:           int(redisConf.Db),
		PoolSize:     int(redisConf.PoolSize),
		MinIdleConns: int(redisConf.MinIdleConns),
		MaxRetries:   int(redisConf.MaxRetries),
	})

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stats := RedisClient.PoolStats()
	recordRedisPoolStats(stats)
	recordRedisOptions(*(RedisClient.Options()))
	_, err = RedisClient.Ping().Result()
	return err
}

func recordRedisPoolStats(stats *redis.PoolStats) {
	if stats != nil {
		log.Infof("Hits: %d, Misses: %d, Timeout: %d, TotalConns: %d, IdleConns: %d, StaleConns: %d,",
			stats.Hits, stats.Misses, stats.Timeouts, stats.TotalConns, stats.IdleConns, stats.StaleConns)
	}
}

func recordRedisOptions(option redis.Options) {
	optionMap := make(map[string]interface{}, 0)
	t := reflect.TypeOf(option)
	if t.Kind() != reflect.Struct {
		log.Error("check type err not struct")
		return
	}
	for i := 0; i < t.NumField(); i++ {
		v := reflect.ValueOf(option)
		value := v.FieldByName(t.Field(i).Name)
		if value.Kind() == reflect.Func && !value.IsNil() {
			optionMap[t.Field(i).Name] = "default func"
		} else if value.Kind() == reflect.Int || value.Kind() == reflect.Int64 {
			optionMap[t.Field(i).Name] = value.Int()
		} else if value.Kind() == reflect.Bool {
			optionMap[t.Field(i).Name] = value.Bool()
		} else if value.Kind() == reflect.Func {
			optionMap[t.Field(i).Name] = "default func"
		} else if value.Kind() == reflect.Ptr {
			optionMap[t.Field(i).Name] = "default"
		} else {
			optionMap[t.Field(i).Name] = v.FieldByName(t.Field(i).Name).String()
		}
	}
	log.Info("redis_utils option info")
}

func initMysql() error {
	datasourceName, poolConfig := getMysqlDatasource()
	DB, err := gorm.Open("mysql", datasourceName)
	if err != nil {
		log.Errorf("initMysql failed, err is: %+v", err)
		panic("初始化mysql错误")
	}
	DB.SingularTable(true)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return defaultTableName + "_tab"
	//}
	//print sql
	//DB.LogMode(true)
	sqlDB := DB.DB()
	sqlDB.SetMaxOpenConns(int(poolConfig[1]))
	sqlDB.SetMaxIdleConns(int(poolConfig[0]))
	//获得当前的SQL配置情况
	data, _ := json.Marshal(sqlDB.Stats())
	log.Info(string(data))
	DB.LogMode(true)
	DB.SetLogger(log)
	MysqlClient = DB
	return err
}


func getMysqlDatasource() (string, []uint16) {
	datasource := config.GetMysqlConfig()
	dataSourceName := strings.Join([]string{datasource.Username, ":", datasource.Password, "@tcp(", datasource.Host, ":", strconv.Itoa(int(datasource.Port)), ")/", datasource.DbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	poolConfig := []uint16{datasource.MaxIdleConn, datasource.MaxOpenConn}
	return dataSourceName, poolConfig
}