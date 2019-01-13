# Performance Library Comparison

## Purpose

The purpose of this project is to compare a few libraries that pull OS performance stats:

* [gopsutil](https://github.com/shirou/gopsutil)
* [NMON](http://nmon.sourceforge.net) (Binary)


## Comparison

### CPU

#### Linux & MacOS
|  Config | gopsutil  | nmon  |   |   |
|---|---|---|---|---|
| vendorId | x | x |   |   |
| family | x  | x |   |   |
| model | x  | x |   |   |
| stepping | x  | x |   |   |
| physicalId | x  | x |   |   |
| coreId |  x | x |   |   |
| cores (#) | x  | x |   |   |
| mhz |  x | x |   |   |
| cacheSize | x  | x |   |   |
| flags | x  | x |  |   |
| Microcode | x  | x |   |   |
| cpu family | x  | x |   |   |
| siblings |  | x  |  |   |
| apicid |   | x |   |   |
| fpu |  | x |   |   |
| wp |  | x |   |   |
| bugs |  | x |   |   |
| bogomips |  | x |   |   |
| clflush_size |  | x |   |   |
| cache alignment |   | x |   |   |
| address sizes |  | x  |   |   |

Notes: 
* NMON grabs all of /proc/cpuinfo
* gopsutil returns a static list of things it collects. I.E if things get added to /proc/cpuinfo it will not be collected.


|  Metric | gopsutil  | nmon  |   |   |
|---|---|---|---|---|
| User |  x | x |   |   |
| System | x  | x |   |   |
| Idle | x  | x |   |   |
| Nice | x  |   |   |   |
| iowait | x  | x |   |   |
| irq | x |  |   |   |
| steal | x  |   |   |   |
| guest | x |   |   |   |
| guestNice | x  |   |   |   |
| stolen | x  |   |   |   |

### Memory

#### Linux 

Virtual Memory

|  Metric  | gopsutil | nmon |   |   |
|---|---|---|---|---|
| total | x | x |  |  |
| available | x |  |  |  |
| used | x |  |  |  |
| usedPercent | x |  |  |  |
| free | x | x |  |  |
| active | x | x |  |  |
| inactive | x | x |  |  |
| wired | x |  |  |  |
| laundry | x |  |  |  |
| buffers | x | x |  |  |
| cached | x | x |  |  |
| writeback | x |  |  |  |
| dirty | x |  |  |  |
| writebacktmp | x |  |  |  |
| shared | x | x |  |  |
| slab | x |  |  |  |
| sreclaimable | x |  |  |  |
| pagetables | x |  |  |  |
| swapcached | x |  |  |  |
| commitlimit | x |  |  |  |
| committedas | x |  |  |  |
| hightotal | x | x |  |  |
| highfree | x | x |  |  |
| lowtotal | x | x |  |  |
| lowfree | x | x |  |  |
| swaptotal | x |  |  |  |
| swapfree | x |  |  |  |
| mapped | x |  |  |  |
| vmalloctotal | x |  |  |  |
| vmallocused | x |  |  |  |
| vmallocchunk | x |  |  |  |
| hugepagestotal | x |  |  |  |
| hugepagesfree | x |  |  |  |
| hugepagesize | x |  |  |  |
| bigfree |  | x |  |  |

Swap

|  Metric  | gopsutil | nmon |   |   |
|---|---|---|---|---|
| total | x | x |  |  |
| used | x |  |  |  |
| free | x | x |  |  |
| swap cached |  | x |  |  |
| usedPercent | x |  |  |  |
| sin | x |  |  |  |
| sout | x |  |  |  |



Notes:
* NMON grabs all of /proc/meminfo, but NOT as metrics. Once per collection it'll grab the stats, but metric collection is different.
* gopsutil returns a static list of things it collects. I.E if things get added to /proc/meminfo it will not be collected.

### Network

#### Linux
Connections

|  Metric  | gopsutil | nmon |   |   |
|---|---|---|---|---|
| fd | x |  |  |  |
| family | x |  |  |  |
| type | x |  |  |  |
| local ip | x |  |  |  |
| local port | x |  |  |  |
| remote ip | x |  |  |  |
| remote port | x |  |  |  |
| status | x |  |  |  |
| uids | x |  |  |  |
| pid | x |  |  |  |

Counters

|  Metric  | gopsutil | nmon |   |   |
|---|---|---|---|---|
| bytesSent | x | x |  |  |
| bytesRecv | x | x |  |  |
| packetsSent | x | x |  |  |
| packetsRecv | x | x |  |  |
| errin | x | * |  |  |
| errout | x | * |  |  |
| dropin | x | * |  |  |
| dropout | x | * |  |  |
| fifoin | x | * |  |  |
| fifoout | x | * |  |  |

* Denotes that NMON collects it, but not every metric interval. Once per entire collection.

Protocols

IP 

|  Metric  | gopsutil | nmon |   |   |
|---|---|---|---|---|
| DefaultTTL | x |  |  |  |
| ForwDatagrams | x |  |  |  |
| Forwarding | x |  |  |  |
| FragCreates | x |  |  |  |
| FragFails | x |  |  |  |
| FragOKs | x |  |  |  |
| InAddrErrors | x |  |  |  |
| InDelivers | x |  |  |  |
| InDiscards | x |  |  |  |
| InHdrErrors | x |  |  |  |
| InReceives | x |  |  |  |
| InUnknownProtos | x |  |  |  |
| OutDiscards | x |  |  |  |
| OutNoRoutes | x |  |  |  |
| OutRequests | x |  |  |  |
| ReasmFails | x |  |  |  |
| ReasmOKs | x |  |  |  |
| ReasmReqds | x |  |  |  |
| ReasmTimeout | x |  |  |  |

ICMP

|  Metric  | gopsutil | nmon |   |   |
|---|---|---|---|---|
| InAddrMaskReps | x |  |  |  |
| InAddrMasks | x |  |  |  |
| InCsumErrors | x |  |  |  |
| InDestUnreachs | x |  |  |  |
| InEchoReps | x |  |  |  |
| InEchos | x |  |  |  |
| InErrors | x |  |  |  |
| InMsgs | x |  |  |  |
| InParmProbs | x |  |  |  |
| InRedirects | x |  |  |  |
| InSrcQuenchs | x |  |  |  |
| InTimeExcds | x |  |  |  |
| InTimestampReps | x |  |  |  |
| InTimestamps | x |  |  |  |
| OutAddrMaskReps | x |  |  |  |
| OutAddrMasks | x |  |  |  |
| OutDestUnreachs | x |  |  |  |
| OutEchoReps | x |  |  |  |
| OutEchos | x |  |  |  |
| OutErrors | x |  |  |  |
| OutMsgs | x |  |  |  |
| OutParmProbs | x |  |  |  |
| OutRedirects | x |  |  |  |
| OutSrcQuenchs | x |  |  |  |
| OutTimeExcds | x |  |  |  |
| OutTimestampReps | x |  |  |  |
| OutTimestamps | x |  |  |  |

UDP

|  Metric  | gopsutil | nmon |   |   |
|---|---|---|---|---|
| IgnoredMulti | x |  |  |  |
| InCsumErrors | x |  |  |  |
| InDatagrams | x |  |  |  |
| InErrors | x |  |  |  |
| NoPorts | x |  |  |  |
| OutDatagrams | x |  |  |  |
| RcvbufErrors | x |  |  |  |
| SndbufErrors | x |  |  |  |


### Disk
### Other