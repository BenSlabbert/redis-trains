package graph

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"log"
	"time"
)

const createNodes = `
CREATE (a:Station {name: 'Kings Cross', latitude: 51.5308, longitude: -0.1238}),
       (b:Station {name: 'Euston', latitude: 51.5282, longitude: -0.1337}),
       (c:Station {name: 'Camden Town', latitude: 51.5392, longitude: -0.1426}),
       (d:Station {name: 'Kentish Town', latitude: 51.5507, longitude: -0.1402}),
       (a)-[:CONNECTION {distance: 1.0}]->(b),
       (b)-[:CONNECTION {distance: 1.0}]->(c),
       (c)-[:CONNECTION {distance: 1.0}]->(d),

       (b)-[:CONNECTION {distance: 1.0}]->(a),
       (c)-[:CONNECTION {distance: 1.0}]->(b),
       (d)-[:CONNECTION {distance: 1.0}]->(c);
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
MATCH (source:Station {name: $sourceStationName}), (target:Station {name: $targetStationName})
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

const graphExistsQuery = `CALL gds.graph.exists($graphName) YIELD exists`

const allStationsQuery = `MATCH (n:Station) return n;`

const stationByNameQuery = `MATCH (n:Station) where n.name = $stationName return n;`

type RailNetworkClient struct {
	target      string
	username    string
	password    string
	realm       string
	maxPoolSize int
	session     neo4j.Session
	driver      neo4j.Driver
	closingChan chan bool
}

type AStarResult struct {
	TotalCost float64
	Source    string
	Target    string
	Path      []string
}

func NewRailNetworkClient(target, username, password, realm string, maxPoolSize int) (*RailNetworkClient, error) {
	driver, err := neo4j.NewDriver(target, neo4j.BasicAuth(username, password, realm), func(config *neo4j.Config) { config.MaxConnectionPoolSize = maxPoolSize })
	if err != nil {
		return nil, err
	}

	err = driver.VerifyConnectivity()
	if err != nil {
		return nil, err
	}

	rnc := &RailNetworkClient{
		target:      target,
		username:    username,
		password:    password,
		realm:       realm,
		maxPoolSize: maxPoolSize,
		session:     driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}),
		driver:      driver,
		closingChan: make(chan bool),
	}

	err = rnc.init()
	if err != nil {
		return nil, err
	}

	go rnc.heartBeat()
	return rnc, nil
}

func (rnc *RailNetworkClient) init() error {
	exists, err := rnc.graphExists()
	if err != nil {
		return err
	}

	if !exists {
		// todo do these in the same transaction!
		_, err = rnc.queryWithWriteTransaction(createNodes, nil)
		if err != nil {
			return err
		}

		_, err = rnc.queryWithWriteTransaction(createGraphQuery, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rnc *RailNetworkClient) heartBeat() {
	for {
		select {
		case <-rnc.closingChan:
			return
		case <-time.After(1 * time.Second):
			err := rnc.driver.VerifyConnectivity()
			if err != nil {
				// not sure what we do here, learn about the driver
				log.Println(err)
			}
		}
	}
}

func (rnc *RailNetworkClient) Close() error {
	rnc.closingChan <- true
	return rnc.session.Close()
}

func (rnc *RailNetworkClient) FindAllStations() ([]string, error) {
	records, err := rnc.queryWithReadTransaction(allStationsQuery, nil)
	if err != nil {
		return nil, err
	}

	var stations []string
	for _, record := range records {
		node := record.Values[0].(dbtype.Node)
		stations = append(stations, node.Props["name"].(string))
	}

	return stations, nil
}

func (rnc *RailNetworkClient) FindStation(name string) (string, error) {
	record, err := rnc.queryWithReadTransaction(stationByNameQuery, map[string]interface{}{"stationName": name})
	if err != nil {
		return "", err
	}

	if len(record) == 0 || len(record[0].Values) == 0 {
		return "", fmt.Errorf("station %s not found", name)
	}

	node := record[0].Values[0].(dbtype.Node)
	return node.Props["name"].(string), nil
}

func (rnc *RailNetworkClient) AStar(fromStation, toStation string) (*AStarResult, error) {
	records, err := rnc.queryWithReadTransaction(aStarQuery, map[string]interface{}{
		"sourceStationName": fromStation,
		"targetStationName": toStation,
	})
	if err != nil {
		return nil, err
	}

	nodes := records[0].Values[6].([]interface{})
	var stations []string
	for _, n := range nodes {
		node := n.(dbtype.Node)
		stations = append(stations, node.Props["name"].(string))
	}

	return &AStarResult{
		TotalCost: records[0].Values[3].(float64),
		Source:    records[0].Values[1].(string),
		Target:    records[0].Values[2].(string),
		// todo all stations have the same connection distance, we need to vary this so the train can move at different speeds, etc
		Path: stations,
	}, nil
}

func (rnc *RailNetworkClient) graphExists() (bool, error) {
	records, err := rnc.queryWithReadTransaction(graphExistsQuery, map[string]interface{}{"graphName": "myGraph"})
	if err != nil {
		return false, err
	}

	switch t := records[0].Values[0].(type) {
	case bool:
		return t, err
	default:
		return false, fmt.Errorf("failed to convert %t to bool", t)
	}
}

func (rnc *RailNetworkClient) queryWithReadTransaction(query string, params map[string]interface{}) ([]*neo4j.Record, error) {
	out, err := rnc.session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(query, params)

		if err != nil {
			return nil, err
		}

		var records []*neo4j.Record
		for result.Next() {
			records = append(records, result.Record())
		}

		return records, result.Err()
	})
	if err != nil {
		return nil, err
	}

	switch t := out.(type) {
	case []*neo4j.Record:
		return t, nil
	default:
		return nil, nil
	}
}

func (rnc *RailNetworkClient) queryWithWriteTransaction(query string, params map[string]interface{}) (*neo4j.Record, error) {
	out, err := rnc.session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(query, params)

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

	switch t := out.(type) {
	case *neo4j.Record:
		return t, nil
	default:
		return nil, nil
	}
}

func (rnc *RailNetworkClient) addNode(targetStationName, newStationName string, newStationLat, newStationLong, connectionDistance float32) (*neo4j.Record, error) {
	out, err := rnc.session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
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

	switch t := out.(type) {
	case *neo4j.Record:
		return t, nil
	default:
		return nil, nil
	}
}
