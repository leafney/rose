package rose

import (
	"net/url"
	"path"
)

// Merge request link
func UrlJoin(basePath string,paths ...string) (string,error)   {
	u,err:=url.Parse(basePath)
	if err!=nil{
		return "", err
	}
	u.Path = path.Join(u.Path,path.Join(paths...))
	return u.String(), nil
}

// Merge links with URL request parameters
func UrlJoinWithQuery(basePath string,pathAndQueries ...string) (*url.URL,error) {
	u,err:=url.Parse(basePath)
	if err!=nil{
		return nil, err
	}

	pqs,err:=url.Parse(path.Join(pathAndQueries...))
	if err!=nil{
		return nil, err
	}

	return u.ResolveReference(pqs),nil
}