package cloudify_v1_3

import (
	"github.com/tliron/puccini/tosca"
)

//
// RelationshipType
//
// [https://docs.cloudify.co/4.5.5/developer/blueprints/spec-relationships/]
//

type RelationshipType struct {
	*Type `name:"relationship type"`

	SourceInterfaceDefinitions InterfaceDefinitions `read:"source_interfaces,InterfaceDefinition" inherit:"source_interfaces,Parent"`
	TargetInterfaceDefinitions InterfaceDefinitions `read:"target_interfaces,InterfaceDefinition" inherit:"target_interfaces,Parent"`
	PropertyDefinitions        PropertyDefinitions  `read:"properties,PropertyDefinition" inherit:"properties,Parent"`

	Parent *RelationshipType `lookup:"derived_from,ParentName" json:"-" yaml:"-"`
}

func NewRelationshipType(context *tosca.Context) *RelationshipType {
	return &RelationshipType{
		Type:                       NewType(context),
		SourceInterfaceDefinitions: make(InterfaceDefinitions),
		TargetInterfaceDefinitions: make(InterfaceDefinitions),
		PropertyDefinitions:        make(PropertyDefinitions),
	}
}

// tosca.Reader signature
func ReadRelationshipType(context *tosca.Context) interface{} {
	self := NewRelationshipType(context)
	context.ValidateUnsupportedFields(context.ReadFields(self))
	return self
}

// tosca.Hierarchical interface
func (self *RelationshipType) GetParent() interface{} {
	return self.Parent
}