package registry

type repositories struct {
	Repos []string `json:"Repositories"`
}

func (registry *Registry) Catalog() ([]string, error) {
	url := registry.url("/v2/_catalog")
	registry.Logf("registry.catalog url=%s", url)

	var response repositories
	if err := registry.getJson(url, &response); err != nil {
		return nil, err
	}

	return response.Repos, nil
}
