package main

import (
	"github.com/peterbourgon/diskv"
	"gopkg.in/mgo.v2/bson"
)

type Chat struct {
	Id int64
	Info string
}

func (c *Chat) Get() error {
	d := diskv.New(diskv.Options{
		BasePath:     "./db",
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 1024 * 1024, // 1MB
	})

	bytes, err := d.Read(string(c.Id))
	if err != nil {
		return err
	}
	return bson.Unmarshal(bytes, c)
}

func (c *Chat) Save() error {
	d := diskv.New(diskv.Options{
		BasePath:     "./db",
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 1024 * 1024, // 1MB
	})

	b, err := bson.Marshal(c)
	if (err != nil) {
		return err
	}
	return d.Write(string(c.Id), b)
}