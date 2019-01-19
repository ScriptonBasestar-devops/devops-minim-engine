package process

import (
	"fmt"
	"testing"
)

func TestCacheName(t *testing.T) {
	v := cacheName(grp, "aaaa", "bbbb")
	fmt.Println(v)
}
