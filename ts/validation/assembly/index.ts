declare function outlog (out: string, length: usize): void;

export function add (a: i32, b: i32): i32 {
  return a + b + b
}

export function log (a: string): void {
  outlog('aaa', a.length)
}
3
