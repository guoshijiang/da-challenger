module github.com/mantlenetworkio/da-challenger

go 1.18

replace github.com/Layr-Labs/datalayr/common => ../mantle/datalayr-mantle/common

replace github.com/Layr-Labs/datalayr/lib/merkzg => ../mantle/datalayr-mantle/lib/merkzg

replace github.com/Layr-Labs/datalayr/lib/kzgFFT => ../mantle/datalayr-mantle/lib/kzgFFT

replace github.com/Layr-Labs/datalayr/middleware/rollup-example/utils => ../mantle/datalayr-mantle/middleware/rollup-example/utils

require (
	//github.com/Layr-Labs/datalayr/common v0.0.0-00010101000000-000000000000
	//github.com/Layr-Labs/datalayr/middleware/rollup-example/utils v0.0.0-00010101000000-000000000000
	github.com/ethereum/go-ethereum v1.10.25
	github.com/shurcooL/graphql v0.0.0-20220606043923-3cf50f8a0a29
	github.com/urfave/cli v1.22.10
	google.golang.org/grpc v1.49.0
)

require (
	github.com/Layr-Labs/datalayr/common v0.0.0-00010101000000-000000000000
	github.com/Layr-Labs/datalayr/middleware/rollup-example/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/Layr-Labs/datalayr/lib/merkzg v0.0.0-00010101000000-000000000000 // indirect
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.8.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/go-ole/go-ole v1.2.1 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/rjeczalik/notify v0.9.1 // indirect
	github.com/rs/zerolog v1.27.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	github.com/tklauser/numcpus v0.2.2 // indirect
	golang.org/x/crypto v0.0.0-20220817201139-bc19a97f63c8 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220727055044-e65921a090b8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
