package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var (
	// git branch output of $(git rev-parse --abbrev-ref HEAD)
	gitVersion = "v0.0.0-master+$Format:%h$"
	// sha1 from git, output of $(git rev-parse HEAD)
	gitCommit = "$Format:%H$"
	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildDate = "1970-01-01T00:00:00Z"
)

type Version struct {
	//Major        string `json:"major"`
	//Minor        string `json:"minor"`
	GitVersion string `json:"gitVersion"`
	GitCommit  string `json:"gitCommit"`
	//GitTreeState string `json:"gitTreeState"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

// String return a json encoded string with 2 space indents.
func (v Version) String() string {
	bytes, _ := json.MarshalIndent(v, "", " ")
	return string(bytes)
}

// Get version 当前发行的版本信息
func Get() Version {
	return Version{
		GitVersion: gitVersion,
		GitCommit:  gitCommit,
		BuildDate:  buildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
