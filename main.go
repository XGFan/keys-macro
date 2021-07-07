package main

import (
	"context"
	"fmt"
	"github.com/micmonay/keybd_event"
	"golang.design/x/hotkey"
	"golang.design/x/mainthread"
	"log"
	"time"
)

func main() {
	//kb, _ := keybd_event.NewKeyBonding()
	// For linux, it is very important to wait 2 seconds
	//if runtime.GOOS == "linux" {
	//	time.Sleep(2 * time.Second)
	//}
	mainthread.Init(func() {
		//Register a desired hotkey.
		hk, err := hotkey.Register([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModCmd}, hotkey.Key1)
		if err != nil {
			panic("hotkey registration failed")
		}
		ctx := new(context.Context)
		// Start listen hotkey event whenever you feel it is ready.
		triggered := hk.Listen(*ctx)

		hk2, err2 := hotkey.Register([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModCmd}, hotkey.Key2)
		if err2 != nil {
			panic("hotkey registration failed")
		}
		// Start listen hotkey event whenever you feel it is ready.
		triggered2 := hk2.Listen(*ctx)
		for range triggered2 {
			fmt.Println("2")
		}
		for range triggered {
			fmt.Println("1")
		}

		//for x := range triggered2 {
		//	log.Println(x)
		//	time.Sleep(time.Millisecond * 200)
		//	stringToKey(&kb, "gget -r fatedier/frp -n darwin_amd64\n")
		//}
		//for x := range triggered {
		//	log.Println(x)
		//	time.Sleep(time.Millisecond * 200)
		//	stringToKey(&kb, "gget -r fatedier/frp\n")
		//}
	})

}

func stringToKey(kb *keybd_event.KeyBonding, s string) {
	for _, r := range s {
		switch r {
		case ' ':
			kb.SetKeys(keybd_event.VK_SPACE)
		case '\n':
			kb.SetKeys(keybd_event.VK_ENTER)
		case '\t':
			kb.SetKeys(keybd_event.VK_TAB)
		case 'a':
			kb.SetKeys(keybd_event.VK_A)
		case 'b':
			kb.SetKeys(keybd_event.VK_B)
		case 'c':
			kb.SetKeys(keybd_event.VK_C)
		case 'd':
			kb.SetKeys(keybd_event.VK_D)
		case 'e':
			kb.SetKeys(keybd_event.VK_E)
		case 'f':
			kb.SetKeys(keybd_event.VK_F)
		case 'g':
			kb.SetKeys(keybd_event.VK_G)
		case 'h':
			kb.SetKeys(keybd_event.VK_H)
		case 'i':
			kb.SetKeys(keybd_event.VK_I)
		case 'j':
			kb.SetKeys(keybd_event.VK_J)
		case 'k':
			kb.SetKeys(keybd_event.VK_K)
		case 'l':
			kb.SetKeys(keybd_event.VK_L)
		case 'm':
			kb.SetKeys(keybd_event.VK_M)
		case 'n':
			kb.SetKeys(keybd_event.VK_N)
		case 'o':
			kb.SetKeys(keybd_event.VK_O)
		case 'p':
			kb.SetKeys(keybd_event.VK_P)
		case 'q':
			kb.SetKeys(keybd_event.VK_Q)
		case 'r':
			kb.SetKeys(keybd_event.VK_R)
		case 's':
			kb.SetKeys(keybd_event.VK_S)
		case 't':
			kb.SetKeys(keybd_event.VK_T)
		case 'u':
			kb.SetKeys(keybd_event.VK_U)
		case 'v':
			kb.SetKeys(keybd_event.VK_V)
		case 'w':
			kb.SetKeys(keybd_event.VK_W)
		case 'x':
			kb.SetKeys(keybd_event.VK_X)
		case 'y':
			kb.SetKeys(keybd_event.VK_Y)
		case 'z':
			kb.SetKeys(keybd_event.VK_Z)
		case 'A':
			kb.SetKeys(keybd_event.VK_A)
			kb.HasSHIFT(true)
		case 'B':
			kb.SetKeys(keybd_event.VK_B)
			kb.HasSHIFT(true)
		case 'C':
			kb.SetKeys(keybd_event.VK_C)
			kb.HasSHIFT(true)
		case 'D':
			kb.SetKeys(keybd_event.VK_D)
			kb.HasSHIFT(true)
		case 'E':
			kb.SetKeys(keybd_event.VK_E)
			kb.HasSHIFT(true)
		case 'F':
			kb.SetKeys(keybd_event.VK_F)
			kb.HasSHIFT(true)
		case 'G':
			kb.SetKeys(keybd_event.VK_G)
			kb.HasSHIFT(true)
		case 'H':
			kb.SetKeys(keybd_event.VK_H)
			kb.HasSHIFT(true)
		case 'I':
			kb.SetKeys(keybd_event.VK_I)
			kb.HasSHIFT(true)
		case 'J':
			kb.SetKeys(keybd_event.VK_J)
			kb.HasSHIFT(true)
		case 'K':
			kb.SetKeys(keybd_event.VK_K)
			kb.HasSHIFT(true)
		case 'L':
			kb.SetKeys(keybd_event.VK_L)
			kb.HasSHIFT(true)
		case 'M':
			kb.SetKeys(keybd_event.VK_M)
			kb.HasSHIFT(true)
		case 'N':
			kb.SetKeys(keybd_event.VK_N)
			kb.HasSHIFT(true)
		case 'O':
			kb.SetKeys(keybd_event.VK_O)
			kb.HasSHIFT(true)
		case 'P':
			kb.SetKeys(keybd_event.VK_P)
			kb.HasSHIFT(true)
		case 'Q':
			kb.SetKeys(keybd_event.VK_Q)
			kb.HasSHIFT(true)
		case 'R':
			kb.SetKeys(keybd_event.VK_R)
			kb.HasSHIFT(true)
		case 'S':
			kb.SetKeys(keybd_event.VK_S)
			kb.HasSHIFT(true)
		case 'T':
			kb.SetKeys(keybd_event.VK_T)
			kb.HasSHIFT(true)
		case 'U':
			kb.SetKeys(keybd_event.VK_U)
			kb.HasSHIFT(true)
		case 'V':
			kb.SetKeys(keybd_event.VK_V)
			kb.HasSHIFT(true)
		case 'W':
			kb.SetKeys(keybd_event.VK_W)
			kb.HasSHIFT(true)
		case 'X':
			kb.SetKeys(keybd_event.VK_X)
			kb.HasSHIFT(true)
		case 'Y':
			kb.SetKeys(keybd_event.VK_Y)
			kb.HasSHIFT(true)
		case 'Z':
			kb.SetKeys(keybd_event.VK_Z)
			kb.HasSHIFT(true)
		case '0':
			kb.SetKeys(keybd_event.VK_0)
		case '1':
			kb.SetKeys(keybd_event.VK_1)
		case '2':
			kb.SetKeys(keybd_event.VK_2)
		case '3':
			kb.SetKeys(keybd_event.VK_3)
		case '4':
			kb.SetKeys(keybd_event.VK_4)
		case '5':
			kb.SetKeys(keybd_event.VK_5)
		case '6':
			kb.SetKeys(keybd_event.VK_6)
		case '7':
			kb.SetKeys(keybd_event.VK_7)
		case '8':
			kb.SetKeys(keybd_event.VK_8)
		case '9':
			kb.SetKeys(keybd_event.VK_9)
		//§-=[];'\,./`
		case '§':
			kb.SetKeys(keybd_event.VK_SP1)
		case '-':
			kb.SetKeys(keybd_event.VK_SP2)
		case '=':
			kb.SetKeys(keybd_event.VK_SP3)
		case '[':
			kb.SetKeys(keybd_event.VK_SP4)
		case ']':
			kb.SetKeys(keybd_event.VK_SP5)
		case ';':
			kb.SetKeys(keybd_event.VK_SP6)
		case '\'':
			kb.SetKeys(keybd_event.VK_SP7)
		case '\\':
			kb.SetKeys(keybd_event.VK_SP8)
		case ',':
			kb.SetKeys(keybd_event.VK_SP9)
		case '.':
			kb.SetKeys(keybd_event.VK_SP10)
		case '/':
			kb.SetKeys(keybd_event.VK_SP11)
		case '`':
			kb.SetKeys(keybd_event.VK_SP12)
		//±_+{}:"|<>?~
		case '±':
			kb.SetKeys(keybd_event.VK_SP1)
			kb.HasSHIFT(true)
		case '_':
			kb.SetKeys(keybd_event.VK_SP2)
			kb.HasSHIFT(true)
		case '+':
			kb.SetKeys(keybd_event.VK_SP3)
			kb.HasSHIFT(true)
		case '{':
			kb.SetKeys(keybd_event.VK_SP4)
			kb.HasSHIFT(true)
		case '}':
			kb.SetKeys(keybd_event.VK_SP5)
			kb.HasSHIFT(true)
		case ':':
			kb.SetKeys(keybd_event.VK_SP6)
			kb.HasSHIFT(true)
		case '"':
			kb.SetKeys(keybd_event.VK_SP7)
			kb.HasSHIFT(true)
		case '|':
			kb.SetKeys(keybd_event.VK_SP8)
			kb.HasSHIFT(true)
		case '<':
			kb.SetKeys(keybd_event.VK_SP9)
			kb.HasSHIFT(true)
		case '>':
			kb.SetKeys(keybd_event.VK_SP10)
			kb.HasSHIFT(true)
		case '?':
			kb.SetKeys(keybd_event.VK_SP11)
			kb.HasSHIFT(true)
		case '~':
			kb.SetKeys(keybd_event.VK_SP12)
			kb.HasSHIFT(true)
		//!@#$%^&*()
		case '!':
			kb.SetKeys(keybd_event.VK_1)
			kb.HasSHIFT(true)
		case '@':
			kb.SetKeys(keybd_event.VK_2)
			kb.HasSHIFT(true)
		case '#':
			kb.SetKeys(keybd_event.VK_3)
			kb.HasSHIFT(true)
		case '$':
			kb.SetKeys(keybd_event.VK_4)
			kb.HasSHIFT(true)
		case '%':
			kb.SetKeys(keybd_event.VK_5)
			kb.HasSHIFT(true)
		case '^':
			kb.SetKeys(keybd_event.VK_6)
			kb.HasSHIFT(true)
		case '&':
			kb.SetKeys(keybd_event.VK_7)
			kb.HasSHIFT(true)
		case '*':
			kb.SetKeys(keybd_event.VK_8)
			kb.HasSHIFT(true)
		case '(':
			kb.SetKeys(keybd_event.VK_9)
			kb.HasSHIFT(true)
		case ')':
			kb.SetKeys(keybd_event.VK_0)
			kb.HasSHIFT(true)
		default:
			log.Printf("unknown: [ %c ]\n", r)
		}
		time.Sleep(100 * time.Millisecond)
		kb.Press()
		kb.Release()
		kb.HasSHIFT(false)
	}
}
