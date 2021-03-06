package api

import (
	"cmsgo/pkg/api/routers"
	"cmsgo/pkg/common/cache"
	"cmsgo/pkg/common/db"
	"cmsgo/pkg/common/log"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	config   string
	port     string
	cors     bool
	loglevel uint8
	//StartCmd : set up restful api server
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start cmsgo API server",
		Example: "cmsgo server config/in-local.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/in-local.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "12000", "Tcp port server listening on")
	StartCmd.PersistentFlags().Uint8VarP(&loglevel, "loglevel", "l", 0, "Log level")
	StartCmd.PersistentFlags().BoolVarP(&cors, "cors", "x", false, "Enable cors headers")
}

func usage() {
	usageStr := `
                      
  ####  #    #  ####  
 #    # ##  ## #      
 #      # ## #  ####  
 #      #    #      # 
 #    # #    # #    # 
  ####  #    #  ####

`
	fmt.Printf("%s\n", usageStr)
}

func setup() {
	//Set up log level
	zerolog.SetGlobalLevel(zerolog.Level(loglevel))
	//Set up configuration
	viper.SetConfigFile(config)
	content, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
		//log.Fatal().Err(err).Msg("Read config file fail")
	}
	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
		//log.Fatal().Err(err).Msg("Parse config file fail")

	}
	//Set up run mode
	mode := viper.GetString("mode")
	gin.SetMode(mode)
	//Set up database connection
	db.Setup()
	//Set up cache
	cache.SetUp()
	//set up mongo
	//db.SetUpMongo()
	//auto migrate
	db.AutoMigrate()
	//Initialize language
	//middleware.InitLang()
}

func run() error {
	//engine := gin.Default()
	//router.Init(engine, cors)
	//return engine.Run(":" + port)
	router := routers.InitRouter()
	address := fmt.Sprintf(":%s", port)
	s := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()
}
