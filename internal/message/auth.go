package message

import "github.com/kuznetsovin/egts-protocol/libs/egts"

const (
	DispatcherID   = 1
	AuthMessagePID = 1
)

func CreateAuthMessage() ([]byte, error) {
	dispatchIdentity := &egts.SrDispatcherIdentity{
		DispatcherType: 0,
		DispatcherID:   DispatcherID,
	}

	authMessage := egts.Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "10",
		HeaderLength:     11,
		HeaderEncoding:   0,
		PacketIdentifier: uint16(AuthMessagePID),
		PacketType:       egts.PtAppdataPacket,
		ServicesFrameData: &egts.ServiceDataSet{
			egts.ServiceDataRecord{
				RecordNumber:             1,
				SourceServiceOnDevice:    "0",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "10",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "0",
				SourceServiceType:        egts.AuthService,
				RecipientServiceType:     egts.AuthService,
				RecordDataSet: egts.RecordDataSet{
					{
						SubrecordType:   egts.SrDispatcherIdentityType,
						SubrecordLength: dispatchIdentity.Length(),
						SubrecordData:   dispatchIdentity,
					},
				},
			},
		},
	}

	return authMessage.Encode()
}
