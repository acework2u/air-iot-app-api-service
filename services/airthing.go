package services

type AirThinkService interface {
	GetCerts(string2 string) (interface{}, error)
	ThingConnect() error
}
