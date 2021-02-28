package beavergo

type Status struct {
	Status string `json:"status"`
}

type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Channel struct {
	Listeners        []string `json:"listeners"`
	ListenersCount   int      `json:"listeners_count"`
	Name             string   `json:"name"`
	Subscribers      []string `json:"subscribers"`
	SubscribersCount int      `json:"subscribers_count"`
	Type             string   `json:"type"`
	CreatedAt        int      `json:"created_at"`
	UpdatedAt        int      `json:"updated_at"`
}

type ClientResp struct {
	Channels  []string `json:"channels"`
	ID        string   `json:"id"`
	Token     string   `json:"token"`
	CreatedAt int      `json:"created_at"`
	UpdatedAt int      `json:"updated_at"`
}
