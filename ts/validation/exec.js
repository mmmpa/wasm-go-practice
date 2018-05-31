const fs = require('fs');
const util = require('util');
const readFile = util.promisify(fs.readFile);
global.TextEncoder = util.TextEncoder;
global.TextDecoder = util.TextDecoder;
const encoder = new TextEncoder("utf-16");
const decoder = new TextDecoder("utf-16");

(async function main () {
  const wasm = await readFile('build/optimized.wasm')
  const memory = new WebAssembly.Memory({ initial: 1 })
  const {
    instance: {
      exports: {
        add,
        log,
        memory: {
          buffer
        }
      },
    },
  } = await WebAssembly.instantiate(wasm, {
    js: {
      memory,
    },
    env: {
      outlog: function (index) {
        // 先頭の4byte(uint32)に文字列の長さが入っている
        const length = mem.getUint32(index, true);

        // 文字列は長さの後に続けて入っている
        const array = new Uint16Array(mem.buffer, index + 4, length);
        const str = decoder.decode(array);
        //const str = String.fromCharCode(...array);
        console.log(index, str);
      },
      sin: Math.sin
    }
  })

  let mem = new DataView(buffer)
  setString(buffer, "abc", 0)

  console.log(add(1, 2))
  console.log(log('Hello, AssemblyScript'))
})()

function setString (buffer, str, offset = 0) {
  return new Uint8Array(buffer, offset, str.length + 1)
    .set(encoder.encode(str + "\0"))
}
