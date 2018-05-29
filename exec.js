const fs = require('fs');
const util = require('util');
const readFile = util.promisify(fs.readFile);

(async function main () {
  const wasm = await readFile('hello_wasm.wasm')
  const {
    instance: {
      exports: {
        hello
      },
    },
  } = await WebAssembly.instantiate(wasm, {})

  console.log(hello(123))
})()
