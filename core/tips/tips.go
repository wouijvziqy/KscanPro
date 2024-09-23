package tips

import (
	"math/rand"
	"time"
)

var Tips = []string{
	"可以使用--check参数，对给定的ip:port,protocol://ip:port/path等字符串进行检测，且不会进行端口扫描",
	"现在-t参数可以直接输入文件路径，不用再增加file:前缀了",
	"如果要进行暴力破解，可以使用--hydra参数",
	"--spy参数，如果不带参数，将会默认对本机所在的网段进行网关存活性探测，如果再加上--scan参数，则会对存活的网段进行端口扫描",
	"可以使用--spy 10.10.20.1，将会对该网址10.10.20.1/16(B段)进行网关存活性探测",
	"可以使用--spy 10，将会进行10.0.0.1/8(A段)进行网关存活性探测",
	"可以使用--spy 172，将会对172.16.0.1/12进行网关存活性探测",
	"可以使用--spy 192，将会对192.168.0.1/16(B段)进行网关存活性探测",
	"-sV参数，一般不建议在进行端口扫描的时候使用，一般用于扫描某一个端口，因为实在影响效率",
	"可以使用-Pn参数关闭存活性探测，将会将所有主机视为存活",
	"在Windows操作系统下，可以新增环境变量KSCAN_COLOR=TRUE,开启颜色显示",
}

func GetTips() string {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(Tips))
	return Tips[index]
}
