package svc

type Repo interface{ Save(name string) error }

type Service struct{ repo Repo }

func New(repo Repo) Service { return Service{repo: repo} }

func (s Service) Create(name string) error { return s.repo.Save(name) }
