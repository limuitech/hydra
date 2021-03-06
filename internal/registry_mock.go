// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/hydra/driver (interfaces: Registry)

// Package internal is a generated GoMock package.
package internal

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sessions "github.com/gorilla/sessions"
	logrus "github.com/sirupsen/logrus"

	fosite "github.com/ory/fosite"
	openid "github.com/ory/fosite/handler/openid"
	herodot "github.com/ory/herodot"
	client "github.com/ory/hydra/client"
	consent "github.com/ory/hydra/consent"
	driver "github.com/ory/hydra/driver"
	configuration "github.com/ory/hydra/driver/configuration"
	jwk "github.com/ory/hydra/jwk"
	prometheus "github.com/ory/hydra/metrics/prometheus"
	oauth2 "github.com/ory/hydra/oauth2"
	x "github.com/ory/hydra/x"
	healthx "github.com/ory/x/healthx"
	tracing "github.com/ory/x/tracing"
)

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// AccessTokenJWTStrategy mocks base method
func (m *MockRegistry) AccessTokenJWTStrategy() jwk.JWTStrategy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenJWTStrategy")
	ret0, _ := ret[0].(jwk.JWTStrategy)
	return ret0
}

// AccessTokenJWTStrategy indicates an expected call of AccessTokenJWTStrategy
func (mr *MockRegistryMockRecorder) AccessTokenJWTStrategy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenJWTStrategy", reflect.TypeOf((*MockRegistry)(nil).AccessTokenJWTStrategy))
}

// AudienceStrategy mocks base method
func (m *MockRegistry) AudienceStrategy() fosite.AudienceMatchingStrategy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AudienceStrategy")
	ret0, _ := ret[0].(fosite.AudienceMatchingStrategy)
	return ret0
}

// AudienceStrategy indicates an expected call of AudienceStrategy
func (mr *MockRegistryMockRecorder) AudienceStrategy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AudienceStrategy", reflect.TypeOf((*MockRegistry)(nil).AudienceStrategy))
}

// BuildDate mocks base method
func (m *MockRegistry) BuildDate() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildDate")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildDate indicates an expected call of BuildDate
func (mr *MockRegistryMockRecorder) BuildDate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildDate", reflect.TypeOf((*MockRegistry)(nil).BuildDate))
}

// BuildHash mocks base method
func (m *MockRegistry) BuildHash() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildHash")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildHash indicates an expected call of BuildHash
func (mr *MockRegistryMockRecorder) BuildHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildHash", reflect.TypeOf((*MockRegistry)(nil).BuildHash))
}

// BuildVersion mocks base method
func (m *MockRegistry) BuildVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildVersion indicates an expected call of BuildVersion
func (mr *MockRegistryMockRecorder) BuildVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildVersion", reflect.TypeOf((*MockRegistry)(nil).BuildVersion))
}

// CanHandle mocks base method
func (m *MockRegistry) CanHandle(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanHandle", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanHandle indicates an expected call of CanHandle
func (mr *MockRegistryMockRecorder) CanHandle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanHandle", reflect.TypeOf((*MockRegistry)(nil).CanHandle), arg0)
}

// ClientHandler mocks base method
func (m *MockRegistry) ClientHandler() *client.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientHandler")
	ret0, _ := ret[0].(*client.Handler)
	return ret0
}

// ClientHandler indicates an expected call of ClientHandler
func (mr *MockRegistryMockRecorder) ClientHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientHandler", reflect.TypeOf((*MockRegistry)(nil).ClientHandler))
}

// ClientHasher mocks base method
func (m *MockRegistry) ClientHasher() fosite.Hasher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientHasher")
	ret0, _ := ret[0].(fosite.Hasher)
	return ret0
}

// ClientHasher indicates an expected call of ClientHasher
func (mr *MockRegistryMockRecorder) ClientHasher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientHasher", reflect.TypeOf((*MockRegistry)(nil).ClientHasher))
}

// ClientManager mocks base method
func (m *MockRegistry) ClientManager() client.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientManager")
	ret0, _ := ret[0].(client.Manager)
	return ret0
}

// ClientManager indicates an expected call of ClientManager
func (mr *MockRegistryMockRecorder) ClientManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientManager", reflect.TypeOf((*MockRegistry)(nil).ClientManager))
}

// ClientValidator mocks base method
func (m *MockRegistry) ClientValidator() *client.Validator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientValidator")
	ret0, _ := ret[0].(*client.Validator)
	return ret0
}

// ClientValidator indicates an expected call of ClientValidator
func (mr *MockRegistryMockRecorder) ClientValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientValidator", reflect.TypeOf((*MockRegistry)(nil).ClientValidator))
}

// ConsentHandler mocks base method
func (m *MockRegistry) ConsentHandler() *consent.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsentHandler")
	ret0, _ := ret[0].(*consent.Handler)
	return ret0
}

// ConsentHandler indicates an expected call of ConsentHandler
func (mr *MockRegistryMockRecorder) ConsentHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentHandler", reflect.TypeOf((*MockRegistry)(nil).ConsentHandler))
}

// ConsentManager mocks base method
func (m *MockRegistry) ConsentManager() consent.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsentManager")
	ret0, _ := ret[0].(consent.Manager)
	return ret0
}

// ConsentManager indicates an expected call of ConsentManager
func (mr *MockRegistryMockRecorder) ConsentManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentManager", reflect.TypeOf((*MockRegistry)(nil).ConsentManager))
}

// ConsentStrategy mocks base method
func (m *MockRegistry) ConsentStrategy() consent.Strategy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsentStrategy")
	ret0, _ := ret[0].(consent.Strategy)
	return ret0
}

// ConsentStrategy indicates an expected call of ConsentStrategy
func (mr *MockRegistryMockRecorder) ConsentStrategy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsentStrategy", reflect.TypeOf((*MockRegistry)(nil).ConsentStrategy))
}

// CookieStore mocks base method
func (m *MockRegistry) CookieStore() sessions.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CookieStore")
	ret0, _ := ret[0].(sessions.Store)
	return ret0
}

// CookieStore indicates an expected call of CookieStore
func (mr *MockRegistryMockRecorder) CookieStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CookieStore", reflect.TypeOf((*MockRegistry)(nil).CookieStore))
}

// HealthHandler mocks base method
func (m *MockRegistry) HealthHandler() *healthx.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthHandler")
	ret0, _ := ret[0].(*healthx.Handler)
	return ret0
}

// HealthHandler indicates an expected call of HealthHandler
func (mr *MockRegistryMockRecorder) HealthHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthHandler", reflect.TypeOf((*MockRegistry)(nil).HealthHandler))
}

// Init mocks base method
func (m *MockRegistry) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRegistryMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRegistry)(nil).Init))
}

// KeyCipher mocks base method
func (m *MockRegistry) KeyCipher() *jwk.AEAD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyCipher")
	ret0, _ := ret[0].(*jwk.AEAD)
	return ret0
}

// KeyCipher indicates an expected call of KeyCipher
func (mr *MockRegistryMockRecorder) KeyCipher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyCipher", reflect.TypeOf((*MockRegistry)(nil).KeyCipher))
}

// KeyGenerators mocks base method
func (m *MockRegistry) KeyGenerators() map[string]jwk.KeyGenerator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyGenerators")
	ret0, _ := ret[0].(map[string]jwk.KeyGenerator)
	return ret0
}

// KeyGenerators indicates an expected call of KeyGenerators
func (mr *MockRegistryMockRecorder) KeyGenerators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyGenerators", reflect.TypeOf((*MockRegistry)(nil).KeyGenerators))
}

// KeyHandler mocks base method
func (m *MockRegistry) KeyHandler() *jwk.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyHandler")
	ret0, _ := ret[0].(*jwk.Handler)
	return ret0
}

// KeyHandler indicates an expected call of KeyHandler
func (mr *MockRegistryMockRecorder) KeyHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyHandler", reflect.TypeOf((*MockRegistry)(nil).KeyHandler))
}

// KeyManager mocks base method
func (m *MockRegistry) KeyManager() jwk.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyManager")
	ret0, _ := ret[0].(jwk.Manager)
	return ret0
}

// KeyManager indicates an expected call of KeyManager
func (mr *MockRegistryMockRecorder) KeyManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyManager", reflect.TypeOf((*MockRegistry)(nil).KeyManager))
}

// Logger mocks base method
func (m *MockRegistry) Logger() logrus.FieldLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(logrus.FieldLogger)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockRegistryMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockRegistry)(nil).Logger))
}

// OAuth2Handler mocks base method
func (m *MockRegistry) OAuth2Handler() *oauth2.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuth2Handler")
	ret0, _ := ret[0].(*oauth2.Handler)
	return ret0
}

// OAuth2Handler indicates an expected call of OAuth2Handler
func (mr *MockRegistryMockRecorder) OAuth2Handler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuth2Handler", reflect.TypeOf((*MockRegistry)(nil).OAuth2Handler))
}

// OAuth2Provider mocks base method
func (m *MockRegistry) OAuth2Provider() fosite.OAuth2Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuth2Provider")
	ret0, _ := ret[0].(fosite.OAuth2Provider)
	return ret0
}

// OAuth2Provider indicates an expected call of OAuth2Provider
func (mr *MockRegistryMockRecorder) OAuth2Provider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuth2Provider", reflect.TypeOf((*MockRegistry)(nil).OAuth2Provider))
}

// OAuth2Storage mocks base method
func (m *MockRegistry) OAuth2Storage() x.FositeStorer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuth2Storage")
	ret0, _ := ret[0].(x.FositeStorer)
	return ret0
}

// OAuth2Storage indicates an expected call of OAuth2Storage
func (mr *MockRegistryMockRecorder) OAuth2Storage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuth2Storage", reflect.TypeOf((*MockRegistry)(nil).OAuth2Storage))
}

// OpenIDConnectRequestValidator mocks base method
func (m *MockRegistry) OpenIDConnectRequestValidator() *openid.OpenIDConnectRequestValidator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenIDConnectRequestValidator")
	ret0, _ := ret[0].(*openid.OpenIDConnectRequestValidator)
	return ret0
}

// OpenIDConnectRequestValidator indicates an expected call of OpenIDConnectRequestValidator
func (mr *MockRegistryMockRecorder) OpenIDConnectRequestValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenIDConnectRequestValidator", reflect.TypeOf((*MockRegistry)(nil).OpenIDConnectRequestValidator))
}

// OpenIDJWTStrategy mocks base method
func (m *MockRegistry) OpenIDJWTStrategy() jwk.JWTStrategy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenIDJWTStrategy")
	ret0, _ := ret[0].(jwk.JWTStrategy)
	return ret0
}

// OpenIDJWTStrategy indicates an expected call of OpenIDJWTStrategy
func (mr *MockRegistryMockRecorder) OpenIDJWTStrategy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenIDJWTStrategy", reflect.TypeOf((*MockRegistry)(nil).OpenIDJWTStrategy))
}

// Ping mocks base method
func (m *MockRegistry) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockRegistryMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRegistry)(nil).Ping))
}

// PrometheusManager mocks base method
func (m *MockRegistry) PrometheusManager() *prometheus.MetricsManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrometheusManager")
	ret0, _ := ret[0].(*prometheus.MetricsManager)
	return ret0
}

// PrometheusManager indicates an expected call of PrometheusManager
func (mr *MockRegistryMockRecorder) PrometheusManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrometheusManager", reflect.TypeOf((*MockRegistry)(nil).PrometheusManager))
}

// RegisterRoutes mocks base method
func (m *MockRegistry) RegisterRoutes(arg0 *x.RouterAdmin, arg1 *x.RouterPublic) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRoutes", arg0, arg1)
}

// RegisterRoutes indicates an expected call of RegisterRoutes
func (mr *MockRegistryMockRecorder) RegisterRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRoutes", reflect.TypeOf((*MockRegistry)(nil).RegisterRoutes), arg0, arg1)
}

// ScopeStrategy mocks base method
func (m *MockRegistry) ScopeStrategy() fosite.ScopeStrategy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScopeStrategy")
	ret0, _ := ret[0].(fosite.ScopeStrategy)
	return ret0
}

// ScopeStrategy indicates an expected call of ScopeStrategy
func (mr *MockRegistryMockRecorder) ScopeStrategy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScopeStrategy", reflect.TypeOf((*MockRegistry)(nil).ScopeStrategy))
}

// SubjectIdentifierAlgorithm mocks base method
func (m *MockRegistry) SubjectIdentifierAlgorithm() map[string]consent.SubjectIdentifierAlgorithm {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubjectIdentifierAlgorithm")
	ret0, _ := ret[0].(map[string]consent.SubjectIdentifierAlgorithm)
	return ret0
}

// SubjectIdentifierAlgorithm indicates an expected call of SubjectIdentifierAlgorithm
func (mr *MockRegistryMockRecorder) SubjectIdentifierAlgorithm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubjectIdentifierAlgorithm", reflect.TypeOf((*MockRegistry)(nil).SubjectIdentifierAlgorithm))
}

// Tracer mocks base method
func (m *MockRegistry) Tracer() *tracing.Tracer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tracer")
	ret0, _ := ret[0].(*tracing.Tracer)
	return ret0
}

// Tracer indicates an expected call of Tracer
func (mr *MockRegistryMockRecorder) Tracer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracer", reflect.TypeOf((*MockRegistry)(nil).Tracer))
}

// WithBuildInfo mocks base method
func (m *MockRegistry) WithBuildInfo(arg0, arg1, arg2 string) driver.Registry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithBuildInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(driver.Registry)
	return ret0
}

// WithBuildInfo indicates an expected call of WithBuildInfo
func (mr *MockRegistryMockRecorder) WithBuildInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithBuildInfo", reflect.TypeOf((*MockRegistry)(nil).WithBuildInfo), arg0, arg1, arg2)
}

// WithConfig mocks base method
func (m *MockRegistry) WithConfig(arg0 configuration.Provider) driver.Registry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithConfig", arg0)
	ret0, _ := ret[0].(driver.Registry)
	return ret0
}

// WithConfig indicates an expected call of WithConfig
func (mr *MockRegistryMockRecorder) WithConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithConfig", reflect.TypeOf((*MockRegistry)(nil).WithConfig), arg0)
}

// WithLogger mocks base method
func (m *MockRegistry) WithLogger(arg0 logrus.FieldLogger) driver.Registry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithLogger", arg0)
	ret0, _ := ret[0].(driver.Registry)
	return ret0
}

// WithLogger indicates an expected call of WithLogger
func (mr *MockRegistryMockRecorder) WithLogger(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithLogger", reflect.TypeOf((*MockRegistry)(nil).WithLogger), arg0)
}

// Writer mocks base method
func (m *MockRegistry) Writer() herodot.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Writer")
	ret0, _ := ret[0].(herodot.Writer)
	return ret0
}

// Writer indicates an expected call of Writer
func (mr *MockRegistryMockRecorder) Writer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Writer", reflect.TypeOf((*MockRegistry)(nil).Writer))
}
