package types

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	for _, tc := range []struct {
		Hash string `json:"transaction_hash"`
	}{{
		Hash: "315e364b162653e5c7b23efd34f8da27ba9c069b68e3042b7d76ce1df890313",
	}, {
		Hash: "8d6955e1bc0d5ba78b04630375f962ce6e5c91115173bc5f6e7843c6ee1269",
	}, {
		Hash: "680b0616e65633dfaf06d5a5ee5f1d1d4b641396009f00a67c0d18dc0f9638",
	}} {
		var th TransactionHash
		
		if err := json.Unmarshal([]byte(fmt.Sprintf(`{"transaction_hash":"%s"}`, tc.Hash)), &th); err != nil {
			t.Fatalf("Unmarshalling text: %v", err)
		}
		// h := th.TransactionHash
		// h2 := HexToHash(tc.Hash)

		// if h != h2 {
		// 	t.Fatalf("Hashes not equal: %s %s", h, h2)
		// }

		// m, err := h.MarshalText()
		// if err != nil {
		// 	t.Fatalf("Marshalling text: %v", err)
		// }

		// m2, err := json.Marshal(h)
		// if err != nil {
		// 	t.Fatalf("Marshalling json: %v", err)
		// }

		// if tc.Hash != string(m) {
		// 	t.Errorf("Hash mismatch, want: %s, got: %s", tc.Hash, m)
		// }

		// if tc.Hash != string(m2) {
		// 	t.Errorf("Hash mismatch, want: %s, got: %s", tc.Hash, m2)
		// }
	}
}
