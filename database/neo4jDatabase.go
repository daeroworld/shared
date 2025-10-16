package database

import (
	"context"
	"crave/shared/configuration"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Record neo4j.Record

type Neo4jWrapper struct {
	Driver neo4j.DriverWithContext
}

func ConnectNeo4jDatabase(database *configuration.Database) *Neo4jWrapper {
	driver, err := neo4j.NewDriverWithContext(database.Uri, neo4j.BasicAuth(database.Username, database.Password, ""))
	if err != nil {
		log.Fatalf("failed to connect to Neo4j: %w", err)
	}
	neo4jWrapper := &Neo4jWrapper{
		Driver: driver,
	}
	createUniqueConstraint(neo4jWrapper)
	return neo4jWrapper
}

func createUniqueConstraint(neo4jWrapper *Neo4jWrapper) error {
	// session, err := Neo4jWrapper.NewSession()
	// if err != nil {
	// 	return err
	// }
	// defer session.Close()
	query := `CREATE CONSTRAINT unique_node_id IF NOT EXISTS 
                  FOR (n:Node) REQUIRE n.id IS UNIQUE;`
	ctx := context.Background()
	result, _ := neo4j.ExecuteQuery(ctx, neo4jWrapper.Driver,
		query,
		nil, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	for _, record := range result.Records {
		fmt.Println(record.AsMap())
	}
	log.Println("[Neo4j] Unique constraint on 'id' created successfully!")
	return nil
}

func (n *Neo4jWrapper) CloseDatabbase() {
	if n.Driver != nil {
		ctx := context.Background()
		n.Driver.Close(ctx)
		log.Println("[Neo4j] Connection closed")
	}
}

func (n *Neo4jWrapper) NewSession(ctx context.Context) neo4j.SessionWithContext {

	return n.Driver.NewSession(ctx, neo4j.SessionConfig{})
}
