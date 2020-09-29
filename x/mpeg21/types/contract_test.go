package types

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var testData = []byte(`
{
  "@id": "contract3",
  "@type": "mco-core:Contract",
  "hasParty": [
  	{ "@id": "Publisher", "label": "PUBLISHER", "@type": ["mco-core:Party"] },
	{
	  "@id": "MechanicalLicenseAgent",
	  "label": "MECHANICAL LICENSE AGENT",
	  "@type": ["mco-core:Party", "mvco:Instantiator"]
	}
  ],
  "issues": [
  {
      "@id": "permission1",
      "label": "Consumer can play a song",
      "@type": ["mvco:Permission"],
      "issuedBy": {
        "@id": "StreamingService",
        "label": "STREAMING SERVICE",
        "@type": ["mco-core:Party", "mvco:Distributor"]
      },
      "permitsAction": {
        "@id": "action1",
        "@type": ["mvco:play"],
        "label": "Consumer plays a song",
        "actedBy": {
          "@id": "Consumer",
          "label": "CONSUMER",
          "@type": ["mco-core:Party", "mvco:EndUser"]
        },
        "actedOver": { "@id": "Song", "label": "SONG", "@type": ["mvco:Work"] }
      }
    },
    {
      "@id": "obligation1",
      "label": "Publisher provide a song to streaming",
      "@type": ["mco-core:Obligation"],
      "issuedBy": {
        "@id": "Publisher",
        "label": "PUBLISHER",
        "@type": ["mco-core:Party"]
      },
      "obligatesAction": {
        "@id": "action5",
        "@type": ["mvco-ipre:provideMaterial"],
        "label": "Streaming provides a song",
        "actedBy": {
          "@id": "Publisher",
          "label": "PUBLISHER",
          "@type": ["mco-core:Party"]
        },
        "actedOver": { "@id": "Song", "label": "SONG", "@type": ["mvco:Work"] },
        "actedTo": {
          "@id": "StreamingService",
          "label": "STREAMING SERVICE",
          "@type": ["mco-core:Party", "mvco:Distributor"]
        }
      }
    }
  ]
}
`)

func TestNewContract(t *testing.T) {
	contract := Contract{}
	err := json.Unmarshal(testData, &contract)
	require.NoError(t, err)

	fmt.Println(contract.ID)
	fmt.Println(contract.Type)
	fmt.Println(contract.Parties)
	fmt.Println(contract.Issues)
	for _, issue := range contract.Issues {
		fmt.Println(issue.ID)
		fmt.Println(issue.Type[0])
		fmt.Println(issue.IssuedBy)
		switch issue.Type[0] {
		case "mvco:Permission":
			fmt.Println("permission")
			fmt.Println(issue.PermitsAction.Type)
		case "mco-core:Obligation":
			fmt.Println("obligation")
			fmt.Println(issue.ObligatesAction.Type)
		}
	}
}
