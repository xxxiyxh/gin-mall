package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%d", strconv.Itoa(int(id)))
}
