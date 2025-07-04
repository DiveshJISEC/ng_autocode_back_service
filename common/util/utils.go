package util

import (
	"math/rand"
	cmn "ng_autocode_back_service/common/model"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type IntRange struct {
	Min int
	Max int
}

var (
	rnd *rand.Rand
	ir  IntRange
)

func Init() error {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	ir = IntRange{0, 9}
	return nil
}

func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.Max-ir.Min+1) + ir.Min
}

func RandomInt(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return rnd.Intn(max-min+1) + min
}

func RandomIntRange(ir IntRange) int {
	if ir.Min > ir.Max {
		ir.Min, ir.Max = ir.Max, ir.Min
	}
	return rnd.Intn(ir.Max-ir.Min+1) + ir.Min
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rnd.Intn(len(charset))]
	}
	return string(b)
}

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func DaySeconds(t time.Time) int {
	return t.Hour()*3600 + t.Minute()*60 + t.Second()
}

func ReadTextFile(filePath string) string {
	filePath = UpdatePathSeparator(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return "error reading file: " + err.Error()
	}
	defer file.Close()
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "error reading file: " + err.Error()
	}
	return string(content)
}

func WriteToFile(filePath, text string) error {
	filePath = UpdatePathSeparator(filePath)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if len(text) > 0 {
		_, err = f.WriteString(text)
		if err != nil {
			return err
		}
	}

	return err
}

func CreateFolders(path string) error {
	path = UpdatePathSeparator(path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func CraeteFileWithPath(sPath string) error {
	dirPath := filepath.Dir(sPath)
	sPath = UpdatePathSeparator(sPath)
	err := CreateFolders(dirPath)
	if err != nil {
		return err
	}

	err = WriteToFile(sPath, "")
	if err != nil {
		return err
	}

	return err
}

func GetCurrentMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetPathSeparator() rune {
	switch os := runtime.GOOS; os {
	case "windows":
		return cmn.PATH_SEPARATOR_WINDOWS
	case "linux":
		return cmn.PATH_SEPARATOR_LINUX
	}
	return cmn.PATH_SEPARATOR_LINUX
}

func UpdatePathSeparator(path string) string {
	sep := GetPathSeparator()
	if sep == cmn.PATH_SEPARATOR_WINDOWS {
		path = strings.ReplaceAll(path, string(cmn.PATH_SEPARATOR_LINUX), string(cmn.PATH_SEPARATOR_WINDOWS))
	} else {
		path = strings.ReplaceAll(path, string(cmn.PATH_SEPARATOR_WINDOWS), string(cmn.PATH_SEPARATOR_LINUX))
	}
	return path
}
