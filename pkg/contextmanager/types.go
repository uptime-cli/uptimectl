package contextmanager

type ContextManager interface {
	CurrentContext() Context
	HasCurrentContext() bool
	SetCurrentContext(name string) error
	AddOrMergeContext(context Context) error

	RenameAPI(oldName, newName string) error
	RenameContext(oldName, newName string) error
	RenameUser(oldName, newName string) error

	Load() error
	Save() error
	Config() Config
}

type Context struct {
	Name         string
	Organisation string
	API          APIs
	User         Users
}

type ContextRefs struct {
	Name    string     `yaml:"name"`
	Context ContextRef `yaml:"context"`
}

type ContextRef struct {
	API          string `yaml:"api"`
	User         string `yaml:"user"`
	Organisation string `yaml:"organisation"`
}

type APIs struct {
	Name string `yaml:"name"`
	API  API    `yaml:"api"`
}

type API struct {
	URL string `yaml:"url"`
}

type Users struct {
	Name string `yaml:"name"`
	User User   `yaml:"user"`
}

type User struct {
	BetterUptimeToken string `yaml:"token"`
}
