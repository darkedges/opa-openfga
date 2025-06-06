package openfga

import (
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
	"github.com/openfga/go-sdk/client"
)

func RegisterBatchCheck(fnName string, fgaClient *client.OpenFgaClient) (*rego.Function, rego.Builtin1) {
	return &rego.Function{
			Name:             fnName,
			Decl:             types.NewFunction(types.Args(types.A), types.A),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, op1 *ast.Term) (*ast.Term, error) {
			var args client.ClientBatchCheckRequest
			if err := ast.As(op1.Value, &args); err != nil {
				return nil, err
			}
			checkResponse, err := fgaClient.BatchCheck(bctx.Context).Body(args).Execute()
			if err != nil {
				traceError(&bctx, fnName, err)
				return nil, err
			}
			v, err := ast.InterfaceToValue(checkResponse)
			if err != nil {
				return nil, err
			}
			return ast.NewTerm(v), nil
		}
}
