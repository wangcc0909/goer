module goer

require (
	github.com/chanxuehong/rand v0.0.0-20180830053958-4b3aff17f488 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/core v0.6.0
	github.com/go-xorm/xorm v0.7.1
	github.com/gorilla/mux v1.7.2
	github.com/lonng/nano v0.4.1-0.20190704005402-15209d995681
	github.com/lonnng/nex v1.4.1
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/sirupsen/logrus v1.2.0
	github.com/smartystreets/goconvey v0.0.0-20190604192920-68dc04aab96a
	github.com/spf13/viper v1.4.0
	github.com/urfave/cli v1.20.1-0.20190203184040-693af58b4d51
	github.com/xxtea/xxtea-go v0.0.0-20170828040851-35c4b17eecf6
	golang.org/x/tools v0.0.0-20190511041617-99f201b6807e
	gopkg.in/chanxuehong/wechat.v2 v2.0.0-20180924084534-7e0579cb5377
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.40.1-0.20190612163021-8a8d2c2fb096
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net => github.com/golang/net v0.0.0-20190318221613-d196dffd7c2b
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190523182746-aaccbc9213b0
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190318195719-6c81ef8f67ca
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190529010454-aa71c3f32488
	google.golang.org/appengine => github.com/golang/appengine v1.6.1-0.20190515044707-311d3c5cf937
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190522204451-c2c4e71fbf69
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.0
)
