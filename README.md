# Event Ticketing Solutions

## Problem Statement
Event Ticketing Solutions aims to revolutionize the event management industry by addressing key challenges faced by event organizers and ticket providers. Traditional systems suffer from inefficiencies like ticket scalping, lack of transparency, and centralized control, leading to inflated prices, counterfeit tickets, and poor user experiences.

### Challenges
1. **Ticket Scalping:** Scalpers exploit demand by reselling tickets at exorbitant prices.
2. **Transparency and Trust:** Customers lack visibility into ticketing processes.
3. **Centralized Control:** Limits competition, innovation, and fair pricing.

## Blockchain Network Design (Hyperledger Fabric - HLF)
- **Organizations:**
  - Event Ticketing Solutions (Single peer organization)
  - Orderer Organization (OrdererOrg)

- **Peer Nodes:**
  - Event Ticketing Solutions (EventOrg1): 1 peer node
  - Ticketing Platform Providers (TicketProviderOrg): 1 peer node

- **Orderer Nodes:**
  - OrdererOrg: 1 orderer node (orderer.eventtickets.com)

- **Channels:**
  - Event Ticketing Channel: Connects EventOrg1 and TicketProviderOrg securely.

## Assignment Two: Wallet Application and Backend Code
- **Wallet Application:**
  - **Functionality:**
    - Identity management: Add, list, and export identities stored in the wallet.
    - Connection to the blockchain network: Load network configuration and establish connection using stored identities.
    - Interacting with smart contracts: Submit transactions to execute smart contract functions.
    - Error handling and security: Ensure robust error handling and implement security measures to protect user identities and transaction data.

- **Backend Code:**
  - **Functionality:**
    - Event management: Create, update, and delete events with details like ID, name, date, and location.
    - Ticket management: Generate tickets for events, update ticket status, and handle transactions related to event and ticket management.
    - Transaction handling: Execute transactions conforming to business requirements using Hyperledger Fabric contract API.
    - Integration with Hyperledger Fabric: Utilize Fabric contract API for interaction with the world state and transaction handling.

Together, the wallet application and backend code form a robust solution for managing events and tickets while addressing the challenges faced by Event Ticketing Solutions. This solution leverages the power of blockchain technology to ensure transparency, security, and efficiency in event ticketing processes.
