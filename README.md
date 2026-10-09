# conf2026-demo

[TinyGo Conference 2026](https://conf.tinygo-keeb.org/2026/) のデモ用リポジトリです。

conf2025badge (XIAO BLE) を、[sago35/tinygo-keyboard](https://github.com/sago35/tinygo-keyboard) の `ble-hid-2` branch を使って BLE キーボードとして動かす最小構成のコードです。

| キー | 動作 |
| ---- | ---- |
| 下側のボタン (D2) | A |
| 上側のボタン (D3) | B |
| ロータリーエンコーダー押し込み (D8) | BLE unpair (ペアリング情報を削除して、ペアリング待ちに戻す) |

ピンは `main.go` の `gpioPins` で切り替えます (conf2025badge / nicebadge)。

## TinyGo について

ビルドには、[sago35/tinygo](https://github.com/sago35/tinygo) の `nrf-fix-20261009` branch の TinyGo が必要です。

この branch には、nRF52840 + SoftDevice でしばらく使っているとメインループが止まる問題 (`rtc_sleep` が起床を取りこぼし、最大 512 秒眠り続ける) の修正が入っています。ただし、修正後の長期試験はまだ終わっていないので、問題が残っている可能性があります。対策として `main.go` では watchdog (1 秒) を有効にしています。止まった場合も自動でリセットされ、再接続します。

```
$ git clone -b nrf-fix-20261009 https://github.com/sago35/tinygo
```

TinyGo のソースからのビルド方法は、[TinyGo の公式ドキュメント](https://tinygo.org/docs/guides/build/) を参照してください。

## ビルドと書き込み

```
$ make          # out/conf2025badge-xiao-ble.uf2 を作成
$ make flash    # XIAO BLE に書き込み
$ make monitor  # シリアルモニタ
```

書き込み後、PC などの Bluetooth 設定から `conf2025badge ble` を選んでペアリングしてください。ペアリングに失敗する場合 (`pairing failed (status 0x85)`) は、キーボード側に古いペアリング情報が残っています。ロータリーエンコーダーを押し込んで unpair してから、もう一度ペアリングしてください。
