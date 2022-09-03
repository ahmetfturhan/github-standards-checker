package output

import (
	"fmt"

	"github.com/r3labs/diff/v3"
)

//For printing the changelog
func PrintDiffChangeLog(changelog diff.Changelog, RET_VALUE int) int {
	isEmpty := true
	for _, v := range changelog {
		if v.From != nil && v.To != nil {
			RET_VALUE++
			isEmpty = false
			fmt.Printf("\nRule Name: %v\nYAML Value: %v\nGitHub API Value: %v\n\n", v.Path, v.From, v.To)
			// retValue++
		}
	}
	if isEmpty {
		fmt.Printf("\nNo differences\n")
	}
	return RET_VALUE
}
