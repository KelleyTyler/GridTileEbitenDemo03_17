package mypkgs

/*


 */

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

type AudioThing struct {
	BaseFreq   int //220
	SampleRate int //48000
	AudContext *audio.Context
	Sound      []byte
}

func (aud *AudioThing) ToString() string {
	outstrng := fmt.Sprintf("%s:\n BF: %6d\n SR:%6d \n", "AUDIO THING:", aud.BaseFreq, aud.SampleRate)

	return outstrng
}
func (aud *AudioThing) Init00(sRate, bFreq int, AC *audio.Context) {
	aud.BaseFreq = bFreq
	aud.SampleRate = sRate
	aud.AudContext = AC
	aud.Sound = aud.Init_Sub(0, 110, []float32{1.0, 0.25}, []float32{4.0, 2.0, 1.0, 0.5})
}
func (aud *AudioThing) Init01(sRate, bFreq, note, refFreq int) {
	aud.BaseFreq = bFreq
	aud.SampleRate = sRate
	aud.AudContext = audio.NewContext(sRate)
	//[]float32{2.0, 1.0, 0.5, 0.25, 0.125, 0.075}
	// []float32{2.0, 1.0, 0.05, 0.025, 0.0125, 0.0075}
	//[]float32{1.0, 0.05}, []float32{1.0, 0.05}

	//aud.Init_Sub(0, 110, []float32{2.0}, []float32{0.250}) //<- with srate being 4800, and bfreq being 220
	aud.Sound = aud.Init_Sub(note, refFreq, []float32{2.0}, []float32{0.150})
}
func (aud *AudioThing) Init_Sub(q, refFreq int, decayAmp, decayX []float32) []byte {
	// const refFreq = 110
	dd := 5    //5
	ee := 12.0 //12
	length := dd * aud.SampleRate * aud.BaseFreq / refFreq
	refData := make([]float32, length)
	for i := 0; i < length; i++ {
		refData[i] = aud.NoiseAt(i, float32(refFreq), 5.0, decayAmp, decayX)
	}

	freq := float64(aud.BaseFreq) * math.Exp2(float64(q-1)/ee) //12.0

	// Calculate the wave data for the freq.
	length02 := dd * aud.SampleRate * aud.BaseFreq / int(freq)
	l := make([]float32, length02)
	r := make([]float32, length02)
	for i := 0; i < length02; i++ {
		idx := int(float64(i) * freq / float64(refFreq))
		if len(refData) <= idx {
			break
		}
		l[i] = refData[idx]
	}
	copy(r, l)
	n := aud.ToBytes(l, r)
	return n
}
func (aud *AudioThing) PlayThing(freq float32) {
	// f := int(freq)
	p := aud.AudContext.NewPlayerF32FromBytes(aud.Sound)
	p.Play()

}

// --This is a copy of  the ebiten examples "PianoAt" function;
func (aud *AudioThing) NoiseAt(i int, freq, divBy float32, amp, x []float32) float32 {
	// Create piano-like waves with multiple sin waves.
	// amp := []float32{1.0, 0.8, 0.6, 0.4, 0.2}
	// x := []float32{4.0, 2.0, 1.0, 0.5, 0.25}
	// amp := []float32{1.0, 0.5, 0.25}
	// x := []float32{4.0, 2.0, 1.0}
	// amp := []float32{1.0, 0.5, 0.25}
	// x := []float32{1.0, 0.25, 0.125}
	var v float32
	for j := 0; j < len(amp); j++ {
		// Decay
		a := amp[j] * float32(math.Exp(float64(-5*float32(i)*freq/float32(aud.BaseFreq)/(x[j]*float32(aud.SampleRate)))))
		v += a * float32(math.Sin(2.0*math.Pi*float64(i)*float64(freq)*float64(j+1)/float64(aud.SampleRate)))
	}
	return v / divBy
}

func (aud *AudioThing) ToBytes(l, r []float32) []byte {
	if len(l) != len(r) {
		panic("len(l) must equal to len(r)")
	}
	b := make([]byte, len(l)*8)
	for i := range l {
		lv := math.Float32bits(l[i])
		rv := math.Float32bits(r[i])
		b[8*i] = byte(lv)
		b[8*i+1] = byte(lv >> 8)
		b[8*i+2] = byte(lv >> 16)
		b[8*i+3] = byte(lv >> 24)
		b[8*i+4] = byte(rv)
		b[8*i+5] = byte(rv >> 8)
		b[8*i+6] = byte(rv >> 16)
		b[8*i+7] = byte(rv >> 24)
	}
	return b
}
