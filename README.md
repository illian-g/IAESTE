# Save this as README.txt

Hyperledger Fabric Network - Docker Swarm Deployment
====================================================

QUICK START
-----------
1. Initialize Swarm:          sudo docker swarm init --advertise-addr <IP>
2. Create network:            sudo docker network create --driver overlay --attachable fabric_network
3. Deploy stack:              sudo docker stack deploy -c docker/compose.yaml fabric
4. Verify:                    sudo docker stack ps fabric

PROJECT STRUCTURE
-----------------
bin/                          # Fabric binaries (configtxgen, peer, orderer, etc.)
ca.org1.example.com/          # Certificate Authority
orderer.example.com/          # Ordering service
peer0.org1.example.com/       # Peer node
cli.example.com/              # Admin CLI
tools/                        # Chaincode & configs
docker/compose.yaml           # Swarm stack file

SERVICES
--------
• ca.org1.example.com:7054    - Certificate Authority
• orderer.example.com:7050    - Ordering Service
• peer0.org1.example.com:7051 - Peer Node (API:7051, Events:7052)
• cli.example.com             - Admin CLI (no ports)

ESSENTIAL COMMANDS
------------------
# Management
sudo docker service ls                   # List all services
sudo docker service logs -f fabric_peer  # View logs
sudo docker stack ps fabric              # Check stack status

# Troubleshooting
sudo docker node ls                      # Check swarm nodes
sudo docker network inspect fabric_network
sudo docker service inspect fabric_peer

# Maintenance
sudo docker service scale fabric_peer=2  # Scale service
sudo docker stack rm fabric              # Remove stack
sudo docker stack deploy -c docker/compose.yaml fabric  # Update

HEALTH CHECKS
-------------
CA:       https://localhost:7054/healthz
Orderer:  http://localhost:7050/healthz  
Peer:     http://localhost:7051/healthz

SECURITY NOTES
--------------
• All .sk files contain private keys - protect them
• TLS certificates in each component's tls/ directory
• Regenerate all crypto material for production use
• Backup ca.org1.example.com/fabric-ca-server.db regularly

NEXT STEPS
----------
After successful deployment:
1. Bootstrap admin identities
2. Create channels (separate operation)
3. Deploy chaincode (separate operation)

NOTE: This setup deploys only the infrastructure. The CLI container
contains tools for subsequent channel and chaincode operations.

--------------------------------------------------------------------
The keys of the office are in the TVbox in the top drawer next to t-
he desk where I was working at. Sayonara 
--------------------------------------------------------------------
