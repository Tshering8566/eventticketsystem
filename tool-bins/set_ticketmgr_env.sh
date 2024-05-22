#!/bin/bash
export ORG_CONTEXT=eventtickets
export ORG_NAME=EventOrg
export CORE_PEER_LOCALMSPID=EventOrgMSP
# Logging specifications
export FABRIC_LOGGING_SPEC=INFO
# Location of the core.yaml
export FABRIC_CFG_PATH=/workspaces/eventticketconsortium/config/ticketmgr
# Address of the peer
export CORE_PEER_ADDRESS=ticketmgr.eventtickets.com:7051
# Local MSP for the admin - Commands need to be executed as org admin
export CORE_PEER_MSPCONFIGPATH=/workspaces/eventticketconsortium/config/crypto-config/peerOrganizations/eventtickets.com/users/Admin@eventtickets.com/msp
# Address of the orderer
export ORDERER_ADDRESS=orderer.eventtickets.com:7050
export CORE_PEER_TLS_ENABLED=false

#### Chaincode related properties
export CC_NAME="eventticketmgt"
export CC_PATH="./chaincodes/eventticketmgt/"
export CC_CHANNEL_ID="eventticketingchannel"
export CC_LANGUAGE="golang"

# Properties of Chaincode
export INTERNAL_DEV_VERSION="1.0"
export CC_VERSION="1.1"
export CC2_PACKAGE_FOLDER="./chaincodes/packages/"
export CC2_SEQUENCE=1
export CC2_INIT_REQUIRED="--init-required"

# Create the package with this name
export CC_PACKAGE_FILE="$CC2_PACKAGE_FOLDER$CC_NAME.$CC_VERSION-$INTERNAL_DEV_VERSION.tar.gz"

# Extracts the package ID for the installed chaincode
export CC_LABEL="$CC_NAME.$CC_VERSION-$INTERNAL_DEV_VERSION"

peer channel list
