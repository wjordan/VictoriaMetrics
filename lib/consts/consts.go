package consts

// MaxInsertPacketSizeForVMStorage is the maximum packet size in bytes vmstorage can accept from vmstorage.
// It cannot be reduced due to backwards compatibility :(
const MaxInsertPacketSizeForVMStorage = 100 * 1024 * 1024

// MaxInsertPacketSizeForVMInsert is the maximum packet size in bytes vminsert may send to vmstorage.
// It is smaller than MaxInsertPacketSizeForVMStorage in order to reduce
// max memory usage occupied by buffers at vmstorage and
// total data retransmission during network interruptions or congestion.
const MaxInsertPacketSizeForVMInsert = 1 * 1024 * 1024

// MaxInsertBufferSize is the maximum buffer size in bytes for vminsert per storage node.
// It limits memory usage occupied by buffers at vminsert.
const MaxInsertBufferSize = 80 * 1024 * 1024
