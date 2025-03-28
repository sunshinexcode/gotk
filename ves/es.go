package ves

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"

	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vreflect"
	"github.com/sunshinexcode/gotk/vsafe"
)

type (
	Es struct {
		Client  *elastic.Client
		Options *Options
	}

	Options struct {
		Url string `mapstructure:",omitempty"`
	}
)

var (
	defaultOptions = map[string]any{
		"Url": "http://127.0.0.1:9200/test?shards=1&replicas=0&sniff=false",
	}
)

// New create new mongodb
func New(options map[string]any) (esS *Es, err error) {
	esS = &Es{Options: &Options{}}
	err = esS.SetConfig(options)

	return
}

func (esS *Es) C() *elastic.Client {
	return esS.Client
}

func (esS *Es) Ping() (*elastic.PingResult, int, error) {
	return esS.Client.Ping(esS.Options.Url).Do(context.Background())
}

// SetConfig set config
func (esS *Es) SetConfig(options map[string]any) (err error) {
	if err = vreflect.SetAttrs(esS.Options, vmap.Merge(defaultOptions, options)); err != nil {
		return
	}

	cfg, err := config.Parse(esS.Options.Url)
	if err != nil {
		return
	}

	esS.Client, err = elastic.NewClientFromConfig(cfg)
	return
}

func (o *Options) String() (data string) {
	data, _ = vjson.Encode(o)

	return vsafe.MaskUrl(data)
}
