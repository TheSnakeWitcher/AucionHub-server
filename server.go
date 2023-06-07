package main

import (
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/abi/bind"
	"github.com/labstack/echo/v4"
    AuctionHub "AuctionHub-server/AuctionHub"
)

const sellerParamName string = "seller"

func newContract(config Configuration,client *ethclient.Client) *AuctionHub.AuctionHub {
    contractAddress := common.HexToAddress(config.ContractAddress)
    contract,err := AuctionHub.NewAuctionHub(contractAddress,client)
    if err != nil {
        panic("could not instantiate contract")
    }
    return contract
}

/**
 * NOTE: indexing is better done using TheGraph protocol
 *       but as this is a didactical task....
 */
func NewServer(svc Service,client *ethclient.Client,config Configuration) *echo.Echo {
    srv := echo.New()
    contract := newContract(config,client)

    srv.GET("/config/period",func (ctx echo.Context) error {
        period,err := contract.GetAuctionPeriod(nil)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError,err)
        }
        return ctx.JSON(http.StatusOK,period)
    })

    srv.GET("/config/fee",func (ctx echo.Context) error {
        fee,err := contract.GetAuctionFee(nil)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError,err)
        }   
        return ctx.JSON(http.StatusOK,fee)
    })

    srv.GET("/auctions",func (ctx echo.Context) error {
        auctions := svc.ListAuction()
        return ctx.JSON(http.StatusOK,auctions)
    })
    srv.GET("/auctions/:seller",func (ctx echo.Context) error {
        seller := ctx.Param(sellerParamName)
        auction,err := svc.GetAuction(seller)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError,err)
        }
        return ctx.JSON(http.StatusOK,auction)
    })
    srv.POST("/auctions",func (ctx echo.Context) error {
        // TODO: transfer ERC721 token
        // TODO: if result ok register in DB else return error
        seller := common.HexToAddress("some")
        auctionData,err := contract.GetAuctionData(nil,seller)
        result := svc.InsertAuction(auctionData)
        return ctx.JSON(http.StatusOK,result)
    })

    srv.GET("/bids",func (ctx echo.Context) error {
        bids := svc.ListBids()
        return ctx.JSON(http.StatusOK,bids)
    })
    srv.GET("/bids/:seller",func (ctx echo.Context) error {
        seller := ctx.Param(sellerParamName)
        bid,err := svc.GetBid(seller)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError,err)
        }
        return ctx.JSON(http.StatusOK,bid)
    })
    srv.POST("/bids",func (ctx echo.Context) error {
        return ctx.String(http.StatusOK,"add a bid")
    })

    return srv
}
