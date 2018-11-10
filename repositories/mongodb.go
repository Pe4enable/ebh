package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/BankEx/ebh/config"
	"github.com/BankEx/bankex-tokensale-common/mgoconn"
	"github.com/BankEx/ebh/models"
	"time"
)

type MongoRepository struct {
	conf         *config.Mongo
	session      *mgo.Session
	valuesC      *mgo.Collection
}

func New(config *config.Mongo) (repository *MongoRepository, err error) {
	repository = new(MongoRepository)
	repository.conf = config

	repository.session = mgoconn.DialLoop(repository.conf.URL)
	mgoconn.TimeoutReconnectFunc = mgoconn.ReconnectFuncForSessions([]*mgo.Session{repository.session})

	db := repository.session.DB(repository.conf.DB)
	repository.valuesC = db.C(repository.conf.Collection)

	return
}

func (r *MongoRepository) GetAllSertificates() (result []models.Sertificate, err error) {

	err = mgoconn.DoResistant(
		func() error {
			return r.valuesC.Find(nil).All(&result)
		},
	)
	if err != nil {
		return nil, err
	}

	return
}

func (r *MongoRepository) CreateSertificate (data models.Sertificate) (err error){

	err = mgoconn.DoResistant(
		func() error {
			data.ID=bson.NewObjectId()
			data.CreateDate=time.Now()
			return r.valuesC.Insert(data)
		},
	)
	return
}