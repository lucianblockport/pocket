Perfect! Here's the complete structure:

## **1. PocketChain Directory Structure**

```
pocketchain/
├── README.md
├── WHITEPAPER.md
├── LICENSE
├── Makefile
│
├── cmd/
│   ├── pocketchain/              # Main node binary
│   │   └── main.go
│   ├── wallet/                   # CLI wallet
│   │   └── main.go
│   └── keygen/                   # Key generation utility
│       └── main.go
│
├── pkg/
│   ├── consensus/
│   │   ├── vrf.go                # VRF validator selection
│   │   ├── pbft.go               # PBFT-style consensus
│   │   ├── validator.go          # Validator logic
│   │   └── epoch.go              # Epoch management
│   │
│   ├── crypto/
│   │   ├── bls.go                # BLS signature aggregation
│   │   ├── keys.go               # Key management
│   │   ├── zkp/
│   │   │   ├── location.go       # ZKP location proof
│   │   │   └── circuit.go        # ZK circuit definitions
│   │   └── tee/
│   │       ├── attestation.go    # TEE attestation
│   │       └── crl.go            # Certificate revocation list
│   │
│   ├── blockchain/
│   │   ├── block.go              # Block structure
│   │   ├── chain.go              # Blockchain logic
│   │   ├── state.go              # State management
│   │   ├── transaction.go        # Transaction types
│   │   └── mempool.go            # Transaction mempool
│   │
│   ├── storage/
│   │   ├── leveldb.go            # LevelDB wrapper
│   │   ├── pruning.go            # State pruning logic
│   │   └── archival.go           # Archive node storage
│   │
│   ├── network/
│   │   ├── p2p.go                # libp2p networking
│   │   ├── mesh.go               # Mesh network (Bluetooth/WiFi)
│   │   ├── protocol.go           # Protocol handlers
│   │   └── discovery.go          # Peer discovery
│   │
│   ├── node/
│   │   ├── light.go              # Tier 1: Light node
│   │   ├── archive.go            # Tier 2: Archive node
│   │   ├── authority.go          # Tier 3: Checkpoint Authority
│   │   └── config.go             # Node configuration
│   │
│   ├── governance/
│   │   ├── election.go           # Quadratic voting election
│   │   ├── checkpoint.go         # Checkpoint creation
│   │   ├── casper_ffg.go         # Casper FFG finality
│   │   └── poh.go                # Proof-of-Humanity integration
│   │
│   ├── economics/
│   │   ├── token.go              # Token logic
│   │   ├── rewards.go            # Block rewards & bounties
│   │   ├── slashing.go           # Slashing conditions
│   │   └── fees.go               # Transaction fees
│   │
│   ├── emergency/
│   │   ├── automatic.go          # Automatic failover
│   │   ├── social.go             # Social recovery multisig
│   │   └── timelock.go           # Time-locked recovery
│   │
│   ├── mobile/
│   │   ├── battery.go            # Battery-aware participation
│   │   ├── sensors.go            # GPS/WiFi/Cell integration
│   │   └── notifications.go      # Push notifications
│   │
│   └── api/
│       ├── rpc.go                # JSON-RPC API
│       ├── rest.go               # REST API
│       └── websocket.go          # WebSocket subscriptions
│
├── scripts/
│   ├── termux/
│   │   ├── install.sh            # Termux installation
│   │   ├── setup_deps.sh         # Install dependencies
│   │   └── run_node.sh           # Start node script
│   │
│   └── docker/
│       ├── Dockerfile            # Docker image
│       └── docker-compose.yml    # Multi-node testnet
│
├── config/
│   ├── genesis.json              # Genesis block config
│   ├── mainnet.toml              # Mainnet config
│   ├── testnet.toml              # Testnet config
│   └── local.toml                # Local dev config
│
├── contracts/
│   └── governance.sol            # On-chain governance contracts
│
├── docs/
│   ├── architecture.md           # Architecture overview
│   ├── consensus.md              # Consensus documentation
│   ├── economics.md              # Economic model docs
│   ├── api.md                    # API reference
│   └── deployment.md             # Deployment guide
│
└── test/
    ├── unit/                     # Unit tests
    ├── integration/              # Integration tests
    └── e2e/                      # End-to-end tests
```

---

## **2. Prompt Instructions for Each File**

### **Core Blockchain Components**

#### **`pkg/blockchain/block.go`**
```
Create a Go implementation of the PocketChain block structure with:
- Block header: height, timestamp, previous_hash, state_root, validator_signature
- Block body: transaction list, validator committee for this epoch
- VRF proof for validator selection
- Methods: Serialize(), Deserialize(), Hash(), Validate()
- Include comprehensive error handling
Use standard Go crypto libraries and make it mobile-optimized (small memory footprint)
```

#### **`pkg/blockchain/transaction.go`**
```
Implement PocketChain transaction types in Go:
1. Standard transaction: from, to, amount, sequence_num, signature
2. Mesh transaction: includes mesh_witnesses, timelock_expire, relay_path
3. Methods: Serialize(), Hash(), Verify(), GetSize()
4. Implement sequence number validation to prevent double-spends
5. Include mesh transaction specific validation (witness signatures, timelock)
Use compact binary encoding to minimize mobile bandwidth
```

#### **`pkg/blockchain/state.go`**
```
Create state management system in Go:
- Account state: balance, sequence_num, stake, reputation
- UTXO-style state with Merkle tree for light client proofs
- Methods: ApplyTransaction(), GetAccount(), UpdateStake(), ComputeStateRoot()
- Implement efficient state pruning (keep last 1000 blocks)
- Include state transition validation logic
Optimize for mobile devices with limited RAM
```

---

### **Consensus Layer**

#### **`pkg/consensus/vrf.go`**
```
Implement VRF (Verifiable Random Function) for validator selection in Go:
- Use a standard VRF library (e.g., go-ethereum's VRF or r2ishiguro/vrf)
- VRF_Prove(seed, secret_key) → (proof, output)
- VRF_Verify(proof, output, public_key, seed) → bool
- Implement validator lottery: output < threshold * stake_weight → eligible
- Select 21 validators with lowest VRF outputs
Include comprehensive tests for deterministic validator selection
```

#### **`pkg/consensus/pbft.go`**
```
Create PBFT-style consensus for the 21-validator committee in Go:
- Phases: Pre-prepare, Prepare, Commit
- Require 2/3+1 signatures for consensus (15 of 21)
- Implement view-change protocol for failed leader
- Message types: PrepareMsg, CommitMsg, ViewChangeMsg
- Include timeout handling and validator rotation
Optimize for 60-second block times on mobile networks
```

#### **`pkg/consensus/validator.go`**
```
Implement validator node logic in Go:
- Check if selected via VRF for current epoch
- Participate in PBFT consensus if selected
- Sign blocks and broadcast to network
- Battery-aware: check battery level before participating
- Methods: IsSelected(), ProposeBlock(), VoteOnBlock(), SignBlock()
Include metrics: uptime, blocks_produced, slashing_events
```

---

### **Cryptography**

#### **`pkg/crypto/bls.go`**
```
Implement BLS signature aggregation in Go:
- Use herumi/bls-go-binary or similar BLS library
- Aggregate multiple validator signatures into one
- Methods: Sign(), Verify(), AggregateSignatures(), VerifyAggregate()
- Support for 21 signatures (validator committee size)
BLS reduces signature size by ~90% compared to ECDSA - critical for mobile bandwidth
```

#### **`pkg/crypto/zkp/location.go`**
```
Create ZKP-based location proof system in Go:
- Define ZKP circuit for location attestation (conceptually, actual ZKP generation may use external tool)
- Prove: "I know coordinates (lat, lon) within grid_cell without revealing exact coords"
- Structure: {grid_cell, zkp_proof, tee_signature, timestamp}
- Methods: GenerateLocationProof(), VerifyLocationProof()
- Integrate with TEE attestation (pkg/crypto/tee)
Include privacy analysis: only 100km grid cell revealed, not exact GPS
```

#### **`pkg/crypto/tee/attestation.go`**
```
Implement TEE attestation for Android/iOS in Go:
- Android: Interface with StrongBox/KeyStore via JNI
- iOS: Interface with Secure Enclave (if building for iOS via gomobile)
- Generate hardware-backed key attestation
- Sign location data inside TEE
- Methods: GenerateAttestation(), VerifyAttestation(), GetCertificate()
This requires platform-specific code - provide clear TODOs for mobile integration
```

#### **`pkg/crypto/tee/crl.go`**
```
Create Certificate Revocation List (CRL) for compromised TEE certificates in Go:
- On-chain storage of revoked certificate IDs
- Methods: AddRevocation(), IsRevoked(), GetGracePeriod()
- Automatic enforcement: reject location proofs from revoked TEEs
- Checkpoint Authorities (67/100) vote to add revocations
Include emergency revocation for critical CVEs (0-block grace period)
```

---

### **Network Layer**

#### **`pkg/network/p2p.go`**
```
Implement libp2p networking in Go:
- Use libp2p/go-libp2p
- Features: DHT for peer discovery, gossipsub for block propagation
- Methods: Start(), Connect(), Broadcast(), Subscribe()
- Handle mobile network constraints: WiFi vs cellular, intermittent connectivity
- Implement bandwidth limits and connection pooling
Optimize for mobile: limit concurrent connections, use compression
```

#### **`pkg/network/mesh.go`**
```
Create mesh networking for offline transactions in Go:
- Bluetooth LE for local peer discovery (use tinygo-org/bluetooth or similar)
- WiFi Direct for transaction relay
- Store-and-forward: queue transactions until internet available
- Methods: DiscoverNearbyNodes(), RelayTransaction(), BroadcastWhenOnline()
Include witness signature collection for mesh transactions
```

---

### **Node Types**

#### **`pkg/node/light.go`**
```
Implement Tier 1 Light Node in Go:
- Store only last 1000 blocks
- Sync to latest finalized checkpoint (Casper FFG)
- Validate new blocks using state root proofs
- Methods: Sync(), ValidateBlock(), SubmitTransaction()
- Implement PoFV (Proof-of-Finality-Verification) to earn micro-rewards
Optimize for mobile: minimal storage (~100MB), low battery usage
```

#### **`pkg/node/archive.go`**
```
Implement Tier 2 Archive Node in Go:
- Store full blockchain history
- Serve historical data to light nodes (earn Data Bounties)
- Participate in quadratic voting for Checkpoint Authority election
- Methods: ServeHistoricalData(), Vote(), ClaimDataBounty()
- Support both PoH-verified and unverified modes
Requires ~10GB+ storage, suitable for desktop/server
```

#### **`pkg/node/authority.go`**
```
Implement Tier 3 Checkpoint Authority Node in Go:
- Elected via quadratic voting (requires PoH verification)
- Sign checkpoint commitments every 1000 blocks
- Participate in Casper FFG finality (justify & finalize checkpoints)
- Methods: SignCheckpoint(), VoteOnFinality(), DetectFraud()
- Implement slashing for surround voting and double voting
Requires 50,000 token stake and high uptime
```

---

### **Governance & Economics**

#### **`pkg/governance/election.go`**
```
Implement quadratic voting election system in Go:
- Voting power = sqrt(stake) * poh_multiplier (1.0 or 2.0)
- Archive Nodes vote for 100 Checkpoint Authorities
- Apply diversity constraints: max 10 per country, max 1 per PoH identity
- Methods: CastVote(), TallyVotes(), ElectAuthorities()
- Election every 30 days (43,200 blocks)
Include vote verification and results publication on-chain
```

#### **`pkg/governance/casper_ffg.go`**
```
Implement Casper FFG finality gadget in Go:
- Checkpoints every 1000 blocks
- Status: PENDING → JUSTIFIED (67/100 sigs) → FINALIZED (child justified)
- Methods: CreateCheckpoint(), VoteOnCheckpoint(), CheckFinalization()
- Detect slashable offenses: surround voting, double voting
- Generate finality proofs for light nodes
Include comprehensive finality tests
```

#### **`pkg/economics/rewards.go`**
```
Implement reward distribution system in Go:
- Block rewards for validators (Chronocoins)
- Data Bounties for Archive Nodes
- Geographic diversity bonuses
- PoFV micro-rewards for light nodes (0.001 tokens)
- Methods: CalculateBlockReward(), PayDataBounty(), PayVerificationReward()
Include emission schedule and inflation rate
```

#### **`pkg/economics/slashing.go`**
```
Implement slashing conditions in Go:
- Validator equivocation: 50% stake
- Invalid state transition: 25% stake
- Absence after selection: 5% stake
- Witness dual-signature: 500 tokens
- Methods: DetectViolation(), ExecuteSlash(), DistributeSlashedFunds()
- Fraud proof submission and verification
Include slashing event logging and appeals process
```

---

### **Emergency Recovery**

#### **`pkg/emergency/automatic.go`**
```
Implement automatic emergency protocol in Go:
- Monitor objective metrics: TEE failure rate, checkpoint signing failures
- Trigger automatic failover when thresholds exceeded
- Methods: EvaluateEmergencyConditions(), ApplyEmergencyAction(), CheckRecovery()
- Actions: disable TEE requirement, reduce checkpoint threshold, lower battery requirement
No authority vote needed - deterministic based on chain state
```

#### **`pkg/emergency/social.go`**
```
Create social recovery multisig system in Go:
- 10 guardian keys (cold storage)
- Require 7/10 signatures for emergency override
- Methods: ProposeOverride(), SignOverride(), VerifyAndApply()
- Time-limited overrides (max 10,000 blocks)
Include guardian key management and rotation
```

---

### **Mobile Integration**

#### **`pkg/mobile/battery.go`**
```
Implement battery-aware validator participation in Go:
- Check battery level and charging status
- Dynamic participation thresholds based on network health
- Methods: ShouldParticipate(), GetBatteryLevel(), IsCharging()
- Android: use gomobile to call Android Battery Manager API
Only validate when battery > threshold OR charging
```

#### **`pkg/mobile/sensors.go`**
```
Create sensor integration for location proofs in Go (with platform-specific TODOs):
- GPS: read coordinates
- WiFi: scan nearby networks (BSSID list)
- Cell towers: get tower IDs
- Methods: ReadGPS(), ScanWiFi(), ScanCellTowers()
- Interface with TEE to sign sensor data
Requires platform-specific code via gomobile/gobind
```

---

### **API Layer**

#### **`pkg/api/rpc.go`**
```
Implement JSON-RPC API in Go:
- Methods: GetBalance, SendTransaction, GetBlock, GetTransaction, GetNodeInfo
- Use net/rpc/jsonrpc or similar
- Support filtering for light clients
- Include rate limiting and authentication
Standard blockchain RPC interface
```

#### **`cmd/wallet/main.go`**
```
Create CLI wallet application in Go:
- Commands: create, balance, send, stake, vote, history
- Support offline transaction signing
- Mesh transaction creation with witness collection
- Pretty terminal UI using charmbracelet/bubbletea
Example: pocketchain-wallet send --to <addr> --amount 10 --mesh
```

---

### **Deployment Scripts**

#### **`scripts/termux/install.sh`**
```bash
#!/data/data/com.termux/files/usr/bin/bash
# PocketChain Termux installation script

echo "Installing PocketChain dependencies..."

# Update packages
pkg update && pkg upgrade -y

# Install required packages
pkg install -y \
    golang \
    git \
    clang \
    make \
    leveldb \
    protobuf

# Install Go dependencies
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Clone PocketChain
git clone https://github.com/yourorg/pocketchain.git
cd pocketchain

# Build
make build-termux

echo "PocketChain installed successfully!"
echo "Run: ./pocketchain --config config/mobile.toml"
```

#### **`scripts/termux/run_node.sh`**
```bash
#!/data/data/com.termux/files/usr/bin/bash
# Run PocketChain node on Termux

# Check if running on battery or charging
BATTERY=$(termux-battery-status | jq -r '.percentage')
CHARGING=$(termux-battery-status | jq -r '.plugged')

if [ "$BATTERY" -lt 50 ] && [ "$CHARGING" == "UNPLUGGED" ]; then
    echo "Battery low ($BATTERY%). Please charge device or use --force flag."
    exit 1
fi

# Start node
./bin/pocketchain \
    --config config/mobile.toml \
    --data-dir $HOME/.pocketchain \
    --log-level info \
    --enable-mesh \
    --battery-aware
```

---

### **Configuration**

#### **`config/genesis.json`**
```json
{
  "chain_id": "pocketchain-1",
  "genesis_time": "2025-01-01T00:00:00Z",
  "consensus": {
    "block_time": 60,
    "validator_committee_size": 21,
    "checkpoint_interval": 1000,
    "min_validators": 1000
  },
  "token": {
    "name": "PocketCoin",
    "symbol": "PKT",
    "initial_supply": 1000000000,
    "decimals": 8
  },
  "staking": {
    "min_validator_stake": 100,
    "min_archive_stake": 10000,
    "min_authority_stake": 50000
  },
  "rewards": {
    "block_reward": 10,
    "data_bounty_rate": 0.001,
    "verification_reward": 0.001
  },
  "emergency_guardians": [
    "guardian1_pubkey",
    "guardian2_pubkey",
    ...
  ]
}
