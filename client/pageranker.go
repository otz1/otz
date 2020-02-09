package client

type PageRankerClient struct{}

func NewPageRankerClient() *PageRankerClient {
	return &PageRankerClient{}
}

func (pr *PageRankerClient) GetRanking(query string) {
}
