/*
Package parser provides config parse funcs.
*/
package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Parse parses config file and returns a Config.
func Parse[T any](filename string) (*T, error) {
	var c T
	buf, err := os.ReadFile(filename)
	if err != nil {
		return &c, err
	}
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return &c, fmt.Errorf("failed to parse file %s: %v", filename, err)
	}
	return &c, nil
}
