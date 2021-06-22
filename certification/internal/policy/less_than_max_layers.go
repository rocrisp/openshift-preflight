package policy

import "github.com/komish/preflight/certification"

type UnderLayerMaxPolicy struct {
}

func (p UnderLayerMaxPolicy) Validate(image string) (bool, error) {
	return false, nil
}

func (p UnderLayerMaxPolicy) Name() string {
	return "MaximumLayerPolicy"
}

func (p UnderLayerMaxPolicy) Metadata() certification.Metadata {
	return certification.Metadata{
		Description:      "Checking if container has less than 40 layers",
		Level:            "better",
		KnowledgeBaseURL: "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
		PolicyURL:        "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
	}
}

func (p UnderLayerMaxPolicy) Help() certification.HelpText {
	return certification.HelpText{
		Message:    "Uncompressed container images should have less than 40 layers. Too many layers within the container images can degrade container performance.",
		Suggestion: "Optimize your Dockerfile to consolidate and minimize the number of layers. Each RUN command will produce a new layer. Try combining RUN commands using && where possible.",
	}
}
