package utils

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func InitStoreSess() {
	Store = session.New()
}

func SetStore(key string, value interface{}, exp_time time.Duration, c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return err
	}
	sess.SetExpiry(exp_time)
	sess.Set(key, value)
	return sess.Save()
}

func GetValue(key string, c *fiber.Ctx) (interface{}, error) {
	sess, err := Store.Get(c)
	if err != nil {
		return nil, err
	}
	return sess.Get(key), nil
}

func ResetValue(key string, c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return err
	}
	sess.Delete(key)
	return nil
}

func StoreRoute(c *fiber.Ctx) error {
	value, err := GetValue("original-route", c)
	if err != nil {
		return err
	}
	if value != nil {
		err := SetStore("original-route", c.Route().Path, time.Minute*2, c)
		if err != nil {
			return err
		}
	}
	return nil
}

// Custom encoding function for gob
func gobEncode(val interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(val); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Custom decoding function for gob
func gobDecode(data []byte, val interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(val); err != nil {
		return err
	}
	return nil
}
