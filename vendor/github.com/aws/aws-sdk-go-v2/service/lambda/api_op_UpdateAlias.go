// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateAliasInput struct {
	_ struct{} `type:"structure"`

	// A description of the alias.
	Description *string `type:"string"`

	// The name of the Lambda function.
	//
	// Name formats
	//
	//    * Function name - MyFunction.
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//    * Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// The function version that the alias invokes.
	FunctionVersion *string `min:"1" type:"string"`

	// The name of the alias.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"Name" min:"1" type:"string" required:"true"`

	// Only update the alias if the revision ID matches the ID that's specified.
	// Use this option to avoid modifying an alias that has changed since you last
	// read it.
	RevisionId *string `type:"string"`

	// The routing configuration (https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html)
	// of the alias.
	RoutingConfig *AliasRoutingConfiguration `type:"structure"`
}

// String returns the string representation
func (s UpdateAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAliasInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.FunctionVersion != nil && len(*s.FunctionVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionVersion", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAliasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionVersion != nil {
		v := *s.FunctionVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoutingConfig != nil {
		v := s.RoutingConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RoutingConfig", v, metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Provides configuration information about a Lambda function alias (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
type UpdateAliasOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the alias.
	AliasArn *string `type:"string"`

	// A description of the alias.
	Description *string `type:"string"`

	// The function version that the alias invokes.
	FunctionVersion *string `min:"1" type:"string"`

	// The name of the alias.
	Name *string `min:"1" type:"string"`

	// A unique identifier that changes when you update the alias.
	RevisionId *string `type:"string"`

	// The routing configuration (https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html)
	// of the alias.
	RoutingConfig *AliasRoutingConfiguration `type:"structure"`
}

// String returns the string representation
func (s UpdateAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AliasArn != nil {
		v := *s.AliasArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AliasArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionVersion != nil {
		v := *s.FunctionVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoutingConfig != nil {
		v := s.RoutingConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RoutingConfig", v, metadata)
	}
	return nil
}

const opUpdateAlias = "UpdateAlias"

// UpdateAliasRequest returns a request value for making API operation for
// AWS Lambda.
//
// Updates the configuration of a Lambda function alias (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
//
//    // Example sending a request using UpdateAliasRequest.
//    req := client.UpdateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/UpdateAlias
func (c *Client) UpdateAliasRequest(input *UpdateAliasInput) UpdateAliasRequest {
	op := &aws.Operation{
		Name:       opUpdateAlias,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/aliases/{Name}",
	}

	if input == nil {
		input = &UpdateAliasInput{}
	}

	req := c.newRequest(op, input, &UpdateAliasOutput{})
	return UpdateAliasRequest{Request: req, Input: input, Copy: c.UpdateAliasRequest}
}

// UpdateAliasRequest is the request type for the
// UpdateAlias API operation.
type UpdateAliasRequest struct {
	*aws.Request
	Input *UpdateAliasInput
	Copy  func(*UpdateAliasInput) UpdateAliasRequest
}

// Send marshals and sends the UpdateAlias API request.
func (r UpdateAliasRequest) Send(ctx context.Context) (*UpdateAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAliasResponse{
		UpdateAliasOutput: r.Request.Data.(*UpdateAliasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAliasResponse is the response type for the
// UpdateAlias API operation.
type UpdateAliasResponse struct {
	*UpdateAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAlias request.
func (r *UpdateAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
