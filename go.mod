module github.com/rPniu/all

go 1.25.0

require (
	github.com/JGLTechnologies/gin-rate-limit v1.5.6
	github.com/gin-gonic/gin v1.10.1
	github.com/go-ini/ini v1.67.0
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.42.0
)

require (
	github.com/bytedance/sonic v1.13.2 // indirect
	github.com/bytedance/sonic/loader v0.2.4 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.26.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/redis/go-redis/v9 v9.7.3 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.16.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/rPniu/all => ../all
	github.com/rPniu/all/conf => ../all/conf
	github.com/rPniu/all/middleware => ../all/middleware
	github.com/rPniu/all/models => ../all/models
	github.com/rPniu/all/pkg/ => ../all/pkg
	github.com/rPniu/all/pkg/checkin => ../all/pkg/checkin
	github.com/rPniu/all/pkg/error => ./pkg/e
	github.com/rPniu/all/pkg/setting => ../all/pkg/setting
	github.com/rPniu/all/pkg/util => ../all/pkg/util
	github.com/rPniu/all/pkg/util/uid => ../all/pkg/util/uid
	github.com/rPniu/all/routers => ../all/routers
	github.com/rPniu/all/routers/api => ../all/routers/api
)
