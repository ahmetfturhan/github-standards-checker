package fileops

import (
	"os"
	"path"
)

func DeleteTempFiles(dirPath string) {
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		panic("Cannot read")
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{dirPath, d.Name()}...))
	}
}

func CreateEmptyYaml() {
	emptyF := "branchName:\nProtection:"
	if err := os.WriteFile("/empty-yaml/empty.yaml", []byte(emptyF), 0777); err != nil {
		panic(err)
	}
}
