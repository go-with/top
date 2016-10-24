package alidayu

// 短信发送请求参数
// http://open.taobao.com/docs/api.htm?apiId=25450
type SMSNumSendRequest struct {
	Extend          string `url:"extend"`
	SMSType         string `url:"sms_type"`
	SMSFreeSignName string `url:"sms_free_sign_name"`
	SMSParam        string `url:"sms_param"`
	RecNum          string `url:"rec_num"`
	SMSTemplateCode string `url:"sms_template_code"`
}

func (r *SMSNumSendRequest) Method() string {
	return "alibaba.aliqin.fc.sms.num.send"
}
