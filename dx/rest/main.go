package main

import (
	"Dexun/config"
	"Dexun/controller"
	"Dexun/middlewares"
	"Dexun/utils"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	httpPort := flag.Int("port", 8080, "port")
	flag.Parse()

	db := config.InitDB()
	defer db.Close()
	config.RedisInit()

	app := gin.Default()
	app.Use(utils.Cors())

	controller.GetToken()
	time.AfterFunc(24*time.Hour, controller.GetToken)
	controller.OpenCert()
	controller.GetSCDNComboLists()
	time.AfterFunc(24*time.Hour, controller.GetSCDNComboLists)
	controller.GetDDoSComboLists()
	time.AfterFunc(24*time.Hour, controller.GetDDoSComboLists)
	controller.AllowDexunDDoSSell()
	time.AfterFunc(24*time.Hour, controller.AllowDexunDDoSSell)
	controller.GetDexunSSLList()
	time.AfterFunc(24*time.Hour, controller.GetDexunSSLList)
	controller.CopyDexunSSLList()
	time.AfterFunc(24*time.Hour, controller.CopyDexunSSLList)
	controller.GetCloudEyeComboLists()
	time.AfterFunc(24*time.Hour, controller.GetCloudEyeComboLists)
	controller.CopyDexunCloudEyeList()
	time.AfterFunc(24*time.Hour, controller.CopyDexunCloudEyeList)
	controller.GetPackageLists()
	time.AfterFunc(24*time.Hour, controller.GetPackageLists)
	controller.GetDexunAtkLogs()
	time.AfterFunc(5*time.Minute, controller.GetDexunAtkLogs)
	controller.GetDexunAccLogs()
	time.AfterFunc(5*time.Minute, controller.GetDexunAccLogs)
	controller.GetDexunFlowStats()
	time.AfterFunc(5*time.Minute, controller.GetDexunFlowStats)
	controller.GetDexunQueryStats()
	time.AfterFunc(5*time.Minute, controller.GetDexunQueryStats)
	controller.GetDexunInterceptStats()
	time.AfterFunc(1*time.Hour, controller.GetDexunInterceptStats)
	controller.GetDexunAtkStats()
	time.AfterFunc(1*time.Hour, controller.GetDexunAtkStats)
	controller.GetAreaStats()
	time.AfterFunc(1*time.Hour, controller.GetAreaStats)
	controller.GetAreaRankStats()
	time.AfterFunc(1*time.Hour, controller.GetAreaRankStats)
	controller.GetAtkInterStats()
	time.AfterFunc(1*time.Hour, controller.GetAtkInterStats)
	controller.GetHttpPackStat()
	time.AfterFunc(1*time.Hour, controller.GetHttpPackStat)
	controller.GetLineChartStats()
	time.AfterFunc(1*time.Hour, controller.GetLineChartStats)
	controller.GetBWStats()
	time.AfterFunc(1*time.Hour, controller.GetBWStats)
	controller.GetTotalFlow()
	time.AfterFunc(1*time.Hour, controller.GetTotalFlow)
	controller.GetAccCDNIPRank()
	time.AfterFunc(1*time.Hour, controller.GetAccCDNIPRank)

	v1 := app.Group("v1")
	{
		v1.POST("/register", controller.Register)
		v1.POST("/login", controller.Login)
		v1.GET("/invite", controller.InviteCode)
		v1.GET("/getcode", controller.GetCode)
		v1.GET("/logout", controller.Logout)
		v1.POST("/role", controller.GetAccountRole)
		v1.POST("/certtype", controller.GetAccountCertType)

		v1.GET("/account", controller.Info)

		v1.POST("/accounts", controller.AccountList)
		v1.POST("/agents", controller.AgentList)

		v1.POST("/cert/person", controller.PersonalCert)
		v1.POST("/cert/corp", controller.CorpCert)
		v1.POST("/cert/card", controller.CardCert)
		v1.POST("/cert/person/status", controller.GetPersonStatus)
		v1.POST("/cert/corp/status", controller.GetCorpStatus)

		v1.POST("/certs/person", middlewares.AdminRequired, controller.PersonalCertList)
		v1.POST("/certs/corp", middlewares.AdminRequired, controller.CorpCertList)
		v1.POST("/cert/corp/:id/status", middlewares.AdminRequired, controller.UpdateCorpCertStatus)
		v1.POST("/cert/person/:id/status", middlewares.AdminRequired, controller.UpdatePersonCertStatus)

		v1.DELETE("/user/:id", controller.AccountDelete)

		//用户充值
		v1.POST("/user/recharge", controller.Recharge)

		//查询订单，user查询自己的订单，agent查询名下用户的订单，admin查所有
		v1.POST("/scdn/orders", controller.GetSCDNLists)
		//用户下单Scdn套餐
		v1.POST("/scdn/purchase", controller.AddScdnOrder)

		//同步Scdn表单
		v1.POST("/scdn/allow", middlewares.AdminRequired, controller.AllowDexunScdnSell)
		//更新Scdn状态
		v1.POST("/scdn/:scdn/sellstatus", middlewares.AdminRequired, controller.UpdateScdnCombosStatus)
		//获取德迅Scdn表单
		v1.POST("/scdn/dx", middlewares.AdminRequired, controller.DexunScdnLists)
		//普通用户/管理员/代理商获取Scdn表单
		v1.POST("/scdn", controller.ScdnLists)

		//同步Scdn包套餐表单
		v1.POST("/scdn/packages/allow", middlewares.AdminRequired, controller.AllowDexunSPSell)
		//更新包套餐状态
		v1.POST("/scdn/packages/:packages/sellstatus", middlewares.AdminRequired, controller.UpdateDexunSPStatus)
		//获取德迅包套餐表单
		v1.POST("/scdn/packages/dx", middlewares.AdminRequired, controller.DexunPackageLists)
		//获取包套餐表单
		v1.POST("/scdn/packages", controller.PackageLists)

		//管理员查看所有用户账单
		v1.POST("/bills", middlewares.AdminRequired, controller.AllBillLists)
		//用户查看自己的账单
		v1.POST("/user/bills", controller.BillLists)
		//用户查看自己的订单
		v1.POST("/user/orders", controller.OrdersLists)

		//TODO
		//用户查看自己的所有域名
		v1.POST("/scdn/users/domains", controller.GetSCDNUserDomains)
		//用户更新域名状态，假删除
		v1.POST("/scdn/domain/:domain", controller.UpdateDomainsStatus)
		//用户添加域名
		v1.POST("/scdn/domain", controller.AddDomain)
		//管理员查看所有的域名
		v1.POST("/scdn/domains", middlewares.AdminRequired, controller.DomainsLists)
		//用户清理域名缓存
		v1.POST("/scdn/domains/:domain/clean", controller.SweepCache)
		//获取域名SNI信息
		v1.POST("/scdn/domain/:domain/sni", controller.GetDomainSni)

		//用户增加域名配置，包括ConfigList和SourceAddress
		v1.POST("/scdn/domain/:domain/domainconfig", controller.AddDomainConfig)
		//用户新增ConfigList
		v1.POST("/scdn/domain/:domain/config", controller.AddConfigLists)
		//用户修改ConfigList内容和状态
		v1.POST("/scdn/domain/:domain/config/:config", controller.UpdateConfigLists)
		//用户获取自己的ConfigList
		v1.POST("/scdn/domain/:domain/configs", controller.ConfigLists)
		//用户新增SourceAddres
		v1.POST("/scdn/domain/:domain/config/:config/sa", controller.AddSourceAddress)
		//用户修改SourceAddress内容和状态
		v1.POST("/scdn/domain/:domain/config/:config/sa/:sa", controller.UpdateSourceAddress)
		//用户获取自己域名下ConfigList的所有SourceAddress
		v1.POST("/scdn/domain/:domain/config/:config/sas", controller.SourceAddresses)

		//代理商注册
		v1.POST("/agent", controller.AgentRegister)
		//代理商登录
		v1.POST("/agent/login", controller.AgentLogin)
		//代理商查看自己的客户
		v1.POST("/agent/users", middlewares.AgentRequired, controller.GetUsers)
		//变更代理商状态，包含通过申请、踢出等
		v1.POST("/agent/:agent/status", middlewares.AdminRequired, controller.UpdateAgentStatus)

		//增加申请码
		v1.POST("/invite/add", middlewares.AdminRequired, controller.AddInviteCode)
		//手动更新申请码状态，1已使用、0未使用
		v1.PATCH("/invite/:invite/status", middlewares.AdminRequired, controller.UpdateInviteCodeStatus)

		//获取SCDN流量总计的大屏数据
		v1.GET("/scdn/analysis/totalflow", controller.GetTotalFlowStats)
		//获取SCDN攻击记录统计的大屏数据
		v1.GET("/scdn/analysis/attackcount", controller.GetAttackStats)
		//获取SCDN流量统计折线图的大屏数据
		v1.GET("/scdn/analysis/linechart", controller.GetLineChart)
		//获取SCDN拦截记录的大屏数据
		v1.GET("/scdn/analysis/intercept", controller.GetInterceptStats)
		//临时数据
		v1.GET("/mock/analysis/total")

		//代理商获取用户总数的大屏数据
		v1.GET("/analysis/totalaccount", controller.GetTotalAccounts)
		//代理商获取产品销售占比的大屏数据
		v1.GET("/analysis/sellstats", controller.GetSellStats)
		//代理商获取现存订单数的大屏数据
		v1.GET("/analysis/existorders", controller.GetExistOrders)

		//缓存规则模版的增删查改
		v1.POST("/cache/model/add", controller.CacheModelAdd)
		v1.PATCH("/cache/model/:model", controller.CacheModelUpdate)
		v1.POST("/cache/models", controller.CacheModelList)

		//缓存规则的增删查改
		v1.POST("/cache/:domain", controller.CacheAdd)
		v1.PATCH("/cache/:cache", controller.UpdateCache)
		v1.POST("/caches", controller.CacheLists)

		//缓存预热开启关闭或路由编辑
		v1.POST("/domain/:domain/heat/switch", controller.SwitchPreHeat)
		v1.POST("/domain/:domain/heat/update", controller.UpdatePreHeat)
		//获取缓存预热相关情况
		v1.POST("/domain/:domain/heats", controller.GetPreHeatLists)

		//证书增删查改
		v1.POST("/domain/:domain/cert", controller.AddCertification)
		v1.PATCH("/domain/cert/:cert", controller.UpdateCertification)
		v1.POST("/domain/certs", controller.CertificationLists)

		//黑白名单(实例)增删查改
		v1.POST("/bwins", controller.AddBWInstance)
		v1.PATCH("/bwins/:bwins", controller.UpdateBWInstance)
		v1.POST("/bwinses", controller.BWInstanceLists)

		//黑白名单(单域名)增删查改
		v1.POST("/bwsin", controller.AddBWSingle)
		v1.PATCH("/bwsin/:bwsin", controller.UpdateBWSingle)
		v1.POST("/bwsines", controller.BWSingleLists)

		//URL黑白名单增删查改
		v1.POST("/domain/:domain/urlbw", controller.AddUrlBW)
		v1.PATCH("/urlbw/:urlbw", controller.UpdateUrlBW)
		v1.POST("urlbws", controller.UrlBWLists)

		//CC增删查改
		v1.POST("/domain/:domain/cc", controller.AddCC)
		v1.PATCH("/cc/:cc", controller.UpdateCC)
		v1.POST("/ccs", controller.CCLists)

		//防盗链增删查改
		v1.POST("/domain/:domain/leechlink", controller.AddLeechLink)
		v1.PATCH("/leechlink/:leechlink", controller.UpdateLeechLink)
		v1.POST("/leechlinks", controller.LeechLinkLists)

		//精准访问控制增删查改
		v1.POST("/domain/:domain/preacc", controller.AddPreAcc)
		v1.PATCH("/preacc/:preacc", controller.UpdatePreAcc)
		v1.POST("/preaccs", controller.PreAccLists)

		//区域访问控制增删查改
		v1.POST("/domain/:domain/areaacc", controller.AddAreaAcc)
		v1.PATCH("/areaacc/:areaacc", controller.UpdateAreaAcc)
		v1.POST("/areaaccs", controller.AreaAccLists)

		//文字内容风控增删查改
		v1.POST("/domain/:domain/wordrc", controller.AddWordRC)
		v1.PATCH("/wordrc/:wordrc", controller.UpdateWordRC)
		v1.POST("/wordrcs", controller.WordRCLists)

		//图片内容风控增删查改
		v1.POST("/domain/:domain/picrc", controller.AddPicRC)
		v1.PATCH("/picrc/:picrc", controller.UpdatePicRC)
		v1.POST("/picrcs", controller.PicRCLists)

		//安全访问控制增删查改
		v1.POST("/domain/:domain/safeacc", controller.AddSafeAcc)
		v1.PATCH("/safeacc/:safeacc", controller.UpdateSafeAcc)
		v1.POST("/safeaccs", controller.SafeAccLists)

		//WAF防火墙设置
		v1.POST("/domain/:domain/waf", controller.UpdateWaf)

		//查询DDoS订单，user查询自己的订单，agent查询名下用户的订单，admin查所有
		v1.POST("/ddos/orders", controller.GetDDoSLists)
		//用户下单DDoS套餐
		v1.POST("/ddos/purchase", controller.AddDDoSOrders)
		//普通用户/管理员/代理商获取DDoS表单
		v1.POST("/ddos", controller.DDoSLists)
		//更新DDoS状态
		v1.POST("/ddos/:ddos/sellstatus", middlewares.AdminRequired, controller.UpdateDDoSCombosStatus)
		//用户登录控制面板
		v1.POST("/ddos/panel/login", controller.DDoSLoginPanel)
		//用户添加DDoS域名
		v1.POST("/ddos/domain", controller.DDoSAddDomain)
		//TODO
		//用户查看自己DDoS的域名
		v1.POST("/ddos/users/domains", controller.GetDDoSDomains)

		//获取SSL证书
		v1.POST("/ssl", controller.SSLList)
		//生成CSR
		v1.POST("/ssl/csr", controller.GetSSLCSR)
		//SSL证书验证方式
		v1.POST("/ssl/confirm", controller.UpdateSSLDoaminConfirm)
		//用户下单SSL
		v1.POST("/ssl/purchase", controller.CreateSSLOrders)
		//查询订单，user查询自己的订单，agent查询名下用户的订单，admin查所有
		v1.POST("/ssl/orders", controller.GetSSLLists)
		//用户查看SSL证书详情
		v1.POST("/ssl/detail", controller.GetSSLDetail)

		//普通用户/管理员/代理商获取Scdn表单
		v1.POST("/cloudeye", controller.CloudEyeLists)
		//用户下单云眼
		v1.POST("/cloudeye/purchase", controller.AddCloudEyeOrders)

		//v1.POST("/test", controller.Test)
	}
	app.Run(fmt.Sprintf(":%d", *httpPort))
}
