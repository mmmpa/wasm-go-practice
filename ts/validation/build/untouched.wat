(module
 (type $iii (func (param i32 i32) (result i32)))
 (type $iv (func (param i32)))
 (type $iiv (func (param i32 i32)))
 (type $v (func))
 (import "env" "outlog" (func $assembly/index/outlog (param i32 i32)))
 (global $HEAP_BASE i32 (i32.const 20))
 (memory $0 1)
 (data (i32.const 8) "\03\00\00\00a\00a\00a\00")
 (export "add" (func $assembly/index/add))
 (export "log" (func $assembly/index/log))
 (export "memory" (memory $0))
 (start $start)
 (func $assembly/index/add (; 1 ;) (type $iii) (param $0 i32) (param $1 i32) (result i32)
  ;;@ assembly/index.ts:3:42
  (return
   ;;@ assembly/index.ts:4:9
   (i32.add
    (i32.add
     (get_local $0)
     ;;@ assembly/index.ts:4:13
     (get_local $1)
    )
    ;;@ assembly/index.ts:4:17
    (get_local $1)
   )
  )
 )
 (func $assembly/index/log (; 2 ;) (type $iv) (param $0 i32)
  ;;@ assembly/index.ts:7:38
  (call $assembly/index/outlog
   ;;@ assembly/index.ts:8:9
   (i32.const 8)
   ;;@ assembly/index.ts:8:16
   (i32.load
    (get_local $0)
   )
  )
 )
 (func $start (; 3 ;) (type $v)
  ;;@ assembly/index.ts:10:0
  (drop
   (i32.const 3)
  )
 )
)
