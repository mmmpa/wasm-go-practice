(module
  (import "console" "log" (func $log (param i32)))
  (func (export "add")
    i32.const 13
    call $log))
