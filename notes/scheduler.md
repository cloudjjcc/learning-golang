# G(Groutine)

Goroutine 是 Go 语言调度器中待执行的任务，它在运行时调度器中的地位与线程在操作系统中差不多，但是它占用了更小的内存空间，也降低了上下文切换的开销。

```go
// go 1.15 
type g struct {
   // Stack parameters.
   // stack describes the actual stack memory: [stack.lo, stack.hi).
   // stackguard0 is the stack pointer compared in the Go stack growth prologue.
   // It is stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
   // stackguard1 is the stack pointer compared in the C stack growth prologue.
   // It is stack.lo+StackGuard on g0 and gsignal stacks.
   // It is ~0 on other goroutine stacks, to trigger a call to morestackc (and crash).
   stack       stack   // offset known to runtime/cgo
   stackguard0 uintptr // offset known to liblink
   stackguard1 uintptr // offset known to liblink

   _panic       *_panic // innermost panic - offset known to liblink
   _defer       *_defer // innermost defer
   m            *m      // current m; offset known to arm liblink
   sched        gobuf
   syscallsp    uintptr        // if status==Gsyscall, syscallsp = sched.sp to use during gc
   syscallpc    uintptr        // if status==Gsyscall, syscallpc = sched.pc to use during gc
   stktopsp     uintptr        // expected sp at top of stack, to check in traceback
   param        unsafe.Pointer // passed parameter on wakeup
   atomicstatus uint32
   stackLock    uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
   goid         int64
   schedlink    guintptr
   waitsince    int64      // approx time when the g become blocked
   waitreason   waitReason // if status==Gwaiting

   preempt       bool // preemption signal, duplicates stackguard0 = stackpreempt
   preemptStop   bool // transition to _Gpreempted on preemption; otherwise, just deschedule
   preemptShrink bool // shrink stack at synchronous safe point

   // asyncSafePoint is set if g is stopped at an asynchronous
   // safe point. This means there are frames on the stack
   // without precise pointer information.
   asyncSafePoint bool

   paniconfault bool // panic (instead of crash) on unexpected fault address
   gcscandone   bool // g has scanned stack; protected by _Gscan bit in status
   throwsplit   bool // must not split stack
   // activeStackChans indicates that there are unlocked channels
   // pointing into this goroutine's stack. If true, stack
   // copying needs to acquire channel locks to protect these
   // areas of the stack.
   activeStackChans bool
   // parkingOnChan indicates that the goroutine is about to
   // park on a chansend or chanrecv. Used to signal an unsafe point
   // for stack shrinking. It's a boolean value, but is updated atomically.
   parkingOnChan uint8

   raceignore     int8     // ignore race detection events
   sysblocktraced bool     // StartTrace has emitted EvGoInSyscall about this goroutine
   sysexitticks   int64    // cputicks when syscall has returned (for tracing)
   traceseq       uint64   // trace event sequencer
   tracelastp     puintptr // last P emitted an event for this goroutine
   lockedm        muintptr
   sig            uint32
   writebuf       []byte
   sigcode0       uintptr
   sigcode1       uintptr
   sigpc          uintptr
   gopc           uintptr         // pc of go statement that created this goroutine
   ancestors      *[]ancestorInfo // ancestor information goroutine(s) that created this goroutine (only used if debug.tracebackancestors)
   startpc        uintptr         // pc of goroutine function
   racectx        uintptr
   waiting        *sudog         // sudog structures this g is waiting on (that have a valid elem ptr); in lock order
   cgoCtxt        []uintptr      // cgo traceback context
   labels         unsafe.Pointer // profiler labels
   timer          *timer         // cached timer for time.Sleep
   selectDone     uint32         // are we participating in a select and did someone win the race?

   // Per-G GC state

   // gcAssistBytes is this G's GC assist credit in terms of
   // bytes allocated. If this is positive, then the G has credit
   // to allocate gcAssistBytes bytes without assisting. If this
   // is negative, then the G must correct this by performing
   // scan work. We track this in bytes to make it fast to update
   // and check for debt in the malloc hot path. The assist ratio
   // determines how this corresponds to scan work debt.
   gcAssistBytes int64
}
```





结构体 [`runtime.g`](https://draveness.me/golang/tree/runtime.g) 的 `atomicstatus` 字段存储了当前 Goroutine 的状态。除了几个已经不被使用的以及与 GC 相关的状态之外，Goroutine 可能处于以下 9 种状态：

| 状态          | 描述                                                         |
| ------------- | ------------------------------------------------------------ |
| `_Gidle`      | 刚刚被分配并且还没有被初始化                                 |
| `_Grunnable`  | 没有执行代码，没有栈的所有权，存储在运行队列中               |
| `_Grunning`   | 可以执行代码，拥有栈的所有权，被赋予了内核线程 M 和处理器 P  |
| `_Gsyscall`   | 正在执行系统调用，拥有栈的所有权，没有执行用户代码，被赋予了内核线程 M 但是不在运行队列上 |
| `_Gwaiting`   | 由于运行时而被阻塞，没有执行用户代码并且不在运行队列上，但是可能存在于 Channel 的等待队列上 |
| `_Gdead`      | 没有被使用，没有执行代码，可能有分配的栈                     |
| `_Gcopystack` | 栈正在被拷贝，没有执行代码，不在运行队列上                   |
| `_Gpreempted` | 由于抢占而被阻塞，没有执行用户代码并且不在运行队列上，等待唤醒 |
| `_Gscan`      | GC 正在扫描栈空间，没有执行代码，可以与其他状态同时存在      |

## g0

运行时协程，拥有较大的栈

`g0` 作为一个特殊的 goroutine，为 scheduler 执行调度循环提供了场地（栈）。对于一个线程来说，g0 总是它第一个创建的 goroutine。之后，它会不断地寻找其他普通的 goroutine 来执行，直到进程退出。

## g的创建

当我们在程序中使用go关键字时，编译器会转化成`runtime.newproc`方法的调用

[`runtime.newproc`](https://draveness.me/golang/tree/runtime.newproc) 的入参是参数大小和表示函数的指针 `funcval`，它会获取 Goroutine 以及调用方的程序计数器，然后调用 [`runtime.newproc1`](https://draveness.me/golang/tree/runtime.newproc1) 函数获取新的 Goroutine 结构体、将其加入处理器的运行队列并在满足条件时调用 [`runtime.wakep`](https://draveness.me/golang/tree/runtime.wakep) 唤醒新的处理执行 Goroutine：

```go
func newproc(siz int32, fn *funcval) {
	argp := add(unsafe.Pointer(&fn), sys.PtrSize)
	gp := getg()
	pc := getcallerpc()
	systemstack(func() {
		newg := newproc1(fn, argp, siz, gp, pc)

		_p_ := getg().m.p.ptr()
		runqput(_p_, newg, true)

		if mainStarted {
			wakep()
		}
	})
}
```

[`runtime.newproc1`](https://draveness.me/golang/tree/runtime.newproc1) 会根据传入参数初始化一个 `g` 结构体，我们可以将该函数分成以下几个部分介绍它的实现：

1. 获取或者创建新的 Goroutine 结构体；
2. 将传入的参数移到 Goroutine 的栈上；
3. 更新 Goroutine 调度相关的属性；

### 获取g结构体的两种方式

`runtime.gfget`函数从p的缓存或者调度器的全局缓存中获取

`runtime.malg`函数创建新的g

### 获取当前线程运行的g

getg方法



# M(Machine)

Go 语言并发模型中的 M 是操作系统线程。调度器最多可以创建 10000 个线程，但是其中大多数的线程都不会执行用户代码（可能陷入系统调用），最多只会有 `GOMAXPROCS` 个活跃线程能够正常运行。

在默认情况下，运行时会将 `GOMAXPROCS` 设置成当前机器的核数，我们也可以在程序中使用 [`runtime.GOMAXPROCS`](https://draveness.me/golang/tree/runtime.GOMAXPROCS) 来改变最大的活跃线程数。

```go
type m struct {
   g0      *g     // goroutine with scheduling stack
   morebuf gobuf  // gobuf arg to morestack
   divmod  uint32 // div/mod denominator for arm - known to liblink

   // Fields not known to debuggers.
   procid        uint64       // for debuggers, but offset not hard-coded
   gsignal       *g           // signal-handling g
   goSigStack    gsignalStack // Go-allocated signal handling stack
   sigmask       sigset       // storage for saved signal mask
   tls           [6]uintptr   // thread-local storage (for x86 extern register)
   mstartfn      func()
   curg          *g       // current running goroutine
   caughtsig     guintptr // goroutine running during fatal signal
   p             puintptr // attached p for executing go code (nil if not executing go code)
   nextp         puintptr
   oldp          puintptr // the p that was attached before executing a syscall
   id            int64
   mallocing     int32
   throwing      int32
   preemptoff    string // if != "", keep curg running on this m
   locks         int32
   dying         int32
   profilehz     int32
   spinning      bool // m is out of work and is actively looking for work
   blocked       bool // m is blocked on a note
   newSigstack   bool // minit on C thread called sigaltstack
   printlock     int8
   incgo         bool   // m is executing a cgo call
   freeWait      uint32 // if == 0, safe to free g0 and delete m (atomic)
   fastrand      [2]uint32
   needextram    bool
   traceback     uint8
   ncgocall      uint64      // number of cgo calls in total
   ncgo          int32       // number of cgo calls currently in progress
   cgoCallersUse uint32      // if non-zero, cgoCallers in use temporarily
   cgoCallers    *cgoCallers // cgo traceback if crashing in cgo call
   doesPark      bool        // non-P running threads: sysmon and newmHandoff never use .park
   park          note
   alllink       *m // on allm
   schedlink     muintptr
   lockedg       guintptr
   createstack   [32]uintptr // stack that created this thread.
   lockedExt     uint32      // tracking for external LockOSThread
   lockedInt     uint32      // tracking for internal lockOSThread
   nextwaitm     muintptr    // next m waiting for lock
   waitunlockf   func(*g, unsafe.Pointer) bool
   waitlock      unsafe.Pointer
   waittraceev   byte
   waittraceskip int
   startingtrace bool
   syscalltick   uint32
   freelink      *m // on sched.freem

   // mFixup is used to synchronize OS related m state (credentials etc)
   // use mutex to access.
   mFixup struct {
      lock mutex
      fn   func(bool) bool
   }

   // these are here because they are too large to be on the stack
   // of low-level NOSPLIT functions.
   libcall   libcall
   libcallpc uintptr // for cpu profiler
   libcallsp uintptr
   libcallg  guintptr
   syscall   libcall // stores syscall parameters on windows

   vdsoSP uintptr // SP for traceback while in VDSO call (0 if not in call)
   vdsoPC uintptr // PC for traceback while in VDSO call

   // preemptGen counts the number of completed preemption
   // signals. This is used to detect when a preemption is
   // requested, but fails. Accessed atomically.
   preemptGen uint32

   // Whether this is a pending preemption signal on this M.
   // Accessed atomically.
   signalPending uint32

   dlogPerM

   mOS

   // Up to 10 locks held by this m, maintained by the lock ranking code.
   locksHeldLen int
   locksHeld    [10]heldLockInfo
}
```



## m0

程序主线程





## m的创建

### 创建时机



### 创建方法



# P(Processor)

调度器中的处理器 P 是线程和 Goroutine 的中间层，它能提供线程需要的上下文环境，也会负责调度线程上的等待队列，通过处理器 P 的调度，每一个内核线程都能够执行多个 Goroutine，它能在 Goroutine 进行一些 I/O 操作时及时让出计算资源，提高线程的利用率。

```go
type p struct {
   id          int32
   status      uint32 // one of pidle/prunning/...
   link        puintptr
   schedtick   uint32     // incremented on every scheduler call
   syscalltick uint32     // incremented on every system call
   sysmontick  sysmontick // last tick observed by sysmon
   m           muintptr   // back-link to associated m (nil if idle)
   mcache      *mcache
   pcache      pageCache
   raceprocctx uintptr

   deferpool    [5][]*_defer // pool of available defer structs of different sizes (see panic.go)
   deferpoolbuf [5][32]*_defer

   // Cache of goroutine ids, amortizes accesses to runtime·sched.goidgen.
   goidcache    uint64
   goidcacheend uint64

   // Queue of runnable goroutines. Accessed without lock.
   runqhead uint32
   runqtail uint32
   runq     [256]guintptr
   // runnext, if non-nil, is a runnable G that was ready'd by
   // the current G and should be run next instead of what's in
   // runq if there's time remaining in the running G's time
   // slice. It will inherit the time left in the current time
   // slice. If a set of goroutines is locked in a
   // communicate-and-wait pattern, this schedules that set as a
   // unit and eliminates the (potentially large) scheduling
   // latency that otherwise arises from adding the ready'd
   // goroutines to the end of the run queue.
   runnext guintptr

   // Available G's (status == Gdead)
   gFree struct {
      gList
      n int32
   }

   sudogcache []*sudog
   sudogbuf   [128]*sudog

   // Cache of mspan objects from the heap.
   mspancache struct {
      // We need an explicit length here because this field is used
      // in allocation codepaths where write barriers are not allowed,
      // and eliminating the write barrier/keeping it eliminated from
      // slice updates is tricky, moreso than just managing the length
      // ourselves.
      len int
      buf [128]*mspan
   }

   tracebuf traceBufPtr

   // traceSweep indicates the sweep events should be traced.
   // This is used to defer the sweep start event until a span
   // has actually been swept.
   traceSweep bool
   // traceSwept and traceReclaimed track the number of bytes
   // swept and reclaimed by sweeping in the current sweep loop.
   traceSwept, traceReclaimed uintptr

   palloc persistentAlloc // per-P to avoid mutex

   _ uint32 // Alignment for atomic fields below

   // The when field of the first entry on the timer heap.
   // This is updated using atomic functions.
   // This is 0 if the timer heap is empty.
   timer0When uint64

   // The earliest known nextwhen field of a timer with
   // timerModifiedEarlier status. Because the timer may have been
   // modified again, there need not be any timer with this value.
   // This is updated using atomic functions.
   // This is 0 if the value is unknown.
   timerModifiedEarliest uint64

   // Per-P GC state
   gcAssistTime         int64 // Nanoseconds in assistAlloc
   gcFractionalMarkTime int64 // Nanoseconds in fractional mark worker (atomic)

   // gcMarkWorkerMode is the mode for the next mark worker to run in.
   // That is, this is used to communicate with the worker goroutine
   // selected for immediate execution by
   // gcController.findRunnableGCWorker. When scheduling other goroutines,
   // this field must be set to gcMarkWorkerNotWorker.
   gcMarkWorkerMode gcMarkWorkerMode
   // gcMarkWorkerStartTime is the nanotime() at which the most recent
   // mark worker started.
   gcMarkWorkerStartTime int64

   // gcw is this P's GC work buffer cache. The work buffer is
   // filled by write barriers, drained by mutator assists, and
   // disposed on certain GC state transitions.
   gcw gcWork

   // wbBuf is this P's GC write barrier buffer.
   //
   // TODO: Consider caching this in the running G.
   wbBuf wbBuf

   runSafePointFn uint32 // if 1, run sched.safePointFn at next safe point

   // statsSeq is a counter indicating whether this P is currently
   // writing any stats. Its value is even when not, odd when it is.
   statsSeq uint32

   // Lock for timers. We normally access the timers while running
   // on this P, but the scheduler can also do it from a different P.
   timersLock mutex

   // Actions to take at some time. This is used to implement the
   // standard library's time package.
   // Must hold timersLock to access.
   timers []*timer

   // Number of timers in P's heap.
   // Modified using atomic instructions.
   numTimers uint32

   // Number of timerModifiedEarlier timers on P's heap.
   // This should only be modified while holding timersLock,
   // or while the timer status is in a transient state
   // such as timerModifying.
   adjustTimers uint32

   // Number of timerDeleted timers in P's heap.
   // Modified using atomic instructions.
   deletedTimers uint32

   // Race context used while executing timer functions.
   timerRaceCtx uintptr

   // preempt is set to indicate that this P should be enter the
   // scheduler ASAP (regardless of what G is running on it).
   preempt bool

   pad cpu.CacheLinePad
}
```



## p的创建

### 创建时机

调度器初始化时

### 创建方法

`runtime.procresize` 



# 调度器启动

运行时通过 [`runtime.schedinit`](https://draveness.me/golang/tree/runtime.schedinit) 初始化调度器：

```go
func schedinit() {
	_g_ := getg()
	...

	sched.maxmcount = 10000

	...
	sched.lastpoll = uint64(nanotime())
	procs := ncpu
	if n, ok := atoi32(gogetenv("GOMAXPROCS")); ok && n > 0 {
		procs = n
	}
	if procresize(procs) != nil {
		throw("unknown runnable goroutine during bootstrap")
	}
}
```

在调度器初始函数执行的过程中会将 `maxmcount` 设置成 10000，这也就是一个 Go 语言程序能够创建的最大线程数，虽然最多可以创建 10000 个线程，但是可以同时运行的线程还是由 `GOMAXPROCS`变量控制。

我们从环境变量 `GOMAXPROCS` 获取了程序能够同时运行的最大处理器数之后就会调用 [`runtime.procresize`](https://draveness.me/golang/tree/runtime.procresize) 更新程序中处理器的数量，在这时整个程序不会执行任何用户 Goroutine，调度器也会进入锁定状态，[`runtime.procresize`](https://draveness.me/golang/tree/runtime.procresize) 的执行过程如下：

1. 如果全局变量 `allp` 切片中的处理器数量少于期望数量，会对切片进行扩容；
2. 使用 `new` 创建新的处理器结构体并调用 [`runtime.p.init`](https://draveness.me/golang/tree/runtime.p.init) 初始化刚刚扩容的处理器；
3. 通过指针将线程 m0 和处理器 `allp[0]` 绑定到一起；
4. 调用 [`runtime.p.destroy`](https://draveness.me/golang/tree/runtime.p.destroy) 释放不再使用的处理器结构；
5. 通过截断改变全局变量 `allp` 的长度保证与期望处理器数量相等；
6. 将除 `allp[0]` 之外的处理器 P 全部设置成 `_Pidle` 并加入到全局的空闲队列中；

调用 [`runtime.procresize`](https://draveness.me/golang/tree/runtime.procresize) 是调度器启动的最后一步，在这一步过后调度器会完成相应数量处理器的启动，等待用户创建运行新的 Goroutine 并为 Goroutine 调度处理器资源。



# WSS(工作窃取)





# runtime.main



# main.main