# Pirate
A CLI for the Digital Ocean API. [Work In Process]

- [Getting Started](#getting-started)
- [Droplets](#droplets)
    - **Use** [Get](#get), [Create](#create), [Delete](#delete)
- [SSH Keys](#ssh-keys)
    - **Use** [Get](#get), [Create](#create), [Delete](#delete)
- [Load Balancers](#load-balancers)
    - **Use** [Get](#get), [Create](#create), [Delete](#delete), [Add Droplets](#add-droplets),[Remove Droplets](#remove-droplets)

## Getting Started 

Must have a working Go [environment](https://golang.org/doc/install) set up.

1. Run `go get -v github.com/mikaelm1/pirate`. If everything worked, running `pirate -h` should display the help page.
2. Inside the project directory, run `touch config.yaml` and fill it in the same way the sample `config.yaml.devexample` is organized. The `output` variable can be either `text` or `json`. The `token` variable will hold your personal access token to use Digital Ocean's API. Instructions for generating the token can be found [here](https://www.digitalocean.com/community/tutorials/how-to-use-the-digitalocean-api-v2).
3. Check to make sure it's working by running `pirate user`. If you get back your account information, it means you're good to go. 

## Droplets

### Get 

1. To fetch all droplets owned: `pirate droplet -l`
2. To fetch single droplet with id: `pirate droplet -s=12345`

### Create

To create single droplet: `pirate droplet create --name=your_droplet_name --key=ssh_id_1,ssh_id_2`

Run `pirate droplet create -h` to see a list of all other flags you can set when creating a droplet.

### Delete

To delete single droplet: `pirate droplet delete --droplet-id=123456`

## SSH Keys

### Get

1. To fetch all ssh keys: `pirate ssh_key -l`
1. To fetch single ssh key: `pirate ssh_key -s --id=123456`. Or you can use the ssh key fingerprint: `pirate ssh_key -f=0f:0f:0f...`

### Create 

To create a new ssh key: `pirate ssh_key create -n=key_name -k=path_to_public_key`

### Delete

To remove an ssh key: `pirate ssh_key delete -i=key_id`

## Load Balancers

### Get

To fetch all load balancers: `pirate balancers`

### Create

There are many flags that allow you set options on the load balancer, but most of them have default values and some are not required to create a load balancer. Run `pirate balancers create -h` to see all the available flags. By defualt, a load balancer will have a single forwarindg rules object added to it. If you wish to add more than one forwarding rule, set the `--num-rules=#` to the number of rules you want to add. If you set this flag, you must also set the flags for the forwarding rules with enough values to match the number of rules. For example, if you set `--num-rules=3`, then you must give `--entry-protocols=http,http,https` three entry protocols, along with the other forwarding rule flags.

To create single load balancer: `pirate balancers create --name=somename`

### Delete

To delete a load balancer: `pirate balancers delete --balancer-id=theID`

### Add Droplets

You can add one or more droplets to a single load balancer. Running `pirate balancers add-droplets --balancer-id=theID --droplet-ids=12345,23456` will add two droplets with the given IDs to the load balancer.

### Remove Droplets

You can remove one or more droplets from a single load balancer. Running `pirate balancers remove-droplets --balancer-id=theID --droplet-ids=12345` will remove one droplet with the given ID from the load balancer.

## Images

### Get

1. To list all images: `pirate image -l`
2. To list all distro images: `pirate image -d`