package curd

import (
	"github.com/spf13/viper"
	"strings"
)

func GetOSSDomain(link string) string {
	if strings.HasPrefix(link, "http") {
		if !strings.Contains(link, "?x-oss-process=") {
			return link + "?x-oss-process=image/interlace,1/resize,p_60/quality,q_75/format,jpeg"
		}
		return link
	}

	link = strings.TrimSuffix(link, "/")
	domain := strings.TrimSuffix(viper.GetString("oss.domain"), "/")

	return domain + "/" + link + "?x-oss-process=image/interlace,1/resize,p_60/quality,q_75/format,jpeg"
}
