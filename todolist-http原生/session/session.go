package session

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type Session struct {
	contents map[string]interface{}
}

func NewSession() *Session {
	return &Session{make(map[string]interface{})}
}

func (s *Session) Set(key string, value interface{}) {
	s.contents[key] = value
}

func (s *Session) Get(key string) (interface{}, bool) {
	value, ok := s.contents[key]
	return value, ok
}

func (s *Session) Delete(key string) {
	delete(s.contents, key)
}

type Manager struct {
	sessions   map[string]*Session
	cookieName string //sid
	timeout    int
}

func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) *Session {
	var session *Session
	var sid string
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		sid := uuid.NewV4().String()
		session = NewSession()
		m.sessions[sid] = session
	} else if _, ok := m.sessions[cookie.Value]; !ok {
		sid := uuid.NewV4().String()
		session = NewSession()
		m.sessions[sid] = session
	} else {
		// 正常情况
		return m.sessions[cookie.Value]
	}

	cookie = &http.Cookie{
		Name:     m.cookieName,
		Path:     "/",
		Value:    sid,
		MaxAge:   m.timeout,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return session
}

var DefaultManager *Manager

func init() {
	DefaultManager = &Manager{
		sessions:   make(map[string]*Session),
		cookieName: "sid",
		timeout:    3600,
	}
}
