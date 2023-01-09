package models

var Tables = newTableRegistry()

func newTableRegistry() *tableRegistry {
	return &tableRegistry{
		Applications:   "applications",
		Repositories:   "repositories",
		RepositoryAuth: "repository_auth",
	}
}

type tableRegistry struct {
	Applications   string
	Repositories   string
	RepositoryAuth string
}

func (t *tableRegistry) All() []string {
	return []string{
		t.Applications,
		t.Repositories,
		t.RepositoryAuth,
	}
}

func (t *tableRegistry) Get(name string) string {
	switch name {
	case "applications":
		return t.Applications
	case "repositories":
		return t.Repositories
	case "repository_auth":
		return t.RepositoryAuth
	}
	return ""
}
