package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var config Conf

type Conf struct {
	MysqlConfigInfo     MysqlConfig `yaml:"mysql_config_info"`
	RedisConfigInfo     RedisConfig `yaml:"redis_config_info"`
	JwtSecretSand       string      `yaml:"jwt_secret_sand"`
	ParseIpUrl          string      `yaml:"parse_ip_url"`
	ParseIpUrl2         string      `yaml:"parse_ip_url2"`
	BrowserDetailUrl    string      `yaml:"browser_detail_url"`
	IpDetailUrl         string      `yaml:"ip_detail_url"`
	QQImageUrl          string      `yaml:"qq_image_url"`
	GithubToken         string      `yaml:"github_token"`
	GithubUploadAPi     string      `yaml:"github_upload_api"`
	CDNUrlGithub        string      `yaml:"cdn_url_github"`
	GithubUsername      string      `yaml:"github_username"`
	GithubRepos         string      `yaml:"github_repos"`
	GithubReposPath     string      `yaml:"github_repos_path"`
	InternalServiceHost string      `yaml:"internal_service_host"`
	DefaultAvatarUrl    string      `json:"default_avatar_url"`
	DefaultIp           string      `yaml:"default_ip"`
	RegionMsgUrl        string      `yaml:"region_msg_url"`
	LogConfig           LogConf     `yaml:"log_config"`
}

type LogConf struct {
	LogPath  string `yaml:"log_path"`
	LogLevel string `yaml:"log_level"`
}

type RedisConfig struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         uint16 `yaml:"port"`
	Db           uint8  `yaml:"db"`
	PoolSize     uint16 `yaml:"pool_size"`
	MaxRetries   uint8  `yaml:"max_retries"`
	MinIdleConns uint8  `yaml:"min_idle_conns"`
}

type MysqlConfig struct {
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        uint16 `yaml:"port"`
	DbName      string `yaml:"db_name"`
	MaxIdleConn uint16 `yaml:"max_idle_conn"`
	MaxOpenConn uint16 `yaml:"max_open_conn"`
}

func GetRedisConfig() RedisConfig {
	return config.RedisConfigInfo
}

func GetMysqlConfig() MysqlConfig {
	return config.MysqlConfigInfo
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	filename := basePath + string(os.PathSeparator) + "config.yaml"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
}

func GetJwtSecretSand() string {
	return config.JwtSecretSand
}

func GetParseIpUrl() string {
	return config.ParseIpUrl
}

func GetParseIpUrl2() string {
	return config.ParseIpUrl2
}

func GetBrowserDetailUrl() string {
	return config.BrowserDetailUrl
}

func GetIpDetailUrl() string {
	return config.IpDetailUrl
}

func GetQQImageURL() string {
	return config.QQImageUrl
}

func GetGithubToken() string {
	return config.GithubToken
}

func GetGithubUploadAPI() string {
	return config.GithubUploadAPi
}

func GetCDNUrlGithub() string {
	return config.CDNUrlGithub
}

func GetGithubUsername() string {
	return config.GithubUsername
}

func GetGithubRepos() string {
	return config.GithubRepos
}

func GetGithubReposPath() string {
	return config.GithubReposPath
}

func GetInternalServiceHost() string {
	return config.InternalServiceHost
}

func GetDefaultAvatarUrl() string {
	return config.DefaultAvatarUrl
}

func GetDefaultIp() string {
	return config.DefaultIp
}

func GetRegionMsgUrl() string {
	return config.RegionMsgUrl
}

func LogConfig() LogConf {
	return config.LogConfig
}
