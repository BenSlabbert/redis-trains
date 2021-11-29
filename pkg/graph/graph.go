package graph

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	rg "github.com/redislabs/redisgraph-go"
	"log"
)

const relationConnection = "connection"
const propertyDistance = "distance"
const propertyName = "name"
const labelStation = "Station"
const graphVersion = "1.0"
const graphName = "map"
const shortestPathParameterizedQuery = "MATCH (a:Station {name: $fromStation}), (b:Station {name: $toStation}) RETURN shortestPath((a)-[:connection*0..]->(b))"

var ErrNoResults = fmt.Errorf("no results found")

type RailNetworkClient struct {
	conn  redis.Conn
	graph *rg.Graph
}

func NewRailNetworkClient(url string) (*RailNetworkClient, error) {
	conn, err := redis.Dial("tcp", url)
	if err != nil {
		return nil, err
	}

	graph := newGraph(graphName, conn)
	rnc := &RailNetworkClient{conn: conn, graph: graph}

	if err = rnc.init(); err != nil {
		return nil, err
	}

	return rnc, nil
}

func (rnc *RailNetworkClient) FindPath(origin, destination string) ([]string, error) {
	// min hops 0
	// max hops infinite
	query, err := rnc.graph.ParameterizedQuery(shortestPathParameterizedQuery, map[string]interface{}{"fromStation": origin, "toStation": destination})
	if err != nil {
		log.Fatalln(err)
	}

	next := query.Next()
	if !next {
		return nil, ErrNoResults
	}

	record := query.Record()
	values := record.Values()
	path := values[0].(rg.Path)

	stations := make([]string, len(path.Nodes), len(path.Nodes))
	for i := 0; i < len(stations); i++ {
		stations[i] = path.Nodes[i].Properties[propertyName].(string)
	}

	return stations, nil
}

func (rnc *RailNetworkClient) init() error {
	return rnc.createGraph()
}

func (rnc *RailNetworkClient) Close() error {
	return rnc.conn.Close()
}

func newGraph(name string, conn redis.Conn) *rg.Graph {
	graph := rg.GraphNew(name, conn)
	return &graph
}

func (rnc *RailNetworkClient) createGraph() error {
	graphKey := fmt.Sprintf("%s-graph", rnc.graph.Id)
	do, err := rnc.graph.Conn.Do("HGET", graphKey, "version")
	if err != nil {
		return err
	}

	if do != nil {
		res := string(do.([]uint8))

		if res == graphVersion {
			return nil
		}
	}

	// graph does not exist/wrong version, create
	err = rnc.graph.Delete()
	if err != nil && err.Error() != "ERR Invalid graph operation on empty key" {
		return err
	}

	stations := []string{"A1", "A2", "A3", "A4", "A5"}
	var nodes []*rg.Node

	for _, station := range stations {
		node := &rg.Node{
			Alias: station,
			Label: labelStation,
			Properties: map[string]interface{}{
				propertyName: station,
			},
		}
		rnc.graph.AddNode(node)
		nodes = append(nodes, node)
	}

	for i := 0; i < len(nodes)-1; i++ {
		// add bidirectional edges
		edge := &rg.Edge{
			Source:      nodes[i],
			Relation:    relationConnection,
			Destination: nodes[i+1],
			Properties: map[string]interface{}{
				propertyDistance: "1.0",
			},
		}
		if e := rnc.graph.AddEdge(edge); e != nil {
			return e
		}

		edge = &rg.Edge{
			Source:      nodes[i+1],
			Relation:    relationConnection,
			Destination: nodes[i],
			Properties: map[string]interface{}{
				propertyDistance: "1.0",
			},
		}
		if e := rnc.graph.AddEdge(edge); e != nil {
			return e
		}
	}

	_, err = rnc.graph.Commit()
	if err != nil {
		return err
	}

	_, err = rnc.graph.Conn.Do("HSET", graphKey, "version", graphVersion)
	if err != nil {
		return err
	}

	return err
}
