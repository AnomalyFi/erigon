package lightrpc

import "github.com/ledgerwatch/erigon/cmd/lightclient/sentinel/proto"

func (*SignedBeaconBlockBellatrix) Clone() proto.Packet {
	return &SignedBeaconBlockBellatrix{}
}

func (*LightClientFinalityUpdate) Clone() proto.Packet {
	return &LightClientFinalityUpdate{}
}

func (*LightClientOptimisticUpdate) Clone() proto.Packet {
	return &LightClientOptimisticUpdate{}
}

func (*MetadataV1) Clone() proto.Packet {
	return &MetadataV1{}
}

func (*MetadataV2) Clone() proto.Packet {
	return &MetadataV2{}
}
