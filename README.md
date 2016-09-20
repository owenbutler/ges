# ges

command line es cluster stats and health

# Installation

    go get github.com/owenbutler/ges

# Usage

    $ ges
    command line es cluster stats and health

    Usage:
      ges [command]

    Available Commands:
      health      cluster health
      master      list master node
      nodes       node stats

    Flags:
      -u, --url string   elastic url (default "http://localhost:9200")
      -v, --verbose      print column headers

    Use "ges [command] --help" for more information about a command.

# Health

    $ ges -v health
    cluster	status	nodes	data	pri	shards	relo	init	unassign
    ex1   	green 	3    	3   	5  	10    	0   	0   	0

# Nodes

    $ ges -v nodes
    id                    	address                    	master?	name
    KV59Xm7wRXOBqKxAQUu1nQ	inet[/192.168.1.11:9300]	*      	foo.ex
    lVTsvcZbTgSdn0JTH0fGmg	inet[/192.168.1.13:9300]	       	bar.ex
    hAmSThzMQbmkUpXUvxeLTw	inet[/192.168.1.17:9300]	       	baz.ex
