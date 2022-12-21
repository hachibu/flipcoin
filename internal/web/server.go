package web

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hachibu/flipcoin/internal/config"
)

var (
	priceCents = int32(100)
	priceChart = NewPriceChart()
)

func generatePriceData() {
	rand.Seed(time.Now().UnixNano())
	for {
		t := time.Now().UTC()
		n := rand.Intn(2)
		if n == 1 {
			priceCents++
		} else if priceCents > 0 {
			priceCents--
		}
		priceChart.X = append(priceChart.X, t)
		priceChart.Y = append(priceChart.Y, priceCents)
		time.Sleep(1 * time.Second)
	}
}

func NewServer(cfg *config.Config) *http.Server {
	addr := fmt.Sprintf(":%s", cfg.HttpServerConfig.Port)
	router := gin.Default()

	if cfg.Env == config.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Printf("%s\n", cfg.ToJSON())
	}

	router.ForwardedByClientIP = true
	router.MaxMultipartMemory = 5 << 20 // 5 MB

	router.StaticFile("/favicon.ico", "internal/web/assets/favicon.ico")
	router.LoadHTMLGlob("internal/web/views/*")

	// Middleware
	router.Use(rateLimiterMiddleware())
	router.Use(sessionsMiddleware())

	// Route Handlers
	router.GET("/", homeHandler)
	router.GET("/api/price-chart", priceChartHandler)

	go generatePriceData()

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
