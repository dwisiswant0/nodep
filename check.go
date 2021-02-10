package main

import "fmt"

func check(reg string, dep string) (bool, error) {
	switch reg {
	case "npm":
		endpoint = fmt.Sprintf("https://api.npms.io/v2/package/%s/", dep)
	case "pip":
		endpoint = fmt.Sprintf("https://pypi.org/simple/%s/", dep)
	case "gem":
		endpoint = fmt.Sprintf("https://rubygems.org/api/v1/gems/%s.json", dep)
	default:
		return false, nil
	}

	resp, err := do.Get(endpoint)
	if err != nil {
		return false, err
	}

	if resp.StatusCode == 200 {
		return true, nil
	}

	return false, nil
}
