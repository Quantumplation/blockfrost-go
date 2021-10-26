package blockfrost

import "fmt"

const (
	CardanoMainNet = "https://cardano-mainnet.blockfrost.io/api/v0"
	CardanoTestNet = "https://cardano-testnet.blockfrost.io/api/v0"
	IPFSNet        = "https://ipfs.blockfrost.io/api/v0"
)

// resource paths
const (
	resourceHealth                   = "health"
	resourceHealthClock              = "health/clock"
	resourceMetrics                  = "metrics"
	resourceMetricsEndpoint          = "metrics/endpoints"
	resourceBlock                    = "blocks"
	resourceBlocksLatest             = "blocks/latest"
	resourceBlocksLatestTransactions = "blocks/latest/txs"
	resourceBlocksSlot               = "blocks/slot"
	resourceBlocksEpoch              = "blocks/epoch"
)

// APIError is used to describe errors from the API.
// See https://docs.blockfrost.io/#section/Errors
type APIError struct {
	Response interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error, %+v", e.Response)
}

// Autobanned defines model for HTTP `418` (Auto Banned).
type AutoBanned struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// BadRequest defines model for HTTP `400` (Bad Request)
type BadRequest struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// InternalServerError defines model for HTTP `500` (Internal Server Error)
type InternalServerError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// NotFound defines model for HTTP `404` (Resource Not Found).
type NotFound struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// OverusageLimit defines model for HTTP `429` (Over Usage).
type OverusageLimit struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// UnauthorizedEror defines model for HTTP `403` (Unauthorized).
type UnauthorizedError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// Block defines content of a block
type Block struct {
	Time          int    `json:"time,omitempty"`
	Height        int    `json:"height,omitempty"`
	Hash          string `json:"hash,omitempty"`
	Slot          int    `json:"slot,omitempty"`
	Epoch         int    `json:"epoch,omitempty"`
	EpochSlot     int    `json:"epoch_slot,omitempty"`
	SlotLeader    string `json:"slot_leader"`
	Size          int    `json:"size,omitempty"`
	TxCount       int    `json:"tx_count,omitempty"`
	Output        string `json:"output,omitempty"`
	Fees          string `json:"fees,omitempty"`
	BlockVRF      string `json:"block_vrf,omitempty"`
	PreviousBlock string `json:"previous_block,omitempty"`
	NextBlock     string `json:"next_block,omitempty"`
	Confirmations int    `json:"confirmations,omitempty"`
}

type Transaction string

// APIQueryParams contains query parameters. Marshalled to
// "count", "page", "order", "from", "to".
type APIQueryParams struct {
	Count int
	Page  int
	Order string
	From  string
	To    string
}
