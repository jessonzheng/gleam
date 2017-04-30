package pb

func (m *FlowExecutionStatus) GetDataset(datasetId int32) *FlowExecutionStatus_Dataset {
	if m != nil {
		for _, t := range m.Datasets {
			if t.Id == datasetId {
				return t
			}
		}
	}
	return nil
}

func (m *FlowExecutionStatus) GetDatasetShard(datasetId, datasetShardId int32) *FlowExecutionStatus_DatasetShard {
	if m != nil {
		for _, t := range m.DatasetShards {
			if t.DatasetId == datasetId && t.Id == datasetShardId {
				return t
			}
		}
	}
	return nil
}

func (m *FlowExecutionStatus) GetTask(stepId, taskId int32) *FlowExecutionStatus_Task {
	if m != nil {
		for _, t := range m.Tasks {
			if t.StepId == stepId && t.Id == taskId {
				return t
			}
		}
	}
	return nil
}

func (m *FlowExecutionStatus) GetStep(stepId int32) *FlowExecutionStatus_Step {
	if m != nil {
		for _, t := range m.Steps {
			if t.Id == stepId {
				return t
			}
		}
	}
	return nil
}
