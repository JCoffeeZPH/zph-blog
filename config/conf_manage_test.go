package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetBrowserDetailUrl(t *testing.T) {
	fmt.Println(GetIpDetailUrl())
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	filename := basePath + string(os.PathSeparator) + "config.yaml"
	fmt.Println(filename)
}
