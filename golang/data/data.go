package data

import (
	"fmt"
	"net/url"

	"github.com/guionardo/auth-service/golang/setup"
)

var repository Repository

func GetRepository() (Repository, error) {
	if repository == nil {
		if err := getRepository(); err != nil {
			return nil, err
		}
	}
	return repository, nil
}

func getRepository() error {
	cs := setup.GetConfiguration().REPOSITORY_CONNECTION_STRING
	csUrl, err := url.Parse(cs)
	if err != nil {
		return err
	}
	if csUrl.Scheme == "memory" {
		repository = &RepositoryMemory{}

	} else {
		return fmt.Errorf("Invalid repository scheme %s", cs)
	}
	repository.Setup(cs)
	return nil
}
