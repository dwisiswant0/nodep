package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	flag.Parse()

	opt = &options{}
	opt.reg = flag.Arg(0)
	opt.dep = flag.Arg(1)

	if err := opt.validate(); err != nil {
		fmt.Fprintf(os.Stderr, "%sError: %s!%s\n\n%s", cred, err, cend, usage)
		os.Exit(1)
	}

	do = client()
}

func main() {
	wg.Add(1)

	go func() {
		for _, dep := range opt.list {
			reg, err := check(opt.reg, dep)
			if err != nil {
				fmt.Printf("%s[error] %s%s\n", cyel, cend, dep)
			}

			if reg {
				fmt.Printf("%s[regis] %s%s\n", cred, cend, dep)
			} else {
				fmt.Printf("%s[avail] %s%s\n", cgrn, cend, dep)
			}
		}

		defer wg.Done()
	}()

	wg.Wait()
}
