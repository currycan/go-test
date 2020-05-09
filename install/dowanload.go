package install

import (
	"fmt"
	"github.com/currycan/go-test/pkg/sshcmd/cmd"
	"net/url"
	"path"
)

//
func DownloadFile(location string) (filePATH, md5 string) {
	if _, isUrl := isUrl(location); isUrl {
		absPATH := "/" + path.Base(location)
		if !cmd.IsFileExist(absPATH) {
			//generator download sshcmd
			dwnCmd := downloadCmd(location)
			//os exec download command
			cmd.Cmd("/bin/sh", "-c", "mkdir -p /tmp/sealos && cd /tmp/sealos && "+dwnCmd)
		}
		location = absPATH
	}
	//file md5
	md5 = md5sum.FromLocal(location)
	return location, md5
}

//根据url 获取command
func downloadCmd(url string) string {
	//only http
	u, isHttp := isUrl(url)
	var c = ""
	if isHttp {
		param := ""
		if u.Scheme == "https" {
			param = "--no-check-certificate"
		}
		c = fmt.Sprintf(" wget -c %s %s", param, url)
	}
	return c
}

func isUrl(u string) (url.URL, bool) {
	if uu, err := url.Parse(u); err == nil && uu != nil && uu.Host != "" {
		return *uu, true
	}
}
