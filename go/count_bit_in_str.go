package main

import (
	"fmt"
	"time"
)

func main() {
	var a = []byte{1, 1, 1, 1}
	fmt.Println(countBit(a, 0, 23))
}

// 计算当前时间与当天 00:00:00 的分钟偏移数
func getMinuteOffset(t time.Time) int {
	dayStartTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	minuteCnt := t.Sub(dayStartTime) / time.Duration(time.Second*60)

	return int(minuteCnt)
}

func countBitForTimeRange(bArr []byte, startMinute, endMinute time.Time) int {
	startMinOffset := getMinuteOffset(startMinute)
	endMinOffset := getMinuteOffset(endMinute)
	return countBit(bArr, startMinOffset, endMinOffset)
}

func countBit(bArr []byte, startBit int, endBit int) int {
	// fmt.Println(bArr, startBit, endBit)
	var (
		startByteIdx    = startBit / 8
		startByteOffset = startBit % 8
		startByteMask   = 1<<(uint(8-startByteOffset)) - 1 // 00001111

		endByteIdx    = endBit / 8
		endByteOffset = endBit % 8
		endByteMask   = 255 << uint(8-endByteOffset-1) & 255 // 11110000

		totalBit = 0
	)

	if startByteIdx < len(bArr) {
		startByte := int(bArr[startByteIdx]) & startByteMask
		totalBit += countSingleByte(byte(startByte))
	}

	if endByteIdx < len(bArr) && endByteIdx != startByteIdx {
		endByte := int(bArr[endByteIdx]) & endByteMask
		totalBit += countSingleByte(byte(endByte))
	}

	for i := startByteIdx + 1; i <= endByteIdx-1 && i < len(bArr); i++ {
		totalBit += countSingleByte(bArr[i])
	}

	return totalBit
}

// 别看了，从 redis 抄的
func countSingleByte(b byte) int {
	var aux = int(b)
	aux = aux - ((aux >> 1) & 0x55555555)
	aux = (aux & 0x33333333) + ((aux >> 2) & 0x33333333)
	return (aux + (aux >> 4)) & 0x0F0F0F0F
}
