package geoscore

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"

	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

var _ framework.ScorePlugin = &Plugin{}

type Plugin struct {
	handle framework.Handle
}

var GeoScoreMap = map[string]int64{"DE": 100, "BE": 79, "NL": 76, "FR": 71, "IT": 68, "UK": 61, "PL": 46, "ES": 12, "FI": 0}

// Name is the name of the plugin used in the plugin registry and configurations.
const Name = "GeoScore"

// Name returns name of the plugin. It is used in logs, etc.
func (pl *Plugin) Name() string {
	return Name
}

func (pl *Plugin) Score(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (int64, *framework.Status) {
	nodeInfo, err := pl.handle.SnapshotSharedLister().NodeInfos().Get(nodeName)
	if err != nil {
		return 0, framework.AsStatus(fmt.Errorf("getting node %q from Snapshot: %w", nodeName, err))
	}

	score := calculateScores(nodeInfo)
	return score, nil
}

func (pl *Plugin) ScoreExtensions() framework.ScoreExtensions {
	return nil
}

func calculateScores(nodeInfo *framework.NodeInfo) int64 {
	region := nodeInfo.Node().Annotations["node.kubernetes.io/region"]
	return GeoScoreMap[region]
}

// New initializes a new plugin and returns it.
func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &Plugin{handle: h}, nil
}
