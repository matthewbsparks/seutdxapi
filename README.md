# About

This is a program, written in Go, for interfacing with the St. Edward's Team Dynamix web API. According to Team Dynamix, it is not possible to maintain authenticated state if you have SSO enabled. However, Andrew and I have developed a workaround using [PlayWright](https://github.com/playwright-community/playwright-go) that allows us to grab an API key after a successful SSO login, store it in this program, and use it as an auth header for each API call.

Right now, the program simply connects, gets your API key, and pulls data from a dummy ticket I've created in TDNext. Much more functionality will be added, such as interaction with knowledge base articles, support tickets, people, asssets, locations, etc. Whatever is decided to be useful and worth our time to build is now on the table with our auth solution working.

# TODO

- [ ] Decide on a UI representation to render all the .json data to ([BubbleTea](https://github.com/charmbracelet/bubbletea) vs web UI?)
- [ ] Determine project structure for separating KB articles and SC services from each other
- [ ] Create JSON marshaler functions to handle the JSON DateTime format
- [ ] Update `DoRequest()` function to deal with requests that need body

## Tickets

- [ ] Decide on remaining fields for ticket struct and add them
- [ ] Create some ticket print functions inside ticket.go
- [ ] Create first function for changing ticket `StatusName` via POST to test (turns out this must be done via `StatusID`!)

## KB

- [ ] Decide on relevant article fields and create article struct
- [ ] Create first `GetArticle()` function and associated early print functions

## Services

- [ ] Decide on relevant service fields and create service struct
- [ ] Create first `GetService()` function and associated early print functions
