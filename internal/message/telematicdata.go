package message

import (
	"github.com/kuznetsovin/egts-protocol/libs/egts"
	"time"
)

func CreateTelematicDataMessage(pid int) ([]byte, error) {
	telematicDataMessage := egts.Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "10",
		HeaderLength:     11,
		HeaderEncoding:   0,
		PacketIdentifier: uint16(pid),
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
				ObjectIDFieldExists:      "1",
				ObjectIdentifier:         uint32(999),
				SourceServiceType:        egts.TeledataService,
				RecipientServiceType:     egts.TeledataService,
				RecordDataSet: egts.RecordDataSet{
					egts.RecordData{
						SubrecordType: egts.SrPosDataType,
						SubrecordData: &egts.SrPosData{
							NavigationTime:      time.Now(),
							Latitude:            55.62752532903746,
							Longitude:           37.782409656276556,
							ALTE:                "0",
							LOHS:                "0",
							LAHS:                "0",
							MV:                  "0",
							BB:                  "0",
							CS:                  "0",
							FIX:                 "0",
							VLD:                 "0",
							DirectionHighestBit: 0,
							AltitudeSign:        0,
							Speed:               0,
							Direction:           0,
							Odometer:            0,
							DigitalInputs:       0,
						},
					},
				},
			},
		},
	}

	return telematicDataMessage.Encode()
}
