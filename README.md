```
nodep — check available dependency packages across npmjs, PyPI or RubyGems registry.

Installation (with Go):
$ go get -u github.com/dwisiswant0/nodep

Or download pre-built binary from releases page (https://github.com/dwisiswant0/nodep/releases/latest).

Usage of nodep:
  nodep <registry> <package_name/dependencies.txt>

Available registry options:
  - npm (npmjs)
  - pip (PyPI)
  - gem (RubyGems)

Examples:
  nodep pip reqeusts
  nodep npm package.txt

Supporting Materials:
- Birsan, Alex. “Dependency Confusion: How I Hacked Into Apple, Microsoft and Dozens of Other Companies.” Medium, February 9, 2021, https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610.
- Tschacher, Nikolai. “Typosquatting Programming Language Package Managers.” incolumitas.com, June 8, 2016, https://incolumitas.com/2016/06/08/typosquatting-package-managers/.
```