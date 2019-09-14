package gopay

import (
	"github.com/iGoogle-ink/gopay/alipay"
	"github.com/iGoogle-ink/gopay/union"
	"github.com/iGoogle-ink/gopay/wechat"
)

//初始化微信客户端 ok
//    appId：应用ID
//    mchId：商户ID
//    apiKey：API秘钥值
//    isProd：是否是正式环境
func NewWeChatClient(appId, mchId, apiKey string, isProd bool) (client *wechat.WeChatClient) {
	client = new(wechat.WeChatClient)
	client.AppId = appId
	client.MchId = mchId
	client.ApiKey = apiKey
	client.IsProd = isProd
	return client
}

//初始化支付宝客户端
//    注意：如果使用支付宝公钥证书验签，请设置 支付宝根证书SN（client.SetAlipayRootCertSN()）、应用公钥证书SN（client.SetAppCertSN()）
//    appId：应用ID
//    privateKey：应用私钥
//    isProd：是否是正式环境
func NewAliPayClient(appId, privateKey string, isProd bool) (client *alipay.AliPayClient) {
	client = new(alipay.AliPayClient)
	client.AppId = appId
	client.PrivateKey = privateKey
	client.IsProd = isProd
	return client
}

//初始化银联支付客户端
//    merId：商户ID
//    isProd：是否是正式环境
func NewUnionPayClient(merId string, isProd bool) (client *union.UnionPayClient) {
	client = new(union.UnionPayClient)
	client.MerId = merId
	client.IsProd = isProd
	return client
}
