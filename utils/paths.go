package utils

import (
	"net/url"
	"path/filepath"
	"regexp"
)

type PathType int

const (
	INVALID_PATH PathType = iota
	FILE_PATH
	URL_PATH
)

func IdentifyPath(path string) PathType {
	// Check if the input string is a URL
	u, err := url.Parse(path)
	if err == nil && u.Scheme != "" && u.Opaque == "" {
		return URL_PATH
	}

	if len(path) > 0 && string(path[0]) == "/" && filepath.Clean(path) != "." {
		return FILE_PATH
	}

	// Check if the input string is a Windows file path
	if match, _ := regexp.MatchString(`^[a-zA-Z]:\\`, path); match {
		return FILE_PATH
	}

	// If the input string is neither a URL nor a file path, it is something else
	return INVALID_PATH
}

func IsURL(path string) bool {
	return IdentifyPath(path) == URL_PATH
}

func IsFile(path string) bool {
	return IdentifyPath(path) == FILE_PATH
}

func IsValidPath(path string) bool {
	return IdentifyPath(path) != INVALID_PATH
}
