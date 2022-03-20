package pay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/limitcool/blog/global"
	"log"
)

func InitAlipay() {

	// 初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient(global.PaySetting.AlipayAppId, global.PaySetting.AlipayPrivateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn
	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	//client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
	//						SetCharset(alipay.UTF8).             // 设置字符编码，不设置默认 utf-8
	//						SetSignType(alipay.RSA2).            // 设置签名类型，不设置默认 RSA2
	//						SetReturnUrl("https://www.fmm.ink"). // 设置返回URL
	//						SetNotifyUrl("https://www.fmm.ink"). // 设置异步通知URL
	//						SetAppAuthToken()                    // 设置第三方应用授权

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", "条码支付").
		Set("scene", "bar_code").
		Set("auth_code", "2088622958163411").
		Set("out_trade_no", "ddfgdgs").
		Set("total_amount", "0.01").
		Set("timeout_express", "2m")

	aliRsp, err := client.TradePay(context.Background(), bm)
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	ok, err := alipay.VerifySyncSign(global.PaySetting.AlipayPublicKey, aliRsp.SignData, aliRsp.Sign)
	if ok {
		log.Println("支付成功")
	} else {
		log.Println("支付失败:", err)
	}
}
