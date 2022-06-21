package randomly

import (
	"fmt"
	"math/rand"
)

func SetSeed(seed int64) {
	rand.Seed(seed)
}

func RandIntGap(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandInt32() int32 {
	return rand.Int31()
}

func RandInt64() int64 {
	return rand.Int63()
}

func RandUnInt32() uint32 {
	return rand.Uint32()
}

func RandUnInt64() uint64 {
	return rand.Uint64()
}

func RandFloat() float32 {
	return float32(RandIntGap(-1000, 1000)) + rand.Float32()
}

func RandDouble() float64 {
	return float64(RandIntGap(-1000, 1000)) + rand.Float64()
}

func RandBytesLen(l int) []byte {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandIntGap(0, 255))
	}
	return bytes
}

func RandDateStr() string {
	year := RandIntGap(1970, 2038)
	month := RandIntGap(1, 13)
	date := RandIntGap(1, 28)
	return fmt.Sprintf("%d-%02d-%02d", year, month, date)
}

func RandTimeStr() string {
	hour := RandIntGap(0, 24)
	minute := RandIntGap(0, 60)
	second := RandIntGap(0, 60)
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func RandDateTimeStr() string {
	return RandDateStr() + " " + RandTimeStr()
}

func RandBool() bool {
	return rand.Intn(2) == 0
}
