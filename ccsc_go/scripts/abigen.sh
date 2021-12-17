#!/bin/sh

cd "$(dirname "$0")" || exit 1

if ! which abigen >/dev/null; then
  echo "error: abigen not installed" >&2
  exit 1
fi

abigen --abi ../../ccsc_contracts/artifacts/contracts/examples/Client.sol/Client.abi --pkg ccsc --type ClientContract --out ../pkg/ccsc/clientcontract.go
abigen --abi ../../ccsc_contracts/artifacts/contracts/ccsc/Proxy.sol/Proxy.abi --pkg ccsc --type ProxyContract --out ../pkg/ccsc/proxycontract.go
abigen --abi ../../ccsc_contracts/artifacts/contracts/ccsc/Server.sol/Server.abi --pkg ccsc --type ServerContract --out ../pkg/ccsc/servercontract.go