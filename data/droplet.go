package data

type Droplet struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Memory    int     `json:"memory"`
	Disk      int     `json:"disk"`
	Locked    bool    `json:"locked"`
	Status    string  `json:"status"`
	DKernel   Kernel  `json:"kernel"`
	CreatedAt string  `json:"created_at"`
	DNetwork  Network `json:"networks"`
}

type Kernel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Droplets struct {
	DropletsList []Droplet `json:"droplets"`
}

type Network struct {
	V4Networks `json:"v4"`
}

type V4Network struct {
	IPAddress string `json:"ip_address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
	Type      string `json:"type"`
}

type V4Networks []V4Network
