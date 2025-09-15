# About

This is a program, written in Go, for interfacing with the St. Edward's Team Dynamix web API. According to Team Dynamix, it is not possible to maintain authenticated state if you have SSO enabled. However, Andrew and I have developed a workaround using [PlayWright](https://github.com/playwright-community/playwright-go) that allows us to grab an API key after a successful SSO login, store it in this program, and use it as an auth header for each API call.

Right now, the program simply connects, gets your API key, and pulls data from a dummy ticket I've created in TDNext. Much more functionality will be added, such as interaction with knowledge base articles, support tickets, people, asssets, locations, etc. Whatever is decided to be useful and worth our time to build is now on the table with our auth solution working.

# TODO

- [ ] Decide on a UI representation to render all the .json data to (BubbleTea vs web UI?)
- [ ] Complete the Ticket struct with all relevant fields from the API and add additional endpoints and calls
- [ ] Implement article struct and calls for interfacing with the knowledge base
