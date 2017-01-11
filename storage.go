package main

import (
	"github.com/peterbourgon/diskv"
	"gopkg.in/mgo.v2/bson"
)

type Chat struct {
	Id       int64
	UserName string
}

func GetStorage(path string) *diskv.Diskv {
	return diskv.New(diskv.Options{
		BasePath:     path,
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 1024 * 1024, // 1MB
	})
}

func (c *Chat) Get(d *diskv.Diskv) error {
	bytes, err := d.Read(string(c.Id))
	if err != nil {
		return err
	}
	return bson.Unmarshal(bytes, c)
}

func (c *Chat) Save(d *diskv.Diskv) error {
	b, err := bson.Marshal(c)
	if (err != nil) {
		return err
	}
	return d.Write(string(c.Id), b)
}

func GetAllChats(d *diskv.Diskv) ([]Chat, error) {
	var chats []Chat
	for key := range d.Keys(nil) {
		val, err := d.Read(key)
		if err != nil {
			return nil, err
		}
		c := Chat{}
		bson.Unmarshal(val, &c)
		chats = append(chats, c)
	}
	return chats, nil
}