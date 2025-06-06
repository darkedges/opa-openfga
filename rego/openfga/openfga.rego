package openfga

import rego.v1

allow if {
	openfga.check({
		"user": input.user_1,
		"relation": "viewer",
		"object": input.object,
	})
}

allow2 if {
	red := openfga.batchcheck({"checks": [
		{
			"user": input.user_1,
			"relation": "viewer",
			"object": input.object,
			"correlation_id": input.correlation_id_1,
		},
		{
			"user": input.user_2,
			"relation": "viewer",
			"object": input.object,
			"correlation_id": input.correlation_id_2,
		},
	]})
	red.result[input.correlation_id_1].allowed == true
	red.result[input.correlation_id_2].allowed == false
}
