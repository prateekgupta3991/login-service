package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("172.18.9.140:9042")
	cluster.Keyspace = "godemo"
	Session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("cassandra init done")
}
