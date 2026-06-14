// WASM での GIF 生成はメインスレッドをブロックするため、ワーカースレッドで実行する。
// これにより生成中も UI（文字入力・色変更）が固まらない。
importScripts('wasm_exec.js');

const go = new Go();
let ready = false;
let latest = null;
let scheduled = false;

// 連続入力時は最新のリクエストだけを処理してバックログを防ぐ。
function process() {
  scheduled = false;
  if (!ready || latest === null) {
    return;
  }
  const { word, color } = latest;
  latest = null;
  const base64 = generate(word, color);
  postMessage({ type: 'result', base64 });
}

function schedule() {
  if (!scheduled) {
    scheduled = true;
    setTimeout(process, 0);
  }
}

onmessage = (e) => {
  latest = e.data;
  schedule();
};

WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then((result) => {
  go.run(result.instance);
  ready = true;
  schedule(); // WASM 準備前に届いたリクエストがあれば処理する
});
