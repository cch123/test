package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/spaolacci/murmur3"
	"github.com/stretchr/testify/assert"
)

func getOriginalBackends() []string {
	var servers = make([]string, 1000)
	for i := 0; i < 1000; i++ {
		servers[i] = fmt.Sprint(i)
	}
	return servers
}

func TestBalanced(t *testing.T) {
	assert.Equal(t, 1, 1)

	// 假设每个 client 向后端建 80 条连接
	subsetSize := 100
	count := map[string]int{}
	// 随机生成 1000 个 host
	rand.Seed(time.Now().Unix())
	for i := 0; i < 50000; i++ {
		ip := fmt.Sprintf("%v.%v.%v.%v", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
		// TODO farmhash
		// 这里 client id 用 hash 类的值是比较难保证均匀的
		clientID := murmur3.Sum32([]byte(ip))
		//clientID = uint32(i)
		subset := Subset(getOriginalBackends(), int(clientID), subsetSize)
		for _, server := range subset {
			count[server]++
		}
	}

	max, min := 0, math.MaxUint32
	for _, v := range count {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max, min)
}
