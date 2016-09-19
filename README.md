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

    Flags:
      -u, --url string   elastic url (default "http://localhost:9200")
      -v, --verbose      print column headers

    Use "ges [command] --help" for more information about a command.

    $ ./ges -v health
    cluster      	status	nodes	data	pri	shards	relo	init	unassign
    elasticsearch	green 	1    	1   	0  	0     	0   	0   	0
