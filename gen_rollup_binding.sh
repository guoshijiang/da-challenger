now_path="../mantle/packages/contracts/artifacts/contracts/da"


rm -rf ${now_path}/BVM_EigenDataLayrChain.sol/abi.json
rm -rf ${now_path}/BVM_EigenDataLayrChain.sol/bin.json

cat ${now_path}/BVM_EigenDataLayrChain.sol/BVM_EigenDataLayrChain.json | jq -r '.abi' > ${now_path}/BVM_EigenDataLayrChain.sol/abi.json
cat ${now_path}/BVM_EigenDataLayrChain.sol/BVM_EigenDataLayrChain.json | jq -r '.bytecode' > ${now_path}/BVM_EigenDataLayrChain.sol/bin.json

abigen --bin=${now_path}/BVM_EigenDataLayrChain.sol/bin.json --abi=${now_path}/BVM_EigenDataLayrChain.sol/abi.json --pkg binding --out ./common/BVM_EigenDataLayrChain/binding.go
