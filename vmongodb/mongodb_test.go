package vmongodb_test

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vmongodb"
	"github.com/sunshinexcode/gotk/vstr"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewMock(t *testing.T) {
	var err error
	mongodb := &vmongodb.Mongodb{Options: &vmongodb.Options{}}
	patch := vmock.ApplyMethod(reflect.TypeOf(mongodb), "SetConfig", func(mongodb *vmongodb.Mongodb, options map[string]any) error {
		mongodb.Options = &vmongodb.Options{Uri: "mongodb://test:test@localhost:27017/test?replicaSet=test"}
		return nil
	})
	defer vmock.Reset(patch)

	vtest.Equal(t, "", mongodb.Options.Uri)
	vtest.Equal(t, readpref.Mode(0x0), mongodb.Options.ReadPrefMode)
	vtest.Equal(t, readpref.Mode(0x1), readpref.PrimaryMode)
	vtest.Equal(t, "1", vstr.S("%d", readpref.PrimaryMode))
	vtest.Equal(t, readpref.Mode(0x4), readpref.SecondaryPreferredMode)
	vtest.Equal(t, "4", vstr.S("%d", readpref.SecondaryPreferredMode))

	mongodb, err = vmongodb.New(nil)

	vtest.Nil(t, err)
	vtest.Equal(t, "mongodb://test:test@localhost:27017/test?replicaSet=test", mongodb.Options.Uri)
}

func TestC(t *testing.T) {
	mongodb := vmongodb.Mongodb{Client: &mongo.Client{}}
	mongodb.SetDb("test")

	vtest.Equal(t, "test", mongodb.C("test").Name())
}

func TestGetCol(t *testing.T) {
	mongodb := vmongodb.Mongodb{Client: &mongo.Client{}}
	mongodb.SetDb("test")

	vtest.Equal(t, "test", mongodb.GetCol("test").Name())
}

func TestPing(t *testing.T) {
	mongodb := vmongodb.Mongodb{Client: &mongo.Client{}}
	mongodb.SetDb("test")
	err := mongodb.Ping()

	vtest.NotNil(t, err)
	vtest.Equal(t, "the Command operation must have a Deployment set before Execute can be called", err.Error())
}

func TestSetConfig(t *testing.T) {
	mongodb := vmongodb.Mongodb{Client: &mongo.Client{}, Options: &vmongodb.Options{}}
	mongodb.SetDb("test")
	err := mongodb.SetConfig(nil)

	vtest.Nil(t, err)

	err = mongodb.SetConfig(vmap.M{"Test": ""})

	vtest.NotNil(t, err)
	vtest.Equal(t, "no attr, attr:Test", err.Error())

	err = mongodb.SetConfig(vmap.M{"ReadPrefMode": ""})

	vtest.NotNil(t, err)
	vtest.Equal(t, "error type, attr:ReadPrefMode, wrong:readpref.Mode, correct:string", err.Error())

	err = mongodb.SetConfig(vmap.M{"ReadPrefMode": readpref.Mode(4)})

	vtest.Nil(t, err)

	err = mongodb.SetConfig(vmap.M{"ReadPrefMode": readpref.SecondaryPreferredMode})

	vtest.Nil(t, err)
}

func TestString(t *testing.T) {
	vtest.Equal(t, `{"Db":"","Limit":0,"MaxPoolSize":0,"ReadPrefMode":0,"Timeout":0,"Uri":"mongodb://test:***@localhost:27017/test?replicaSet=test"}`, vstr.S("%+v", &vmongodb.Options{Uri: "mongodb://test:test@localhost:27017/test?replicaSet=test"}))
}
