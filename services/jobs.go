package services

type JobsService interface {
	JobsThingsHandler(userId string) (interface{}, error)
}
