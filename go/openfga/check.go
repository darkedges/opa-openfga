package openfga

import (
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
	"github.com/openfga/go-sdk/client"
)

func RegisterCheck(fnName string, fgaClient *client.OpenFgaClient) (*rego.Function, rego.Builtin1) {
	return &rego.Function{
			Name:    fnName,
			Decl:    types.NewFunction(types.Args(types.A), types.B),
			Memoize: true,
		},
		func(bctx rego.BuiltinContext, op1 *ast.Term) (*ast.Term, error) {
			var args client.ClientCheckRequest
			if err := ast.As(op1.Value, &args); err != nil {
				return nil, err
			}
			checkResponse, err := fgaClient.Check(bctx.Context).Body(args).Execute()
			if err != nil {
				traceError(&bctx, fnName, err)
				return nil, err
			}
			return ast.BooleanTerm(*checkResponse.Allowed), nil
		}
}
