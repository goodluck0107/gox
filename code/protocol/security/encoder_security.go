package security

// import (
// 	"gitee.com/andyxt/gox/code/protocol/protocolCoder"
// 	"gitee.com/andyxt/gox/code/protocol/protocolCoderImpl"
// 	"gitee.com/andyxt/gox/code/protocol/protocolDefine"

// 	"gitee.com/andyxt/gona/boot/channel"
// 	"gitee.com/andyxt/gona/logger"
// )

// const (
// 	ChannelSecurityType string = "ChannelSecurityType" //string
// )

// // *ProtocolBuffer-->*ProtocolBuffer
// type SecurityEncoderHandler struct {
// 	securitierMap map[int8]protocolCoder.Securitier
// }

// func NewSecurityEncoderHandler() (this *SecurityEncoderHandler) {
// 	this = new(SecurityEncoderHandler)
// 	this.securitierMap = this.createSerializierMap()
// 	return
// }

// func (encoder *SecurityEncoderHandler) createSerializierMap() map[int8]protocolCoder.Securitier {
// 	mSecurityMap := make(map[int8]protocolCoder.Securitier)
// 	mSecurityMap[protocolDefine.CommonSecurityType] = protocolCoderImpl.NewDefualtSecuritier()
// 	return mSecurityMap
// }

// func (encoder *SecurityEncoderHandler) Write(ctx channel.ChannelContext, e interface{}) (ret interface{}) {
// 	//logger.Debug("SecurityEncoder Write-0")
// 	buf := e.([]byte)
// 	encryptType := ctx.ContextAttr().GetInt8(ChannelSecurityType)
// 	if encryptType <= 0 { //不加密
// 		ret = buf
// 		//logger.Debug("SecurityEncoder Write1")
// 	} else {
// 		//logger.Debug("SecurityEncoder Write2")
// 		securitier := encoder.securitierMap[encryptType]
// 		if securitier == nil {
// 			logger.Error("协议安全类型无效:encryptType=", encryptType)
// 			ret = buf
// 			//logger.Debug("SecurityEncoder Write3")
// 		} else {
// 			securityBuf := securitier.Encrypt(buf)
// 			securityTypeBuf := []byte{byte(encryptType)}
// 			buf = append(securityTypeBuf, securityBuf...)
// 			ret = buf
// 			//logger.Debug("SecurityEncoder Write4")
// 		}

// 	}
// 	return
// }

// func (encoder *SecurityEncoderHandler) Close(ctx channel.ChannelContext) {
// 	//	logger.Debug("SecurityEncoder Close")
// }
// func (encoder *SecurityEncoderHandler) ExceptionCaught(ctx channel.ChannelContext, err error) {
// 	//	logger.Debug("SecurityEncoder ExceptionCaught")
// }
