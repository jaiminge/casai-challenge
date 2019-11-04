package main

import (
	"errors"
	"reflect"
	"testing"

	"casai.com/posada/lib"
	"github.com/go-pg/pg/orm"
)

type mockPgDb struct {
	calls []reflect.Type
	opts  []*orm.CreateTableOptions
}

func (m *mockPgDb) CreateTable(model interface{}, opt *orm.CreateTableOptions) error {
	if reflect.TypeOf(model) == nil {
		return errors.New("INFO")
	}
	m.calls = append(m.calls, reflect.TypeOf(model))
	m.opts = append(m.opts, opt)
	return nil
}

func TestMaybeCreateSchemas(t *testing.T) {
	db := &mockPgDb{}

	// now we execute our method
	maybeCreateSchemas(db)

	emptyOpts := &orm.CreateTableOptions{}

	// verify calls were made.
	for i := range db.calls {
		if db.calls[i] != reflect.TypeOf(lib.Tables[i]) {
			t.Errorf("Calls Differ %v %v", db.calls[i], reflect.TypeOf(lib.Tables[i]))
		}
		if *db.opts[i] != *emptyOpts {
			t.Errorf("Options Differ %v %v", *db.opts[i], *emptyOpts)
		}
	}
}

func TestLoadDevConfig(t *testing.T) {
	conf := loadConfig(false)
	if conf.Port == 0 || conf.DbName == "" {
		t.Fatal("Missing required parameters")
	}

}
