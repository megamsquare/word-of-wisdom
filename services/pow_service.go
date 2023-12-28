package services

import (
	"crypto/sha1"
	"fmt"
)

type HashcashData struct {
	Version    int
	ZerosCount int
	Date       int64
	Resource   string
	Rand       string
	Counter    int
}

func sha1Hash(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	return fmt.Sprintf("%v", sum)
}

func IsHash(hash string, counts int) bool {
	if counts > len(hash) {
		return false
	}

	for _, v := range hash[:counts] {
		if v != 48 {
			return false
		}
	}
	return true
}

func (h HashcashData) ConvertToString() string {
	return fmt.Sprintf("%d:%d:%d:%s::%s:%d", h.Version, h.ZerosCount, h.Date, h.Resource, h.Rand, h.Counter)
}

func (h HashcashData) CalCulateHashcash(maxLoop int) (HashcashData, error)  {
	for h.Counter <= maxLoop || maxLoop <= 0 {
		header := h.ConvertToString()
		hash := sha1Hash(header)

		if IsHash(hash, h.ZerosCount) {
			return h, nil
		}
		h.Counter++
	}
	return h, fmt.Errorf("exceed max loop")
}