package babel

import (
	"testing"
)

func TestText(t *testing.T) {
	largeText := ""
	for range 800 {
		largeText += "this is more language as you can see "
	}
	tests := []struct {
		in             string
		wantScript     Script
		wantLang       Language
		wantConfidence float64
	}{
		{"", ScriptUnknown, LanguageUnknown, 1},
		{"wep lvna eeii vl jkk azc nmn iuah ppl zccl c%l aa1z", ScriptLatin, LanguageUnknown, 0.75},
		{"this is the language", ScriptLatin, LanguageEnglish, 0.75},
		{largeText, ScriptLatin, LanguageEnglish, 0.999},

		{"We hold these truths to be self-evident, that all men are created equal", ScriptLatin, LanguageEnglish, .95},
		{"Wir halten diese Wahrheiten für ausgemacht, daß alle Menschen gleich erschaffen worden", ScriptLatin, LanguageGerman, .95},
		{"If you wanted to greet someone in this language, you'd say 'wie geht es'", ScriptLatin, LanguageEnglish, 0.35},
		{"the best thing to say is своїй гідності in my opinon of this.", ScriptLatin, LanguageEnglish, 0.55},
		{"Sostenemos como evidentes estas verdades: que los hombres son creados iguales", ScriptLatin, LanguageSpanish, 0.75},
		{"Consideramos estas verdades como autoevidentes, que todos os homens são criados iguais", ScriptLatin, LanguagePortuguese, 0.95},
		{"Wszyscy ludzie rodzą się wolni i równi w swojej godności i prawach", ScriptLatin, LanguagePolish, 0.95},
		{"ਮੇਰਾ ਨਾਮ ਭਰਤ ਹੈ.", ScriptGurmukhi, LanguagePunjabi, 0.95},
		{"Minden emberi lény szabadon születik és egyenlő méltósága és joga van", ScriptLatin, "hu", 0.95},
		{"Tutti gli esseri umani nascono liberi ed eguali in dignità e diritti", ScriptLatin, "it", 0.95},
		{"Все люди рождаются свободными и равными в своем достоинстве и правах", ScriptCyrillic, "ru", 0.55},
		{"Всі люди народжуються вільними і рівними у своїй гідності та правах", ScriptCyrillic, "uk", 0.80},
		{"Tous les êtres humains naissent libres et égaux", ScriptLatin, "fr", 0.95},
		{"원래 AB형 사람이 똑똑해", ScriptHangul, "ko", 0.95},
		{"何を食べますか", ScriptHiragana, "ja", 0.85},
		{"球的采编网络,记者遍布", ScriptHan, "zh", 0.95},
		{"اهتمامًا بذلك المشروع. المجموعة الوحيدة التي ", ScriptArabic, "ar", 0.55},
		// {"এই গবেষণায় রত, তাঁদেরকে বলা হয় ভাষাবিজ্ঞানী।ভাষাবিজ্ঞানীরা নৈর্ব্যক্তিক", ScriptBengali, "bn", 0.85},
		// {"ब तक लगातार चल रहा है। इसका प्रसारण प्रत्येक शनिवार और रविवार को रात 10 बजे होता है। इसका पुनः प्रसारण सोनी पल चैनल पर रात 9 बजे होता", ScriptBengali, "hi", 0.75},
		{"Ολοι οι άνθρωποι γεννιούνται ελεύθεροι και ίσοι στην αξιοπρέπεια και στα δικαιώματα", ScriptGreek, LanguageGreek, 0.95},
		{"כראוי. בִּדקו את כותרת הדף", ScriptHebrew, LanguageHebrew, 0.95},
		{"ગુજરાતી", ScriptGujarati, "gu", 0.95},
		{"ไทย ไทยไทย", ScriptThai, "th", 0.95},
		{"ըստ Գրիգորյան օրացույցի", ScriptArmenian, "hy", 0.95},
		{"ljudi ne znaju jer me uglavnom vide", ScriptLatin, "sr", 0.85},
		{"Код животиња су ове реакције посебно важне при зарастању рана", ScriptCyrillic, "sr", 0.95},
		{"Truyền thông Việt Nam vào dịp này đăng bài ký tên ông", ScriptLatin, "vi", 0.95},
		{"భారతదేశంలోని దక్షిణ", ScriptTelugu, "te", 0.95},
		{" நீளமான, கிளைக்காத", ScriptTamil, "ta", 0.95},
		{"ano ang nangyayari sa iyo at ang mah-ina mo ay hindi mo", ScriptLatin, "tl", 0.95},
		{"Een ieder heeft, waar hij zich ook bevindt, het recht als persoon erkend te worden voor de wet", ScriptLatin, "nl", 0.95},
		{"ನನ್ನ ಹೆಸರು ಭಾರತ್.", ScriptKannada, "kn", 0.95},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			haveScript, haveLang, haveConfidence := Detect(tt.in)
			if haveScript != tt.wantScript {
				t.Errorf("wrong script\nin:   %q\nhave: %q\nwant: %q", tt.in, haveScript, tt.wantScript)
			}
			if haveLang != tt.wantLang {
				t.Errorf("wrong language\nin:   %q\nhave: %q\nwant: %q", tt.in, haveLang, tt.wantLang)
			}
			if haveConfidence < tt.wantConfidence {
				t.Errorf("confidence too low\nin:   %q\nhave: %v\nwant: %v", tt.in, haveConfidence, tt.wantConfidence)
			}
		})
	}
}

func TestScript(t *testing.T) {
	tests := []struct {
		in             string
		want           Script
		wantConfidence float64
	}{
		{"All – Lätin 🤷...", ScriptLatin, 1},
		{"Latin text with a bit of Cyrillic in there: Слава Україні!", ScriptLatin, 0.7},
		{"More Cyrillic than Latin: Слава Україні Слава Україні Слава Україні", ScriptCyrillic, 0.6},
		{"Equal amount of Latin and Cyrillic in here x: Слава Україні Слава Україні Слава Україні", ScriptUnknown, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			have, haveConfidence := detectScript(tt.in)
			if have != tt.want {
				t.Errorf("wrong script\nhave: %q\nwant: %q", have, tt.want)
			}
			if haveConfidence < tt.wantConfidence {
				t.Errorf("confidence too low\nhave: %v\nwant: %v", haveConfidence, tt.wantConfidence)
			}
		})
	}
}
