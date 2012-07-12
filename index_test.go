package wikiparse

import (
	"io"
	"strings"
	"testing"
)

const testData = `499:10:AccessibleComputing
499:12:Anarchism
499:13:AfghanistanHistory
499:14:AfghanistanGeography
499:15:AfghanistanPeople
499:18:AfghanistanCommunications
499:19:AfghanistanTransportations
499:20:AfghanistanMilitary
499:21:AfghanistanTransnationalIssues
499:23:AssistiveTechnology
2147418907:2638569:William Earl Brown
2147418907:2638570:Lebuhraya Persekutuan
2147418907:2638571:St Francis of Paola
2147418907:2638573:Francesco di Paula
2147418907:2638575:Arapahoe Community College
2147418907:2638583:Francesco Borgia
-2147469295:2638585:Philadelphia Bulletin
-2147469295:2638588:Zrínyi Miklós
-2147469295:2638602:Privatize
-2147469295:2638604:Island of Montréal
`

const lastChunk = 2147498001

func TestIndexReader(t *testing.T) {
	r := strings.NewReader(testData)
	ir := NewIndexReader(r)

	e, err := ir.Next()
	if err != nil {
		t.Fatalf("Error parsing first entry: %v", err)
	}

	for {
		var tmp IndexEntry
		tmp, err = ir.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Error reading stream:  %v", err)
		}
		e = tmp
	}
	if e.StreamOffset != lastChunk {
		t.Fatalf("Expected %v, got %v for the last chunk offset",
			int64(lastChunk), e.StreamOffset)
	}

}