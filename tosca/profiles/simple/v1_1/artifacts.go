// This file was auto-generated from a YAML file

package v1_1

func init() {
	Profile["/tosca/simple/1.1/artifacts.yaml"] = `
# Modified from a file that was distributed with this NOTICE:
#
#   Apache AriaTosca
#   Copyright 2016-2017 The Apache Software Foundation
#
#   This product includes software developed at
#   The Apache Software Foundation (http://www.apache.org/).

tosca_definitions_version: tosca_simple_yaml_1_1

artifact_types:

  tosca.artifacts.Root:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      citation_location: 5.4.1
    description: >-
      This is the default (root) TOSCA Artifact Type definition that all other TOSCA base Artifact
      Types derive from.

  tosca.artifacts.File:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      citation_location: 5.4.2
    description: >-
      This artifact type is used when an artifact definition needs to have its associated file
      simply treated as a file and no special handling/handlers are invoked (i.e., it is not treated
      as either an implementation or deployment artifact type).
    derived_from: tosca.artifacts.Root

  #
  # Deployments
  #

  tosca.artifacts.Deployment:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      citation_location: 5.4.3.1
    description: >-
      This artifact type represents the parent type for all deployment artifacts in TOSCA. This
      class of artifacts typically represents a binary packaging of an application or service that
      is used to install/create or deploy it as part of a node's lifecycle.
    derived_from: tosca.artifacts.Root

  tosca.artifacts.Deployment.Image:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      citation_location: 5.4.3.3
    description: >-
      This artifact type represents a parent type for any "image" which is an opaque packaging of a
      TOSCA Node's deployment (whether real or virtual) whose contents are typically already
      installed and pre-configured (i.e., "stateful") and prepared to be run on a known target
      container.
    derived_from: tosca.artifacts.Deployment

  tosca.artifacts.Deployment.Image.VM:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      citation_location: 5.4.3.4
    description: >-
      This artifact represents the parent type for all Virtual Machine (VM) image and container
      formatted deployment artifacts. These images contain a stateful capture of a machine (e.g.,
      server) including operating system and installed software along with any configurations and
      can be run on another machine using a hypervisor which virtualizes typical server (i.e.,
      hardware) resources.
    derived_from: tosca.artifacts.Deployment

  #
  # Implementations
  #

  tosca.artifacts.Implementation:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      citation_location: 5.4.4.1
    description: >-
      This artifact type represents the parent type for all implementation artifacts in TOSCA. These
      artifacts are used to implement operations of TOSCA interfaces either directly (e.g., scripts)
      or indirectly (e.g., config. files).
    derived_from: tosca.artifacts.Root

  tosca.artifacts.Implementation.Bash:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      citation_location: 5.4.4.3
    description: >-
      This artifact type represents a Bash script type that contains Bash commands that can be
      executed on the Unix Bash shell.
    derived_from: tosca.artifacts.Implementation
    mime_type: application/x-sh
    file_ext: [ sh ]

  tosca.artifacts.Implementation.Python:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      citation_location: 5.4.4.4
    description: >-
      This artifact type represents a Python file that contains Python language constructs that can
      be executed within a Python interpreter.
    derived_from: tosca.artifacts.Implementation
    mime_type: application/x-python
    file_ext: [ py ]
`
}
