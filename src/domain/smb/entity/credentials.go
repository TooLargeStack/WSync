package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Credentials struct {
	Username      string
	Password      string
	ServerAddress string       `valid:"ipv4"`
	ServerPort    uint8        `valid:"range(1|99999)"`
	ShareSettings ShareSetting `valid:"json"`
}

type ShareSetting struct {
	RemoteShareName string
	WatchersMapping []WatcherMapping
}

type WatcherMapping struct {
	From string
	To   string
}

func NewCredentials(user string, pass string, addr string, port uint8, share string, watchers []WatcherMapping) *Credentials {
	credentials := &Credentials{
		Username:      user,
		Password:      pass,
		ServerAddress: addr,
		ServerPort:    port,
		ShareSettings: ShareSetting{
			RemoteShareName: share,
			WatchersMapping: watchers,
		},
	}

	credentials.validate()
	return credentials
}

func (c *Credentials) validate() {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		panic("error: " + err.Error())
	}
}
