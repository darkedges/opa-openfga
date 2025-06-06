# opa-openfga

## Pre-Requisistes

Requires `go1.24.2`

```console
cd /tmp
wget https://go.dev/dl/go1.24.3.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.24.3.linux-amd64.tar.gz
sudo chmod 777 /usr/local/go/bin
cat <<EOF>>~/.bash_profile
export PATH=/usr/local/go/bin/:$PATH
EOF
source ~/.bash_profile
go env -w GOBIN=/tmp/go
go install github.com/openfga/openfga/cmd/openfga@latest
go install github.com/openfga/cli/cmd/fga@latest
```

## Configure and start openFGA

1. start `openfga` in one console

   ```console
   openfga run
   ```

2. Configure `openfga` in another console

    ```console
    export FGA_API_URL=http://localhost:8080
    export FGA_STORE_ID=$(fga store create --name "FGA Demo Store" | jq -r .store.id)
    export FGA_MODEL_ID=$(fga model write --file rego/openfga/model.fga | jq -r .authorization_model_id)
    fga tuple write 'user:auth0|682a7b2a796219fd71d9441e' viewer 'document:003dM000003m357QAA'
    fga query check 'user:auth0|682a7b2a796219fd71d9441e' viewer 'document:003dM000003m357QAA'
    ```

## Build

```console
cd go
go build -o opa++
```

## Test 1

```console
./opa++ eval -d ../rego/openfga/openfga.rego -i ../rego/input.json  'data' -f raw | jq .
```

should return

```json
{
  "openfga": {
    "allow": true,
    "allow2": true
  }
}
```

## Test 2

```console
./opa++ test ../rego/ -v
```

should return

```console
../rego/openfga/openfga_test.rego:
data.openfga_test.test_get_user_allowed: PASS (38.574804ms)
data.openfga_test.test_get_another_user_denied: PASS (1.214779ms)
data.openfga_test.test_get_batch_user_allowed: PASS (1.503659ms)
--------------------------------------------------------------------------------
PASS: 3/3
```

## Manual Run

1. Open a OPA Console

   ```console
   ./opa++ run
   ```

1. Issue the following query

    ```console
    openfga.check({"user":"user:auth0|682a7b2a796219fd71d9441e","relation":"viewer","object":"document:003dM000003m357QAA"})
    ```

    should return

    ```console
    true
    ```

1. Issue the following query

    ```console
    openfga.batchcheck({"checks": [{"user": "user:auth0|682a7b2a796219fd71d9441e","relation": "viewer","object": "document:003dM000003m357QAA","correlation_id": "886224f6-04ae-4b13-bd8e-559c7d3754e1"},{"user": "user:auth0|1234567890a1b2c3d4e5f6g7","relation": "viewer","object": "document:003dM000003m357QAA","correlation_id": "3ac7ab9-36de-471f-a2ee-d14bccad0e3d"},]})
    ```

    should return

    ```console
    {
      "result": {
        "3ac7ab9-36de-471f-a2ee-d14bccad0e3d": {
          "allowed": false
        },
        "886224f6-04ae-4b13-bd8e-559c7d3754e1": {
          "allowed": true
        }
      }
    }
    ```

1. Enter `exit` to quit.