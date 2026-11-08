package getlang

import (
	"strings"
	"testing"
)

func TestText(t *testing.T) {
	largeText := ""
	for range 800 {
		largeText += "this is more language as you can see "
	}
	tests := []struct {
		in             string
		wantLang       string
		wantConfidence float64
	}{
		{"", "und", 1},
		{"wep lvna eeii vl jkk azc nmn iuah ppl zccl c%l aa1z", "und", 0.75},
		{"this is the language", "en", 0.75},
		{largeText, "en", 0.999},

		{"We hold these truths to be self-evident, that all men are created equal", "en", .95},
		{"Wir halten diese Wahrheiten für ausgemacht, daß alle Menschen gleich erschaffen worden", "de", .95},
		{"If you wanted to greet someone in this language, you'd say 'wie geht es'", "en", 0.35},
		{"the best thing to say is своїй гідності in my opinon of this.", "en", 0.55},
		{"Sostenemos como evidentes estas verdades: que los hombres son creados iguales", "es", 0.75},
		{"Consideramos estas verdades como autoevidentes, que todos os homens são criados iguais", "pt", 0.95},
		{"Wszyscy ludzie rodzą się wolni i równi w swojej godności i prawach", "pl", 0.95},
		{"ਮੇਰਾ ਨਾਮ ਭਰਤ ਹੈ.", "pa", 0.95},
		{"Minden emberi lény szabadon születik és egyenlő méltósága és joga van", "hu", 0.95},
		{"Tutti gli esseri umani nascono liberi ed eguali in dignità e diritti", "it", 0.95},
		{"Все люди рождаются свободными и равными в своем достоинстве и правах", "ru", 0.55},
		{"Всі люди народжуються вільними і рівними у своїй гідності та правах", "uk", 0.80},
		{"Tous les êtres humains naissent libres et égaux", "fr", 0.95},
		{"원래 AB형 사람이 똑똑해", "ko", 0.95},
		{"何を食べますか", "ja", 0.90},
		{"球的采编网络,记者遍布", "zh", 0.95},
		{"اهتمامًا بذلك المشروع. المجموعة الوحيدة التي ", "ar", 0.55},
		{"এই গবেষণায় রত, তাঁদেরকে বলা হয় ভাষাবিজ্ঞানী।ভাষাবিজ্ঞানীরা নৈর্ব্যক্তিক", "bn", 0.85},
		{"ब तक लगातार चल रहा है। इसका प्रसारण प्रत्येक शनिवार और रविवार को रात 10 बजे होता है। इसका पुनः प्रसारण सोनी पल चैनल पर रात 9 बजे होता", "hi", 0.75},
		{"Ολοι οι άνθρωποι γεννιούνται ελεύθεροι και ίσοι στην αξιοπρέπεια και στα δικαιώματα", "el", 0.95},
		{"כראוי. בִּדקו את כותרת הדף", "he", 0.95},
		{"ગુજરાતી", "gu", 0.95},
		{"ไทย ไทยไทย", "th", 0.95},
		{"ըստ Գրիգորյան օրացույցի", "hy", 0.95},
		{"ljudi ne znaju jer me uglavnom vide", "sr", 0.85},
		{"Код животиња су ове реакције посебно важне при зарастању рана", "sr", 0.95},
		{"Truyền thông Việt Nam vào dịp này đăng bài ký tên ông", "vi", 0.95},
		{"భారతదేశంలోని దక్షిణ", "te", 0.95},
		{" நீளமான, கிளைக்காத", "ta", 0.95},
		{"ano ang nangyayari sa iyo at ang mah-ina mo ay hindi mo", "tl", 0.95},
		{"Een ieder heeft, waar hij zich ook bevindt, het recht als persoon erkend te worden voor de wet", "nl", 0.95},
		{"ನನ್ನ ಹೆಸರು ಭಾರತ್.", "kn", 0.95},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			test := func(have Info) {
				if have.LanguageCode() != tt.wantLang {
					t.Errorf("wrong language\nhave: %q\nwant: %q", have.LanguageCode(), tt.wantLang)
				}
				if have.Confidence() < tt.wantConfidence {
					t.Errorf("confidence too lower\nhave: %v\nwant: %v", have.Confidence(), tt.wantConfidence)
				}
			}

			test(FromString(tt.in))
			r, err := FromReader(strings.NewReader(tt.in))
			if err != nil {
				t.Fatal(err)
			}
			test(r)

		})
	}
}
