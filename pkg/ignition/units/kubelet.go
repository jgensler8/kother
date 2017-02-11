package units

import "github.com/lann/builder"

type KubeletUnit struct {
	Name string
	HyperkubeImage string
	HyperkubeTag string
	APIServerDNS string
	ClusterDNS string
	ClusterDomain string
	ManifestPath string
}

type kubeletUnitBuilder builder.Builder

func (b kubeletUnitBuilder) Name(n string) kubeletUnitBuilder {
	return builder.Set(b, "Name", n).(kubeletUnitBuilder)
}

func (b kubeletUnitBuilder) HyperkubeImage(i string) kubeletUnitBuilder {
	return builder.Set(b, "HyperkubeImage", i).(kubeletUnitBuilder)
}

func (b kubeletUnitBuilder) HyperkubeTag(t string) kubeletUnitBuilder {
	return builder.Set(b, "HyperkubeTag", t).(kubeletUnitBuilder)
}

func (b kubeletUnitBuilder) APIServerDNS(d string) kubeletUnitBuilder {
	return builder.Set(b, "APIServerDNS", d).(kubeletUnitBuilder)
}

func (b kubeletUnitBuilder) ClusterDNS(d string) kubeletUnitBuilder {
	return builder.Set(b, "ClusterDNS", d).(kubeletUnitBuilder)
}

func (b kubeletUnitBuilder) ClusterDomain(d string) kubeletUnitBuilder {
	return builder.Set(b, "ClusterDomain", d).(kubeletUnitBuilder)
}

func (b kubeletUnitBuilder) ManifestPath(p string) kubeletUnitBuilder {
	return builder.Set(b, "ManifestPath", p).(kubeletUnitBuilder)
}

func (b kubeletUnitBuilder) Build() KubeletUnit {
	return builder.GetStruct(b).(KubeletUnit)
}

var KubeletUnitBuilder = builder.Register(kubeletUnitBuilder{}, KubeletUnit{}).(kubeletUnitBuilder)