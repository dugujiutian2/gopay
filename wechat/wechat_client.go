package wechat

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/iGoogle-ink/gopay"
	"io/ioutil"
	"strings"
)

type WeChatClient struct {
	AppId   string
	MchId   string
	ApiKey  string
	baseURL string
	IsProd  bool
}

//提交付款码支付 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1
func (this *WeChatClient) Micropay(body gopay.BodyMap) (wxRsp *WeChatMicropayResponse, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		bytes, err = this.doWeChat(body, wx_Micropay)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wx_SanBox_Micropay)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatMicropayResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return wxRsp, nil
}

//统一下单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
func (this *WeChatClient) UnifiedOrder(body gopay.BodyMap) (wxRsp *WeChatUnifiedOrderResponse, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		bytes, err = this.doWeChat(body, wx_UnifiedOrder)
		if err != nil {
			return nil, err
		}
	} else {
		body.Set("total_fee", 101)
		bytes, err = this.doWeChat(body, wx_SanBox_UnifiedOrder)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatUnifiedOrderResponse)
	//fmt.Println("bytes:", string(bytes))
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		//fmt.Println("xml.Unmarshal.Err:", err)
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return wxRsp, nil
}

//查询订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
func (this *WeChatClient) QueryOrder(body gopay.BodyMap) (wxRsp *WeChatQueryOrderResponse, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		bytes, err = this.doWeChat(body, wx_OrderQuery)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wx_SanBox_OrderQuery)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatQueryOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return wxRsp, nil
}

//关闭订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
func (this *WeChatClient) CloseOrder(body gopay.BodyMap) (wxRsp *WeChatCloseOrderResponse, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		bytes, err = this.doWeChat(body, wx_CloseOrder)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wx_SanBox_CloseOrder)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatCloseOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return wxRsp, nil
}

//撤销订单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (this *WeChatClient) Reverse(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatReverseResponse, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		pkcsPool := x509.NewCertPool()
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
		if err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
		}
		tlsConfig := new(tls.Config)
		tlsConfig.Certificates = []tls.Certificate{certificate}
		tlsConfig.RootCAs = pkcsPool
		tlsConfig.InsecureSkipVerify = true

		bytes, err = this.doWeChat(body, wx_Reverse, tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wx_SanBox_Reverse)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatReverseResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return wxRsp, nil
}

//申请退款 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (this *WeChatClient) Refund(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatRefundResponse, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		pkcsPool := x509.NewCertPool()
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
		if err != nil {
			return nil, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
		}
		tlsConfig := new(tls.Config)
		tlsConfig.Certificates = []tls.Certificate{certificate}
		tlsConfig.RootCAs = pkcsPool
		tlsConfig.InsecureSkipVerify = true

		bytes, err = this.doWeChat(body, wx_Refund, tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wx_SanBox_Refund)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatRefundResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return wxRsp, nil
}

//查询退款 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
func (this *WeChatClient) QueryRefund(body gopay.BodyMap) (wxRsp *WeChatQueryRefundResponse, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		bytes, err = this.doWeChat(body, wx_RefundQuery)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wx_SanBox_RefundQuery)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatQueryRefundResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return wxRsp, nil
}

//下载对账单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
func (this *WeChatClient) DownloadBill(body gopay.BodyMap) (wxRsp string, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		bytes, err = this.doWeChat(body, wx_DownloadBill)
	} else {
		bytes, err = this.doWeChat(body, wx_SanBox_DownloadBill)
	}
	wxRsp = string(bytes)
	if err != nil {
		return wxRsp, err
	}
	return wxRsp, nil
}

//下载资金账单 ok
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7
//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (this *WeChatClient) DownloadFundFlow(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		pkcsPool := x509.NewCertPool()
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return null, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
		if err != nil {
			return null, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
		}
		tlsConfig := new(tls.Config)
		tlsConfig.Certificates = []tls.Certificate{certificate}
		tlsConfig.RootCAs = pkcsPool
		tlsConfig.InsecureSkipVerify = true

		bytes, err = this.doWeChat(body, wx_DownloadFundFlow, tlsConfig)
	} else {
		bytes, err = this.doWeChat(body, wx_SanBox_DownloadFundFlow)
	}

	if err != nil {
		return null, err
	}
	wxRsp = string(bytes)
	return wxRsp, nil
}

//拉取订单评价数据
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11
//    好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (this *WeChatClient) BatchQueryComment(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var bytes []byte
	if this.IsProd {
		//正式环境
		body.Set("sign_type", SignType_HMAC_SHA256)

		pkcsPool := x509.NewCertPool()
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return null, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
		if err != nil {
			return null, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
		}
		tlsConfig := new(tls.Config)
		tlsConfig.Certificates = []tls.Certificate{certificate}
		tlsConfig.RootCAs = pkcsPool
		tlsConfig.InsecureSkipVerify = true

		bytes, err = this.doWeChat(body, wx_BatchQueryComment, tlsConfig)
	} else {
		bytes, err = this.doWeChat(body, wx_SanBox_BatchQueryComment)
	}

	if err != nil {
		return null, err
	}

	wxRsp = string(bytes)
	return wxRsp, nil
}

//企业向微信用户个人付款
//    文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_1
//    注意：此方法未支持沙箱环境，默认正式环境，转账请慎重
func (this *WeChatClient) Transfer(body gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatTransfersResponse, err error) {
	var bytes []byte
	var sign string
	body.Set("mch_appid", this.AppId)
	body.Set("mchid", this.MchId)

	agent := gopay.HttpAgent()

	//正式环境
	pkcsPool := x509.NewCertPool()
	pkcs, err := ioutil.ReadFile(pkcs12FilePath)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile：%v", err.Error())
	}
	pkcsPool.AppendCertsFromPEM(pkcs)
	certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return nil, fmt.Errorf("tls.LoadX509KeyPair：%v", err.Error())
	}
	tlsConfig := new(tls.Config)
	tlsConfig.Certificates = []tls.Certificate{certificate}
	tlsConfig.RootCAs = pkcsPool
	tlsConfig.InsecureSkipVerify = true

	agent.TLSClientConfig(tlsConfig)

	//本地计算Sign
	sign = getWeChatReleaseSign(this.ApiKey, SignType_MD5, body)

	body.Set("sign", sign)
	reqXML := generateXml(body)

	if this.baseURL != null {
		agent.Post(this.baseURL + wx_Transfers)
	} else {
		agent.Post(wx_base_url_ch + wx_Transfers)
	}
	agent.Type("xml")
	agent.SendString(reqXML)

	_, bytes, errs := agent.EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	wxRsp = new(WeChatTransfersResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, fmt.Errorf("xml.Unmarshal：%v", err.Error())
	}
	return wxRsp, nil
}

//向微信发送请求 ok
func (this *WeChatClient) doWeChat(body gopay.BodyMap, path string, tlsConfig ...*tls.Config) (bytes []byte, err error) {
	var sign string
	body.Set("appid", this.AppId)
	body.Set("mch_id", this.MchId)
	//===============生成参数===================
	if body.Get("sign") != null {
		goto GoRequest
	}

	//计算Sign值
	if !this.IsProd {
		//沙箱环境
		body.Set("sign_type", SignType_MD5)
		//沙箱环境Sign值
		sign, err = getWeChatSignBoxSign(this.MchId, this.ApiKey, body)
		if err != nil {
			//fmt.Println("getWeChatSignBoxSign:", err)
			return nil, err
		}
	} else {
		//正式环境
		sign = getWeChatReleaseSign(this.ApiKey, body.Get("sign_type"), body)
	}
	body.Set("sign", sign)

GoRequest:
	reqXML := generateXml(body)
	//fmt.Println("reqXML:",reqXML)
	//===============发起请求===================
	agent := gopay.HttpAgent()

	if this.IsProd && tlsConfig != nil {
		agent.TLSClientConfig(tlsConfig[0])
	}

	if this.baseURL != null {
		agent.Post(this.baseURL + path)
	} else {
		agent.Post(wx_base_url_ch + path)
	}
	agent.Type("xml")
	agent.SendString(reqXML)
	res, bytes, errs := agent.EndBytes()
	if len(errs) > 0 {
		//fmt.Println("errs[0]:", errs[0])
		return nil, errs[0]
	}
	//fmt.Println("res:", res.StatusCode)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %v", res.StatusCode)
	}
	if strings.Contains(string(bytes), "HTML") {
		return nil, errors.New(string(bytes))
	}
	//fmt.Println("bytes:", string(bytes))
	return bytes, nil
}
