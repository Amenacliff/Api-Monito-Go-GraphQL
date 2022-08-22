package User

type User struct {
	ID string `json:"id" gorm:"type:uuid;column:id"`
	EmailAddress string `json:"emailAddress" gorm:"type:varchar;column:emailAddress"`
	HashedUserId string `json:"hashedUserId" gorm:"type:uuid;column:hashedUserId"`
	PasswordHash string `json:"passwordHash" gorm:"type:varchar;column:passwordHash"`
	Projects []string `json:"projects"  gorm:"type:text[];column:projects"`
	ApiKey string `json:"apiKey"  gorm:"type:varchar;column:apiKey"`
	TimeZone string `json:"timeZone" gorm:"type:varchar;column:timeZone"`
	NotificationTurnedOn bool `json:"notificationTurnedOn" gorm:"type:bool;column:notificationTurnedOn"`
}