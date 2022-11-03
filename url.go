package rose

import (
	"net/url"
	"path"
)

// UrlJoin Merge request link
func UrlJoin(baseUrl string, elem ...string) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	if len(elem) > 0 {
		elem = append([]string{u.Path}, elem...)
		u.Path = path.Join(elem...)
	}
	return u.String(), nil
}

// UrlJoinWithQuery Merge links with URL request parameters
func UrlJoinWithQuery(baseUrl string, pathOrQueries ...string) (*url.URL, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	pqs, err := url.Parse(path.Join(pathOrQueries...))
	if err != nil {
		return nil, err
	}

	return u.ResolveReference(pqs), nil
}

// UrlJoinPath Url路径合并
func UrlJoinPath(baseUrl string, elem ...string) (*url.URL, error) {
	// [proposal: net/url: add JoinPath, URL.JoinPath · Issue #47005 · golang/go](https://github.com/golang/go/issues/47005)
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	if len(elem) > 0 {
		elem = append([]string{u.Path}, elem...)
		u.Path = path.Join(elem...)
	}
	return u, nil
}
