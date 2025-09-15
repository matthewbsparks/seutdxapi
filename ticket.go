package main

import (
	"encoding/json"
	"fmt"
)

type Ticket struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	StatusName  string `json:"StatusName"`
	Description string `json:"Description"`
}

func (c *Client) GetTicket(ticketID int) (*Ticket, error) {
	path := fmt.Sprintf("/%d/tickets/%d", AppIDTickets, ticketID)
	body, err := c.DoRequest("GET", path)
	if err != nil {
		return nil, err
	}

	var ticket Ticket
	if err := json.Unmarshal(body, &ticket); err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (t *Ticket) String() string {
	return fmt.Sprintf("Ticket %d: %s\nStatus: %s \nDescription: %s", t.ID, t.Title, t.StatusName, t.Description)
}
