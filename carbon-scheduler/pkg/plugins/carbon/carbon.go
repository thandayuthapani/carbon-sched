package carbon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type Plugin struct {
	handle framework.Handle
}

var _ framework.ScorePlugin = &Plugin{}
var lastRetrieved time.Time

var EmissionRank = map[string]int64{}

// Name is the name of the plugin used in the plugin registry and configurations.
const Name = "carbon"

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
	if lastRetrieved.IsZero() {
		lastRetrieved = time.Now()
	}
	emissionRank, err := getEmissionRanking(lastRetrieved)
	if err != nil {
		klog.Info(err)
		return 0
	}
	//region := nodeInfo.Node().Labels["node.kubernetes.io/region"]
	region := nodeInfo.Node().Annotations["node.kubernetes.io/region"]
	score := emissionRank[region]
	return int64(10 * score)
}

func getEmissionRanking(lastRetrieved time.Time) (map[string]int64, error) {
	if lastRetrieved.IsZero() || time.Now().Sub(lastRetrieved) >= 5*time.Minute || len(EmissionRank) == 0 {
		err := queryDataFromServer()
		if err != nil {
			return map[string]int64{}, err
		}
		return EmissionRank, nil
	}
	return EmissionRank, nil
}

func queryDataFromServer() error {
	url := "http://metrics-collector.default.svc.cluster.local:8080/getemission"
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		klog.Info(err)
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		klog.Info(err)
		return err
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&EmissionRank)
	if err != nil {
		klog.Info(err)
		return err
	}
	return nil
}

// New initializes a new plugin and returns it.
func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &Plugin{handle: h}, nil
}
