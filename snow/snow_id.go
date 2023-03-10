package snow

import (
	"math"
	"sync"
	"time"
)

/*
首位：第一个bit作为符号位，正数为0。
时间戳：占用41bit，精确到毫秒。41位最好可以表示2^41-1毫秒，转化成单位年为69年。
机器号：占用10bit，最多可以容纳1024个节点。
序列号：占用12bit，每个节点每毫秒从0开始不断累加，最多可以累加到4095，一共可以产生4096个ID。
*/
const (
	workerBits     = uint(10)                         // 机器号位数
	sequenceBits   = uint(12)                         // 序列号位数
	workerMax      = int64(-1 ^ (-1 << workerBits))   // 机器号最大值（即1023）
	sequenceMax    = int64(-1 ^ (-1 << sequenceBits)) // 序列号最大值（即4095）
	workerShift    = sequenceBits                     // 机器码偏移量
	timeStampShift = workerBits + sequenceBits        // 时间戳偏移量
	epoch          = int64(946656000000)              // 起始常量时间戳（毫秒）,此处选取的时间是2000-01-01 00:00:00
)

type Snow struct {
	sync.Mutex
	WorkerId  int64 // 机器号,0~1023
	TimeStamp int64 // 时间戳
	Sequence  int64 // 序列号
}

var NewSnow *Snow

func InitSnowFlake(workerId int64) {
	if NewSnow == nil {
		if workerId < 0 || workerId > workerMax {
			workerId = int64(math.Abs(float64(workerId % workerMax)))
		}
		NewSnow = &Snow{WorkerId: workerId, TimeStamp: 0, Sequence: 0}
	}
}

func (snow *Snow) GetId() int64 {
	snow.Lock()
	defer snow.Unlock()
	now := time.Now().UnixNano() / 1e6
	if snow.TimeStamp == now {
		snow.Sequence++
		if snow.Sequence > sequenceMax {
			for now <= snow.TimeStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		snow.Sequence = 0
		snow.TimeStamp = now
	}
	return (now-epoch)<<timeStampShift | (snow.WorkerId << workerShift) | (snow.Sequence)
}
