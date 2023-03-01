/**
* @description :
* @author : Jarick
* @Date : 2022-07-14
* @Url : http://CloudWebOps
 */

package util

import (
	"fmt"
	"runtime"
)

var (
	version      string
	gitBranch    string
	gitTag       string
	gitCommit    string
	gitTreeState string
	gitAuthor    string
	buildDate    string
)

type Info struct {
	Version      string `json:"version"`
	GitBranch    string `json:"gitBranch"`
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitAuthor    string `json:"gitAuthor"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	return info.GitCommit
}

func GetVersion() Info {
	return Info{
		Version:      version,
		GitBranch:    gitBranch,
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitAuthor:    gitAuthor,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func VersionFucn() {
	v := GetVersion()
	fmt.Printf("Version: \t%s\nCommitId: \t%s\nGitBranch: \t%s\nGitAuthor: \t%s\nBuild Date: \t%s\nGo Version: \t%s\nOS/Arch: \t%s\n", v.Version, v.GitCommit, v.GitBranch, v.GitAuthor, v.BuildDate, v.GoVersion, v.Platform)
}
