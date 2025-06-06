# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [rpc.proto](#rpc.proto)
    - [RPCError](#protowire.RPCError)
    - [RpcBlock](#protowire.RpcBlock)
    - [RpcBlockHeader](#protowire.RpcBlockHeader)
    - [RpcBlockLevelParents](#protowire.RpcBlockLevelParents)
    - [RpcBlockVerboseData](#protowire.RpcBlockVerboseData)
    - [RpcTransaction](#protowire.RpcTransaction)
    - [RpcTransactionInput](#protowire.RpcTransactionInput)
    - [RpcScriptPublicKey](#protowire.RpcScriptPublicKey)
    - [RpcTransactionOutput](#protowire.RpcTransactionOutput)
    - [RpcOutpoint](#protowire.RpcOutpoint)
    - [RpcUtxoEntry](#protowire.RpcUtxoEntry)
    - [RpcTransactionVerboseData](#protowire.RpcTransactionVerboseData)
    - [RpcTransactionInputVerboseData](#protowire.RpcTransactionInputVerboseData)
    - [RpcTransactionOutputVerboseData](#protowire.RpcTransactionOutputVerboseData)
    - [GetCurrentNetworkRequestMessage](#protowire.GetCurrentNetworkRequestMessage)
    - [GetCurrentNetworkResponseMessage](#protowire.GetCurrentNetworkResponseMessage)
    - [SubmitBlockRequestMessage](#protowire.SubmitBlockRequestMessage)
    - [SubmitBlockResponseMessage](#protowire.SubmitBlockResponseMessage)
    - [GetBlockTemplateRequestMessage](#protowire.GetBlockTemplateRequestMessage)
    - [GetBlockTemplateResponseMessage](#protowire.GetBlockTemplateResponseMessage)
    - [NotifyBlockAddedRequestMessage](#protowire.NotifyBlockAddedRequestMessage)
    - [NotifyBlockAddedResponseMessage](#protowire.NotifyBlockAddedResponseMessage)
    - [BlockAddedNotificationMessage](#protowire.BlockAddedNotificationMessage)
    - [GetPeerAddressesRequestMessage](#protowire.GetPeerAddressesRequestMessage)
    - [GetPeerAddressesResponseMessage](#protowire.GetPeerAddressesResponseMessage)
    - [GetPeerAddressesKnownAddressMessage](#protowire.GetPeerAddressesKnownAddressMessage)
    - [GetSelectedTipHashRequestMessage](#protowire.GetSelectedTipHashRequestMessage)
    - [GetSelectedTipHashResponseMessage](#protowire.GetSelectedTipHashResponseMessage)
    - [GetMempoolEntryRequestMessage](#protowire.GetMempoolEntryRequestMessage)
    - [GetMempoolEntryResponseMessage](#protowire.GetMempoolEntryResponseMessage)
    - [GetMempoolEntriesRequestMessage](#protowire.GetMempoolEntriesRequestMessage)
    - [GetMempoolEntriesResponseMessage](#protowire.GetMempoolEntriesResponseMessage)
    - [MempoolEntry](#protowire.MempoolEntry)
    - [GetConnectedPeerInfoRequestMessage](#protowire.GetConnectedPeerInfoRequestMessage)
    - [GetConnectedPeerInfoResponseMessage](#protowire.GetConnectedPeerInfoResponseMessage)
    - [GetConnectedPeerInfoMessage](#protowire.GetConnectedPeerInfoMessage)
    - [AddPeerRequestMessage](#protowire.AddPeerRequestMessage)
    - [AddPeerResponseMessage](#protowire.AddPeerResponseMessage)
    - [SubmitTransactionRequestMessage](#protowire.SubmitTransactionRequestMessage)
    - [SubmitTransactionResponseMessage](#protowire.SubmitTransactionResponseMessage)
    - [NotifyVirtualSelectedParentChainChangedRequestMessage](#protowire.NotifyVirtualSelectedParentChainChangedRequestMessage)
    - [NotifyVirtualSelectedParentChainChangedResponseMessage](#protowire.NotifyVirtualSelectedParentChainChangedResponseMessage)
    - [VirtualSelectedParentChainChangedNotificationMessage](#protowire.VirtualSelectedParentChainChangedNotificationMessage)
    - [GetBlockRequestMessage](#protowire.GetBlockRequestMessage)
    - [GetBlockResponseMessage](#protowire.GetBlockResponseMessage)
    - [GetSubnetworkRequestMessage](#protowire.GetSubnetworkRequestMessage)
    - [GetSubnetworkResponseMessage](#protowire.GetSubnetworkResponseMessage)
    - [GetVirtualSelectedParentChainFromBlockRequestMessage](#protowire.GetVirtualSelectedParentChainFromBlockRequestMessage)
    - [AcceptedTransactionIds](#protowire.AcceptedTransactionIds)
    - [GetVirtualSelectedParentChainFromBlockResponseMessage](#protowire.GetVirtualSelectedParentChainFromBlockResponseMessage)
    - [GetBlocksRequestMessage](#protowire.GetBlocksRequestMessage)
    - [GetBlocksResponseMessage](#protowire.GetBlocksResponseMessage)
    - [GetBlockCountRequestMessage](#protowire.GetBlockCountRequestMessage)
    - [GetBlockCountResponseMessage](#protowire.GetBlockCountResponseMessage)
    - [GetBlockDagInfoRequestMessage](#protowire.GetBlockDagInfoRequestMessage)
    - [GetBlockDagInfoResponseMessage](#protowire.GetBlockDagInfoResponseMessage)
    - [ResolveFinalityConflictRequestMessage](#protowire.ResolveFinalityConflictRequestMessage)
    - [ResolveFinalityConflictResponseMessage](#protowire.ResolveFinalityConflictResponseMessage)
    - [NotifyFinalityConflictsRequestMessage](#protowire.NotifyFinalityConflictsRequestMessage)
    - [NotifyFinalityConflictsResponseMessage](#protowire.NotifyFinalityConflictsResponseMessage)
    - [FinalityConflictNotificationMessage](#protowire.FinalityConflictNotificationMessage)
    - [FinalityConflictResolvedNotificationMessage](#protowire.FinalityConflictResolvedNotificationMessage)
    - [ShutDownRequestMessage](#protowire.ShutDownRequestMessage)
    - [ShutDownResponseMessage](#protowire.ShutDownResponseMessage)
    - [GetHeadersRequestMessage](#protowire.GetHeadersRequestMessage)
    - [GetHeadersResponseMessage](#protowire.GetHeadersResponseMessage)
    - [NotifyUtxosChangedRequestMessage](#protowire.NotifyUtxosChangedRequestMessage)
    - [NotifyUtxosChangedResponseMessage](#protowire.NotifyUtxosChangedResponseMessage)
    - [UtxosChangedNotificationMessage](#protowire.UtxosChangedNotificationMessage)
    - [UtxosByAddressesEntry](#protowire.UtxosByAddressesEntry)
    - [StopNotifyingUtxosChangedRequestMessage](#protowire.StopNotifyingUtxosChangedRequestMessage)
    - [StopNotifyingUtxosChangedResponseMessage](#protowire.StopNotifyingUtxosChangedResponseMessage)
    - [GetUtxosByAddressesRequestMessage](#protowire.GetUtxosByAddressesRequestMessage)
    - [GetUtxosByAddressesResponseMessage](#protowire.GetUtxosByAddressesResponseMessage)
    - [GetBalanceByAddressRequestMessage](#protowire.GetBalanceByAddressRequestMessage)
    - [GetBalanceByAddressResponseMessage](#protowire.GetBalanceByAddressResponseMessage)
    - [GetBalancesByAddressesRequestMessage](#protowire.GetBalancesByAddressesRequestMessage)
    - [BalancesByAddressEntry](#protowire.BalancesByAddressEntry)
    - [GetBalancesByAddressesResponseMessage](#protowire.GetBalancesByAddressesResponseMessage)
    - [GetVirtualSelectedParentBlueScoreRequestMessage](#protowire.GetVirtualSelectedParentBlueScoreRequestMessage)
    - [GetVirtualSelectedParentBlueScoreResponseMessage](#protowire.GetVirtualSelectedParentBlueScoreResponseMessage)
    - [NotifyVirtualSelectedParentBlueScoreChangedRequestMessage](#protowire.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage)
    - [NotifyVirtualSelectedParentBlueScoreChangedResponseMessage](#protowire.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage)
    - [VirtualSelectedParentBlueScoreChangedNotificationMessage](#protowire.VirtualSelectedParentBlueScoreChangedNotificationMessage)
    - [NotifyVirtualDaaScoreChangedRequestMessage](#protowire.NotifyVirtualDaaScoreChangedRequestMessage)
    - [NotifyVirtualDaaScoreChangedResponseMessage](#protowire.NotifyVirtualDaaScoreChangedResponseMessage)
    - [VirtualDaaScoreChangedNotificationMessage](#protowire.VirtualDaaScoreChangedNotificationMessage)
    - [NotifyPruningPointUTXOSetOverrideRequestMessage](#protowire.NotifyPruningPointUTXOSetOverrideRequestMessage)
    - [NotifyPruningPointUTXOSetOverrideResponseMessage](#protowire.NotifyPruningPointUTXOSetOverrideResponseMessage)
    - [PruningPointUTXOSetOverrideNotificationMessage](#protowire.PruningPointUTXOSetOverrideNotificationMessage)
    - [StopNotifyingPruningPointUTXOSetOverrideRequestMessage](#protowire.StopNotifyingPruningPointUTXOSetOverrideRequestMessage)
    - [StopNotifyingPruningPointUTXOSetOverrideResponseMessage](#protowire.StopNotifyingPruningPointUTXOSetOverrideResponseMessage)
    - [BanRequestMessage](#protowire.BanRequestMessage)
    - [BanResponseMessage](#protowire.BanResponseMessage)
    - [UnbanRequestMessage](#protowire.UnbanRequestMessage)
    - [UnbanResponseMessage](#protowire.UnbanResponseMessage)
    - [GetInfoRequestMessage](#protowire.GetInfoRequestMessage)
    - [GetInfoResponseMessage](#protowire.GetInfoResponseMessage)
    - [EstimateNetworkHashesPerSecondRequestMessage](#protowire.EstimateNetworkHashesPerSecondRequestMessage)
    - [EstimateNetworkHashesPerSecondResponseMessage](#protowire.EstimateNetworkHashesPerSecondResponseMessage)
    - [NotifyNewBlockTemplateRequestMessage](#protowire.NotifyNewBlockTemplateRequestMessage)
    - [NotifyNewBlockTemplateResponseMessage](#protowire.NotifyNewBlockTemplateResponseMessage)
    - [NewBlockTemplateNotificationMessage](#protowire.NewBlockTemplateNotificationMessage)
    - [MempoolEntryByAddress](#protowire.MempoolEntryByAddress)
    - [GetMempoolEntriesByAddressesRequestMessage](#protowire.GetMempoolEntriesByAddressesRequestMessage)
    - [GetMempoolEntriesByAddressesResponseMessage](#protowire.GetMempoolEntriesByAddressesResponseMessage)
    - [GetCoinSupplyRequestMessage](#protowire.GetCoinSupplyRequestMessage)
    - [GetCoinSupplyResponseMessage](#protowire.GetCoinSupplyResponseMessage)
  
    - [SubmitBlockResponseMessage.RejectReason](#protowire.SubmitBlockResponseMessage.RejectReason)
  
- [Scalar Value Types](#scalar-value-types)



<a name="rpc.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rpc.proto
RPC-related types. Request messages, response messages, and dependant types.

Clients are expected to build RequestMessages and wrap them in WaglayladMessage. (see messages.proto)

Having received a RequestMessage, (wrapped in a WaglayladMessage) the RPC server will respond with a
ResponseMessage (likewise wrapped in a WaglayladMessage) respective to the original RequestMessage.

**IMPORTANT:** This API is a work in progress and is subject to break between versions.


<a name="protowire.RPCError"></a>

### RPCError
RPCError represents a generic non-internal error.

Receivers of any ResponseMessage are expected to check whether its error field is not null.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="protowire.RpcBlock"></a>

### RpcBlock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [RpcBlockHeader](#protowire.RpcBlockHeader) |  |  |
| transactions | [RpcTransaction](#protowire.RpcTransaction) | repeated |  |
| verboseData | [RpcBlockVerboseData](#protowire.RpcBlockVerboseData) |  |  |






<a name="protowire.RpcBlockHeader"></a>

### RpcBlockHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [uint32](#uint32) |  |  |
| parents | [RpcBlockLevelParents](#protowire.RpcBlockLevelParents) | repeated |  |
| hashMerkleRoot | [string](#string) |  |  |
| acceptedIdMerkleRoot | [string](#string) |  |  |
| utxoCommitment | [string](#string) |  |  |
| timestamp | [int64](#int64) |  |  |
| bits | [uint32](#uint32) |  |  |
| nonce | [uint64](#uint64) |  |  |
| daaScore | [uint64](#uint64) |  |  |
| blueWork | [string](#string) |  |  |
| pruningPoint | [string](#string) |  |  |
| blueScore | [uint64](#uint64) |  |  |






<a name="protowire.RpcBlockLevelParents"></a>

### RpcBlockLevelParents



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parentHashes | [string](#string) | repeated |  |






<a name="protowire.RpcBlockVerboseData"></a>

### RpcBlockVerboseData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  |  |
| difficulty | [double](#double) |  |  |
| selectedParentHash | [string](#string) |  |  |
| transactionIds | [string](#string) | repeated |  |
| isHeaderOnly | [bool](#bool) |  |  |
| blueScore | [uint64](#uint64) |  |  |
| childrenHashes | [string](#string) | repeated |  |
| mergeSetBluesHashes | [string](#string) | repeated |  |
| mergeSetRedsHashes | [string](#string) | repeated |  |
| isChainBlock | [bool](#bool) |  |  |






<a name="protowire.RpcTransaction"></a>

### RpcTransaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [uint32](#uint32) |  |  |
| inputs | [RpcTransactionInput](#protowire.RpcTransactionInput) | repeated |  |
| outputs | [RpcTransactionOutput](#protowire.RpcTransactionOutput) | repeated |  |
| lockTime | [uint64](#uint64) |  |  |
| subnetworkId | [string](#string) |  |  |
| gas | [uint64](#uint64) |  |  |
| payload | [string](#string) |  |  |
| verboseData | [RpcTransactionVerboseData](#protowire.RpcTransactionVerboseData) |  |  |






<a name="protowire.RpcTransactionInput"></a>

### RpcTransactionInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| previousOutpoint | [RpcOutpoint](#protowire.RpcOutpoint) |  |  |
| signatureScript | [string](#string) |  |  |
| sequence | [uint64](#uint64) |  |  |
| sigOpCount | [uint32](#uint32) |  |  |
| verboseData | [RpcTransactionInputVerboseData](#protowire.RpcTransactionInputVerboseData) |  |  |






<a name="protowire.RpcScriptPublicKey"></a>

### RpcScriptPublicKey



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [uint32](#uint32) |  |  |
| scriptPublicKey | [string](#string) |  |  |






<a name="protowire.RpcTransactionOutput"></a>

### RpcTransactionOutput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [uint64](#uint64) |  |  |
| scriptPublicKey | [RpcScriptPublicKey](#protowire.RpcScriptPublicKey) |  |  |
| verboseData | [RpcTransactionOutputVerboseData](#protowire.RpcTransactionOutputVerboseData) |  |  |






<a name="protowire.RpcOutpoint"></a>

### RpcOutpoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactionId | [string](#string) |  |  |
| index | [uint32](#uint32) |  |  |






<a name="protowire.RpcUtxoEntry"></a>

### RpcUtxoEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [uint64](#uint64) |  |  |
| scriptPublicKey | [RpcScriptPublicKey](#protowire.RpcScriptPublicKey) |  |  |
| blockDaaScore | [uint64](#uint64) |  |  |
| isCoinbase | [bool](#bool) |  |  |






<a name="protowire.RpcTransactionVerboseData"></a>

### RpcTransactionVerboseData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactionId | [string](#string) |  |  |
| hash | [string](#string) |  |  |
| mass | [uint64](#uint64) |  |  |
| blockHash | [string](#string) |  |  |
| blockTime | [uint64](#uint64) |  |  |






<a name="protowire.RpcTransactionInputVerboseData"></a>

### RpcTransactionInputVerboseData







<a name="protowire.RpcTransactionOutputVerboseData"></a>

### RpcTransactionOutputVerboseData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scriptPublicKeyType | [string](#string) |  |  |
| scriptPublicKeyAddress | [string](#string) |  |  |






<a name="protowire.GetCurrentNetworkRequestMessage"></a>

### GetCurrentNetworkRequestMessage
GetCurrentNetworkRequestMessage requests the network waglaylad is currently running against.

Possible networks are: Mainnet, Testnet, Simnet, Devnet






<a name="protowire.GetCurrentNetworkResponseMessage"></a>

### GetCurrentNetworkResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentNetwork | [string](#string) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.SubmitBlockRequestMessage"></a>

### SubmitBlockRequestMessage
SubmitBlockRequestMessage requests to submit a block into the DAG.
Blocks are generally expected to have been generated using the getBlockTemplate call.

See: GetBlockTemplateRequestMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block | [RpcBlock](#protowire.RpcBlock) |  |  |
| allowNonDAABlocks | [bool](#bool) |  |  |






<a name="protowire.SubmitBlockResponseMessage"></a>

### SubmitBlockResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rejectReason | [SubmitBlockResponseMessage.RejectReason](#protowire.SubmitBlockResponseMessage.RejectReason) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetBlockTemplateRequestMessage"></a>

### GetBlockTemplateRequestMessage
GetBlockTemplateRequestMessage requests a current block template.
Callers are expected to solve the block template and submit it using the submitBlock call

See: SubmitBlockRequestMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payAddress | [string](#string) |  | Which waglayla address should the coinbase block reward transaction pay into |
| extraData | [string](#string) |  |  |






<a name="protowire.GetBlockTemplateResponseMessage"></a>

### GetBlockTemplateResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block | [RpcBlock](#protowire.RpcBlock) |  |  |
| isSynced | [bool](#bool) |  | Whether waglaylad thinks that it&#39;s synced. Callers are discouraged (but not forbidden) from solving blocks when waglaylad is not synced. That is because when waglaylad isn&#39;t in sync with the rest of the network there&#39;s a high chance the block will never be accepted, thus the solving effort would have been wasted. |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.NotifyBlockAddedRequestMessage"></a>

### NotifyBlockAddedRequestMessage
NotifyBlockAddedRequestMessage registers this connection for blockAdded notifications.

See: BlockAddedNotificationMessage






<a name="protowire.NotifyBlockAddedResponseMessage"></a>

### NotifyBlockAddedResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.BlockAddedNotificationMessage"></a>

### BlockAddedNotificationMessage
BlockAddedNotificationMessage is sent whenever a blocks has been added (NOT accepted)
into the DAG.

See: NotifyBlockAddedRequestMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block | [RpcBlock](#protowire.RpcBlock) |  |  |






<a name="protowire.GetPeerAddressesRequestMessage"></a>

### GetPeerAddressesRequestMessage
GetPeerAddressesRequestMessage requests the list of known waglaylad addresses in the
current network. (mainnet, testnet, etc.)






<a name="protowire.GetPeerAddressesResponseMessage"></a>

### GetPeerAddressesResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [GetPeerAddressesKnownAddressMessage](#protowire.GetPeerAddressesKnownAddressMessage) | repeated |  |
| bannedAddresses | [GetPeerAddressesKnownAddressMessage](#protowire.GetPeerAddressesKnownAddressMessage) | repeated |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetPeerAddressesKnownAddressMessage"></a>

### GetPeerAddressesKnownAddressMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Addr | [string](#string) |  |  |






<a name="protowire.GetSelectedTipHashRequestMessage"></a>

### GetSelectedTipHashRequestMessage
GetSelectedTipHashRequestMessage requests the hash of the current virtual&#39;s
selected parent.






<a name="protowire.GetSelectedTipHashResponseMessage"></a>

### GetSelectedTipHashResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selectedTipHash | [string](#string) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetMempoolEntryRequestMessage"></a>

### GetMempoolEntryRequestMessage
GetMempoolEntryRequestMessage requests information about a specific transaction
in the mempool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| txId | [string](#string) |  | The transaction&#39;s TransactionID. |
| includeOrphanPool | [bool](#bool) |  |  |
| filterTransactionPool | [bool](#bool) |  |  |






<a name="protowire.GetMempoolEntryResponseMessage"></a>

### GetMempoolEntryResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [MempoolEntry](#protowire.MempoolEntry) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetMempoolEntriesRequestMessage"></a>

### GetMempoolEntriesRequestMessage
GetMempoolEntriesRequestMessage requests information about all the transactions
currently in the mempool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| includeOrphanPool | [bool](#bool) |  |  |
| filterTransactionPool | [bool](#bool) |  |  |






<a name="protowire.GetMempoolEntriesResponseMessage"></a>

### GetMempoolEntriesResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [MempoolEntry](#protowire.MempoolEntry) | repeated |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.MempoolEntry"></a>

### MempoolEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fee | [uint64](#uint64) |  |  |
| transaction | [RpcTransaction](#protowire.RpcTransaction) |  |  |
| isOrphan | [bool](#bool) |  |  |






<a name="protowire.GetConnectedPeerInfoRequestMessage"></a>

### GetConnectedPeerInfoRequestMessage
GetConnectedPeerInfoRequestMessage requests information about all the p2p peers
currently connected to this waglaylad.






<a name="protowire.GetConnectedPeerInfoResponseMessage"></a>

### GetConnectedPeerInfoResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| infos | [GetConnectedPeerInfoMessage](#protowire.GetConnectedPeerInfoMessage) | repeated |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetConnectedPeerInfoMessage"></a>

### GetConnectedPeerInfoMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| address | [string](#string) |  |  |
| lastPingDuration | [int64](#int64) |  | How long did the last ping/pong exchange take |
| isOutbound | [bool](#bool) |  | Whether this waglaylad initiated the connection |
| timeOffset | [int64](#int64) |  |  |
| userAgent | [string](#string) |  |  |
| advertisedProtocolVersion | [uint32](#uint32) |  | The protocol version that this peer claims to support |
| timeConnected | [int64](#int64) |  | The timestamp of when this peer connected to this waglaylad |
| isIbdPeer | [bool](#bool) |  | Whether this peer is the IBD peer (if IBD is running) |






<a name="protowire.AddPeerRequestMessage"></a>

### AddPeerRequestMessage
AddPeerRequestMessage adds a peer to waglaylad&#39;s outgoing connection list.
This will, in most cases, result in waglaylad connecting to said peer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| isPermanent | [bool](#bool) |  | Whether to keep attempting to connect to this peer after disconnection |






<a name="protowire.AddPeerResponseMessage"></a>

### AddPeerResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.SubmitTransactionRequestMessage"></a>

### SubmitTransactionRequestMessage
SubmitTransactionRequestMessage submits a transaction to the mempool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction | [RpcTransaction](#protowire.RpcTransaction) |  |  |
| allowOrphan | [bool](#bool) |  |  |






<a name="protowire.SubmitTransactionResponseMessage"></a>

### SubmitTransactionResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactionId | [string](#string) |  | The transaction ID of the submitted transaction |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.NotifyVirtualSelectedParentChainChangedRequestMessage"></a>

### NotifyVirtualSelectedParentChainChangedRequestMessage
NotifyVirtualSelectedParentChainChangedRequestMessage registers this connection for virtualSelectedParentChainChanged notifications.

See: VirtualSelectedParentChainChangedNotificationMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| includeAcceptedTransactionIds | [bool](#bool) |  |  |






<a name="protowire.NotifyVirtualSelectedParentChainChangedResponseMessage"></a>

### NotifyVirtualSelectedParentChainChangedResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.VirtualSelectedParentChainChangedNotificationMessage"></a>

### VirtualSelectedParentChainChangedNotificationMessage
VirtualSelectedParentChainChangedNotificationMessage is sent whenever the DAG&#39;s selected parent
chain had changed.

See: NotifyVirtualSelectedParentChainChangedRequestMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| removedChainBlockHashes | [string](#string) | repeated | The chain blocks that were removed, in high-to-low order |
| addedChainBlockHashes | [string](#string) | repeated | The chain blocks that were added, in low-to-high order |
| acceptedTransactionIds | [AcceptedTransactionIds](#protowire.AcceptedTransactionIds) | repeated | Will be filled only if `includeAcceptedTransactionIds = true` in the notify request. |






<a name="protowire.GetBlockRequestMessage"></a>

### GetBlockRequestMessage
GetBlockRequestMessage requests information about a specific block


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  | The hash of the requested block |
| includeTransactions | [bool](#bool) |  | Whether to include transaction data in the response |






<a name="protowire.GetBlockResponseMessage"></a>

### GetBlockResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block | [RpcBlock](#protowire.RpcBlock) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetSubnetworkRequestMessage"></a>

### GetSubnetworkRequestMessage
GetSubnetworkRequestMessage requests information about a specific subnetwork

Currently unimplemented


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subnetworkId | [string](#string) |  |  |






<a name="protowire.GetSubnetworkResponseMessage"></a>

### GetSubnetworkResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gasLimit | [uint64](#uint64) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetVirtualSelectedParentChainFromBlockRequestMessage"></a>

### GetVirtualSelectedParentChainFromBlockRequestMessage
GetVirtualSelectedParentChainFromBlockRequestMessage requests the virtual selected
parent chain from some startHash to this waglaylad&#39;s current virtual


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startHash | [string](#string) |  |  |
| includeAcceptedTransactionIds | [bool](#bool) |  |  |






<a name="protowire.AcceptedTransactionIds"></a>

### AcceptedTransactionIds



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| acceptingBlockHash | [string](#string) |  |  |
| acceptedTransactionIds | [string](#string) | repeated |  |






<a name="protowire.GetVirtualSelectedParentChainFromBlockResponseMessage"></a>

### GetVirtualSelectedParentChainFromBlockResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| removedChainBlockHashes | [string](#string) | repeated | The chain blocks that were removed, in high-to-low order |
| addedChainBlockHashes | [string](#string) | repeated | The chain blocks that were added, in low-to-high order |
| acceptedTransactionIds | [AcceptedTransactionIds](#protowire.AcceptedTransactionIds) | repeated | The transactions accepted by each block in addedChainBlockHashes. Will be filled only if `includeAcceptedTransactionIds = true` in the request. |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetBlocksRequestMessage"></a>

### GetBlocksRequestMessage
GetBlocksRequestMessage requests blocks between a certain block lowHash up to this
waglaylad&#39;s current virtual.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lowHash | [string](#string) |  |  |
| includeBlocks | [bool](#bool) |  |  |
| includeTransactions | [bool](#bool) |  |  |






<a name="protowire.GetBlocksResponseMessage"></a>

### GetBlocksResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blockHashes | [string](#string) | repeated |  |
| blocks | [RpcBlock](#protowire.RpcBlock) | repeated |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetBlockCountRequestMessage"></a>

### GetBlockCountRequestMessage
GetBlockCountRequestMessage requests the current number of blocks in this waglaylad.
Note that this number may decrease as pruning occurs.






<a name="protowire.GetBlockCountResponseMessage"></a>

### GetBlockCountResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blockCount | [uint64](#uint64) |  |  |
| headerCount | [uint64](#uint64) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetBlockDagInfoRequestMessage"></a>

### GetBlockDagInfoRequestMessage
GetBlockDagInfoRequestMessage requests general information about the current state
of this waglaylad&#39;s DAG.






<a name="protowire.GetBlockDagInfoResponseMessage"></a>

### GetBlockDagInfoResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| networkName | [string](#string) |  |  |
| blockCount | [uint64](#uint64) |  |  |
| headerCount | [uint64](#uint64) |  |  |
| tipHashes | [string](#string) | repeated |  |
| difficulty | [double](#double) |  |  |
| pastMedianTime | [int64](#int64) |  |  |
| virtualParentHashes | [string](#string) | repeated |  |
| pruningPointHash | [string](#string) |  |  |
| virtualDaaScore | [uint64](#uint64) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.ResolveFinalityConflictRequestMessage"></a>

### ResolveFinalityConflictRequestMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| finalityBlockHash | [string](#string) |  |  |






<a name="protowire.ResolveFinalityConflictResponseMessage"></a>

### ResolveFinalityConflictResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.NotifyFinalityConflictsRequestMessage"></a>

### NotifyFinalityConflictsRequestMessage







<a name="protowire.NotifyFinalityConflictsResponseMessage"></a>

### NotifyFinalityConflictsResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.FinalityConflictNotificationMessage"></a>

### FinalityConflictNotificationMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| violatingBlockHash | [string](#string) |  |  |






<a name="protowire.FinalityConflictResolvedNotificationMessage"></a>

### FinalityConflictResolvedNotificationMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| finalityBlockHash | [string](#string) |  |  |






<a name="protowire.ShutDownRequestMessage"></a>

### ShutDownRequestMessage
ShutDownRequestMessage shuts down this waglaylad.






<a name="protowire.ShutDownResponseMessage"></a>

### ShutDownResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetHeadersRequestMessage"></a>

### GetHeadersRequestMessage
GetHeadersRequestMessage requests headers between the given startHash and the
current virtual, up to the given limit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startHash | [string](#string) |  |  |
| limit | [uint64](#uint64) |  |  |
| isAscending | [bool](#bool) |  |  |






<a name="protowire.GetHeadersResponseMessage"></a>

### GetHeadersResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [string](#string) | repeated |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.NotifyUtxosChangedRequestMessage"></a>

### NotifyUtxosChangedRequestMessage
NotifyUtxosChangedRequestMessage registers this connection for utxoChanged notifications
for the given addresses.

This call is only available when this waglaylad was started with `--utxoindex`

See: UtxosChangedNotificationMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [string](#string) | repeated | Leave empty to get all updates |






<a name="protowire.NotifyUtxosChangedResponseMessage"></a>

### NotifyUtxosChangedResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.UtxosChangedNotificationMessage"></a>

### UtxosChangedNotificationMessage
UtxosChangedNotificationMessage is sent whenever the UTXO index had been updated.

See: NotifyUtxosChangedRequestMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| added | [UtxosByAddressesEntry](#protowire.UtxosByAddressesEntry) | repeated |  |
| removed | [UtxosByAddressesEntry](#protowire.UtxosByAddressesEntry) | repeated |  |






<a name="protowire.UtxosByAddressesEntry"></a>

### UtxosByAddressesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| outpoint | [RpcOutpoint](#protowire.RpcOutpoint) |  |  |
| utxoEntry | [RpcUtxoEntry](#protowire.RpcUtxoEntry) |  |  |






<a name="protowire.StopNotifyingUtxosChangedRequestMessage"></a>

### StopNotifyingUtxosChangedRequestMessage
StopNotifyingUtxosChangedRequestMessage unregisters this connection for utxoChanged notifications
for the given addresses.

This call is only available when this waglaylad was started with `--utxoindex`

See: UtxosChangedNotificationMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [string](#string) | repeated |  |






<a name="protowire.StopNotifyingUtxosChangedResponseMessage"></a>

### StopNotifyingUtxosChangedResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetUtxosByAddressesRequestMessage"></a>

### GetUtxosByAddressesRequestMessage
GetUtxosByAddressesRequestMessage requests all current UTXOs for the given waglaylad addresses

This call is only available when this waglaylad was started with `--utxoindex`


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [string](#string) | repeated |  |






<a name="protowire.GetUtxosByAddressesResponseMessage"></a>

### GetUtxosByAddressesResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [UtxosByAddressesEntry](#protowire.UtxosByAddressesEntry) | repeated |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetBalanceByAddressRequestMessage"></a>

### GetBalanceByAddressRequestMessage
GetBalanceByAddressRequest returns the total balance in unspent transactions towards a given address

This call is only available when this waglaylad was started with `--utxoindex`


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="protowire.GetBalanceByAddressResponseMessage"></a>

### GetBalanceByAddressResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [uint64](#uint64) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetBalancesByAddressesRequestMessage"></a>

### GetBalancesByAddressesRequestMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [string](#string) | repeated |  |






<a name="protowire.BalancesByAddressEntry"></a>

### BalancesByAddressEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| balance | [uint64](#uint64) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetBalancesByAddressesResponseMessage"></a>

### GetBalancesByAddressesResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [BalancesByAddressEntry](#protowire.BalancesByAddressEntry) | repeated |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetVirtualSelectedParentBlueScoreRequestMessage"></a>

### GetVirtualSelectedParentBlueScoreRequestMessage
GetVirtualSelectedParentBlueScoreRequestMessage requests the blue score of the current selected parent
of the virtual block.






<a name="protowire.GetVirtualSelectedParentBlueScoreResponseMessage"></a>

### GetVirtualSelectedParentBlueScoreResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blueScore | [uint64](#uint64) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage"></a>

### NotifyVirtualSelectedParentBlueScoreChangedRequestMessage
NotifyVirtualSelectedParentBlueScoreChangedRequestMessage registers this connection for
virtualSelectedParentBlueScoreChanged notifications.

See: VirtualSelectedParentBlueScoreChangedNotificationMessage






<a name="protowire.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage"></a>

### NotifyVirtualSelectedParentBlueScoreChangedResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.VirtualSelectedParentBlueScoreChangedNotificationMessage"></a>

### VirtualSelectedParentBlueScoreChangedNotificationMessage
VirtualSelectedParentBlueScoreChangedNotificationMessage is sent whenever the blue score
of the virtual&#39;s selected parent changes.

See NotifyVirtualSelectedParentBlueScoreChangedRequestMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtualSelectedParentBlueScore | [uint64](#uint64) |  |  |






<a name="protowire.NotifyVirtualDaaScoreChangedRequestMessage"></a>

### NotifyVirtualDaaScoreChangedRequestMessage
NotifyVirtualDaaScoreChangedRequestMessage registers this connection for
virtualDaaScoreChanged notifications.

See: VirtualDaaScoreChangedNotificationMessage






<a name="protowire.NotifyVirtualDaaScoreChangedResponseMessage"></a>

### NotifyVirtualDaaScoreChangedResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.VirtualDaaScoreChangedNotificationMessage"></a>

### VirtualDaaScoreChangedNotificationMessage
VirtualDaaScoreChangedNotificationMessage is sent whenever the DAA score
of the virtual changes.

See NotifyVirtualDaaScoreChangedRequestMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| virtualDaaScore | [uint64](#uint64) |  |  |






<a name="protowire.NotifyPruningPointUTXOSetOverrideRequestMessage"></a>

### NotifyPruningPointUTXOSetOverrideRequestMessage
NotifyPruningPointUTXOSetOverrideRequestMessage registers this connection for
pruning point UTXO set override notifications.

This call is only available when this waglaylad was started with `--utxoindex`

See: NotifyPruningPointUTXOSetOverrideResponseMessage






<a name="protowire.NotifyPruningPointUTXOSetOverrideResponseMessage"></a>

### NotifyPruningPointUTXOSetOverrideResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.PruningPointUTXOSetOverrideNotificationMessage"></a>

### PruningPointUTXOSetOverrideNotificationMessage
PruningPointUTXOSetOverrideNotificationMessage is sent whenever the UTXO index
resets due to pruning point change via IBD.

See NotifyPruningPointUTXOSetOverrideRequestMessage






<a name="protowire.StopNotifyingPruningPointUTXOSetOverrideRequestMessage"></a>

### StopNotifyingPruningPointUTXOSetOverrideRequestMessage
StopNotifyingPruningPointUTXOSetOverrideRequestMessage unregisters this connection for
pruning point UTXO set override notifications.

This call is only available when this waglaylad was started with `--utxoindex`

See: PruningPointUTXOSetOverrideNotificationMessage






<a name="protowire.StopNotifyingPruningPointUTXOSetOverrideResponseMessage"></a>

### StopNotifyingPruningPointUTXOSetOverrideResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.BanRequestMessage"></a>

### BanRequestMessage
BanRequestMessage bans the given ip.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip | [string](#string) |  |  |






<a name="protowire.BanResponseMessage"></a>

### BanResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.UnbanRequestMessage"></a>

### UnbanRequestMessage
UnbanRequestMessage unbans the given ip.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip | [string](#string) |  |  |






<a name="protowire.UnbanResponseMessage"></a>

### UnbanResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetInfoRequestMessage"></a>

### GetInfoRequestMessage
GetInfoRequestMessage returns info about the node.






<a name="protowire.GetInfoResponseMessage"></a>

### GetInfoResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| p2pId | [string](#string) |  |  |
| mempoolSize | [uint64](#uint64) |  |  |
| serverVersion | [string](#string) |  |  |
| isUtxoIndexed | [bool](#bool) |  |  |
| isSynced | [bool](#bool) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.EstimateNetworkHashesPerSecondRequestMessage"></a>

### EstimateNetworkHashesPerSecondRequestMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| windowSize | [uint32](#uint32) |  |  |
| startHash | [string](#string) |  |  |






<a name="protowire.EstimateNetworkHashesPerSecondResponseMessage"></a>

### EstimateNetworkHashesPerSecondResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| networkHashesPerSecond | [uint64](#uint64) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.NotifyNewBlockTemplateRequestMessage"></a>

### NotifyNewBlockTemplateRequestMessage
NotifyNewBlockTemplateRequestMessage registers this connection for
NewBlockTemplate notifications.

See: NewBlockTemplateNotificationMessage






<a name="protowire.NotifyNewBlockTemplateResponseMessage"></a>

### NotifyNewBlockTemplateResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.NewBlockTemplateNotificationMessage"></a>

### NewBlockTemplateNotificationMessage
NewBlockTemplateNotificationMessage is sent whenever a new updated block template is
available for miners.

See NotifyNewBlockTemplateRequestMessage






<a name="protowire.MempoolEntryByAddress"></a>

### MempoolEntryByAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| sending | [MempoolEntry](#protowire.MempoolEntry) | repeated |  |
| receiving | [MempoolEntry](#protowire.MempoolEntry) | repeated |  |






<a name="protowire.GetMempoolEntriesByAddressesRequestMessage"></a>

### GetMempoolEntriesByAddressesRequestMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [string](#string) | repeated |  |
| includeOrphanPool | [bool](#bool) |  |  |
| filterTransactionPool | [bool](#bool) |  |  |






<a name="protowire.GetMempoolEntriesByAddressesResponseMessage"></a>

### GetMempoolEntriesByAddressesResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [MempoolEntryByAddress](#protowire.MempoolEntryByAddress) | repeated |  |
| error | [RPCError](#protowire.RPCError) |  |  |






<a name="protowire.GetCoinSupplyRequestMessage"></a>

### GetCoinSupplyRequestMessage







<a name="protowire.GetCoinSupplyResponseMessage"></a>

### GetCoinSupplyResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| maxLeor | [uint64](#uint64) |  | note: this is a hard coded maxSupply, actual maxSupply is expected to deviate by upto -5%, but cannot be measured exactly. |
| circulatingLeor | [uint64](#uint64) |  |  |
| error | [RPCError](#protowire.RPCError) |  |  |





 


<a name="protowire.SubmitBlockResponseMessage.RejectReason"></a>

### SubmitBlockResponseMessage.RejectReason


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| BLOCK_INVALID | 1 |  |
| IS_IN_IBD | 2 |  |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

