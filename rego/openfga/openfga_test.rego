package openfga_test

import data.openfga
import rego.v1

test_get_user_allowed if {
	openfga.allow with input as {
		"user_1": "user:auth0|682a7b2a796219fd71d9441e",
		"object": "document:003dM000003m357QAA",
	}
}

test_get_another_user_denied if {
	not openfga.allow with input as {
		"user_1": "user:auth0|682a7b2a796219fd71d9441e",
		"object": "document:003dM000003m33VQAQ",
	}
}

test_get_batch_user_allowed if {
	openfga.allow2 with input as {
		"user_1": "user:auth0|682a7b2a796219fd71d9441e",
		"object": "document:003dM000003m357QAA",
		"correlation_id_1": "886224f6-04ae-4b13-bd8e-559c7d3754e1",
		"user_2": "user:auth0|1234567890a1b2c3d4e5f6g7",
		"correlation_id_2": "3ac7ab9-36de-471f-a2ee-d14bccad0e3d",
	}
}
