# Zappos Core Value Server

__A simple Core Value Server just to learn GoLang and Heroku__

## API for Core Values

### Overview

[Zappos](http://www.zappos.com) has Core Values. I wanted to build a simple API server to emulate the current Core Value API by using GoLang and Heroku (maybe I'll migrate it to AWS eventually).

### Usage

This is a public API because the data provided by these endpoints is _freely_ available everywhere: [Zappos Core Values](http://about.zappos.com/our-unique-culture/zappos-core-values)

[http://zapposcorevalues.herokuapp.com/](http://zapposcorevalues.herokuapp.com/)

API healthcheck

[http://zapposcorevalues.herokuapp.com/CoreValue](http://zapposcorevalues.herokuapp.com/CoreValue)

Get All of the Core Values

[http://zapposcorevalues.herokuapp.com/CoreValue/1](http://zapposcorevalues.herokuapp.com/CoreValue/1)

Get a specific core value (1-10)

[http://zapposcorevalues.herokuapp.com/CoreValue/random](http://zapposcorevalues.herokuapp.com/CoreValue/random)

Need a random core value to get through the day? Here you go!

### For Developers

Make sure you have go setup and pull down the repo.

run `PORT=8080 go run main.go`, visit `http://localhost:8080` and you should get a healthcheck JSON response.

As all good projects, this one started out being written in a janky fashion just to see how quickly I could get it up and running. If you find issues, be sure to submit them. I'm new to Go, so I'm also open to suggestions on how to make it better.

### Disclaimer

I built this because I wanted to play with GoLang and Heroku on my own computer on my own time. In no shape, form or fashion is Zappos.com on the hook for anything you see or experience here.