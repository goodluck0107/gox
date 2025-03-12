package security

// import (
// 	"fmt"

// 	"gitee.com/andyxt/gox/code/protocol/protocolCoder"
// 	"gitee.com/andyxt/gox/code/protocol/protocolCoderImpl"
// 	"gitee.com/andyxt/gox/code/protocol/protocolDefine"

// 	"gitee.com/andyxt/gona/boot/channel"
// 	"gitee.com/andyxt/gona/logger"
// )

// // []byte-->*ProtocolBuffer
// type SecurityDecoderHandler struct {
// 	securitierMap map[int8]protocolCoder.Securitier
// }

// func NewSecurityDecoderHandler() (this *SecurityDecoderHandler) {
// 	this = new(SecurityDecoderHandler)
// 	this.securitierMap = this.createSerializierMap()
// 	return
// }

// func (decoder *SecurityDecoderHandler) createSerializierMap() map[int8]protocolCoder.Securitier {
// 	mSecurityMap := make(map[int8]protocolCoder.Securitier)
// 	mSecurityMap[protocolDefine.CommonSecurityType] = protocolCoderImpl.NewDefualtSecuritier()
// 	return mSecurityMap
// }
// func (decoder *SecurityDecoderHandler) MessageReceived(ctx channel.ChannelContext, e interface{}) (ret interface{}, goonNext bool) {
// 	fmt.Println("SecurityDecoder MessageReceived")
// 	byteSlice := e.([]byte)
// 	fmt.Println("SecurityDecoder BeforeDecrypt:", byteSlice)
// 	encryptType := int8(byteSlice[0]) //加密类型
// 	if encryptType <= 0 {             //不加密
// 		fmt.Println("SecurityDecoder 不加密")
// 		goonNext = true
// 		ret = byteSlice[1:] //去掉加密类型字节
// 	} else {
// 		fmt.Println("SecurityDecoder 加密类型：", encryptType)
// 		securitier := decoder.securitierMap[encryptType]
// 		if securitier == nil {
// 			logger.Info("关闭连接：", " 关闭原因：协议安全类型无效:IP=", ctx.RemoteAddr(), "加密类型：", encryptType)
// 			ctx.Close()
// 			return
// 		}
// 		//logger.Error(ctx.GetInt32(channelConsts.ChannelId)," before Decrypt:",byteSlice)
// 		byteSlice = byteSlice[1:] //去掉加密类型字节
// 		valid, newByteSlice := securitier.Decrypt(byteSlice)
// 		fmt.Println("SecurityDecoder AfterDecrypt:", newByteSlice)
// 		//structType := int8(byteSlice[0])
// 		//ctx.Set(channel.ChannelStructType, structType)//格式类型
// 		if valid {
// 			goonNext = true
// 			ret = newByteSlice
// 			//logger.Error(ctx.GetInt32(channelConsts.ChannelId),"after Decrypt:",byteSlice)
// 		} else {
// 			logger.Info("关闭连接：", " 关闭原因：协议安全验证失败:IP=", ctx.RemoteAddr())
// 			ctx.Close()
// 		}
// 	}
// 	return
// }

// func (decoder *SecurityDecoderHandler) ChannelActive(ctx channel.ChannelContext) (goonNext bool) {
// 	//	logger.Debug("SecurityDecoder ChannelActive")
// 	return true
// }
// func (decoder *SecurityDecoderHandler) ChannelInactive(ctx channel.ChannelContext) (goonNext bool) {
// 	//	logger.Debug("SecurityDecoder ChannelInactive")
// 	return true
// }

// func (decoder *SecurityDecoderHandler) ExceptionCaught(ctx channel.ChannelContext, err error) {
// 	//	logger.Debug("SecurityDecoder ExceptionCaught")
// }
