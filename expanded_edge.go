package osm2ch

import "github.com/paulmach/osm"

// ExpandedGraph represents an edge in expanded graph
type ExpandedEdge struct {
	ID              int64
	Source          EdgeID
	Target          EdgeID
	SourceOSMWayID  osm.WayID
	TargetOSMWayID  osm.WayID
	SourceComponent ExpandedEdgeComponent
	TargetComponent ExpandedEdgeComponent
	WasOneway       bool
	CostMeters      float64
	/* CostSeconds  float64 */ //@todo: consider cost customization
	Geom                       []GeoPoint
}

// ExpandedEdgeComponent represents former Way
type ExpandedEdgeComponent struct {
	SourceNodeID osm.NodeID
	TargetNodeID osm.NodeID
	Tags         osm.Tags
	CostMeters   float64
}

// restrictionComponent represents member of restriction relation. Could be either way or node.
type restrictionComponent struct {
	ID   int64
	Type string
}
