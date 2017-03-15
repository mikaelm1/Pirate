# Pirate
A CLI for the Digital Ocean API. [Work In Process]

## Getting Started 

Must have a working Go [environment](https://golang.org/doc/install) set up.

1. Run `go get -v github.com/mikaelm1/pirate`. If everything worked, running `pirate -h` should display the help page.
2. Inside the project directory, run `touch config.yaml` and fill it in the same way the sample `config.yaml.devexample` is organized. The `output` variable can be either `text` or `json`. The `token` variable will hold your personal access token to use Digital Ocean's API. Instructions for generating the token can be found [here](https://www.digitalocean.com/community/tutorials/how-to-use-the-digitalocean-api-v2).


