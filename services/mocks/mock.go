package mocks

import (
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
)

type UserReqMock struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MockThinkService struct {
	mock.Mock
}

func NewMockThinkService() *MockThinkService {
	return &MockThinkService{}
}

func (m *MockThinkService) GetCerts() (interface{}, error) {
	panic("mo action")
}
func (m *MockThinkService) GetUserCert(reqMock *UserReqMock) (interface{}, error) {
	panic("mock action")
}
func (m *MockThinkService) UploadToS3(file *multipart.FileHeader) (interface{}, error) {
	panic("Mock action")
}
func (m *MockThinkService) ThingRegister(idToken string) (interface{}, error) {
	panic("Mock action")
}
func (m *MockThinkService) ThingsConnected(idToken string, thing string) (*iotdataplane.PublishOutput, error) {
	panic("Mock Action")
}
func (m *MockThinkService) ThingsCert(idToken string) (interface{}, error) {
	panic("Mock Action")
}
func (m *MockThinkService) ThinksShadows(idToken string, rs string) (interface{}, error) {
	panic("Mock Action")
}
func (m *MockThinkService) NewAwsMqttConnect(cognitoId string) error { panic("Mock no action") }
func (m *MockThinkService) PubGetShadows(thinkName string, shadowName string) (interface{}, error) {
	panic("Mock no action")
}
func (m *MockThinkService) PubUpdateShadows(thinkName string, payload string) (interface{}, error) {
	panic("Mock no action")
}

//UploadToS3(file *multipart.FileHeader) (interface{}, error)
//ThingRegister(idToken string) (interface{}, error)
//ThingsConnected(idToken string, thing string) (*iotdataplane.PublishOutput, error)
//ThingsCert(idToken string) (interface{}, error)
//ThinksShadows(idToken string, rs string) (*ShadowsValue, error)
//NewAwsMqttConnect(cognitoId string) error
//PubGetShadows(thinkName string, shadowName string) (*IndoorInfo, error)
//PubUpdateShadows(thinkName string, payload string) (*IndoorInfo, error)
