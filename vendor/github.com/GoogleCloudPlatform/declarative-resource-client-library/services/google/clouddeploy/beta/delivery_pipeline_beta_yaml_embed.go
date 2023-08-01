// Copyright 2023 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package beta -var YAML_delivery_pipeline blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/clouddeploy/beta/delivery_pipeline.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/clouddeploy/beta/delivery_pipeline.yaml
var YAML_delivery_pipeline = []byte("info:\n  title: Clouddeploy/DeliveryPipeline\n  description: The Cloud Deploy `DeliveryPipeline` resource\n  x-dcl-struct-name: DeliveryPipeline\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: REST API\n    url: https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines\npaths:\n  get:\n    description: The function used to get information about a DeliveryPipeline\n    parameters:\n    - name: deliveryPipeline\n      required: true\n      description: A full instance of a DeliveryPipeline\n  apply:\n    description: The function used to apply information about a DeliveryPipeline\n    parameters:\n    - name: deliveryPipeline\n      required: true\n      description: A full instance of a DeliveryPipeline\n  delete:\n    description: The function used to delete a DeliveryPipeline\n    parameters:\n    - name: deliveryPipeline\n      required: true\n      description: A full instance of a DeliveryPipeline\n  deleteAll:\n    description: The function used to delete all DeliveryPipeline\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many DeliveryPipeline\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    DeliveryPipeline:\n      title: DeliveryPipeline\n      x-dcl-id: projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: User annotations. These attributes can only be set and used\n            by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations\n            for more details such as format and size limitations.\n        condition:\n          type: object\n          x-dcl-go-name: Condition\n          x-dcl-go-type: DeliveryPipelineCondition\n          readOnly: true\n          description: Output only. Information around the state of the Delivery Pipeline.\n          properties:\n            pipelineReadyCondition:\n              type: object\n              x-dcl-go-name: PipelineReadyCondition\n              x-dcl-go-type: DeliveryPipelineConditionPipelineReadyCondition\n              description: Details around the Pipeline's overall status.\n              properties:\n                status:\n                  type: boolean\n                  x-dcl-go-name: Status\n                  description: True if the Pipeline is in a valid state. Otherwise\n                    at least one condition in `PipelineCondition` is in an invalid\n                    state. Iterate over those conditions and see which condition(s)\n                    has status = false to find out what is wrong with the Pipeline.\n                updateTime:\n                  type: string\n                  format: date-time\n                  x-dcl-go-name: UpdateTime\n                  description: Last time the condition was updated.\n            targetsPresentCondition:\n              type: object\n              x-dcl-go-name: TargetsPresentCondition\n              x-dcl-go-type: DeliveryPipelineConditionTargetsPresentCondition\n              description: Details around targets enumerated in the pipeline.\n              properties:\n                missingTargets:\n                  type: array\n                  x-dcl-go-name: MissingTargets\n                  description: The list of Target names that are missing. For example,\n                    projects/{project_id}/locations/{location_name}/targets/{target_name}.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n                    x-dcl-references:\n                    - resource: Clouddeploy/Target\n                      field: selfLink\n                status:\n                  type: boolean\n                  x-dcl-go-name: Status\n                  description: True if there aren't any missing Targets.\n                updateTime:\n                  type: string\n                  format: date-time\n                  x-dcl-go-name: UpdateTime\n                  description: Last time the condition was updated.\n            targetsTypeCondition:\n              type: object\n              x-dcl-go-name: TargetsTypeCondition\n              x-dcl-go-type: DeliveryPipelineConditionTargetsTypeCondition\n              description: Details on the whether the targets enumerated in the pipeline\n                are of the same type.\n              properties:\n                errorDetails:\n                  type: string\n                  x-dcl-go-name: ErrorDetails\n                  description: Human readable error message.\n                status:\n                  type: boolean\n                  x-dcl-go-name: Status\n                  description: True if the targets are all a comparable type. For\n                    example this is true if all targets are GKE clusters. This is\n                    false if some targets are Cloud Run targets and others are GKE\n                    clusters.\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Time at which the pipeline was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Description of the `DeliveryPipeline`. Max length is 255 characters.\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: This checksum is computed by the server based on the value\n            of other fields, and may be sent on update and delete requests to ensure\n            the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: 'Labels are attributes that can be set and used by both the\n            user and by Google Cloud Deploy. Labels must meet the following constraints:\n            * Keys and values can contain only lowercase letters, numeric characters,\n            underscores, and dashes. * All characters must use UTF-8 encoding, and\n            international characters are allowed. * Keys must start with a lowercase\n            letter or international character. * Each resource is limited to a maximum\n            of 64 labels. Both keys and values are additionally constrained to be\n            <= 128 bytes.'\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the `DeliveryPipeline`. Format is [a-z][a-z0-9\\-]{0,62}.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        serialPipeline:\n          type: object\n          x-dcl-go-name: SerialPipeline\n          x-dcl-go-type: DeliveryPipelineSerialPipeline\n          description: SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.\n          properties:\n            stages:\n              type: array\n              x-dcl-go-name: Stages\n              description: Each stage specifies configuration for a `Target`. The\n                ordering of this list defines the promotion flow.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: DeliveryPipelineSerialPipelineStages\n                properties:\n                  profiles:\n                    type: array\n                    x-dcl-go-name: Profiles\n                    description: Skaffold profiles to use when rendering the manifest\n                      for this stage's `Target`.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: string\n                      x-dcl-go-type: string\n                  strategy:\n                    type: object\n                    x-dcl-go-name: Strategy\n                    x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategy\n                    description: Optional. The strategy to use for a `Rollout` to\n                      this stage.\n                    properties:\n                      canary:\n                        type: object\n                        x-dcl-go-name: Canary\n                        x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyCanary\n                        description: Canary deployment strategy provides progressive\n                          percentage based deployments to a Target.\n                        properties:\n                          canaryDeployment:\n                            type: object\n                            x-dcl-go-name: CanaryDeployment\n                            x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment\n                            description: Configures the progressive based deployment\n                              for a Target.\n                            x-dcl-conflicts:\n                            - customCanaryDeployment\n                            required:\n                            - percentages\n                            properties:\n                              percentages:\n                                type: array\n                                x-dcl-go-name: Percentages\n                                description: Required. The percentage based deployments\n                                  that will occur as a part of a `Rollout`. List is\n                                  expected in ascending order and each integer n is\n                                  0 <= n < 100.\n                                x-dcl-send-empty: true\n                                x-dcl-list-type: list\n                                items:\n                                  type: integer\n                                  format: int64\n                                  x-dcl-go-type: int64\n                              verify:\n                                type: boolean\n                                x-dcl-go-name: Verify\n                                description: Whether to run verify tests after each\n                                  percentage deployment.\n                          customCanaryDeployment:\n                            type: object\n                            x-dcl-go-name: CustomCanaryDeployment\n                            x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment\n                            description: Configures the progressive based deployment\n                              for a Target, but allows customizing at the phase level\n                              where a phase represents each of the percentage deployments.\n                            x-dcl-conflicts:\n                            - canaryDeployment\n                            required:\n                            - phaseConfigs\n                            properties:\n                              phaseConfigs:\n                                type: array\n                                x-dcl-go-name: PhaseConfigs\n                                description: Required. Configuration for each phase\n                                  in the canary deployment in the order executed.\n                                x-dcl-send-empty: true\n                                x-dcl-list-type: list\n                                items:\n                                  type: object\n                                  x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs\n                                  required:\n                                  - phaseId\n                                  - percentage\n                                  properties:\n                                    percentage:\n                                      type: integer\n                                      format: int64\n                                      x-dcl-go-name: Percentage\n                                      description: Required. Percentage deployment\n                                        for the phase.\n                                    phaseId:\n                                      type: string\n                                      x-dcl-go-name: PhaseId\n                                      description: 'Required. The ID to assign to\n                                        the `Rollout` phase. This value must consist\n                                        of lower-case letters, numbers, and hyphens,\n                                        start with a letter and end with a letter\n                                        or a number, and have a max length of 63 characters.\n                                        In other words, it must match the following\n                                        regex: `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.'\n                                    profiles:\n                                      type: array\n                                      x-dcl-go-name: Profiles\n                                      description: Skaffold profiles to use when rendering\n                                        the manifest for this phase. These are in\n                                        addition to the profiles list specified in\n                                        the `DeliveryPipeline` stage.\n                                      x-dcl-send-empty: true\n                                      x-dcl-list-type: list\n                                      items:\n                                        type: string\n                                        x-dcl-go-type: string\n                                    verify:\n                                      type: boolean\n                                      x-dcl-go-name: Verify\n                                      description: Whether to run verify tests after\n                                        the deployment.\n                          runtimeConfig:\n                            type: object\n                            x-dcl-go-name: RuntimeConfig\n                            x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig\n                            description: Optional. Runtime specific configurations\n                              for the deployment strategy. The runtime configuration\n                              is used to determine how Cloud Deploy will split traffic\n                              to enable a progressive deployment.\n                            properties:\n                              cloudRun:\n                                type: object\n                                x-dcl-go-name: CloudRun\n                                x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun\n                                description: Cloud Run runtime configuration.\n                                x-dcl-conflicts:\n                                - kubernetes\n                                properties:\n                                  automaticTrafficControl:\n                                    type: boolean\n                                    x-dcl-go-name: AutomaticTrafficControl\n                                    description: Whether Cloud Deploy should update\n                                      the traffic stanza in a Cloud Run Service on\n                                      the user's behalf to facilitate traffic splitting.\n                                      This is required to be true for CanaryDeployments,\n                                      but optional for CustomCanaryDeployments.\n                              kubernetes:\n                                type: object\n                                x-dcl-go-name: Kubernetes\n                                x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes\n                                description: Kubernetes runtime configuration.\n                                x-dcl-conflicts:\n                                - cloudRun\n                                properties:\n                                  gatewayServiceMesh:\n                                    type: object\n                                    x-dcl-go-name: GatewayServiceMesh\n                                    x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh\n                                    description: Kubernetes Gateway API service mesh\n                                      configuration.\n                                    x-dcl-conflicts:\n                                    - serviceNetworking\n                                    required:\n                                    - httpRoute\n                                    - service\n                                    - deployment\n                                    properties:\n                                      deployment:\n                                        type: string\n                                        x-dcl-go-name: Deployment\n                                        description: Required. Name of the Kubernetes\n                                          Deployment whose traffic is managed by the\n                                          specified HTTPRoute and Service.\n                                      httpRoute:\n                                        type: string\n                                        x-dcl-go-name: HttpRoute\n                                        description: Required. Name of the Gateway\n                                          API HTTPRoute.\n                                      service:\n                                        type: string\n                                        x-dcl-go-name: Service\n                                        description: Required. Name of the Kubernetes\n                                          Service.\n                                  serviceNetworking:\n                                    type: object\n                                    x-dcl-go-name: ServiceNetworking\n                                    x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking\n                                    description: Kubernetes Service networking configuration.\n                                    x-dcl-conflicts:\n                                    - gatewayServiceMesh\n                                    required:\n                                    - service\n                                    - deployment\n                                    properties:\n                                      deployment:\n                                        type: string\n                                        x-dcl-go-name: Deployment\n                                        description: Required. Name of the Kubernetes\n                                          Deployment whose traffic is managed by the\n                                          specified Service.\n                                      service:\n                                        type: string\n                                        x-dcl-go-name: Service\n                                        description: Required. Name of the Kubernetes\n                                          Service.\n                      standard:\n                        type: object\n                        x-dcl-go-name: Standard\n                        x-dcl-go-type: DeliveryPipelineSerialPipelineStagesStrategyStandard\n                        description: Standard deployment strategy executes a single\n                          deploy and allows verifying the deployment.\n                        properties:\n                          verify:\n                            type: boolean\n                            x-dcl-go-name: Verify\n                            description: Whether to verify a deployment.\n                  targetId:\n                    type: string\n                    x-dcl-go-name: TargetId\n                    description: The target_id to which this stage points. This field\n                      refers exclusively to the last segment of a target name. For\n                      example, this field would just be `my-target` (rather than `projects/project/locations/location/targets/my-target`).\n                      The location of the `Target` is inferred to be the same as the\n                      location of the `DeliveryPipeline` that contains this `Stage`.\n        suspended:\n          type: boolean\n          x-dcl-go-name: Suspended\n          description: When suspended, no new releases or rollouts can be created,\n            but in-progress ones will complete.\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. Unique identifier of the `DeliveryPipeline`.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Most recent time at which the pipeline was updated.\n          x-kubernetes-immutable: true\n")

// 22297 bytes
// MD5: 6ac8060b3a6d65bfc00e2c1ec0eba06f
