package alidayu

// 多方通话请求参数
// http://open.taobao.com/docs/api.htm?apiId=25443
type VoiceNumDoubleCallRequest struct {
	SessionTimeOut string `url:"session_time_out"`
	Extend         string `url:"extend"`
	CallerNum      string `url:"caller_num"`
	CallerShowNum  string `url:"caller_show_num"`
	CalledNum      string `url:"called_num"`
	CalledShowNum  string `url:"called_show_num"`
}

func (r *VoiceNumDoubleCallRequest) Method() string {
	return "alibaba.aliqin.fc.voice.num.doublecall"
}

// 文本转语音通知请求参数
// http://open.taobao.com/docs/api.htm?apiId=25444
type TTSNumSingleCallRequest struct {
	Extend        string `url:"extend"`
	TTSParam      string `url:"tts_param"`
	CalledNum     string `url:"called_num"`
	CalledShowNum string `url:"called_show_num"`
	TTSCode       string `url:"tts_code"`
}

func (r *TTSNumSingleCallRequest) Method() string {
	return "alibaba.aliqin.fc.tts.num.singlecall"
}

// 语音通知请求参数
// http://open.taobao.com/docs/api.htm?apiId=25445
type VoiceNumSingleCallRequest struct {
	Extend        string `url:"extend"`
	CalledNum     string `url:"called_num"`
	CalledShowNum string `url:"called_show_num"`
	VoiceCode     string `url:"voice_code"`
}

func (r *VoiceNumSingleCallRequest) Method() string {
	return "alibaba.aliqin.fc.voice.num.singlecall"
}
