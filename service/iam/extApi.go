package iam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opGetAccountPolicy = "GetAccountPolicy"

func (c *IAM) GetAccountPolicyRequest(input *GetAccountPolicyInput) (req *request.Request, output *GetAccountPolicyOutput) {
	op := &request.Operation{
		Name:       opGetAccountPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAccountPolicyInput{}
	}

	output = &GetAccountPolicyOutput{}
	req = c.newRequest(op, input, output)
	return
}

func (c *IAM) GetAccountPolicy(input *GetAccountPolicyInput) (*GetAccountPolicyOutput, error) {
	req, out := c.GetAccountPolicyRequest(input)
	return out, req.Send()
}

func (c *IAM) GetAccountPolicyWithContext(ctx aws.Context, input *GetAccountPolicyInput, opts ...request.Option) (*GetAccountPolicyOutput, error) {
	req, out := c.GetAccountPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetAccountPolicyInput struct {
	_ struct{} `type:"structure"`

	DryRun *bool `locationName:"dryRun" type:"boolean"`

	UserId *string `locationName:"userId" type:"string" required:"true"`

	AccessKeyId *string `locationName:"accessKeyId" type:"string" required:"false"`

	AuthParams *string `locationName:"authParams" type:"string" required:"false"`
}

// String returns the string representation
func (s GetAccountPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAccountPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccountPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetAccountPolicyInput"}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDryRun sets the DryRun field's value.
func (s *GetAccountPolicyInput) SetDryRun(v bool) *GetAccountPolicyInput {
	s.DryRun = &v
	return s
}

// Contains the output of ExecuteInstanceTask.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ExecuteInstanceTaskResult
type GetAccountPolicyOutput struct {
	_ struct{} `type:"structure"`

	UserId *string `locationName:"UserId" type:"string""`
	// Information about one or more terminated instances.
	PolicySets []*Policy `locationName:"PolicySets" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s GetAccountPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAccountPolicyOutput) GoString() string {
	return s.String()
}

func (s *GetAccountPolicyOutput) SetPolicySets(v []*Policy) *GetAccountPolicyOutput {
	s.PolicySets = v
	return s
}

//type Policy struct {
//	_ struct{} `type:"structure"`
//
//	PolicyName *string `locationName:"policyName" type:"string"`
//
//	PolicyDocument *string `locationName:"policyDocument" type:"string"`
//}

//String returns the string representation
//func (s Policy) String() string {
//	return awsutil.Prettify(s)
//}
//
//// GoString returns the string representation
//func (s Policy) GoString() string {
//	return s.String()
//}
//
//func (s *Policy) SetPolicyName(v string) *Policy {
//	s.PolicyName = &v
//	return s
//}
//
//func (s *Policy) SetPolicyDocument(v string) *Policy {
//	s.PolicyDocument = &v
//	return s
//}
