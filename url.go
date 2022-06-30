package rose

import (
	"net/url"
	"path"
)

// Merge request link
func UrlJoin(basePath string, paths ...string) (string, error) {
	u, err := url.Parse(basePath)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, path.Join(paths...))
	return u.String(), nil
}

// Merge links with URL request parameters
func UrlJoinWithQuery(basePath string, pathAndQueries ...string) (*url.URL, error) {
	u, err := url.Parse(basePath)
	if err != nil {
		return nil, err
	}

	pqs, err := url.Parse(path.Join(pathAndQueries...))
	if err != nil {
		return nil, err
	}

	return u.ResolveReference(pqs), nil
}

// Url合并路径
func JoinPaths(baseUrl string, elem ...string) (*url.URL, error) {
	// [proposal: net/url: add JoinPath, URL.JoinPath · Issue #47005 · golang/go](https://github.com/golang/go/issues/47005)
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	if len(elem) > 0 {
		elem = append([]string{url.Path}, elem...)
		url.Path = path.Join(elem...)
	}
	return url, nil
}
