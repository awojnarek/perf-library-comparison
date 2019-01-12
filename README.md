# Performance Library Comparison

## Purpose

The purpose of this project is to compare a few libraries that pull OS performance stats:

gopsutil
nmon (binary)

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
|  Metric  | gopsutil | nmon |   |   |
|---|---|---|---|---|
| total | x |  |  |  |
| available | x |  |  |  |
| used | x |  |  |  |
| usedPercent | x |  |  |  |
| free | x |  |  |  |
| active | x |  |  |  |
| inactive | x |  |  |  |
| wired | x |  |  |  |
| laundry | x |  |  |  |
| buffers | x |  |  |  |
| cached | x |  |  |  |
| writeback | x |  |  |  |
| dirty | x |  |  |  |
| writebacktmp | x |  |  |  |
| shared | x |  |  |  |
| slab | x |  |  |  |
| sreclaimable | x |  |  |  |
| pagetables | x |  |  |  |
| swapcached | x |  |  |  |
| commitlimit | x |  |  |  |
| committedas | x |  |  |  |
| hightotal | x |  |  |  |
| highfree | x |  |  |  |
| lowtotal | x |  |  |  |
| lowfree | x |  |  |  |
| swaptotal | x |  |  |  |
| swapfree | x |  |  |  |
| mapped | x |  |  |  |
| vmalloctotal | x |  |  |  |
| vmallocused | x |  |  |  |
| vmallocchunk | x |  |  |  |
| hugepagestotal | x |  |  |  |
| hugepagesfree | x |  |  |  |
| hugepagesize | x |  |  |  |

### Network
### Disk
### Other