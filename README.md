# fizzbuzz-with-voicevox
Fizz Buzzの出力結果（20まで）を[VOICEVOX ENGINE](https://github.com/VOICEVOX/voicevox_engine)を用いて音声出力するサンプル

## 実行方法
Dockerを用いて予めVOICEVOX ENGINEを起動させておきます。

```sh
docker pull voicevox/voicevox_engine:cpu-ubuntu20.04-latest
docker run --rm -it -p '127.0.0.1:50021:50021' voicevox/voicevox_engine:cpu-ubuntu20.04-latest
```

その後、以下のコマンドを叩きます。

```sh
go run main.go
```

すると、カレントに `audio.wav` という音声ファイルが生成されます。