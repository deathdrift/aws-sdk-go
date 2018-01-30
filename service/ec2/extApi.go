package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const opExecuteInstanceTask = "ExecuteInstanceTask"

func (c *EC2) ExecuteInstanceTaskRequest(input *ExecuteInstanceTaskInput) (req *request.Request, output *ExecuteInstanceTaskOutput) {
	op := &request.Operation{
		Name:       opExecuteInstanceTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExecuteInstanceTaskInput{}
	}

	output = &ExecuteInstanceTaskOutput{}
	req = c.newRequest(op, input, output)
	return
}

func (c *EC2) ExecuteInstanceTask(input *ExecuteInstanceTaskInput) (*ExecuteInstanceTaskOutput, error) {
	req, out := c.ExecuteInstanceTaskRequest(input)
	return out, req.Send()
}

func (c *EC2) ExecuteInstanceTaskWithContext(ctx aws.Context, input *ExecuteInstanceTaskInput, opts ...request.Option) (*ExecuteInstanceTaskOutput, error) {
	req, out := c.ExecuteInstanceTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ExecuteInstanceTaskInput struct {
	_ struct{} `type:"structure"`

	DryRun *bool `locationName:"dryRun" type:"boolean"`

	InstanceId *string `locationName:"instanceId" type:"string" required:"true"`

	CommandType *string `locationName:"commandType" type:"string" required:"true"`

	CommandName *string `locationName:"commandName" type:"string" required:"true"`

	CommandArgs *string `locationName:"commandArgs" type:"string"`
}

// String returns the string representation
func (s ExecuteInstanceTaskInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ExecuteInstanceTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExecuteInstanceTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ExecuteInstanceTaskInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if s.CommandType == nil {
		invalidParams.Add(request.NewErrParamRequired("CommandType"))
	}

	if s.CommandName == nil {
		invalidParams.Add(request.NewErrParamRequired("CommandName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDryRun sets the DryRun field's value.
func (s *ExecuteInstanceTaskInput) SetDryRun(v bool) *ExecuteInstanceTaskInput {
	s.DryRun = &v
	return s
}

func (s *ExecuteInstanceTaskInput) SetInstanceId(v *string) *ExecuteInstanceTaskInput {
	s.InstanceId = v
	return s
}

func (s *ExecuteInstanceTaskInput) SetCommandType(v *string) *ExecuteInstanceTaskInput {
	s.CommandType = v
	return s
}

func (s *ExecuteInstanceTaskInput) SetCommandName(v *string) *ExecuteInstanceTaskInput {
	s.CommandName = v
	return s
}

func (s *ExecuteInstanceTaskInput) SetCommandArgs(v *string) *ExecuteInstanceTaskInput {
	s.CommandArgs = v
	return s
}

// Contains the output of ExecuteInstanceTask.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ExecuteInstanceTaskResult
type ExecuteInstanceTaskOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more terminated instances.
	Result *string `locationName:"result" type:"string"`
}

// String returns the string representation
func (s ExecuteInstanceTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ExecuteInstanceTaskOutput) GoString() string {
	return s.String()
}

// SetTerminatingInstances sets the TerminatingInstances field's value.
func (s *ExecuteInstanceTaskOutput) SetTerminatingInstances(v *string) *ExecuteInstanceTaskOutput {
	s.Result = v
	return s
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const opDescribeInstanceTypes = "DescribeInstanceTypes"

func (c *EC2) DescribeInstanceTypesRequest(input *DescribeInstanceTypesInput) (req *request.Request, output *DescribeInstanceTypesOutput) {
	op := &request.Operation{
		Name:       opDescribeInstanceTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstanceTypesInput{}
	}

	output = &DescribeInstanceTypesOutput{}
	req = c.newRequest(op, input, output)
	return
}

func (c *EC2) DescribeInstanceTypes(input *DescribeInstanceTypesInput) (*DescribeInstanceTypesOutput, error) {
	req, out := c.DescribeInstanceTypesRequest(input)
	return out, req.Send()
}

func (c *EC2) DescribeInstanceTypesWithContext(ctx aws.Context, input *DescribeInstanceTypesInput, opts ...request.Option) (*DescribeInstanceTypesOutput, error) {
	req, out := c.DescribeInstanceTypesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeInstanceTypesInput struct {
	_ struct{} `type:"structure"`

	DryRun *bool `locationName:"dryRun" type:"boolean"`

}

// String returns the string representation
func (s DescribeInstanceTypesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceTypesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstanceTypesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeInstanceTypesInput"}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDryRun sets the DryRun field's value.
func (s *DescribeInstanceTypesInput) SetDryRun(v bool) *DescribeInstanceTypesInput {
	s.DryRun = &v
	return s
}

// Contains the output of ExecuteInstanceTask.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ExecuteInstanceTaskResult
type DescribeInstanceTypesOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more terminated instances.
	InstanceTypeInfo []*InstanceType `locationName:"instanceTypeInfo" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeInstanceTypesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesOutput) SetInstanceTypeInfo(v []*InstanceType) *DescribeInstanceTypesOutput {
	s.InstanceTypeInfo = v
	return s
}

type InstanceType struct {
	_ struct{} `type:"structure"`

	InstanceType *string `locationName:"instanceType" type:"string"`

	DisplayName *string `locationName:"displayName" type:"string"`

	Cpu *int64 `locationName:"cpu" type:"integer"`

	Ram *int64 `locationName:"ram" type:"integer"`

	Disk *int64 `locationName:"disk" type:"integer"`

	Max *int64 `locationName:"max" type:"integer"`

	Available *int64 `locationName:"available" type:"integer"`
}

// String returns the string representation
func (s InstanceType) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceType) GoString() string {
	return s.String()
}

func (s *InstanceType) SetInstanceType(v string) *InstanceType {
	s.InstanceType = &v
	return s
}

func (s *InstanceType) SetDisplayName(v string) *InstanceType {
	s.DisplayName = &v
	return s
}

func (s *InstanceType) SetCpu(v int64) *InstanceType {
	s.Cpu = &v
	return s
}

func (s *InstanceType) SetRam(v int64) *InstanceType {
	s.Ram = &v
	return s
}

func (s *InstanceType) SetDisk(v int64) *InstanceType {
	s.Disk = &v
	return s
}

func (s *InstanceType) SetMax(v int64) *InstanceType {
	s.Max = &v
	return s
}

func (s *InstanceType) SetAvailable(v int64) *InstanceType {
	s.Available = &v
	return s
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const opDescribeClusters = "DescribeClusters"

func (c *EC2) DescribeClustersRequest(input *DescribeClustersInput) (req *request.Request, output *DescribeClustersOutput) {
	op := &request.Operation{
		Name:       opDescribeClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeClustersInput{}
	}

	output = &DescribeClustersOutput{}
	req = c.newRequest(op, input, output)
	return
}

func (c *EC2) DescribeClusters(input *DescribeClustersInput) (*DescribeClustersOutput, error) {
	req, out := c.DescribeClustersRequest(input)
	return out, req.Send()
}

func (c *EC2) DescribeClustersWithContext(ctx aws.Context, input *DescribeClustersInput, opts ...request.Option) (*DescribeClustersOutput, error) {
	req, out := c.DescribeClustersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeClustersInput struct {
	_ struct{} `type:"structure"`

	DryRun *bool `locationName:"dryRun" type:"boolean"`

}

// String returns the string representation
func (s DescribeClustersInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeClustersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClustersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeClustersInput"}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDryRun sets the DryRun field's value.
func (s *DescribeClustersInput) SetDryRun(v bool) *DescribeClustersInput {
	s.DryRun = &v
	return s
}

// Contains the output of ExecuteInstanceTask.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ExecuteInstanceTaskResult
type DescribeClustersOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more terminated instances.
	ClusterSet []*Cluster `locationName:"clusterSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeClustersOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeClustersOutput) GoString() string {
	return s.String()
}

func (s *DescribeClustersOutput) SetClusterSet(v []*Cluster) *DescribeClustersOutput {
	s.ClusterSet = v
	return s
}

type Cluster struct {
	_ struct{} `type:"structure"`

	ClusterId *string `locationName:"clusterId" type:"string"`

	ClusterName *string `locationName:"clusterName" type:"string"`

	Status *string `locationName:"status" type:"string"`

	Hypervisor *string `locationName:"hypervisor" type:"string"`

	SchedPolicy *string `locationName:"schedPolicy" type:"string"`

	MaxVolumeStorage *int64 `locationName:"maxVolumeStorage" type:"integer"`

	MaxEbsInstance *int64 `locationName:"maxEbsInstance" type:"integer"`

	MaxAttachedVolume *int64 `locationName:"maxAttachedVolume" type:"integer"`

	ExtendDiskMode *string `locationName:"extendDiskMode" type:"string"`

	CreateVolumeMode *string `locationName:"createVolumeMode" type:"string"`

	//ClusterControllerSet *string `locationName:"clusterControllerSet" type:"string"`
}

// String returns the string representation
func (s Cluster) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Cluster) GoString() string {
	return s.String()
}

func (s *Cluster) SetClusterId(v string) *Cluster {
	s.ClusterId = &v
	return s
}

func (s *Cluster) SetClusterName(v string) *Cluster {
	s.ClusterName = &v
	return s
}

func (s *Cluster) SetStatus(v string) *Cluster {
	s.Status = &v
	return s
}

func (s *Cluster) SetHypervisor(v string) *Cluster {
	s.Hypervisor = &v
	return s
}

func (s *Cluster) SetSchedPolicy(v string) *Cluster {
	s.SchedPolicy = &v
	return s
}

func (s *Cluster) SetMaxVolumeStorage(v int64) *Cluster {
	s.MaxVolumeStorage = &v
	return s
}

func (s *Cluster) SetMaxEbsInstance(v int64) *Cluster {
	s.MaxEbsInstance = &v
	return s
}

func (s *Cluster) SetMaxAttachedVolume(v int64) *Cluster {
	s.MaxAttachedVolume = &v
	return s
}

func (s *Cluster) SetCreateVolumeMode(v string) *Cluster {
	s.CreateVolumeMode = &v
	return s
}

//func (s *Cluster) SetClusterControllerSet(v int64) *InstanceType {
//	s.ClusterControllerSet = &v
//	return s
//}

