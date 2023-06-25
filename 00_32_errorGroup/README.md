# Go言語でのエラーハンドリングと並行処理の最適化

---



今回は、Go言語でのエラーハンドリングと並行処理の最適化について詳しく説明します。Go言語は、エラーハンドリングをシンプルかつ効果的に行うことができる強力な機能を提供しています。また、並行処理を活用することで、アプリケーションのパフォーマンスを向上させることもできます。さっそく、具体的なコード例を見ていきましょう。



まず、以下のコードをご覧ください。

```go
package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime/debug"
)

func main() {
	// エラーグループ起動
	// 実行タスク指定
	result, err := runGoErrGroup(subPro, subPro1)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range result {
		fmt.Println(v)
	}

}

func runGoErrGroup(fns ...func(int) (int, error)) ([]int, error) {
	eg := errgroup.Group{}
	result := make([]int, len(fns))

	for i, fn := range fns {
		i, fn := i, fn
		eg.Go(func() error {
			defer func() {
				if err := recover(); err != nil {
					debug.PrintStack()
				}
			}()
			// タスク実行
			k, err := fn(i)
			// エラーなければ、結果を設定、エラーの場合、処理終了
			if err == nil {
				result[i] = k
			}
			return err
		})
	}
	// エラー終了待ち
	if err := eg.Wait(); err != nil {
		// エラーを無視して、結果を返す
		// return result, err
		// エラー終了 log.Fatal(err)
		// log.Println(err)
		return nil, err // 結果を無視、エラー返す
	}
	return result, nil
}

func subPro(i int) (int, error) {
	if i == 0 {
		return i, errors.New("0 で割ることはできません")
	}
	return i, nil
}

func subPro1(i int) (int, error) {
	return i, nil
}
```

上記のコードは、Go言語でエラーハンドリングと並行処理を実現する方法を示しています。



以下では、コードの各部分について詳しく説明します。

### `main` 関数

`main` 関数は、プログラムのエ

ントリーポイントです。以下の処理を行います。

1. `runGoErrGroup` 関数を呼び出し、実行するタスクとして `subPro` と `subPro1` 関数を指定します。
2. `runGoErrGroup` 関数の戻り値として、結果とエラーを受け取ります。
3. もしエラーが発生した場合は、エラーメッセージを表示します。
4. 結果を順番に表示します。

### `runGoErrGroup` 関数

`runGoErrGroup` 関数は、与えられた関数群を並行して実行し、結果を収集します。以下の処理を行います。

1. `errgroup.Group` を初期化します。
2. 結果を格納するためのスライス `result` を作成します。
3. 各関数とインデックスをループで処理し、ゴルーチンとして実行します。
4. ゴルーチン内では、タスクの実行結果を `result` スライスに格納します。
5. もしエラーが発生した場合、エラーを返して処理を終了します。
6. `eg.Wait()` を呼び出して、すべてのゴルーチンの終了を待ちます。もしエラーが発生した場合は、エラーを返します。

### サブ関数

`subPro` 関数と `subPro1` 関数は、それぞれの関数に特定のルールに基づいた処理を行うためのものです。`subPro` 関数では、引数が0の場合にエラーを返します。



以上が、Go言語でのエラーハンドリングと並行処理の最適化に関する詳細な説明でした。Go言語の強力な機能を活用することで、効果的なエラーハンドリングと高速な並行処理を実現することができます。



以上！