const { BaseTrie: Trie } = require("merkle-patricia-tree");

function encodedHeaderFromBlock(block) {
  return ethers.utils.RLP.encode([
    block.parentHash,
    block.sha3Uncles,
    block.miner,
    block.stateRoot,
    block.transactionsRoot,
    block.receiptsRoot,
    block.logsBloom,
    ethers.utils.hexlify(ethers.BigNumber.from(block.difficulty)),
    ethers.utils.hexlify(ethers.BigNumber.from(block.number)),
    ethers.utils.hexlify(ethers.BigNumber.from(block.gasLimit)),
    ethers.utils.hexlify(ethers.BigNumber.from(block.gasUsed)),
    ethers.utils.hexlify(ethers.BigNumber.from(block.timestamp)),
    block.extraData,
    block.mixHash,
    ethers.utils.hexlify(ethers.BigNumber.from(block.nonce)),
  ]);
}

function encodeTransaction(tx) {
  return ethers.utils.RLP.encode([
    ethers.utils.hexlify(ethers.BigNumber.from(tx.nonce)),
    ethers.utils.hexlify(ethers.BigNumber.from(tx.gasPrice)),
    ethers.utils.hexlify(tx.gasLimit),
    tx.to,
    ethers.utils.hexlify(tx.value),
    tx.data,
    ethers.utils.hexlify(ethers.BigNumber.from(tx.v)),
    tx.r,
    tx.s,
  ]);
}

function encodeLogs(logs) {
  const encodedLogs = [];
  for (let i = 0; i < logs.length; i++) {
    encodedLogs.push([logs[i].address, logs[i].topics, logs[i].data]);
  }
  return encodedLogs;
}

function encodeTransactionReceipt(receipt) {
  return ethers.utils.RLP.encode([
    ethers.utils.hexlify(receipt.status),
    ethers.utils.hexlify(receipt.cumulativeGasUsed),
    receipt.logsBloom,
    encodeLogs(receipt.logs),
  ]);
}

async function encodedTransactionMerkleProof(block, index) {
  const trie = new Trie();

  for (let i = 0; i < block.transactions.length; i++) {
    const rlpTx = encodeTransaction(
      await ethers.provider.getTransaction(block.transactions[i].hash)
    );
    const key = ethers.utils.RLP.encode(
      ethers.utils.hexlify(ethers.BigNumber.from(i))
    );
    await trie.put(key, rlpTx);
  }

  const key = ethers.utils.RLP.encode(
    ethers.utils.hexlify(ethers.BigNumber.from(index))
  );
  return ethers.utils.RLP.encode(await Trie.createProof(trie, key));
}

async function encodedTransactionReceiptMerkleProof(block, index) {
  const trie = new Trie();

  for (let i = 0; i < block.transactions.length; i++) {
    const rlpReceipt = encodeTransactionReceipt(
      await ethers.provider.getTransactionReceipt(block.transactions[i].hash)
    );
    const key = ethers.utils.RLP.encode(
      ethers.utils.hexlify(ethers.BigNumber.from(i))
    );
    await trie.put(key, rlpReceipt);
  }

  const key = ethers.utils.RLP.encode(
    ethers.utils.hexlify(ethers.BigNumber.from(index))
  );
  return ethers.utils.RLP.encode(await Trie.createProof(trie, key));
}

async function getProof(tx) {
  const block = await ethers.provider.send("eth_getBlockByNumber", [
    ethers.utils.hexStripZeros(ethers.utils.hexlify(tx.blockNumber)),
    true,
  ]);
  const receipt = await ethers.provider.getTransactionReceipt(tx.hash);

  const blockHeader = encodedHeaderFromBlock(block);
  const rlpTransaction = encodeTransaction(tx);
  const rlpTransactionReceipt = encodeTransactionReceipt(receipt);
  const path = ethers.utils.RLP.encode(
    ethers.utils.hexlify(ethers.BigNumber.from(tx.transactionIndex))
  );
  const txMerkleProof = await encodedTransactionMerkleProof(
    block,
    tx.transactionIndex
  );
  const receiptMerkleProof = await encodedTransactionReceiptMerkleProof(
    block,
    tx.transactionIndex
  );
  return [
    blockHeader,
    rlpTransaction,
    rlpTransactionReceipt,
    path,
    txMerkleProof,
    receiptMerkleProof,
  ];
}

module.exports = {
  getProof,
};
