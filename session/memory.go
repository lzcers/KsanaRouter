package session

import (
	"sync"
	"time"
)

// MemoryStore 存储 session 的结构
type MemoryStore struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

// Set 设置 Session 值
func (st *MemoryStore) Set(key, value interface{}) error {
	st.value[key] = value
	return nil
}

// Get 获取 Session 值
func (st *MemoryStore) Get(key interface{}) interface{} {
	return st.value[key]
}

// Delete 删除 Session 值
func (st *MemoryStore) Delete(key interface{}) error {
	delete(st.value, key)
	return nil
}

// SessionID 获取当前 SessionID
func (st *MemoryStore) SessionID() string {
	return st.sid
}

// MemoryProvider 内存 Session 存储
type MemoryProvider struct {
	sync.Mutex
	sessions map[string]*MemoryStore
}

// SessionInit 初始化一个 Session
func (mp *MemoryProvider) SessionInit(sid string) (Session, error) {
	mp.Lock()
	defer mp.Unlock()
	if mp.sessions == nil {
		mp.sessions = make(map[string]*MemoryStore)
	}
	v := make(map[interface{}]interface{})
	sess := &MemoryStore{sid: sid, timeAccessed: time.Now(), value: v}
	mp.sessions[sid] = sess
	return sess, nil
}

// SessionRead 读取一个 Session
func (mp *MemoryProvider) SessionRead(sid string) (Session, error) {
	if st, ok := mp.sessions[sid]; ok {
		return st, nil
	}
	// 在内存中取不到 Session 就重新初始化一个
	// 比如服务器重启了，但是客户端 Cookie 中还有 sid 这时就会出现取不到的情况
	sess, _ := mp.SessionInit(sid)
	return sess, nil
}

// SessionDestroy 干掉一个 Session
func (mp *MemoryProvider) SessionDestroy(sid string) error {
	delete(mp.sessions, sid)
	return nil
}

// SessionGC 定时清理失效 Session
func (mp *MemoryProvider) SessionGC(maxLifetime int64) {
	mp.Lock()
	defer mp.Unlock()
	for key, element := range mp.sessions {
		if (element.timeAccessed.Unix() + maxLifetime) < time.Now().Unix() {
			delete(mp.sessions, key)
		}
	}
}

func init() {
	provides = make(map[string]Provider)
	provides["memory"] = new(MemoryProvider)
	sess, err := NewManager("memory", "gosessionid", 3600)
	GlobalSessions = sess
	if err != nil {
		panic(err)
	}
	go GlobalSessions.GC()
}
