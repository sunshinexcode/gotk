package vmongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vreflect"
	"github.com/sunshinexcode/gotk/vsafe"
)

type (
	Mongodb struct {
		Client        *mongo.Client
		ClientOptions *options.ClientOptions
		Database      *mongo.Database
		Options       *Options
	}

	Options struct {
		Db           string        `mapstructure:",omitempty"`
		Limit        int           `mapstructure:",omitempty"` // query limit
		MaxPoolSize  int           `mapstructure:",omitempty"` // max pool size
		ReadPrefMode readpref.Mode `mapstructure:",omitempty"`
		Timeout      int           `mapstructure:",omitempty"` // second
		Uri          string        `mapstructure:",omitempty"`
	}
)

var (
	defaultOptions = map[string]any{
		"Db":           "test",
		"Limit":        10000,
		"MaxPoolSize":  100,
		"ReadPrefMode": readpref.SecondaryPreferredMode,
		"Timeout":      5,
		"Uri":          "mongodb://test:test@localhost:27017/test?replicaSet=test",
	}
)

// New create new mongodb
func New(options map[string]any) (mongodbS *Mongodb, err error) {
	mongodbS = &Mongodb{Options: &Options{}}
	err = mongodbS.SetConfig(options)

	return
}

// C shortcut for GetCol
func (mongodbS *Mongodb) C(col string) *mongo.Collection {
	return mongodbS.GetCol(col)
}

// GetCol select collection
func (mongodbS *Mongodb) GetCol(col string) *mongo.Collection {
	return mongodbS.Database.Collection(col)
}

func (mongodbS *Mongodb) Ping() error {
	return mongodbS.Client.Ping(context.TODO(), readpref.Primary())
}

// SetConfig set config
func (mongodbS *Mongodb) SetConfig(opts map[string]any) (err error) {
	mongodbS.ClientOptions = &options.ClientOptions{}
	if err = vreflect.SetAttrs(mongodbS.Options, vmap.Merge(defaultOptions, opts)); err != nil {
		return
	}

	rp, _ := readpref.New(mongodbS.Options.ReadPrefMode)
	mongodbS.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodbS.Options.Uri).SetMaxPoolSize(uint64(mongodbS.Options.MaxPoolSize)).SetReadPreference(rp))
	mongodbS.SetDb(mongodbS.Options.Db)

	return
}

// SetDb select database
func (mongodbS *Mongodb) SetDb(db string) {
	mongodbS.Database = mongodbS.Client.Database(db)
}

func (o *Options) String() (data string) {
	data, _ = vjson.Encode(o)

	return vsafe.MaskUrl(data)
}
