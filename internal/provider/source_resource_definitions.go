package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	. "github.com/mezmo-inc/terraform-provider-mezmo/internal/provider/models/sources"
)

func NewDemoSourceResource() resource.Resource {
	return &SourceResource[DemoSourceModel]{
		typeName:          "demo",
		fromModelFunc:     DemoSourceFromModel,
		toModelFunc:       DemoSourceToModel,
		getIdFunc:         func(m *DemoSourceModel) basetypes.StringValue { return m.Id },
		getPipelineIdFunc: func(m *DemoSourceModel) basetypes.StringValue { return m.PipelineId },
		getSchemaFunc:     DemoSourceResourceSchema,
	}
}

func NewAgentSourceResource() resource.Resource {
	return &SourceResource[AgentSourceModel]{
		typeName:          "agent",
		fromModelFunc:     AgentSourceFromModel,
		toModelFunc:       AgentSourceToModel,
		getIdFunc:         func(m *AgentSourceModel) basetypes.StringValue { return m.Id },
		getPipelineIdFunc: func(m *AgentSourceModel) basetypes.StringValue { return m.PipelineId },
		getSchemaFunc:     AgentSourceResourceSchema,
	}
}

func NewS3SourceResource() resource.Resource {
	return &SourceResource[S3SourceModel]{
		typeName:          "s3",
		fromModelFunc:     S3SourceFromModel,
		toModelFunc:       S3SourceToModel,
		getIdFunc:         func(m *S3SourceModel) basetypes.StringValue { return m.Id },
		getPipelineIdFunc: func(m *S3SourceModel) basetypes.StringValue { return m.PipelineId },
		getSchemaFunc:     S3SourceResourceSchema,
	}
}

func NewHttpSourceResource() resource.Resource {
	return &SourceResource[HttpSourceModel]{
		typeName:          "http",
		fromModelFunc:     HttpSourceFromModel,
		toModelFunc:       HttpSourceToModel,
		getIdFunc:         func(m *HttpSourceModel) basetypes.StringValue { return m.Id },
		getPipelineIdFunc: func(m *HttpSourceModel) basetypes.StringValue { return m.PipelineId },
		getSchemaFunc:     HttpSourceResourceSchema,
	}
}

func NewSQSSourceResource() resource.Resource {
	return &SourceResource[SQSSourceModel]{
		typeName:          "sqs",
		fromModelFunc:     SQSSourceFromModel,
		toModelFunc:       SQSSourceToModel,
		getIdFunc:         func(m *SQSSourceModel) basetypes.StringValue { return m.Id },
		getPipelineIdFunc: func(m *SQSSourceModel) basetypes.StringValue { return m.PipelineId },
		getSchemaFunc:     SQSSourceResourceSchema,
	}
}

func NewSplunkHecSourceResource() resource.Resource {
	return &SourceResource[SplunkHecSourceModel]{
		typeName:          "splunk_hec",
		fromModelFunc:     SplunkHecSourceFromModel,
		toModelFunc:       SplunkHecSourceToModel,
		getIdFunc:         func(m *SplunkHecSourceModel) basetypes.StringValue { return m.Id },
		getPipelineIdFunc: func(m *SplunkHecSourceModel) basetypes.StringValue { return m.PipelineId },
		getSchemaFunc:     SplunkHecSourceResourceSchema,
	}
}

func NewLogStashSourceResource() resource.Resource {
	return &SourceResource[LogStashSourceModel]{
		typeName:          "logstash",
		fromModelFunc:     LogStashSourceFromModel,
		toModelFunc:       LogStashSourceToModel,
		getIdFunc:         func(m *LogStashSourceModel) basetypes.StringValue { return m.Id },
		getPipelineIdFunc: func(m *LogStashSourceModel) basetypes.StringValue { return m.PipelineId },
		getSchemaFunc:     LogStashSourceResourceSchema,
	}
}
