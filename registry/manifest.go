package registry

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/docker/distribution/manifest/schema1"
)

func (registry *Registry) Manifest(repository, reference string) (*schema1.SignedManifest, error) {
	url := registry.url("/v2/%s/manifests/%s", repository, reference)
	registry.Logf("registry.manifest.get url=%s repository=%s reference=%s", url, repository, reference)

	resp, err := registry.Client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	signedManifest := &schema1.SignedManifest{}
	err = signedManifest.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}

	return signedManifest, nil
}

func (registry *Registry) PutManifest(repository, reference string, signedManifest *schema1.SignedManifest) error {
	url := registry.url("/v2/%s/manifests/%s", repository, reference)
	registry.Logf("registry.manifest.put url=%s repository=%s reference=%s", url, repository, reference)

	body, err := signedManifest.MarshalJSON()
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(body)
	req, err := http.NewRequest("PUT", url, buffer)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", schema1.MediaTypeManifest)
	resp, err := registry.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}
