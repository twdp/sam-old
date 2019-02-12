module tianwei.pro/sam

require (
	github.com/astaxie/beego v1.11.1
	github.com/casbin/casbin v1.8.0 // indirect
	github.com/couchbase/go-couchbase v0.0.0-20181210201043-bd8e99474993 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/go-redis/redis v6.15.1+incompatible // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/goburrow/cache v0.1.0
	github.com/gogo/protobuf v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/syndtr/goleveldb v0.0.0-20181128100959-b001fa50d6b2 // indirect
	github.com/wendal/errors v0.0.0-20181209125328-7f31f4b264ec // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	tianwei.pro/beego-guava v0.0.0-20190108132444-348c37386e91
	tianwei.pro/business v0.0.9
	tianwei.pro/smart v0.0.0-20190113040010-d33287b8e896
)

replace (
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793 => github.com/golang/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/net v0.0.0-20180826012351-8a410e7b638d => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys v0.0.0-20180903190138-2b024373dcd9 => github.com/golang/sys v0.0.0-20180903190138-2b024373dcd9
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
