package types

const ( // trade_type
	BUY  = "BUY"
	SELL = "SELL"

	LIMIT  = "LIMIT"
	MARKET = "MARKET"

	DEPOSIT    = "DEPOSIT"
	WITHDRAW   = "WITHDRAW"
	FEE        = "FEE"
	POSTCOLL   = "POSTCOLL"
	CANCELCOLL = "CANCELCOLL"
	PAYMENT    = "PAYMENT"
	TRANSFER   = "TRANSFER"
)

const ( // major ProductCodes
	BTCJPY   = "BTC_JPY"
	ETHJPY   = "ETH_JPY"
	XRPJPY   = "XRP_JPY"
	ETHBTC   = "ETH_BTC"
	FXBTCJPY = "FX_BTC_JPY"

	// major currencyCodes
	JPY = "JPY"
	USD = "USD"
)

const ( // Health
	NORMAL    = "NORMAL"
	BUSY      = "BUSY"
	VERYBUSY  = "VERY BUSY"
	SUPERBUSY = "SUPER BUSY"
	NOORDER   = "NO ORDER"
	STOP      = "STOP"

	// Status
	COMPLETED = "COMPLETED"
	PENDING   = "PENDING"

	RUNNING      = "RUNNING"
	CLOSE        = "CLOSE"
	STARTING     = "STARTING"
	PREOPEN      = "PREOPEN"
	CIRCUITBREAK = "CIRCUIT BREAK"

	AWAITINGSQ = "AWAITING SQ"
	MATURED    = "MATURED"

	// Lightning Futures の SQ
	SPECIALQUOTATION = "special_quotation"
)

const (
	TIFGTC = "GTC"
	TIFIOC = "IOC"
	TIFFOK = "FOK"
)
