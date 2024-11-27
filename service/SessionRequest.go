package service

func NewSessionRequest(channelContext IChannelContext, reqContext IReqContext) IServiceRequest {
	req := new(Request)
	req.mChannelContext = channelContext
	req.mReqContext = reqContext
	return req
}

type Request struct {
	mChannelContext IChannelContext
	mReqContext     IReqContext
}

func (req *Request) ChannelContext() IChannelContext {
	return req.mChannelContext
}
func (req *Request) ReqContext() IReqContext {
	return req.mReqContext
}
