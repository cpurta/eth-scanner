你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# eth-scanner

An ethereum blockchain scanner that allows you to find transactions associated to
an address. This gives you the capability to search the entire chain for transactions
or just part of the chain. This then outputs all of the transactions to csv containing
all of the pertinent block/transaction information. Think of it as a very limited
etherscan.

## Getting started

To compile the usable binary you can run:

```
$ go build -o ./bin/eth-scanner ./cmd/eth-scanner
```

You can now execute the binary:

```
$ ./bin/eth-scanner -h
```

## Benchmarks

Currently with 25 block workers enables and with the scanner using an infura node,
the scanner can process ~125k blocks per hour. One way that you can speed up the
scanning is to connect to a fully synced node locally. Since the network latency
will not be a major bottleneck when attempting to scan large portions of the blockchain.

## Drawbacks

One major drawback to using this over etherscan.io is the amount of time that it
takes to scan since it will take `O(n)` request to scan all blocks specified. With
the mainnet at 6M+ blocks in order to scan

## Improvements

One major improvement that is being considered is the use of allowing a sync command
to sync all blocks to a relational database which can then be used to easily query
addresses and transactions.

Also dockerizing the environment would substantially help in creating database(s)
and checking the status of the scanner. The specific of which are still in debate.

Increase sync/scan times would be extremely beneficial in the usefulness of this
tool. Having to wait hours to find transactions is not at all helpful when compared
to other existing tools such as etherscan.

# LICENSE

GPL 3.0
