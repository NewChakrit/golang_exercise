package main

import "fmt"

type User struct {
	ID    int
	First string
}

type MockDatastore struct {
	Users map[int]string
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return User{ID: id, First: user}, nil
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if !ok {
		return fmt.Errorf("User %v already exist", u.ID)
	}

	md.Users[u.ID] = u.First
	return nil
}

type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds Datastore
}

func main() {
}
