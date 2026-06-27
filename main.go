package main

import (
	"time"
	"github.com/dice2005x3005/Pokedex/internal/pokecache"
)

func main(){
	conf := &Config{}
	cache := pokecache.NewCache(8 * time.Second)
	repl(conf, cache)
}

