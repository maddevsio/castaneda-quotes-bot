package main

import (
	"github.com/peterbourgon/diskv"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Chat struct {
	id int64
	info string
}

func (c *Chat) Get() error {
	d := diskv.New(diskv.Options{
		BasePath:     "./db",
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 1024 * 1024, // 1MB
	})

	bytes, err := d.Read(string(c.id))
	if err != nil {
		return err
	}
	log.Printf("%v \n", bytes)
	cc := Chat{}
	bson.Unmarshal(bytes, cc)
	log.Printf("%v \n", cc.info)
	return nil
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
	return d.Write(string(c.id), b)
}