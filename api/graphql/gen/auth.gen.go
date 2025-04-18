// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type MutationResolver interface {
	SendConfirmationCode(ctx context.Context, input SendConfirmationCodeInput) (*VoidPayload, error)
	VerifyConfirmationCode(ctx context.Context, input VerifyConfirmationCodeInput) (*VoidPayload, error)
	SecurityImage(ctx context.Context, input SecurityImageInput) (*SecurityImagePayload, error)
	ProcessFirebaseAuth(ctx context.Context, input ProcessFirebaseAuthInput) (*VoidPayload, error)
}
type QueryResolver interface {
	Dummy(ctx context.Context) (*string, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_processFirebaseAuth_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Mutation_processFirebaseAuth_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_processFirebaseAuth_argsInput(
	ctx context.Context,
	rawArgs map[string]any,
) (ProcessFirebaseAuthInput, error) {
	if _, ok := rawArgs["input"]; !ok {
		var zeroVal ProcessFirebaseAuthInput
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNProcessFirebaseAuthInput2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐProcessFirebaseAuthInput(ctx, tmp)
	}

	var zeroVal ProcessFirebaseAuthInput
	return zeroVal, nil
}

func (ec *executionContext) field_Mutation_securityImage_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Mutation_securityImage_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_securityImage_argsInput(
	ctx context.Context,
	rawArgs map[string]any,
) (SecurityImageInput, error) {
	if _, ok := rawArgs["input"]; !ok {
		var zeroVal SecurityImageInput
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNSecurityImageInput2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐSecurityImageInput(ctx, tmp)
	}

	var zeroVal SecurityImageInput
	return zeroVal, nil
}

func (ec *executionContext) field_Mutation_sendConfirmationCode_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Mutation_sendConfirmationCode_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_sendConfirmationCode_argsInput(
	ctx context.Context,
	rawArgs map[string]any,
) (SendConfirmationCodeInput, error) {
	if _, ok := rawArgs["input"]; !ok {
		var zeroVal SendConfirmationCodeInput
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNSendConfirmationCodeInput2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐSendConfirmationCodeInput(ctx, tmp)
	}

	var zeroVal SendConfirmationCodeInput
	return zeroVal, nil
}

func (ec *executionContext) field_Mutation_verifyConfirmationCode_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Mutation_verifyConfirmationCode_argsInput(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["input"] = arg0
	return args, nil
}
func (ec *executionContext) field_Mutation_verifyConfirmationCode_argsInput(
	ctx context.Context,
	rawArgs map[string]any,
) (VerifyConfirmationCodeInput, error) {
	if _, ok := rawArgs["input"]; !ok {
		var zeroVal VerifyConfirmationCodeInput
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
	if tmp, ok := rawArgs["input"]; ok {
		return ec.unmarshalNVerifyConfirmationCodeInput2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐVerifyConfirmationCodeInput(ctx, tmp)
	}

	var zeroVal VerifyConfirmationCodeInput
	return zeroVal, nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Query___type_argsName(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["name"] = arg0
	return args, nil
}
func (ec *executionContext) field_Query___type_argsName(
	ctx context.Context,
	rawArgs map[string]any,
) (string, error) {
	if _, ok := rawArgs["name"]; !ok {
		var zeroVal string
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
	if tmp, ok := rawArgs["name"]; ok {
		return ec.unmarshalNString2string(ctx, tmp)
	}

	var zeroVal string
	return zeroVal, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Mutation_sendConfirmationCode(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_sendConfirmationCode(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().SendConfirmationCode(rctx, fc.Args["input"].(SendConfirmationCodeInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*VoidPayload)
	fc.Result = res
	return ec.marshalNVoidPayload2ᚖgithubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐVoidPayload(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_sendConfirmationCode(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "ok":
				return ec.fieldContext_VoidPayload_ok(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type VoidPayload", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_sendConfirmationCode_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_verifyConfirmationCode(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_verifyConfirmationCode(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().VerifyConfirmationCode(rctx, fc.Args["input"].(VerifyConfirmationCodeInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*VoidPayload)
	fc.Result = res
	return ec.marshalNVoidPayload2ᚖgithubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐVoidPayload(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_verifyConfirmationCode(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "ok":
				return ec.fieldContext_VoidPayload_ok(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type VoidPayload", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_verifyConfirmationCode_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_securityImage(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_securityImage(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().SecurityImage(rctx, fc.Args["input"].(SecurityImageInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*SecurityImagePayload)
	fc.Result = res
	return ec.marshalNSecurityImagePayload2ᚖgithubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐSecurityImagePayload(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_securityImage(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "image":
				return ec.fieldContext_SecurityImagePayload_image(ctx, field)
			case "phrase":
				return ec.fieldContext_SecurityImagePayload_phrase(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type SecurityImagePayload", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_securityImage_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_processFirebaseAuth(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_processFirebaseAuth(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().ProcessFirebaseAuth(rctx, fc.Args["input"].(ProcessFirebaseAuthInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*VoidPayload)
	fc.Result = res
	return ec.marshalNVoidPayload2ᚖgithubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐVoidPayload(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_processFirebaseAuth(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "ok":
				return ec.fieldContext_VoidPayload_ok(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type VoidPayload", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_processFirebaseAuth_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Query__dummy(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query__dummy(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Dummy(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query__dummy(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(fc.Args["name"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "isOneOf":
				return ec.fieldContext___Type_isOneOf(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query___type_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___schema(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___schema(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "description":
				return ec.fieldContext___Schema_description(ctx, field)
			case "types":
				return ec.fieldContext___Schema_types(ctx, field)
			case "queryType":
				return ec.fieldContext___Schema_queryType(ctx, field)
			case "mutationType":
				return ec.fieldContext___Schema_mutationType(ctx, field)
			case "subscriptionType":
				return ec.fieldContext___Schema_subscriptionType(ctx, field)
			case "directives":
				return ec.fieldContext___Schema_directives(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Schema", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _SecurityImagePayload_image(ctx context.Context, field graphql.CollectedField, obj *SecurityImagePayload) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SecurityImagePayload_image(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Image, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SecurityImagePayload_image(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SecurityImagePayload",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _SecurityImagePayload_phrase(ctx context.Context, field graphql.CollectedField, obj *SecurityImagePayload) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_SecurityImagePayload_phrase(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Phrase, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_SecurityImagePayload_phrase(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "SecurityImagePayload",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputProcessFirebaseAuthInput(ctx context.Context, obj any) (ProcessFirebaseAuthInput, error) {
	var it ProcessFirebaseAuthInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"token"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "token":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("token"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Token = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputSecurityImageInput(ctx context.Context, obj any) (SecurityImageInput, error) {
	var it SecurityImageInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"email"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "email":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Email = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputSendConfirmationCodeInput(ctx context.Context, obj any) (SendConfirmationCodeInput, error) {
	var it SendConfirmationCodeInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"method", "recipient"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "method":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("method"))
			data, err := ec.unmarshalNDeliveryMethod2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐDeliveryMethod(ctx, v)
			if err != nil {
				return it, err
			}
			it.Method = data
		case "recipient":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("recipient"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Recipient = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputVerifyConfirmationCodeInput(ctx context.Context, obj any) (VerifyConfirmationCodeInput, error) {
	var it VerifyConfirmationCodeInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"recipient", "code"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "recipient":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("recipient"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Recipient = data
		case "code":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("code"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Code = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "sendConfirmationCode":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_sendConfirmationCode(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "verifyConfirmationCode":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_verifyConfirmationCode(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "securityImage":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_securityImage(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "processFirebaseAuth":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_processFirebaseAuth(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, queryImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "_dummy":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query__dummy(ctx, field)
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx,
					func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return rrm(innerCtx) })
		case "__type":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___type(ctx, field)
			})
		case "__schema":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___schema(ctx, field)
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var securityImagePayloadImplementors = []string{"SecurityImagePayload"}

func (ec *executionContext) _SecurityImagePayload(ctx context.Context, sel ast.SelectionSet, obj *SecurityImagePayload) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, securityImagePayloadImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("SecurityImagePayload")
		case "image":
			out.Values[i] = ec._SecurityImagePayload_image(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "phrase":
			out.Values[i] = ec._SecurityImagePayload_phrase(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNDeliveryMethod2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐDeliveryMethod(ctx context.Context, v any) (DeliveryMethod, error) {
	var res DeliveryMethod
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNDeliveryMethod2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐDeliveryMethod(ctx context.Context, sel ast.SelectionSet, v DeliveryMethod) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNProcessFirebaseAuthInput2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐProcessFirebaseAuthInput(ctx context.Context, v any) (ProcessFirebaseAuthInput, error) {
	res, err := ec.unmarshalInputProcessFirebaseAuthInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNSecurityImageInput2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐSecurityImageInput(ctx context.Context, v any) (SecurityImageInput, error) {
	res, err := ec.unmarshalInputSecurityImageInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNSecurityImagePayload2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐSecurityImagePayload(ctx context.Context, sel ast.SelectionSet, v SecurityImagePayload) graphql.Marshaler {
	return ec._SecurityImagePayload(ctx, sel, &v)
}

func (ec *executionContext) marshalNSecurityImagePayload2ᚖgithubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐSecurityImagePayload(ctx context.Context, sel ast.SelectionSet, v *SecurityImagePayload) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._SecurityImagePayload(ctx, sel, v)
}

func (ec *executionContext) unmarshalNSendConfirmationCodeInput2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐSendConfirmationCodeInput(ctx context.Context, v any) (SendConfirmationCodeInput, error) {
	res, err := ec.unmarshalInputSendConfirmationCodeInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNVerifyConfirmationCodeInput2githubᚗcomᚋezexᚑioᚋezexᚑgatewayᚋapiᚋgraphqlᚋgenᚐVerifyConfirmationCodeInput(ctx context.Context, v any) (VerifyConfirmationCodeInput, error) {
	res, err := ec.unmarshalInputVerifyConfirmationCodeInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
