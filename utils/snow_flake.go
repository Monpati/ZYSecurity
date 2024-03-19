package utils

import (
	"errors"
	"sync"
	"time"
)

const (
	twepoch        = int64(1483228800000)
	workeridBits   = uint(10)
	sequenceBits   = uint(12)
	workeridMax    = int64(-1 ^ (-1 << workeridBits))
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits))
	workeridShift  = sequenceBits
	timestampShift = sequenceBits + workeridBits
)

type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerid  int64
	sequence  int64
}

func NewSnowflake(workerid int64) (*Snowflake, error) {
	if workerid < 0 || workerid > workeridMax {
		return nil, errors.New("workerid must be between 0 .. 1023")
	}
	return &Snowflake{
		timestamp: 0,
		workerid:  workerid,
		sequence:  0,
	}, nil
}

func (s *Snowflake) Generate() int64 {
	s.Lock()
	now := time.Now().UnixNano() / 1000000
	if s.timestamp == now {
		s.sequence = (s.timestamp + 1) & sequenceMask

		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 1
	}
	s.timestamp = now
	r := int64((now-twepoch)<<timestampShift | (s.workerid << workeridShift) | (s.sequence))
	s.Unlock()
	return r
}
