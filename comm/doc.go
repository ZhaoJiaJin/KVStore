/*
    package comm implements membership management, such as
        - leader election.
        - sync logs among nodes.

    params:
    id: node id
    node addresses: address of all the nodes
    join: join flag
    getSnapShot: get snapshot function
    proposeC: commit propose Chan(logs waiting to be commited)
    confChangeC: conf change propose logs(waiting to be commited)
    return:
    commitC: commited Channel. When recover, all the logs after the latest snapshot will also be redirected to this channel, and a nil is sent after all commited logs is sent
    errorC: error Chan
    snapshotterReady: snapshot ready channel(Used for recover when starting). The lastest snapshot will be sent via this channel, and nil if the lastest snapshot doesn't exist

    steps to commit a log:
    1. put log into proposeC
    2. logs in proposeC will be sent to master node
    3. master node will process logs one by one
        3.1. the master sync the log to all the nodes using 2PC

    how many timeout/interval do we need:
        - heartbeat frequency. heartbeat frequency should be much smaller than heartbeat timeout. For example, 50ms
        - heartbeat timeout. a random interval 150ms-300ms
        - election timeout? we would have election timeout in this case for simplicity, but grpc request will have a timeout
*/

package comm
