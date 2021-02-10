package main

import (
	"net/http"
	"sync"
)

var (
	endpoint string

	cpre = "\033["
	cend = cpre + "0m"
	cred = cpre + "31m"
	cgrn = cpre + "32m"
	cyel = cpre + "33m"

	do  *http.Client
	wg  sync.WaitGroup
	opt *options

	usage = `nodep — check available dependency packages across npmjs, PyPI or RubyGems registry.
by @dwisiswant0

Usage:
  nodep <registry> <package_name/dependencies.txt>

Available registry options:
  - npm (npmjs)
  - pip (PyPI)
  - gem (RubyGems)

Examples:
  nodep pip reqeusts
  nodep npm package.txt
`
)
