// h 20180923
//
// Ethereum RPC Client

package eth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Client struct {
	url        string
	httpClient *http.Client
	id         int64
	idLock     sync.Mutex
}

func NewClient(url string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{url: url, httpClient: httpClient}
}

func (c *Client) Web3ClientVersion() (string, error) {
	v := ""
	e := c.CallMethod(&v, "web3_clientVersion")
	return v, e
}

func (c *Client) Web3Sha3(data string) (common.Hash, error) {
	v := common.Hash{}
	e := c.CallMethod(&v, "web3_sha3", data)
	return v, e
}

func (c *Client) NetVersion() (string, error) {
	v := ""
	e := c.CallMethod(&v, "net_version")
	return v, e
}

func (c *Client) NetListening() (bool, error) {
	v := false
	e := c.CallMethod(&v, "net_listening")
	return v, e
}

func (c *Client) NetPeerCount() (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "net_peerCount")
	return &v, e
}

func (c *Client) EthProtocolVersion() (string, error) {
	v := ""
	e := c.CallMethod(&v, "eth_protocolVersion")
	return v, e
}

func (c *Client) EthSyncing() (bool, error) {
	v := false
	e := c.CallMethod(&v, "eth_syncing")
	return v, e
}

func (c *Client) EthMining() (bool, error) {
	v := false
	e := c.CallMethod(&v, "eth_mining")
	return v, e
}

func (c *Client) EthCoinbase() (common.Address, error) {
	v := common.Address{}
	e := c.CallMethod(&v, "eth_coinbase")
	return v, e
}

func (c *Client) EthAccounts() ([]common.Address, error) {
	v := ([]common.Address)(nil)
	e := c.CallMethod(&v, "eth_accounts")
	return v, e
}

func (c *Client) EthHashrate() (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_hashrate")
	return &v, e
}

func (c *Client) EthSubmitHashrate(hashrate, id common.Hash) (bool, error) {
	v := false
	e := c.CallMethod(&v, "eth_submitHashrate", hashrate, id)
	return v, e
}

func (c *Client) EthGetWork() ([]string, error) {
	v := []string(nil)
	e := c.CallMethod(&v, "eth_getWork")
	return v, e
}

func (c *Client) EthSubmitWork(nonce [8]byte, header, mix common.Hash) (bool, error) {
	v := false
	e := c.CallMethod(&v, "eth_submitWork", hexutil.Bytes(nonce[:]), header, mix)
	return v, e
}

func (c *Client) EthGasPrice() (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_gasPrice")
	return &v, e
}

type EstimateGasRequest struct {
	To string `json:"to,omitempty"`
}

var (
	EstimateTransactionGasRequest = &EstimateGasRequest{To: "0x0000000000000000000000000000000000000000"}
	EstimateContractGasRequest    = &EstimateGasRequest{To: ""}
)

func (c *Client) EthEstimateGas(req *EstimateGasRequest) (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_estimateGas", req)
	return &v, e
}

func (c *Client) EthGetBalance(addr, block string) (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_getBalance", addr, block)
	return &v, e
}

func (c *Client) EthSign(addr, msg string) ([]byte, error) {
	v := hexutil.Bytes{}
	e := c.CallMethod(&v, "eth_sign", addr, msg)
	return v, e
}

func (c *Client) EthGetCode(addr, block string) ([]byte, error) {
	v := hexutil.Bytes{}
	e := c.CallMethod(&v, "eth_getCode", addr, block)
	return v, e
}

func (c *Client) EthGetStorageAt(addr, pos, block string) ([]byte, error) {
	v := hexutil.Bytes{}
	e := c.CallMethod(&v, "eth_getStorageAt", addr, pos, block)
	return v, e
}

func (c *Client) EthBlockNumber() (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_blockNumber")
	return &v, e
}

func (c *Client) EthGetBlockByNumber(number string, full bool) ([]byte, error) {
	v := json.RawMessage{}
	e := c.CallMethod(&v, "eth_getBlockByNumber", number, full)
	return v, e
}

func (c *Client) EthGetBlockByHash(hash string, full bool) ([]byte, error) {
	v := json.RawMessage{}
	e := c.CallMethod(&v, "eth_getBlockByHash", hash, full)
	return v, e
}

func (c *Client) EthGetUncleCountByNumber(block *hexutil.Big) (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_getUncleCountByNumber", block)
	return &v, e
}

func (c *Client) EthGetUncleCountByHash(hash common.Hash) (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_getUncleCountByHash", hash)
	return &v, e
}

// func (c *Client) EthGetUncleByBlockNumberAndIndex(block BlockNumber, idx big.Int) (*Block, error) {}
// func (c *Client) EthGetUncleByBlockHashAndIndex(hash [32]byte, idx big.Int) (*Block, error) {}

func (c *Client) EthGetBlockTransactionCountByNumber(block string) (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_getBlockTransactionCountByNumber", block)
	return &v, e
}

func (c *Client) EthGetBlockTransactionCountByHash(hash string) (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_getBlockTransactionCountByHash", hash)
	return &v, e
}

func (c *Client) EthGetTransactionCount(addr, block string) (*hexutil.Big, error) {
	v := hexutil.Big{}
	e := c.CallMethod(&v, "eth_getTransactionCount", addr, block)
	return &v, e
}

func (c *Client) EthGetTransactionByBlockNumberAndIndex(blk, idx string) ([]byte, error) {
	v := json.RawMessage{}
	e := c.CallMethod(&v, "eth_getTransactionByBlockNumberAndIndex", blk, idx)
	return v, e
}

func (c *Client) EthGetTransactionByBlockHashAndIndex(blk, idx string) ([]byte, error) {
	v := json.RawMessage{}
	e := c.CallMethod(&v, "eth_getTransactionByBlockHashAndIndex", blk, idx)
	return v, e
}

func (c *Client) EthGetTransactionByHash(txn string) ([]byte, error) {
	v := json.RawMessage{}
	e := c.CallMethod(&v, "eth_getTransactionByHash", txn)
	return v, e
}

type Topic struct {
	Data []byte
}

type Topics []Topic

type Log struct {
	LogIndex         uint64         `json:"logIndex"`
	BlockNumber      *hexutil.Big   `json:"blockNumber"`
	BlockHash        common.Hash    `json:"blockHash"`
	TransactionHash  common.Hash    `json:"transactionHash"`
	TransactionIndex uint64         `json:"transactionIndex"`
	Address          common.Address `json:"address"`
	Data             []byte         `json:"data"`
	Topics           Topics         `json:"topics"`
}

type TransactionReceipt struct {
	BlockHash         common.Hash    `json:"blockHash"`
	BlockNumber       *hexutil.Big   `json:"blockNumber"`
	ContractAddress   common.Address `json:"contractAddress"`
	CumulativeGasUsed *hexutil.Big   `json:"cumulativeGasUsed"`
	From              common.Address `json:"from"`
	GasUsed           *hexutil.Big   `json:"gasUsed"`
	Logs              []Log          `json:"logs"`
	LogsBloom         [512]byte      `json:"logsBloom"`
	Root              common.Hash    `json:"root"`
	To                common.Address `json:"to"`
	Hash              common.Hash    `json:"transactionHash"`
	TransactionIndex  uint64         `json:"transactionIndex"`
}

func (c *Client) EthGetTransactionReceipt(txn string) ([]byte, error) {
	v := json.RawMessage{}
	e := c.CallMethod(&v, "eth_getTransactionReceipt", txn)
	return v, e
}

type TransactionRequest struct {
	From     common.Address  `json:"from,omitempty"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Big    `json:"gas,omitempty"`
	GasPrice *hexutil.Big    `json:"gasPrice,omitempty"`
	Value    *hexutil.Big    `json:"value"`
	Data     hexutil.Bytes   `json:"data,omitempty"`
	Nonce    *hexutil.Uint64 `json:"nonce,omitempty"`
}

func (c *Client) EthSendTransaction(req *TransactionRequest) (common.Hash, error) {
	v := common.Hash{}
	e := c.CallMethod(&v, "eth_sendTransaction", req)
	return v, e
}

func (c *Client) EthSendRawTransaction(raw string) (common.Hash, error) {
	v := common.Hash{}
	e := c.CallMethod(&v, "eth_sendRawTransaction", raw)
	return v, e
}

// func (c *Client) EthCall() ([]byte, error) {}
// Executes a new message call immediately without creating a transaction on the block chain

func (c *Client) EthGetCompilers() ([]string, error) {
	v := []string(nil)
	e := c.CallMethod(&v, "eth_getCompilers")
	return v, e
}

func (c *Client) EthCompileSolidity(code string) ([]byte, error) {
	v := hexutil.Bytes{}
	e := c.CallMethod(&v, "eth_compileSolidity", code)
	return v, e
}

func (c *Client) EthCompileLLL(code string) ([]byte, error) {
	v := hexutil.Bytes{}
	e := c.CallMethod(&v, "eth_compileLLL", code)
	return v, e
}

func (c *Client) EthCompileSerpent(code string) ([]byte, error) {
	v := hexutil.Bytes{}
	e := c.CallMethod(&v, "eth_compileSerpent", code)
	return v, e
}

// func (c *Client) EthNewFilter(fromBlock BlockNumber, toBlock BlockNumber, addrs [][20]byte, topics []byte) (big.Int, error) {}
// func (c *Client) EthNewBlockFilter() (big.Int, error) {}
// func (c *Client) EthNewPendingTransactionFilter() (big.Int, error) {}
// func (c *Client) EthUninstallFilter(id big.Int) (bool, error) {}
// func (c *Client) EthGetFilterChanges(id big.Int) ([]*Filter, error) {}
// func (c *Client) EthGetLogs() {}
// func (c *Client) EthGetFilterLogs(id big.Int) ([]*Filter, error) {}

// func (c *Client) ShhVersion() (string, error) {}
// func (c *Client) ShhHasIdentity(identity [60]byte) (bool, error) {}
// func (c *Client) ShhNewIdentity() ([60]byte, error) {}
// func (c *Client) ShhNewGroup() ([60]byte, error) {}
// func (c *Client) ShhAddToGroup(identity [60]byte) (bool, error) {}
// func (c *Client) ShhNewFilter(identity [60]byte, topics []byte) (big.Int, error) {}
// func (c *Client) ShhUninstallFilter(filterId big.Int) (bool, error) {}
// func (c *Client) ShhPost() (bool, error) {}
// func (c *Client) ShhGetMessages(filterId big.Int) ( error) {}
// func (c *Client) ShhGetFilterChanges(filterId big.Int) (error) {}

type ObjectError struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *ObjectError) Error() string {
	return e.Message
}

type Request struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	ID      int64         `json:"id"`
	Params  []interface{} `json:"params"`
}

type ResponseBase struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int64           `json:"id"`
	Error   *ObjectError    `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

const _VER_ = "2.0"

func (c *Client) CallMethod(v interface{}, method string, params ...interface{}) error {
	c.idLock.Lock()
	c.id++
	req := Request{JSONRPC: _VER_, ID: c.id, Method: method, Params: params}
	c.idLock.Unlock()
	e := error(nil)
	for {
		b := []byte(nil)
		b, e = json.Marshal(req)
		if e != nil {
			break
		}
		w := (*http.Response)(nil)
		w, e = c.httpClient.Post(c.url, "application/json", bytes.NewReader(b))
		if w != nil {
			defer w.Body.Close()
		}
		if e != nil {
			break
		}
		b, e := ioutil.ReadAll(w.Body)
		if e != nil {
			break
		}
		var parsed ResponseBase
		e = json.Unmarshal(b, &parsed)
		if e != nil {
			break
		}
		if parsed.Error != nil {
			e = parsed.Error
			break
		}
		if parsed.ID != req.ID || parsed.JSONRPC != _VER_ {
			e = errors.New("Error: RPC Specification error")
			break
		}
		e = json.Unmarshal(parsed.Result, v)
		//
		// Finally
		if true {
			break
		}
	}
	return e
}
