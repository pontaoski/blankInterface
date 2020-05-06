package main

import "encoding/xml"

type Arg struct {
	Name          string `xml:"name,attr"`
	ID            string `xml:"type,attr"`
	InterfaceType string `xml:"interface,attr,omitempty"`
}

type Method struct {
	Name  string `xml:"name,attr"`
	Since int    `xml:"since,attr,omitempty"`
	Args  []Arg  `xml:"arg"`
}

type Entry struct {
	Name  string `xml:"name,attr"`
	Value int    `xml:"value,attr"`
}

type Enum struct {
	Name    string  `xml:"name,attr"`
	Entries []Entry `xml:"entry"`
}

type Interface struct {
	Name        string      `xml:"name,attr"`
	Version     int         `xml:"version,attr"`
	Description Description `xml:"description,omitempty"`
	Request     []Method    `xml:"request"`
	Event       []Method    `xml:"event"`
	Enum        []Enum      `xml:"enum"`
}

type Copyright struct {
	Information string `xml:",cdata"`
}

type Description struct {
	Summary     string `xml:"summary,attr"`
	Description string `xml:",innerxml"`
}

type Protocol struct {
	XMLName     xml.Name    `xml:"protocol"`
	Name        string      `xml:"name,attr"`
	Copyright   Copyright   `xml:"copyright"`
	Description Description `xml:"description,omitempty"`
	Interfaces  []Interface `xml:"interface"`
}

func (p ProtocolDescription) ToProtocol() Protocol {
	outputProtocol := Protocol{}
	outputProtocol.Name = *p.ProtocolVersion.Name
	outputProtocol.Copyright.Information = *p.ProtocolCopyright.Data
	outputProtocol.Description.Description = *p.ProtocolDescription.Description
	outputProtocol.Description.Summary = *p.ProtocolDescription.Summary
	for _, inputInterface := range p.Interface {
		outputInterface := Interface{}
		outputInterface.Name = *inputInterface.Name
		outputInterface.Version = *inputInterface.Version
		if inputInterface.Description != nil {
			if inputInterface.Description.Summary != nil {
				outputInterface.Description.Summary = *inputInterface.Description.Summary
			}
			if inputInterface.Description.Description != nil {
				outputInterface.Description.Description = *inputInterface.Description.Description
			}
		}
		for _, inputEvent := range inputInterface.Event {
			outputEvent := Method{}
			outputEvent.Name = *inputEvent.Name
			if inputEvent.Since != nil {
				outputEvent.Since = *inputEvent.Since
			}
			for _, inputArgument := range inputEvent.Arguments {
				outputArgument := Arg{}
				outputArgument.Name = *inputArgument.Name
				if inputArgument.Type.SingleType != nil {
					outputArgument.ID = *inputArgument.Type.SingleType
				} else {
					outputArgument.ID = *inputArgument.Type.DualType.Base
					outputArgument.InterfaceType = *inputArgument.Type.DualType.Sub
				}
				outputEvent.Args = append(outputEvent.Args, outputArgument)
			}
			outputInterface.Event = append(outputInterface.Event, outputEvent)
		}
		for _, inputRequest := range inputInterface.Request {
			outputRequest := Method{}
			outputRequest.Name = *inputRequest.Name
			if inputRequest.Since != nil {
				outputRequest.Since = *inputRequest.Since
			}
			for _, inputArgument := range inputRequest.Arguments {
				outputArgument := Arg{}
				outputArgument.Name = *inputArgument.Name
				if inputArgument.Type.SingleType != nil {
					outputArgument.ID = *inputArgument.Type.SingleType
				} else {
					outputArgument.ID = *inputArgument.Type.DualType.Base
					outputArgument.InterfaceType = *inputArgument.Type.DualType.Sub
				}
				outputRequest.Args = append(outputRequest.Args, outputArgument)
			}
			outputInterface.Request = append(outputInterface.Event, outputRequest)
		}
		for _, inputEnum := range inputInterface.Enums {
			outputEnum := Enum{}
			outputEnum.Name = *inputEnum.Name
			for _, inputEntry := range inputEnum.Values {
				outputEntry := Entry{}
				outputEntry.Name = *inputEntry.Name
				outputEntry.Value = *inputEntry.Value
				outputEnum.Entries = append(outputEnum.Entries, outputEntry)
			}
			outputInterface.Enum = append(outputInterface.Enum, outputEnum)
		}
		outputProtocol.Interfaces = append(outputProtocol.Interfaces, outputInterface)
	}
	return outputProtocol
}
