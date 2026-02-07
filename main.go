package main

import "eksdoctor/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
