package ves_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/olivere/elastic/v7"

	"github.com/sunshinexcode/gotk/ves"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vstr"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewMock(t *testing.T) {
	var err error
	es := &ves.Es{Client: &elastic.Client{}, Options: &ves.Options{}}
	patch := vmock.ApplyMethod(reflect.TypeOf(es), "SetConfig", func(es *ves.Es, options map[string]any) error {
		es.Options = &ves.Options{Url: "localhost"}
		return nil
	})
	defer vmock.Reset(patch)

	vtest.Equal(t, "", es.Options.Url)

	es, err = ves.New(nil)

	vtest.Nil(t, err)
	vtest.Equal(t, "localhost", es.Options.Url)
}

func TestC(t *testing.T) {
	es := ves.Es{Client: &elastic.Client{}, Options: &ves.Options{Url: "http://localhost"}}

	vtest.Equal(t, "", es.C().String())
}

func TestPing(t *testing.T) {
	es := ves.Es{Client: &elastic.Client{}, Options: &ves.Options{Url: "http://localhost"}}
	defer func() {
		r := recover()

		vtest.Equal(t, "runtime error: invalid memory address or nil pointer dereference", vstr.S("%s", r))
	}()
	_, _, _ = es.Ping()
}

func TestSetConfig(t *testing.T) {
	es := ves.Es{Client: &elastic.Client{}, Options: &ves.Options{}}
	err := es.SetConfig(vmap.M{"Url": "http://192.168.1.1:9500"})

	vtest.NotNil(t, err)

	err = es.SetConfig(vmap.M{"Url": "http://test:123@192.168.1.1:9500/test?shards=1&replicas=0&sniff=false"})

	vtest.NotNil(t, err)

	err = es.SetConfig(vmap.M{"Url": "http://192.168.1.1:9500?\u007F"})

	vtest.NotNil(t, err)
	vtest.Equal(t, true, strings.Contains(err.Error(), "error parsing elastic parameter"))

	err = es.SetConfig(vmap.M{"Test": ""})

	vtest.NotNil(t, err)
	vtest.Equal(t, "no attr, attr:Test", err.Error())
}

func TestString(t *testing.T) {
	vtest.Equal(t, `{Url:http://127.0.0.1:9200/test?shards=1&replicas=0&sniff=false}`, vstr.S("%+v", ves.Options{Url: "http://127.0.0.1:9200/test?shards=1&replicas=0&sniff=false"}))
	vtest.Equal(t, `{"Url":"http://test:***@127.0.0.1:9200/test?shards=1\u0026replicas=0\u0026sniff=false"}`, vstr.S("%+v", &ves.Options{Url: "http://test:123@127.0.0.1:9200/test?shards=1&replicas=0&sniff=false"}))
}
