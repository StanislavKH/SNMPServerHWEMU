package rfd12

import "fmt"
import "github.com/slayercat/gosnmp"
import "github.com/slayercat/GoSNMPServer"

var outputTable map[int]int

func buildPDUValueControlItem(index int) []*GoSNMPServer.PDUValueControlItem {
	toRet := []*GoSNMPServer.PDUValueControlItem{
		{
			OID:      fmt.Sprintf("1.3.6.1.4.1.18480.12.2.4.1.1.1.%d", index),
			Type:     gosnmp.Integer,
			OnGet:    func() (value interface{}, err error) { return GoSNMPServer.Asn1IntegerWrap(index), nil },
			Document: fmt.Sprintf("outputIndex.%d", index),
		},
		{
			OID:      fmt.Sprintf("1.3.6.1.4.1.18480.12.2.4.1.1.2.%d", index),
			Type:     gosnmp.Integer,
			OnGet:    func() (value interface{}, err error) { return GoSNMPServer.Asn1IntegerWrap(outputTable[index]), nil },
			OnSet:    func(value interface{}) error {
				outputTable[index] = GoSNMPServer.Asn1IntegerUnwrap(value)
				return nil
			},
			Document: fmt.Sprintf("portSel.%d", index),
		},
		{
			OID:      fmt.Sprintf("1.3.6.1.4.1.18480.12.2.4.1.1.3.%d", index),
			Type:     gosnmp.OctetString,
			OnGet:    func() (value interface{}, err error) { return GoSNMPServer.Asn1OctetStringWrap(fmt.Sprintf("out%d", index)), nil },
			Document: fmt.Sprintf("outputName.%d", index),
		},
	}
	return toRet
}

// DeviceOIDs Returns a list
func DeviceOIDs() []*GoSNMPServer.PDUValueControlItem {
	toRet := []*GoSNMPServer.PDUValueControlItem{}
	outputTable = make(map[int]int)

	for i := 1; i <= 16; i++ {
		outputTable[i] = i
		PDUItem := buildPDUValueControlItem(i)
		toRet = append(toRet, PDUItem...)
	}

	return toRet
}
