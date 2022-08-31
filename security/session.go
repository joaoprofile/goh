package security

type Sessions struct {
	userID   string
	tenantID string
}

var _instance *Sessions

func Session() *Sessions {
	if _instance == nil {
		_instance = &Sessions{}
	}
	return _instance
}

func (s *Sessions) User(userID string) *Sessions {
	s.userID = userID
	return s
}

func (s *Sessions) Tenant(userID string) *Sessions {
	s.userID = userID
	return s
}

func (s *Sessions) GetUser() string {
	return s.userID
}

func (s *Sessions) GetTenant() string {
	return s.tenantID
}
