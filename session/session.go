package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	provides map[string]Provider
	// GlobalSessions 全局 Session 管理器
	GlobalSessions *Manager
)

// Manager 一个 session 管理器
type Manager struct {
	sync.Mutex
	cookieName  string
	provider    Provider
	maxLifetime int64
}

func (m *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionStart 开始一个 Session
func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	m.Lock()
	defer m.Unlock()
	// 看下请求是否带特定的 cookie ，没有就新建一个
	cookie, err := r.Cookie(m.cookieName)
	// 用户没有禁用 Cookie 而且 cookie 为空
	if err != nil || cookie.Value == "" {
		sid := m.sessionID()                     // 创建唯一 ID
		session, _ = m.provider.SessionInit(sid) // 初始化一个 Session
		cookie := http.Cookie{
			Name:     m.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(m.maxLifetime),
		}
		http.SetCookie(w, &cookie) // 将创建的 cookie 回写回客户端
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = m.provider.SessionRead(sid)
	}
	return // 返回 session
}

// SessionDestroy 清除 Session
func (m *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	m.Lock()
	defer m.Unlock()
	// 先干掉 Session
	m.provider.SessionDestroy(cookie.Value)
	expiration := time.Now()
	// 将 cookie 设置为过期
	newCookie := http.Cookie{Name: m.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
	http.SetCookie(w, &newCookie)
}

// GC 定时回收过期 session
func (m *Manager) GC() {
	m.Lock()
	defer m.Unlock()
	m.provider.SessionGC(m.maxLifetime)
	time.AfterFunc(time.Duration(m.maxLifetime), func() { m.GC() })
}

// NewManager 创建一个Session 管理器
func NewManager(provideName, cookieName string, maxLifetime int64) (*Manager, error) {
	provider := provides[provideName]
	if provider == nil {
		return nil, fmt.Errorf("session: Unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifetime: maxLifetime}, nil
}

// Provider 将对 Session 的存取抽象为一个接口，因为 Session 可以存储在任何地方，内存或数据库中
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifetime int64)
}

// Session session 对象
type Session interface {
	Set(key, value interface{}) error // 设置 session 的值，因其可以设置为任意类型数据，所以用空接口
	Get(key interface{}) interface{}  // 获取 session 值
	Delete(key interface{}) error
	SessionID() string
}
