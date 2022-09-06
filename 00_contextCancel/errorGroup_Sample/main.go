package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime/debug"
)

func main() {
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
			k, err := fn(i)

			if err == nil {
				result[i] = k
			}
			return err
		})
	}

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
		return i, errors.New("0 dived")
	}
	return i, nil
}

func subPro1(i int) (int, error) {
	return i, nil
}
