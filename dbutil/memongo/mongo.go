package memgo

import (
	"errors"
	"gopkg.in/mgo.v2"
	"sync"
	"time"
)

var (
	defaultSession *mgo.Session
	sessionLock    = new(sync.RWMutex)
	sessions       = make(map[string]*mgo.Session)
)

// MongoConfig config
type MongoConfig struct {
	Addrs          []string `json:"addrs" toml:"addrs"`
	Source         string   `json:"source" toml:"source"`
	ReplicaSetName string   `json:"replica_set_name" toml:"replica_set_name"`
	Timeout        int      `json:"timeout" toml:"timeout"`
	Username       string   `json:"username" toml:"username"`
	Password       string   `json:"password" toml:"password"`
	Mode           *int     `json:"mode" toml:"mode"`
	Alias          string   `json:"alias" toml:"alias"`
}

func (mc *MongoConfig) getMongoDialInfo() *mgo.DialInfo {
	if mc.Mode == nil {
		m := int(mgo.PrimaryPreferred)
		mc.Mode = &m
	}
	if *mc.Mode > int(mgo.Nearest) {
		m := int(mgo.PrimaryPreferred)
		mc.Mode = &m
	}
	if mc.Timeout == 0 {
		mc.Timeout = 2
	}
	return &mgo.DialInfo{
		Addrs:          mc.Addrs,
		Source:         mc.Source,
		ReplicaSetName: mc.ReplicaSetName,
		Timeout:        time.Duration(mc.Timeout) * time.Second,
		Username:       mc.Username,
		Password:       mc.Password,
	}
}

func InitMongo(mc *MongoConfig) error {
	dialInfo := mc.getMongoDialInfo()
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return err
	}
	s.SetMode(*mc.Mode, true)
	sessionLock.Lock()
	defer sessionLock.Unlock()
	if _, ok := sessions[mc.Alias]; ok {
		return errors.New("duplicate session")
	}
	sessions[mc.Alias] = s
	if mc.Alias == "" {
		defaultSession = s
	}
	return nil
}

func GetSession() *mgo.Session {
	if defaultSession != nil {
		return defaultSession.Copy()
	}
	return nil
}

func GetSessionBy(alias string) *mgo.Session {
	sessionLock.RLock()
	defer sessionLock.RUnlock()
	if s, ok := sessions[alias]; ok && s != nil {
		return s.Copy()
	}
	return nil
}
