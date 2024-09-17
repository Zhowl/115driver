package driver

const (
	UADefalut    = "Mozilla/5.0"
	UA115Browser = "Mozilla/5.0 115Browser/27.0.3.7"
	UA115Disk    = "Mozilla/5.0 115disk/32.7.0"
	UA115Desktop = "Mozilla/5.0 115Desktop/2.0.10.2"
	UAIosApp     = "Mozilla/5.0; Darwin/10.0; UDown/32.3.1"
)

const (
	CookieDomain115 = ".115.com"

	CookieUrl = "https://115.com"

	CookieNameUid  = "UID"
	CookieNameCid  = "CID"
	CookieNameSeid = "SEID"
)

const (
	OSSRegionID = "oss-cn-shenzhen"
	OSSEndpoint = "cn-shenzhen.oss.aliyuncs.com" // 双栈域名

	OSSUserAgent               = "aliyun-sdk-android/2.9.1"
	OssSecurityTokenHeaderName = "X-OSS-Security-Token"
)

const (
	KB = 1 << (10 * (iota + 1))
	MB
	GB
)
