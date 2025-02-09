package hello

import "testing"

func TestSayHello(t *testing.T){
	subtests := []struct{
		items []string
		result string
	}{
		{
			result: "hello world",
		},
		{
			items: []string{"minuda"},
			result: "hello minuda",
		},
		{
			items: []string{"minuda","saswitha"},
			result: "hello minuda,saswitha",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}

}