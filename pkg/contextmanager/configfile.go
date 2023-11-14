package contextmanager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrContextNotFound = errors.New("context not found")
)

type Config struct {
	ConfigVersion  string        `yaml:"configVersion"`
	Contexts       []ContextRefs `yaml:"contexts"`
	CurrentContext string        `yaml:"current-context"`
	APIs           []APIs        `yaml:"apis"`
	Users          []Users       `yaml:"users"`
}

type configFileContextManager struct {
	filename string

	config Config
}

func NewConfigFileContextManager(filename string) *configFileContextManager {
	return &configFileContextManager{
		filename: filename,
	}
}

func (c *configFileContextManager) HasCurrentContext() bool {
	_, ok := c.getContextRef(c.config.CurrentContext)
	return ok
}

func (c *configFileContextManager) CurrentContext() Context {
	if c.config.CurrentContext == "" {
		_, _ = fmt.Fprintf(os.Stderr, "no current context set in config\n")
		os.Exit(1)
	}

	contextRef, ok := c.getContextRef(c.config.CurrentContext)
	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, "missing current context %q in config\n", c.config.CurrentContext)
		os.Exit(1)
	}

	api, ok := c.getAPI(contextRef.Context.API)
	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, "missing api %q in config\n", contextRef.Context.API)
		os.Exit(1)
	}

	user, ok := c.getUser(contextRef.Context.User)
	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, "missing user %q in config\n", contextRef.Context.User)
		os.Exit(1)
	}

	return Context{
		Name:         contextRef.Name,
		Organisation: contextRef.Context.Organisation,
		API:          api,
		User:         user,
	}
}

func (c *configFileContextManager) SetCurrentContext(name string) error {
	if _, ok := c.getContextRef(name); !ok {
		return ErrContextNotFound
	}

	c.config.CurrentContext = name
	return nil
}

func (c *configFileContextManager) Load() error {
	yamlFile, err := ioutil.ReadFile(c.filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &c.config)
	if err != nil {
		return err
	}
	return nil
}

func (c *configFileContextManager) AddOrMergeContext(context Context) error {
	c.replaceUser(context.User)
	c.replaceAPI(context.API)
	contextRef := ContextRefs{
		Name: context.Name,
		Context: ContextRef{
			API:          context.API.Name,
			User:         context.User.Name,
			Organisation: context.Organisation,
		},
	}
	c.replaceContext(contextRef)
	return nil
}

func (c *configFileContextManager) RenameAPI(oldName, newName string) error {
	api, ok := c.getAPI(oldName)
	if !ok {
		return ErrContextNotFound
	}
	c.removeAPI(oldName)

	api.Name = newName
	_, alreadyExists := c.getAPI(newName)
	if alreadyExists {
		c.replaceAPI(api)
	} else {
		c.config.APIs = append(c.config.APIs, api)
	}

	updatedContexts := c.config.Contexts
	for i, contextRef := range c.config.Contexts {
		if contextRef.Context.API == oldName {
			contextRef.Context.API = newName
			updatedContexts = removeContext(updatedContexts, i)
			updatedContexts = append(updatedContexts, contextRef)
		}
	}
	c.config.Contexts = updatedContexts
	return nil
}

func (c *configFileContextManager) RenameContext(oldName, newName string) error {
	if c.config.CurrentContext == oldName {
		c.config.CurrentContext = newName
	}
	contextRef, ok := c.getContextRef(oldName)
	if !ok {
		return ErrContextNotFound
	}
	c.removeContext(oldName)

	contextRef.Name = newName
	_, alreadyExists := c.getContextRef(newName)
	if alreadyExists {
		c.replaceContext(contextRef)
	} else {
		c.config.Contexts = append(c.config.Contexts, contextRef)
	}
	return nil
}

func (c *configFileContextManager) RenameUser(oldName, newName string) error {
	user, ok := c.getUser(oldName)
	if !ok {
		return ErrContextNotFound
	}
	c.removeUser(oldName)

	user.Name = newName
	_, alreadyExists := c.getUser(newName)
	if alreadyExists {
		c.replaceUser(user)
	} else {
		c.config.Users = append(c.config.Users, user)
	}

	updatedContexts := c.config.Contexts
	for i, contextRef := range c.config.Contexts {
		if contextRef.Context.User == oldName {
			contextRef.Context.User = newName
			updatedContexts = removeContext(updatedContexts, i)
			updatedContexts = append(updatedContexts, contextRef)
		}
	}
	c.config.Contexts = updatedContexts
	return nil
}

func (c *configFileContextManager) Save() error {
	c.config.ConfigVersion = "v1"

	data, err := yaml.Marshal(c.config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.filename, data, 0600)
}

func (c *configFileContextManager) Config() Config {
	return c.config
}

// -----------

func (c *configFileContextManager) getContextRef(name string) (ContextRefs, bool) {
	for _, context := range c.config.Contexts {
		if context.Name == name {
			return context, true
		}
	}
	return ContextRefs{}, false
}

func (c *configFileContextManager) getAPI(name string) (APIs, bool) {
	for _, api := range c.config.APIs {
		if api.Name == name {
			return api, true
		}
	}
	return APIs{}, false
}

func (c *configFileContextManager) getUser(name string) (Users, bool) {
	for _, user := range c.config.Users {
		if user.Name == name {
			return user, true
		}
	}
	return Users{}, false
}

func (c *configFileContextManager) replaceContext(contextRef ContextRefs) {
	c.removeContext(contextRef.Name)
	c.config.Contexts = append(c.config.Contexts, contextRef)
}

func (c *configFileContextManager) removeContext(name string) {
	for i, ctx := range c.config.Contexts {
		if ctx.Name == name {
			c.config.Contexts = removeContext(c.config.Contexts, i)
			break
		}
	}
}

func (c *configFileContextManager) replaceAPI(api APIs) {
	c.removeAPI(api.Name)
	c.config.APIs = append(c.config.APIs, api)
}

func (c *configFileContextManager) removeAPI(name string) {
	for i, a := range c.config.APIs {
		if a.Name == name {
			c.config.APIs = removeAPI(c.config.APIs, i)
			break
		}
	}
}

func (c *configFileContextManager) replaceUser(user Users) {
	c.removeUser(user.Name)
	c.config.Users = append(c.config.Users, user)
}

func (c *configFileContextManager) removeUser(name string) {
	for i, u := range c.config.Users {
		if u.Name == name {
			c.config.Users = removeUser(c.config.Users, i)
			break
		}
	}
}

func removeAPI(slice []APIs, s int) []APIs {
	return append(slice[:s], slice[s+1:]...)
}

func removeContext(slice []ContextRefs, s int) []ContextRefs {
	return append(slice[:s], slice[s+1:]...)
}

func removeUser(slice []Users, s int) []Users {
	return append(slice[:s], slice[s+1:]...)
}
