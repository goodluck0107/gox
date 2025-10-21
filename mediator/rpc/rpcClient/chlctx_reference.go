package rpcClient

import (
	"sync"

	"github.com/goodluck0107/gox/service"
)

var chlCtxReference *Reference = new(Reference)

type Reference struct {
	mChlCtxReference sync.Map // 玩家对链接的引用
}

// StoreTableReference 添加链接引用
func (r *Reference) Store(playerID int64, reference service.IChannelContext) {
	r.mChlCtxReference.Store(playerID, reference)
}

// DeleteTableReference 删除链接引用
func (r *Reference) Delete(playerID int64) {
	r.mChlCtxReference.Delete(playerID)
}

// SearchTableReference 查询链接引用
func (r *Reference) Search(playerID int64) service.IChannelContext {
	v, ok := r.mChlCtxReference.Load(playerID)
	if ok {
		return v.(service.IChannelContext)
	}
	return nil
}
