package main

import (
	"time"
	"github.com/dice2005x3005/Pokedex/internal/pokecache"
)

func main(){
	conf := &Config{}
	user := &User{
		capturas: make(map[string]Pokemon),
	}
	cache := pokecache.NewCache(8 * time.Second)
	repl(conf, cache, user)
}

