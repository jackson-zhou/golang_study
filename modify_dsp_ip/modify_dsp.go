package main

//go get github.com/garyburd/redigo
//go get github.com/dlintw/goconf
//go get github.com/cihub/seelog

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/dlintw/goconf"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"strings"
)

/*
函数返回：
redisHost, redisPort, dspList, 目标dsp的ip,seelog配置文件名,error
*/
func getCfgInfo(cfgFileName string) (string, string, []string, string, string, error) {
	//goconf的用法参见$GOPATH/src/github.com/dlintw/goconf/README.rst
	var redisHost string
	var redisPort string
	var dspList []string
	var dspListStr string
	var dstDspIp string
	var logCfg string

	c, err := goconf.ReadConfigFile(cfgFileName)
	if err != nil {
		return redisHost, redisPort, dspList, logCfg, dstDspIp, err
	}
	redisHost, _ = c.GetString("redis", "host")
	redisPort, _ = c.GetString("redis", "port")
	dspListStr, _ = c.GetString("DspWillBeModify", "DspList")
	dstDspIp, _ = c.GetString("DspWillBeModify", "DestHost")
	logCfg, _ = c.GetString("seelog", "logCfg")
	dspList = strings.Split(dspListStr, ",")
	return redisHost, redisPort, dspList, logCfg, dstDspIp, nil
}

// 修改redis里的dsp的bid_url
func modifyDsp(redisClient redis.Conn, dspId, dstDspIp string) {
	log.Debug("Modify dsp:", dspId)
	dsp_txt, err := redis.String(redisClient.Do("get", "dsp."+dspId))
	if err != nil {
		log.Debug("Can not modify:", dspId)
		return
	}

	var dsp_json interface{}
	if err := json.Unmarshal([]byte(dsp_txt), &dsp_json); err != nil {
		log.Debug("cannot parse json for dsp:", dsp_json, " err:", err, "dsp_txt:", dsp_txt)
		return
	}
	dsp_json_obj := dsp_json.(map[string]interface{})
	bid_url := dsp_json_obj["bid_url"].(string)
	bid_url_slice := strings.Split(bid_url, "/")

	old_bid_url := bid_url
	if len(dspId) < 2 {
		bid_url_slice[2] = dstDspIp + ":" + "3800" + dspId
	} else {
		bid_url_slice[2] = dstDspIp + ":" + "380" + dspId
	}
	bid_url = strings.Join(bid_url_slice, "/")
	log.Debug("modify dspid:", dspId, " old bid_url:", old_bid_url, " to new bid_url:", bid_url)

	dsp_json_obj["bid_url"] = bid_url
	new_dsp_txt, err := json.Marshal(dsp_json_obj)
	log.Debugf("set dsp.%s  %s", dspId, ([]byte(new_dsp_txt)))

	reply, err := redisClient.Do("set", "dsp."+dspId, new_dsp_txt)
	log.Debug(reply)
	if err != nil {
		log.Error("Cannot set dsp.", dspId, " err = ", err)
	}
}

func modifyDspList(redisClient redis.Conn, dspList []string, dstDspIp string) {
	log.Debug("Will modify ", dspList)
	for _, dsp := range dspList {
		modifyDsp(redisClient, dsp, dstDspIp)
		log.Debug("====================")
	}
}

func main() {
	// 初始化日志, 日志的帮助文件在https://github.com/cihub/seelog/wiki
	defer log.Flush()

	// 获取配置
	redisHost, redisPort, dspList, logCfg, dstDspIp, err := getCfgInfo("modifyDspIp.conf")
	if err != nil {
		log.Critical("err parsing config log file", err)
		return
	}

	// 配置logger
	logger, err := log.LoggerFromConfigAsFile(logCfg)
	log.ReplaceLogger(logger)

	log.Debug("connecting redis", redisHost, redisPort)
	//连接redis, redigo的帮助文档在https://godoc.org/github.com/garyburd/redigo/redis
	redisClient, err := redis.Dial("tcp", redisHost+":"+redisPort)
	if err != nil {
		log.Critical("Cannot connect to redis", err)
		return
	}
	defer redisClient.Close()

	modifyDspList(redisClient, dspList, dstDspIp)
	return
	reply, err := redis.Values(redisClient.Do("keys", "*"))

	if err != nil {

	}
	for _, key := range reply {
		fmt.Println(reflect.TypeOf(key))
		fmt.Println(redis.String(key, err))
	}

}
