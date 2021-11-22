package difflib_test

import (
	"bytes"
	"testing"

	"github.com/wiggin77/go-difflib/difflib"
)

func TestWriteContextDiffAsMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		original string
		current  string
		want     string
		wantErr  bool
	}{
		{name: "Multi paragraph", original: orig1, current: curr1, want: "", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctxDiff := difflib.ContextDiff{
				A:       difflib.SplitLines(tt.original),
				B:       difflib.SplitLines(tt.current),
				Context: 3,
			}

			writer := &bytes.Buffer{}

			if err := difflib.WriteContextDiffAsMarkdown(writer, ctxDiff); (err != nil) != tt.wantErr {
				t.Errorf("WriteContextDiffAsMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.want {
				t.Errorf("WriteContextDiffAsMarkdown() = %v, want %v", gotWriter, tt.want)
			}
		})
	}
}

const (
	orig1 = `I met a traveller from an antique land who said: 
Two vast and trunkless legs of stone stand in the desert. Near them on the sand,
half sunk, a shattered visage lies, whose frown and wrinkled lip and sneer of cold 
command tell that its sculptor well those passions read which yet survive, stamped 
on these lifeless things, the hand that mocked them and the heart that fed.
And on the pedestal these words appear: "My name is Ozymandias, King of Kings. Look 
on my works, ye mighty, and despair!"
Nothing beside remains. Round the decay of that colossal wreck, boundless and bare,
the lone and level sands stretch far away.`

	curr1 = `I met a traveller from an antique land who said: 
Two vast and trunkless legs of stone stand in the desert. Near them on the sand,
half sunk, a shattered visage lies, whose frown and wrinkled lip and sneer of cold 
command tell that its sculptor well those passions read which yet survive, stamped 
on these lifeless things, the hand that mocked them and the heart that fed.
And on the pedestal these words appear: "My name is McDavid, King of Pucks. Look 
on my works, ye mighty, and despair!"
Nothing beside remains. Round the decay of that colossal wreck, boundless and bare,
the lone and level sands stretch far away.`
)
