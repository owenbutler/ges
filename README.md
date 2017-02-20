# ges

Command line es cluster stats and health

This will only be useful for you if you are using 0.19.x of Elasticsearch.  Use the cat API if you are on a higher version of ES.

# Installation

    go get github.com/owenbutler/ges

# Usage

    $ ges
    command line es cluster stats and health

    Usage:
      ges [command]

    Available Commands:
      allocation  Manage shard allocation
      health      Cluster health
      heap        Heap statistics for nodes
      indices     Show elasticsearch indices
      master      List master node
      nodes       Node stats
      shards      List shards in the cluster 

    Flags:
      -u, --url string   elastic url (default "http://localhost:9200")
      -v, --verbose      print column headers

    Use "ges [command] --help" for more information about a command.

# Health

    $ ges -v health
    cluster	status	nodes	data	pri	shards	relo	init	unassign
    ex1   	green 	3    	3   	5  	10    	0   	0   	0

# Heap

    $ ges -v heap
    id                    	old gen	max  	ratio 	name
    lVTsvcZbTgSdn0JTH0fGmg	306.4mb	1.8gb	16.28%	foo.ex
    hAmSThzMQbmkUpXUvxeLTw	372.8mb	1.8gb	19.82%	bar.ex
    KV59Xm7wRXOBqKxAQUu1nQ	212.7mb	1.8gb	11.30%	baz.ex

# Indices

    $ ges -v indices
    status	name 	pri	rep	size	docs
    green 	examp	5  	5  	56mb	29449

# Master

    $ ges -v master
    nodeid                	address                	name
    aAfe19jhTMyFmMKtjss7YQ	inet[/10.0.58.149:9300]	foo.ex

# Nodes

    $ ges -v nodes
    id                    	address                    	master?	name
    KV59Xm7wRXOBqKxAQUu1nQ	inet[/192.168.1.11:9300]	*      	foo.ex
    lVTsvcZbTgSdn0JTH0fGmg	inet[/192.168.1.13:9300]	       	bar.ex
    hAmSThzMQbmkUpXUvxeLTw	inet[/192.168.1.17:9300]	       	baz.ex


# Shards

    $ ges -v shards
    index	shard	pri/rep	state  	docs	size 	node                  	relocating
    examp	2    	p      	STARTED	5945	5.7mb	vqmL6JCDQqWCjH15tAwXgg	          
    examp	2    	r      	STARTED	5945	5.7mb	9F9gtnrbTQCJn4bF02Snow	          
    examp	3    	r      	STARTED	5780	5.5mb	9F9gtnrbTQCJn4bF02Snow	          
    examp	3    	p      	STARTED	5780	5.5mb	aAfe19jhTMyFmMKtjss7YQ	          
    examp	4    	r      	STARTED	5954	5.6mb	9F9gtnrbTQCJn4bF02Snow	          
    examp	4    	p      	STARTED	5954	5.6mb	aAfe19jhTMyFmMKtjss7YQ	          
    examp	0    	p      	STARTED	5804	5.5mb	vqmL6JCDQqWCjH15tAwXgg	          
    examp	0    	r      	STARTED	5804	5.5mb	aAfe19jhTMyFmMKtjss7YQ	          
    examp	1    	p      	STARTED	5966	5.6mb	vqmL6JCDQqWCjH15tAwXgg	          
    examp	1    	r      	STARTED	5966	5.6mb	aAfe19jhTMyFmMKtjss7YQ

# Allocation

Small utilities to disable and enable shard allocation, useful in rolling restarts.

    $ ges allocation disable
    $ # restart a node
    $ ges allocation enable
