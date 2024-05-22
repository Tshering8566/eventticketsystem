const express = require('express');
const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');

const app = express();
app.use(express.json());

// Static Middleware  
app.use(express.static(path.join(__dirname, 'public')));

async function getContract() {
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    const identity = await wallet.get('Admin@eventtickets.com'); // Update the identity name as per your wallet
    const gateway = new Gateway();

    if (!identity) {
        throw new Error('Identity not found in wallet');
    }
    
    const connectionProfile = JSON.parse(fs.readFileSync(path.resolve(__dirname, 'connection.json'), 'utf8'));
    const connectionOptions = { wallet, identity: identity, discovery: { enabled: false, asLocalhost: true } };

    await gateway.connect(connectionProfile, connectionOptions);

    const network = await gateway.getNetwork('eventorgchannel'); // Update the channel name as per your network
    const contract = network.getContract('eventticketmgt'); // Update the chaincode name as per your network

    return contract;
}

async function submitTransaction(functionName, ...args) {
    const contract = await getContract();
    try {
        const result = await contract.submitTransaction(functionName, ...args);
        console.log(`Transaction ${functionName} result: `, result);
        if (!result) {
            throw new Error(`Transaction ${functionName} returned undefined`);
        }
        return result.toString();
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        throw new Error(`Transaction error: ${error.message}`);
    }
}

async function evaluateTransaction(functionName, ...args) {
    const contract = await getContract();
    const result = await contract.evaluateTransaction(functionName, ...args);
    return result.toString();
}

async function fetchAllEvents() {
    try {
        const contract = await getContract();
        const result = await contract.evaluateTransaction('GetAvailableEvents');
        const events = JSON.parse(result.toString());
        console.log('Fetched events:', events);
        return events;
    } catch (error) {
        console.error('Failed to fetch events:', error);
        throw new Error(`Failed to fetch events: ${error.message}`);
    }
}

// Event CRUD operations
app.post('/events', async (req, res) => {
    try {
        const { id, name, date, location } = req.body;
        const result = await submitTransaction('CreateEvent', id, name, date, location);
        res.status(204).send(result);
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(500).send(`Failed to submit transaction: ${error}`);
    }
});

app.get('/events/:id', async (req, res) => {
    try {
        const { id } = req.params;
        const result = await evaluateTransaction('ReadEvent', id);
        res.status(200).send(result);
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(404).send(`Failed to evaluate transaction: ${error}`);
    }
});

app.get('/allevents', async (req, res) => {
    try {
        const events = await fetchAllEvents(); // Direct call to blockchain or database
        res.status(200).json(events);
    } catch (error) {
        console.error(`Failed to fetch events: ${error}`);
        res.status(500).send(`Failed to fetch events: ${error}`);
    }
});

app.put('/events/:id', async (req, res) => {
    try {
        const { id } = req.params;
        const { name, date, location } = req.body;
        const result = await submitTransaction('UpdateEvent', id, name, date, location);
        res.status(204).send(result);
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(500).send(`Failed to submit transaction: ${error}`);
    }
});

app.delete('/events/:id', async (req, res) => {
    try {
        const { id } = req.params;
        const result = await submitTransaction('DeleteEvent', id);
        res.send(result);
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(500).send(`Failed to submit transaction: ${error}`);
    }
});

// Ticket CRUD operations
app.post('/tickets', async (req, res) => {
    try {
        const { eventId, eventName, count, holder, status } = req.body;
        const result = await submitTransaction('CreateTicket', eventId, eventName, count, holder, status);
        res.status(204).send(result);
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(500).send(`Failed to submit transaction: ${error}`);
    }
});

app.get('/tickets/:id', async (req, res) => {
    try {
        const { id } = req.params;
        const result = await evaluateTransaction('ReadTicket', id);
        res.status(200).send(result);
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(404).send(`Failed to evaluate transaction: ${error}`);
    }
});

app.put('/tickets/:id', async (req, res) => {
    try {
        const { id } = req.params;
        const { status } = req.body;
        const result = await submitTransaction('UpdateTicketStatus', id, status);
        res.status(204).send(result);
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(500).send(`Failed to submit transaction: ${error}`);
    }
});

app.delete('/tickets/:id', async (req, res) => {
    try {
        const { id } = req.params;
        const result = await submitTransaction('DeleteTicket', id);
        res.send(result);
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.status(500).send(`Failed to submit transaction: ${error}`);
    }
});

// Start server
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});
