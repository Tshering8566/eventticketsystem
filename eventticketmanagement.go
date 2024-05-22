package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Event represents an event in the system
type Event struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
}

// Ticket represents a ticket for an event
type Ticket struct {
	ID        string `json:"id"`
	EventID   string `json:"eventId"`
	EventName string `json:"eventName"`
	Holder    string `json:"holder"`
	Status    string `json:"status"`
}

// EventTicketingContract provides functions for managing events and tickets
type EventTicketingContract struct {
	contractapi.Contract
}

// CreateEvent creates a new event
func (c *EventTicketingContract) CreateEvent(ctx contractapi.TransactionContextInterface, id string, name string, date time.Time, location string) error {
	event := Event{
		ID:       id,
		Name:     name,
		Date:     date,
		Location: location,
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(id, eventJSON)
	if err != nil {
		return fmt.Errorf("failed to put event to world state: %v", err)
	}

	eventPayload := fmt.Sprintf("Created event: %s", id)
	err = ctx.GetStub().SetEvent("CreateEvent", []byte(eventPayload))
	if err != nil {
		return fmt.Errorf("event failed to register: %v", err)
	}

	return nil
}

// StoreEvent stores the details of a created event in the world state
func (c *EventTicketingContract) StoreEvent(ctx contractapi.TransactionContextInterface, eventId string, eventName string) error {
	event := Event{
		ID:   eventId,
		Name: eventName,
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Store the event in the world state
	err = ctx.GetStub().PutState(eventId, eventJSON)
	if err != nil {
		return fmt.Errorf("failed to put event to world state: %v", err)
	}

	return nil
}

// GetAvailableEvents retrieves the list of available events from the world state
func (c *EventTicketingContract) GetAvailableEvents(ctx contractapi.TransactionContextInterface) ([]Event, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var events []Event
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var event Event
		err = json.Unmarshal(queryResponse.Value, &event)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// CreateTicket creates new tickets for a specific event
func (c *EventTicketingContract) CreateTicket(ctx contractapi.TransactionContextInterface, eventId string, eventName string, count int, holder string, status string) error {
	// Check if the event exists
	eventJSON, err := ctx.GetStub().GetState(eventId)
	if err != nil {
		return fmt.Errorf("failed to read event from world state: %v", err)
	}
	if eventJSON == nil {
		return fmt.Errorf("event %s does not exist", eventId)
	}

	for i := 0; i < count; i++ {
		// Create a unique ticket ID
		ticketID := fmt.Sprintf("%s-%s-%d", eventId, holder, i+1)

		// Create the ticket
		ticket := Ticket{
			ID:        ticketID,
			EventID:   eventId,
			EventName: eventName,
			Holder:    holder,
			Status:    status,
		}

		ticketJSON, err := json.Marshal(ticket)
		if err != nil {
			return err
		}

		// Store the ticket in the world state
		err = ctx.GetStub().PutState(ticketID, ticketJSON)
		if err != nil {
			return fmt.Errorf("failed to put ticket to world state: %v", err)
		}

		eventPayload := fmt.Sprintf("Created ticket: %s", ticketID)
		err = ctx.GetStub().SetEvent("CreateTicket", []byte(eventPayload))
		if err != nil {
			return fmt.Errorf("event failed to register: %v", err)
		}
	}

	return nil
}

// ReadEvent reads the details of an event by ID
func (c *EventTicketingContract) ReadEvent(ctx contractapi.TransactionContextInterface, id string) (*Event, error) {
	eventJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read event from world state: %v", err)
	}
	if eventJSON == nil {
		return nil, fmt.Errorf("event %s does not exist", id)
	}

	var event Event
	err = json.Unmarshal(eventJSON, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// ReadTicket reads the details of a ticket by ID
func (c *EventTicketingContract) ReadTicket(ctx contractapi.TransactionContextInterface, id string) (*Ticket, error) {
	ticketJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read ticket from world state: %v", err)
	}
	if ticketJSON == nil {
		return nil, fmt.Errorf("ticket %s does not exist", id)
	}

	var ticket Ticket
	err = json.Unmarshal(ticketJSON, &ticket)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

// UpdateTicketStatus updates the status of a ticket
func (c *EventTicketingContract) UpdateTicketStatus(ctx contractapi.TransactionContextInterface, id string, status string) error {
	ticket, err := c.ReadTicket(ctx, id)
	if err != nil {
		return err
	}

	ticket.Status = status

	ticketJSON, err := json.Marshal(ticket)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(id, ticketJSON)
	if err != nil {
		return fmt.Errorf("failed to update ticket: %v", err)
	}

	return nil
}

// UpdateEvent updates an existing event with new details
func (c *EventTicketingContract) UpdateEvent(ctx contractapi.TransactionContextInterface, id string, name string, date time.Time, location string) error {
	// Retrieve the existing event from the world state
	event, err := c.ReadEvent(ctx, id)
	if err != nil {
		return err
	}

	// Update the event details
	event.Name = name
	event.Date = date
	event.Location = location

	// Marshal the updated event
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Put the updated event back into the world state
	err = ctx.GetStub().PutState(id, eventJSON)
	if err != nil {
		return fmt.Errorf("failed to update event: %v", err)
	}

	return nil
}

// DeleteEvent deletes an event by ID
func (c *EventTicketingContract) DeleteEvent(ctx contractapi.TransactionContextInterface, id string) error {
	err := ctx.GetStub().DelState(id)
	if err != nil {
		return fmt.Errorf("failed to delete event: %v", err)
	}

	return nil
}

// DeleteTicket deletes a ticket by ID
func (c *EventTicketingContract) DeleteTicket(ctx contractapi.TransactionContextInterface, id string) error {
	err := ctx.GetStub().DelState(id)
	if err != nil {
		return fmt.Errorf("failed to delete ticket: %v", err)
	}

	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(EventTicketingContract))
	if err != nil {
		fmt.Printf("Error creating event ticketing chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting event ticketing chaincode: %s", err.Error())
	}
}
