package main

import (
	"fmt"
	"./gocql"
	"log"
)

func main() {

	cluster := gocql.NewCluster("10.181.54.160")
	cluster.Keyspace = "logger"
	session, _ := cluster.CreateSession()
	fmt.Println(session)
	defer session.Close()

    var user_agent string
    err := session.Query(`SELECT user_agent FROM article_pv WHERE id = ? and catid = ? and time > 1388505600 and time < 1416458906 limit 1  allow filtering `,177, 755).Consistency(gocql.One).Scan(&user_agent)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("result:", user_agent)

}

