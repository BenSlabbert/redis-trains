package main

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

const createNodes = `
CREATE (a:Station {name: 'Kings Cross', latitude: 51.5308, longitude: -0.1238}),
       (b:Station {name: 'Euston', latitude: 51.5282, longitude: -0.1337}),
       (c:Station {name: 'Camden Town', latitude: 51.5392, longitude: -0.1426}),
       (d:Station {name: 'Mornington Crescent', latitude: 51.5342, longitude: -0.1387}),
       (e:Station {name: 'Kentish Town', latitude: 51.5507, longitude: -0.1402}),
       (a)-[:CONNECTION {distance: 0.7}]->(b),
       (b)-[:CONNECTION {distance: 1.3}]->(c),
       (b)-[:CONNECTION {distance: 0.7}]->(d),
       (d)-[:CONNECTION {distance: 0.6}]->(c),
       (c)-[:CONNECTION {distance: 1.3}]->(e);
`

const createGraphQuery = `
CALL gds.graph.create(
'myGraph',
'Station',
'CONNECTION',
  {
    nodeProperties: ['latitude', 'longitude'],
    relationshipProperties: 'distance'
  }
);
`

const dropGraphQuery = `CALL gds.graph.drop('myGraph');`

const aStarQuery = `
MATCH (source:Station {name: 'Kings Cross'}), (target:Station {name: 'Kentish Town'})
CALL gds.shortestPath.astar.stream('myGraph', {
  sourceNode: source,
  targetNode: target,
  latitudeProperty: 'latitude',
  longitudeProperty: 'longitude',
  relationshipWeightProperty: 'distance'
})
YIELD index, sourceNode, targetNode, totalCost, nodeIds, costs, path
RETURN
  index,
  gds.util.asNode(sourceNode).name AS sourceNodeName,
  gds.util.asNode(targetNode).name AS targetNodeName,
  totalCost,
  [nodeId IN nodeIds | gds.util.asNode(nodeId).name] AS nodeNames,
  costs,
  nodes(path) as path
  ORDER BY index;
`

const addNodeQuery = `
MATCH (target:Station {name: $targetStationName})
CREATE (a:Station {name: $newStationName, latitude: $newStationLat, longitude: $newStationLong}),
       (a)-[:CONNECTION {distance: $connectionDistance}]->(target),
       (target)-[:CONNECTION {distance: $connectionDistance}]->(a);
`

func main() {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "neo4j", ""))
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if e := driver.Close(); e != nil {
			log.Println(e)
		}
	}()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func() {
		if e := session.Close(); e != nil {
			log.Println(e)
		}
	}()

	out, err := queryWithWriteTransaction(session, createNodes)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(out)

	out, err = queryWithWriteTransaction(session, createGraphQuery)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(out)

	out, err = queryWithReadTransaction(session, aStarQuery)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(out)

	out, err = addNode(session, "Kings Cross", "new station", 1.23, 4.56, 1.4)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(out)
}

func queryWithReadTransaction(session neo4j.Session, query string) (*neo4j.Record, error) {
	out, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(query, nil)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record(), nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return nil, err
	}

	switch out.(type) {
	case *neo4j.Record:
		return out.(*neo4j.Record), nil
	default:
		return nil, nil
	}
}

func queryWithWriteTransaction(session neo4j.Session, query string) (*neo4j.Record, error) {
	out, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(query, nil)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record(), nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return nil, err
	}

	switch out.(type) {
	case *neo4j.Record:
		return out.(*neo4j.Record), nil
	default:
		return nil, nil
	}
}

func addNode(session neo4j.Session, targetStationName, newStationName string, newStationLat, newStationLong, connectionDistance float32) (*neo4j.Record, error) {
	out, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(addNodeQuery, map[string]interface{}{
			"targetStationName":  targetStationName,
			"newStationName":     newStationName,
			"newStationLat":      newStationLat,
			"newStationLong":     newStationLong,
			"connectionDistance": connectionDistance,
		})
		if err != nil {
			return nil, err
		}

		// after we add the node, we need to drop the gds graph and recreate it for astar queries
		_, err = tx.Run(dropGraphQuery, nil)
		if err != nil {
			return nil, err
		}

		_, err = tx.Run(createGraphQuery, nil)
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record(), nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return nil, err
	}

	switch out.(type) {
	case *neo4j.Record:
		return out.(*neo4j.Record), nil
	default:
		return nil, nil
	}
}
