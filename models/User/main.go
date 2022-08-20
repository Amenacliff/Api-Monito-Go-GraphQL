package User

type User struct {
	ID string `json:"id"`
	EmailAddress string `json:"emailAddress"`
	HashedUserId string `json:"hashedUserId"`
	PasswordHash string `json:"passwordHash"`
	Projects []string `json:"projects"`
	ApiKey string `json:"apiKey"`
	TimeZone string `json:"timeZone"`
	NotificationTurnedOn bool `json:"notificationTurnedOn"`
}