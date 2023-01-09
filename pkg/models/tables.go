package models

var Tables = newTableRegistry()

func newTableRegistry() *tableRegistry {
	return &tableRegistry{
		Clusters: "clusters",
	}
}

type tableRegistry struct {
	Clusters string
}

func (t *tableRegistry) All() []string {
	return []string{
		t.Clusters,
	}
}

func (t *tableRegistry) Get(name string) string {
	switch name {
	case "cluster":
		return t.Clusters
	}
	return ""
}
