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
	Settings   *GameSettings
	Sound      []byte
	Sounds     [][]byte
}

func (aud *AudioThing) ToString() string {
	outstrng := fmt.Sprintf("%s:\n BF: %6d\n SR:%6d \n", "AUDIO THING:", aud.BaseFreq, aud.SampleRate)

	return outstrng
}
func (aud *AudioThing) Init00(sRate, bFreq int, AC *audio.Context, GS *GameSettings) {
	aud.BaseFreq = bFreq
	aud.SampleRate = sRate
	aud.AudContext = AC
	aud.Settings = GS
	aud.Sound = aud.Init_Sub(0, 110, []float32{1.0, 0.25}, []float32{4.0, 2.0, 1.0, 0.5})
	aud.Sounds = append(aud.Sounds, aud.Init_Sub(0, 110, []float32{1.0, 0.25}, []float32{4.0, 2.0, 1.0, 0.5}))
}
func (aud *AudioThing) Init01(GS *GameSettings, sRate, bFreq, note, refFreq int) {
	aud.BaseFreq = bFreq
	aud.SampleRate = sRate
	aud.AudContext = audio.NewContext(sRate)
	//[]float32{2.0, 1.0, 0.5, 0.25, 0.125, 0.075}
	// []float32{2.0, 1.0, 0.05, 0.025, 0.0125, 0.0075}
	//[]float32{1.0, 0.05}, []float32{1.0, 0.05}
	aud.Settings = GS
	//aud.Init_Sub(0, 110, []float32{2.0}, []float32{0.250}) //<- with srate being 4800, and bfreq being 220 q being 60
	aud.Sound = aud.Init_Sub(note, refFreq, []float32{2.0}, []float32{0.150})
	aud.Sounds = append(aud.Sounds, aud.Init_Sub(note, refFreq, []float32{1.0}, []float32{0.150}))
}

func (aud *AudioThing) AddToAudioThing(note, refFreq int) {
	// aud.Sounds = append(aud.Sounds, aud.Init_Sub(note, refFreq, []float32{1.0}, []float32{0.0750000}))
	aud.Sounds = append(aud.Sounds, Soundwave_CreateSound(aud.SampleRate, aud.BaseFreq, note, refFreq, []float32{1.0}, []float32{0.0750000}))

}

func (aud *AudioThing) PlayThing(num int) {
	// f := int(freq)
	p := aud.AudContext.NewPlayerF32FromBytes(aud.Sounds[num])
	p.SetVolume(float64(aud.Settings.UIAudioVolume) / 100)
	p.Play()
	fmt.Printf("PLAY THING %d --- %5.2f -----\n", len(aud.Sounds[num]), p.Volume())
	fmt.Printf("%v %v %v %v %v %v %v \n", aud.Sounds[num][0], aud.Sounds[num][500], aud.Sounds[num][1000], aud.Sounds[num][1500], aud.Sounds[num][2000], aud.Sounds[num][2500], aud.Sounds[num][3000])
}

func (aud *AudioThing) PlayByte(b []byte) {
	// f := int(freq)
	p := aud.AudContext.NewPlayerF32FromBytes(b)
	p.SetVolume(float64(aud.Settings.UIAudioVolume) / 100)
	p.Play()
	fmt.Printf("PLAY THING %d --- %5.2f -----\n", len(b), p.Volume())
	fmt.Printf("%v %v %v %v %v %v %v \n", b[0], b[500], b[1000], b[1500], b[2000], b[2500], b[3000])
}

func (aud *AudioThing) Init_Sub(q, refFreq int, decayAmp, decayX []float32) []byte {
	// const refFreq = 110
	dd := 5    //5
	ee := 12.0 //12
	length := dd * aud.SampleRate * aud.BaseFreq / refFreq
	refData := make([]float32, length)
	for i := 0; i < length; i++ {
		refData[i] = Soundwave_NoiseAt(aud.SampleRate, aud.BaseFreq, i, float32(refFreq), 5.0, decayAmp, decayX)
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
	n := Soundwave_ToBytes(l, r)
	return n
}

func Soundwave_CreateSound(SampleRate, BaseFreq, q, refFreq int, decayAmp, decayX []float32) []byte {
	// const refFreq = 110
	dd := 5    //5
	ee := 12.0 //12
	length := dd * SampleRate * BaseFreq / refFreq
	refData := make([]float32, length)
	for i := 0; i < length; i++ {
		refData[i] = Soundwave_NoiseAt(SampleRate, BaseFreq, i, float32(refFreq), 5.0, decayAmp, decayX)
	}

	freq := float64(BaseFreq) * math.Exp2(float64(q-1)/ee) //12.0

	// Calculate the wave data for the freq.
	length02 := dd * SampleRate * BaseFreq / int(freq)
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
	n := Soundwave_ToBytes(l, r)
	return n
}

// --This is a copy of  the ebiten examples "PianoAt" function;
func Soundwave_NoiseAt(sRate, BaseFreq, i int, freq, divBy float32, amp, decayX []float32) float32 {
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
		a := amp[j] * float32(math.Exp(float64(-5*float32(i)*freq/float32(BaseFreq)/(decayX[j]*float32(sRate)))))
		v += a * float32(math.Sin(2.0*math.Pi*float64(i)*float64(freq)*float64(j+1)/float64(sRate)))
	}
	return v / divBy
}

func Soundwave_ToBytes(l, r []float32) []byte {
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
