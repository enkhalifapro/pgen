package test

import (
	"testing"

	"github.com/facebookgo/inject"
	"github.com/spf13/viper"
	"gitlab.com/enkhalifapro/ulearn-api/config"
	"gitlab.com/enkhalifapro/ulearn-api/db"
)

// Initialize element initialization.
func Initialize(t *testing.T, objects ...interface{}) (close func()) {
	if err := config.Load("test", "../etc"); err != nil {
		t.Fatal(err)
	}

	// DB
	dbc, err := db.Dial(viper.GetString("db.uri"))
	if err != nil {
		t.Fatal(err)
	}
	close = dbc.Close

	allObjects := []*inject.Object{
		&inject.Object{Value: dbc},
	}

	for _, o := range objects {
		allObjects = append(allObjects, &inject.Object{Value: o})
	}

	graph := &inject.Graph{}
	if err := graph.Provide(allObjects...); err != nil {
		t.Fatalf("prepare DI graph: %v", err)
	}

	if err := graph.Populate(); err != nil {
		t.Fatalf("populate DI graph: %v", err)
	}
	return
}
