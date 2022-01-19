package patch

import (
	"encoding/json"

	sp "github.com/RafaySystems/rcloud-base/components/common/pkg/controller/strategicpatch"
	infrav3 "github.com/RafaySystems/rcloud-base/components/common/proto/types/infrapb/v3"
)

type clusterConditions struct {
	Conditions []*infrav3.ClusterCondition `json:"conditions" patchStrategy:"merge" patchMergeKey:"type"`
}

// ClusterStatus patches existing cluster status with current
func ClusterStatus(existing, current *infrav3.ClusterStatus) error {
	eb, err := json.Marshal(clusterConditions{Conditions: existing.Conditions})
	if err != nil {
		return err
	}

	cb, err := json.Marshal(clusterConditions{Conditions: current.Conditions})
	if err != nil {
		return err
	}

	pb, err := sp.CreateTwoWayMergePatch(eb, cb, (*clusterConditions)(nil))
	if err != nil {
		return err
	}

	fb, err := sp.StrategicMergePatch(eb, pb, (*clusterConditions)(nil))
	if err != nil {
		return err
	}

	err = json.Unmarshal(fb, existing)
	if err != nil {
		return err
	}

	if current.PublishedBlueprint != "" {
		existing.PublishedBlueprint = current.PublishedBlueprint
	}

	return nil
}

// ClusterNodeStatus patches existing cluster node status with current
func ClusterNodeStatus(existing, current *infrav3.ClusterNodeStatus) error {
	eb, err := json.Marshal(existing)
	if err != nil {
		return err
	}

	cb, err := json.Marshal(current)
	if err != nil {
		return err
	}

	pb, err := sp.CreateTwoWayMergePatch(eb, cb, (*infrav3.ClusterNodeStatus)(nil))
	if err != nil {
		return err
	}

	fb, err := sp.StrategicMergePatch(eb, pb, (*infrav3.ClusterNodeStatus)(nil))
	if err != nil {
		return err
	}

	err = json.Unmarshal(fb, existing)
	if err != nil {
		return err
	}

	return nil
}
