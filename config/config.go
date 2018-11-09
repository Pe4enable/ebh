package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"os"
	"strconv"
)

type Config struct {
	Ticker     string // BTC
	Multiplier int64
	Cache           int64
	Interval int64

	NodeConfig *Node
	DBConfig   *Mongo
	Port       string
}

func LoadConfig(filePath string) (conf *Config, err error) {
	viper.SetConfigFile(filePath)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	log.Printf("Read config")

	conf = new(Config)
	conf.Interval = GetIntEnvValue("INTERVAL", viper.GetInt64("interval"))
	conf.Cache = GetIntEnvValue("CACHE", viper.GetInt64("cache"))
	conf.Ticker = GetStringEnvValue("TICKER", viper.GetString("ticker"))
	conf.Multiplier = GetIntEnvValue("DECIMALS", viper.GetInt64("decimals"))
	conf.NodeConfig = &Node{
		GetStringEnvValue("NODE_HOST_URL", viper.GetString("node.host")),
		GetStringEnvValue("NODE_TYPE", viper.GetString("node.type")),
		GetIntEnvValue("START_BLOCK_HEIGHT", viper.GetInt64("block")),
		GetIntEnvValue("CONFIRM_COUNT", viper.GetInt64("confirmations")),
	}

	conf.DBConfig = &Mongo{
		GetStringEnvValue("DB_NAME", viper.GetString("db.name")),
		GetStringEnvValue("DB_URL", viper.GetString("db.url")),
		GetStringEnvValue("DB_COLLECTION", viper.GetString("db.collection")),
	}

	conf.Port = ":" + GetStringEnvValue("PORT", viper.GetString("port"))
	log.Printf("Loaded config: %#v", conf)
	return
}

func GetIntEnvValue(name string, defValue int64) (result int64) {
	if startBlockHeight, ok := os.LookupEnv(name); ok {
		s, err := strconv.ParseInt(startBlockHeight, 10, 64)
		if err != nil {
			log.Printf("ENV-%s not found, used default value %s", name, defValue)
			return defValue
		}
		return s
	} else {
		log.Printf("ENV-%s not found, used default value %s", name, defValue)
		return defValue
	}
}

func GetStringEnvValue(name string, defValue string) (result string) {
	if u, ok := os.LookupEnv(name); ok {
		return u
	} else {
		log.Printf("ENV-%s not found, used default value '%s'", name, defValue)
		return defValue
	}
}


func (conf Config) AmountWithMultiplier(amount float64) *big.Int {
	bigAmount := big.NewFloat(amount)
	mul := new(big.Float).SetInt(conf.BigMultiplier())
	bigAmount = bigAmount.Mul(bigAmount, mul)

	bigAmountInt := new(big.Int)
	bigAmount.Int(bigAmountInt)

	return bigAmountInt
}

func (conf Config) BigMultiplier() *big.Int {
	return bigIntWithMultiplier(conf.Multiplier)
}

func bigIntWithMultiplier(m int64) *big.Int {
	if m <= 1 {
		return big.NewInt(1)
	}

	s := "1"
	var i int64
	for i = 0; i < m; i++ {
		s += "0"
	}

	bi := new(big.Int)
	bi, ok := bi.SetString(s, 10)
	if !ok {
		panic(fmt.Errorf("Can't convert %d to *big.Int: string=`%s`", m, s))
	}
	return bi
}
