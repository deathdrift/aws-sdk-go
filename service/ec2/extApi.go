package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

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
