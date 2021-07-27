package ethereum

import (
	"github.com/trustwallet/blockatlas/platform/bitcoin/blockbook"
	"github.com/trustwallet/blockatlas/platform/ethereum/bounce"
	"github.com/trustwallet/blockatlas/platform/ethereum/opensea"
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/network/middleware"
)

type Platform struct {
	CoinIndex   uint
	client      EthereumClient
	collectible CollectibleClient
}

func InitWithBlockbook(coinType uint, blockbookApi string) *Platform {
	c := client.InitClient(blockbookApi, middleware.SentryErrorHandler)
	c.Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36"
	return &Platform{
		CoinIndex: coinType,
		client:    &blockbook.Client{Request: c},
	}
}

func InitWithOpenSea(coinType uint, blockbookApi, collectionApi, collectionKey string) *Platform {
	platform := InitWithBlockbook(coinType, blockbookApi)
	platform.collectible = opensea.InitClient(collectionApi, collectionKey)
	return platform
}

func InitWithBounce(coinType uint, blockbookApi, collectionApi string) *Platform {
	platform := InitWithBlockbook(coinType, blockbookApi)
	platform.collectible = bounce.InitClient(collectionApi)
	return platform
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
