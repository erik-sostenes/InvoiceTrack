package business

import (
	"encoding/xml"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	maxGoroutines = 10
)

type ECDSettlement struct {
	FUECD      string
	Settlement Settlement
}

type TotalAmountCalculator interface {
	Calculate(dirPath, extension string) (Map, error)
}

var _ TotalAmountCalculator = totalAmountCalculator{}

type totalAmountCalculator struct{}

func NewTotalAmountCalculator() (totalAmountCalculator, error) {
	return totalAmountCalculator{}, nil
}

func (t totalAmountCalculator) Calculate(dirPath, extension string) (Map, error) {
	done := make(chan struct{})
	defer close(done)

	dirFilesCh := reader(dirPath, extension, done)
	dailyStatementsCh := parseXML(done, dirFilesCh)
	ecdSettlementsCh := filterSettlements(done, dailyStatementsCh)

	return sumTotalAmountByFUECDAndType(done, ecdSettlementsCh), nil
}

func reader(dirPath, extension string, done <-chan struct{}) <-chan string {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		slog.Error(err.Error())
	}

	dirFilesCh := make(chan string)

	go func() {
		semaphore := make(chan struct{}, maxGoroutines)
		wg := &sync.WaitGroup{}

		defer func() {
			wg.Wait()
			close(dirFilesCh)
			close(semaphore)
		}()

		for _, entry := range entries {
			wg.Add(1)

			go func(entry fs.DirEntry) {
				semaphore <- struct{}{}

				defer func() {
					<-semaphore
					wg.Done()
				}()

				if entry.IsDir() || !strings.EqualFold(filepath.Ext(entry.Name()), "."+extension) {
					return
				}

				fullPath := filepath.Join(dirPath, entry.Name())

				select {
				case <-done:
					return
				case dirFilesCh <- fullPath:
				}
			}(entry)
		}
	}()

	return dirFilesCh
}

func parseXML(done <-chan struct{}, dirFilesCh <-chan string) <-chan DailyStatement {
	dailyStatementsCh := make(chan DailyStatement)

	go func() {
		semaphore := make(chan struct{}, maxGoroutines)
		wg := &sync.WaitGroup{}

		defer func() {
			wg.Wait()

			close(dailyStatementsCh)
			close(semaphore)
		}()

		for dirFile := range dirFilesCh {
			wg.Add(1)

			go func(dirFile string) {
				semaphore <- struct{}{}

				defer func() {
					<-semaphore
					wg.Done()
				}()

				file, err := os.Open(dirFile)
				if err != nil {
					slog.Error(err.Error())
					return
				}
				defer file.Close()

				decoder := xml.NewDecoder(file)
				var dailyStatement DailyStatement

				if err := decoder.Decode(&dailyStatement); err != nil {
					slog.Error(fmt.Sprintf("Failed to decode XML in %s: %v", dirFile, err))
					return
				}

				select {
				case <-done:
					return
				case dailyStatementsCh <- dailyStatement:
				}
			}(dirFile)
		}
	}()

	return dailyStatementsCh
}

func filterSettlements(done <-chan struct{}, dailyStatementsCh <-chan DailyStatement) <-chan *ECDSettlement {
	settlementsCh := make(chan *ECDSettlement)

	go func() {
		semaphore := make(chan struct{}, maxGoroutines)
		wg := &sync.WaitGroup{}

		defer func() {
			wg.Wait()

			close(settlementsCh)
			close(semaphore)
		}()

		for dailyStatement := range dailyStatementsCh {
			wg.Add(1)

			go func(dailyStatement DailyStatement) {
				semaphore <- struct{}{}

				defer func() {
					<-semaphore
					wg.Done()
				}()

				for _, settlement := range dailyStatement.Settlements {
					if settlement.NumLiq != 0 {
						continue
					}

					ecdSettlement := &ECDSettlement{
						FUECD:      dailyStatement.FUECD,
						Settlement: settlement,
					}

					select {
					case <-done:
						return
					case settlementsCh <- ecdSettlement:
					}
				}
			}(dailyStatement)
		}
	}()

	return settlementsCh
}

func sumTotalAmountByFUECDAndType(done <-chan struct{}, ecdSettlementsCh <-chan *ECDSettlement) Map {
	groupECD := make(Map)

	for {
		select {
		case <-done:
			return groupECD
		case ecdSettlement, ok := <-ecdSettlementsCh:
			if !ok {
				return groupECD
			}

			if _, exists := groupECD[ecdSettlement.FUECD]; !exists {
				groupECD[ecdSettlement.FUECD] = make(map[string]float64)
			}

			for _, invoice := range ecdSettlement.Settlement.Invoices {
				for _, concept := range invoice.Concepts {
					groupECD[ecdSettlement.FUECD][invoice.Type] += concept.TotalAmount
				}
			}
		}
	}
}
