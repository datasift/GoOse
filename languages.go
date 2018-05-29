package goose

// whitelist of valid language codes, according to
// http://www.iana.org/assignments/language-subtag-registry/language-subtag-registry

// ISO 639-1: two-letter codes
// @see https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
var iso6931languages map[string]struct{}

// ISO 639-2: three-letter codes
// @see https://en.wikipedia.org/wiki/List_of_ISO_639-2_codes
var iso6932languages map[string]struct{}

// ISO 639-3: three-letter codes
// @see https://en.wikipedia.org/wiki/ISO_639-3
var iso6933languages map[string]struct{}

// https://iso639-3.sil.org/sites/iso639-3/files/downloads/iso-639-3.tab
var iso6933toiso6931 map[string]string

func init() {
	iso6931languages = map[string]struct{}{
		"aa": struct{}{},
		"ab": struct{}{},
		"af": struct{}{},
		"ak": struct{}{},
		"sq": struct{}{},
		"am": struct{}{},
		"ar": struct{}{},
		"an": struct{}{},
		"hy": struct{}{},
		"as": struct{}{},
		"av": struct{}{},
		"ae": struct{}{},
		"ay": struct{}{},
		"az": struct{}{},
		"ba": struct{}{},
		"bm": struct{}{},
		"eu": struct{}{},
		"be": struct{}{},
		"bn": struct{}{},
		"bh": struct{}{},
		"bi": struct{}{},
		"bs": struct{}{},
		"br": struct{}{},
		"bg": struct{}{},
		"my": struct{}{},
		"ca": struct{}{},
		"ch": struct{}{},
		"ce": struct{}{},
		"zh": struct{}{},
		"cu": struct{}{},
		"cv": struct{}{},
		"kw": struct{}{},
		"co": struct{}{},
		"cr": struct{}{},
		"cs": struct{}{},
		"da": struct{}{},
		"dv": struct{}{},
		"nl": struct{}{},
		"dz": struct{}{},
		"en": struct{}{},
		"eo": struct{}{},
		"et": struct{}{},
		"ee": struct{}{},
		"fo": struct{}{},
		"fj": struct{}{},
		"fi": struct{}{},
		"fr": struct{}{},
		"fy": struct{}{},
		"ff": struct{}{},
		"ka": struct{}{},
		"de": struct{}{},
		"gd": struct{}{},
		"ga": struct{}{},
		"gl": struct{}{},
		"gv": struct{}{},
		"el": struct{}{},
		"gn": struct{}{},
		"gu": struct{}{},
		"ht": struct{}{},
		"ha": struct{}{},
		"he": struct{}{},
		"hz": struct{}{},
		"hi": struct{}{},
		"ho": struct{}{},
		"hr": struct{}{},
		"hu": struct{}{},
		"ig": struct{}{},
		"is": struct{}{},
		"io": struct{}{},
		"ii": struct{}{},
		"iu": struct{}{},
		"ie": struct{}{},
		"ia": struct{}{},
		"id": struct{}{},
		"ik": struct{}{},
		"it": struct{}{},
		"jv": struct{}{},
		"ja": struct{}{},
		"kl": struct{}{},
		"kn": struct{}{},
		"ks": struct{}{},
		"kr": struct{}{},
		"kk": struct{}{},
		"km": struct{}{},
		"ki": struct{}{},
		"rw": struct{}{},
		"ky": struct{}{},
		"kv": struct{}{},
		"kg": struct{}{},
		"ko": struct{}{},
		"kj": struct{}{},
		"ku": struct{}{},
		"lo": struct{}{},
		"la": struct{}{},
		"lv": struct{}{},
		"li": struct{}{},
		"ln": struct{}{},
		"lt": struct{}{},
		"lb": struct{}{},
		"lu": struct{}{},
		"lg": struct{}{},
		"mk": struct{}{},
		"mh": struct{}{},
		"ml": struct{}{},
		"mi": struct{}{},
		"mr": struct{}{},
		"ms": struct{}{},
		"mg": struct{}{},
		"mt": struct{}{},
		"mn": struct{}{},
		"na": struct{}{},
		"nv": struct{}{},
		"nr": struct{}{},
		"nd": struct{}{},
		"ng": struct{}{},
		"ne": struct{}{},
		"nn": struct{}{},
		"nb": struct{}{},
		"no": struct{}{},
		"ny": struct{}{},
		"oc": struct{}{},
		"oj": struct{}{},
		"or": struct{}{},
		"om": struct{}{},
		"os": struct{}{},
		"pa": struct{}{},
		"fa": struct{}{},
		"pi": struct{}{},
		"pl": struct{}{},
		"pt": struct{}{},
		"ps": struct{}{},
		"qu": struct{}{},
		"rm": struct{}{},
		"ro": struct{}{},
		"rn": struct{}{},
		"ru": struct{}{},
		"sg": struct{}{},
		"sa": struct{}{},
		"si": struct{}{},
		"sk": struct{}{},
		"sl": struct{}{},
		"se": struct{}{},
		"sm": struct{}{},
		"sn": struct{}{},
		"sd": struct{}{},
		"so": struct{}{},
		"st": struct{}{},
		"es": struct{}{},
		"sc": struct{}{},
		"sr": struct{}{},
		"ss": struct{}{},
		"su": struct{}{},
		"sw": struct{}{},
		"sv": struct{}{},
		"ty": struct{}{},
		"ta": struct{}{},
		"tt": struct{}{},
		"te": struct{}{},
		"tg": struct{}{},
		"tl": struct{}{},
		"th": struct{}{},
		"bo": struct{}{},
		"ti": struct{}{},
		"to": struct{}{},
		"tn": struct{}{},
		"ts": struct{}{},
		"tk": struct{}{},
		"tr": struct{}{},
		"tw": struct{}{},
		"ug": struct{}{},
		"uk": struct{}{},
		"ur": struct{}{},
		"uz": struct{}{},
		"ve": struct{}{},
		"vi": struct{}{},
		"vo": struct{}{},
		"cy": struct{}{},
		"wa": struct{}{},
		"wo": struct{}{},
		"xh": struct{}{},
		"yi": struct{}{},
		"yo": struct{}{},
		"za": struct{}{},
		"zu": struct{}{},
	}

	iso6932languages = map[string]struct{}{
		"aar": struct{}{},
		"abk": struct{}{},
		"ace": struct{}{},
		"ach": struct{}{},
		"ada": struct{}{},
		"ady": struct{}{},
		"afa": struct{}{},
		"afh": struct{}{},
		"afr": struct{}{},
		"ain": struct{}{},
		"aka": struct{}{},
		"akk": struct{}{},
		"alb": struct{}{},
		"ale": struct{}{},
		"alg": struct{}{},
		"alt": struct{}{},
		"amh": struct{}{},
		"ang": struct{}{},
		"anp": struct{}{},
		"apa": struct{}{},
		"ara": struct{}{},
		"arc": struct{}{},
		"arg": struct{}{},
		"arm": struct{}{},
		"arn": struct{}{},
		"arp": struct{}{},
		"art": struct{}{},
		"arw": struct{}{},
		"asm": struct{}{},
		"ast": struct{}{},
		"ath": struct{}{},
		"aus": struct{}{},
		"ava": struct{}{},
		"ave": struct{}{},
		"awa": struct{}{},
		"aym": struct{}{},
		"aze": struct{}{},
		"bad": struct{}{},
		"bai": struct{}{},
		"bak": struct{}{},
		"bal": struct{}{},
		"bam": struct{}{},
		"ban": struct{}{},
		"baq": struct{}{},
		"bas": struct{}{},
		"bat": struct{}{},
		"bej": struct{}{},
		"bel": struct{}{},
		"bem": struct{}{},
		"ben": struct{}{},
		"ber": struct{}{},
		"bho": struct{}{},
		"bih": struct{}{},
		"bik": struct{}{},
		"bin": struct{}{},
		"bis": struct{}{},
		"bla": struct{}{},
		"bnt": struct{}{},
		"bos": struct{}{},
		"bra": struct{}{},
		"bre": struct{}{},
		"btk": struct{}{},
		"bua": struct{}{},
		"bug": struct{}{},
		"bul": struct{}{},
		"bur": struct{}{},
		"byn": struct{}{},
		"cad": struct{}{},
		"cai": struct{}{},
		"car": struct{}{},
		"cat": struct{}{},
		"cau": struct{}{},
		"ceb": struct{}{},
		"cel": struct{}{},
		"cha": struct{}{},
		"chb": struct{}{},
		"che": struct{}{},
		"chg": struct{}{},
		"chi": struct{}{},
		"chk": struct{}{},
		"chm": struct{}{},
		"chn": struct{}{},
		"cho": struct{}{},
		"chp": struct{}{},
		"chr": struct{}{},
		"chu": struct{}{},
		"chv": struct{}{},
		"chy": struct{}{},
		"cmc": struct{}{},
		"cop": struct{}{},
		"cor": struct{}{},
		"cos": struct{}{},
		"cpe": struct{}{},
		"cpf": struct{}{},
		"cpp": struct{}{},
		"cre": struct{}{},
		"crh": struct{}{},
		"crp": struct{}{},
		"csb": struct{}{},
		"cus": struct{}{},
		"cze": struct{}{},
		"dak": struct{}{},
		"dan": struct{}{},
		"dar": struct{}{},
		"day": struct{}{},
		"del": struct{}{},
		"den": struct{}{},
		"dgr": struct{}{},
		"din": struct{}{},
		"div": struct{}{},
		"doi": struct{}{},
		"dra": struct{}{},
		"dsb": struct{}{},
		"dua": struct{}{},
		"dum": struct{}{},
		"dut": struct{}{},
		"dyu": struct{}{},
		"dzo": struct{}{},
		"efi": struct{}{},
		"egy": struct{}{},
		"eka": struct{}{},
		"elx": struct{}{},
		"eng": struct{}{},
		"enm": struct{}{},
		"epo": struct{}{},
		"est": struct{}{},
		"ewe": struct{}{},
		"ewo": struct{}{},
		"fan": struct{}{},
		"fao": struct{}{},
		"fat": struct{}{},
		"fij": struct{}{},
		"fil": struct{}{},
		"fin": struct{}{},
		"fiu": struct{}{},
		"fon": struct{}{},
		"fre": struct{}{},
		"frm": struct{}{},
		"fro": struct{}{},
		"frr": struct{}{},
		"frs": struct{}{},
		"fry": struct{}{},
		"ful": struct{}{},
		"fur": struct{}{},
		"gaa": struct{}{},
		"gay": struct{}{},
		"gba": struct{}{},
		"gem": struct{}{},
		"geo": struct{}{},
		"ger": struct{}{},
		"gez": struct{}{},
		"gil": struct{}{},
		"gla": struct{}{},
		"gle": struct{}{},
		"glg": struct{}{},
		"glv": struct{}{},
		"gmh": struct{}{},
		"goh": struct{}{},
		"gon": struct{}{},
		"gor": struct{}{},
		"got": struct{}{},
		"grb": struct{}{},
		"grc": struct{}{},
		"gre": struct{}{},
		"grn": struct{}{},
		"gsw": struct{}{},
		"guj": struct{}{},
		"gwi": struct{}{},
		"hai": struct{}{},
		"hat": struct{}{},
		"hau": struct{}{},
		"haw": struct{}{},
		"heb": struct{}{},
		"her": struct{}{},
		"hil": struct{}{},
		"him": struct{}{},
		"hin": struct{}{},
		"hit": struct{}{},
		"hmn": struct{}{},
		"hmo": struct{}{},
		"hrv": struct{}{},
		"hsb": struct{}{},
		"hun": struct{}{},
		"hup": struct{}{},
		"iba": struct{}{},
		"ibo": struct{}{},
		"ice": struct{}{},
		"ido": struct{}{},
		"iii": struct{}{},
		"ijo": struct{}{},
		"iku": struct{}{},
		"ile": struct{}{},
		"ilo": struct{}{},
		"ina": struct{}{},
		"inc": struct{}{},
		"ind": struct{}{},
		"ine": struct{}{},
		"inh": struct{}{},
		"ipk": struct{}{},
		"ira": struct{}{},
		"iro": struct{}{},
		"ita": struct{}{},
		"jav": struct{}{},
		"jbo": struct{}{},
		"jpn": struct{}{},
		"jpr": struct{}{},
		"jrb": struct{}{},
		"kaa": struct{}{},
		"kab": struct{}{},
		"kac": struct{}{},
		"kal": struct{}{},
		"kam": struct{}{},
		"kan": struct{}{},
		"kar": struct{}{},
		"kas": struct{}{},
		"kau": struct{}{},
		"kaw": struct{}{},
		"kaz": struct{}{},
		"kbd": struct{}{},
		"kha": struct{}{},
		"khi": struct{}{},
		"khm": struct{}{},
		"kho": struct{}{},
		"kik": struct{}{},
		"kin": struct{}{},
		"kir": struct{}{},
		"kmb": struct{}{},
		"kok": struct{}{},
		"kom": struct{}{},
		"kon": struct{}{},
		"kor": struct{}{},
		"kos": struct{}{},
		"kpe": struct{}{},
		"krc": struct{}{},
		"krl": struct{}{},
		"kro": struct{}{},
		"kru": struct{}{},
		"kua": struct{}{},
		"kum": struct{}{},
		"kur": struct{}{},
		"kut": struct{}{},
		"lad": struct{}{},
		"lah": struct{}{},
		"lam": struct{}{},
		"lao": struct{}{},
		"lat": struct{}{},
		"lav": struct{}{},
		"lez": struct{}{},
		"lim": struct{}{},
		"lin": struct{}{},
		"lit": struct{}{},
		"lol": struct{}{},
		"loz": struct{}{},
		"ltz": struct{}{},
		"lua": struct{}{},
		"lub": struct{}{},
		"lug": struct{}{},
		"lui": struct{}{},
		"lun": struct{}{},
		"luo": struct{}{},
		"lus": struct{}{},
		"mac": struct{}{},
		"mad": struct{}{},
		"mag": struct{}{},
		"mah": struct{}{},
		"mai": struct{}{},
		"mak": struct{}{},
		"mal": struct{}{},
		"man": struct{}{},
		"mao": struct{}{},
		"map": struct{}{},
		"mar": struct{}{},
		"mas": struct{}{},
		"may": struct{}{},
		"mdf": struct{}{},
		"mdr": struct{}{},
		"men": struct{}{},
		"mga": struct{}{},
		"mic": struct{}{},
		"min": struct{}{},
		"mis": struct{}{},
		"mkh": struct{}{},
		"mlg": struct{}{},
		"mlt": struct{}{},
		"mnc": struct{}{},
		"mni": struct{}{},
		"mno": struct{}{},
		"moh": struct{}{},
		"mon": struct{}{},
		"mos": struct{}{},
		"mul": struct{}{},
		"mun": struct{}{},
		"mus": struct{}{},
		"mwl": struct{}{},
		"mwr": struct{}{},
		"myn": struct{}{},
		"myv": struct{}{},
		"nah": struct{}{},
		"nai": struct{}{},
		"nap": struct{}{},
		"nau": struct{}{},
		"nav": struct{}{},
		"nbl": struct{}{},
		"nde": struct{}{},
		"ndo": struct{}{},
		"nds": struct{}{},
		"nep": struct{}{},
		"new": struct{}{},
		"nia": struct{}{},
		"nic": struct{}{},
		"niu": struct{}{},
		"nno": struct{}{},
		"nob": struct{}{},
		"nog": struct{}{},
		"non": struct{}{},
		"nor": struct{}{},
		"nqo": struct{}{},
		"nso": struct{}{},
		"nub": struct{}{},
		"nwc": struct{}{},
		"nya": struct{}{},
		"nym": struct{}{},
		"nyn": struct{}{},
		"nyo": struct{}{},
		"nzi": struct{}{},
		"oci": struct{}{},
		"oji": struct{}{},
		"ori": struct{}{},
		"orm": struct{}{},
		"osa": struct{}{},
		"oss": struct{}{},
		"ota": struct{}{},
		"oto": struct{}{},
		"paa": struct{}{},
		"pag": struct{}{},
		"pal": struct{}{},
		"pam": struct{}{},
		"pan": struct{}{},
		"pap": struct{}{},
		"pau": struct{}{},
		"peo": struct{}{},
		"per": struct{}{},
		"phi": struct{}{},
		"phn": struct{}{},
		"pli": struct{}{},
		"pol": struct{}{},
		"pon": struct{}{},
		"por": struct{}{},
		"pra": struct{}{},
		"pro": struct{}{},
		"pus": struct{}{},
		"que": struct{}{},
		"raj": struct{}{},
		"rap": struct{}{},
		"rar": struct{}{},
		"roa": struct{}{},
		"roh": struct{}{},
		"rom": struct{}{},
		"rum": struct{}{},
		"run": struct{}{},
		"rup": struct{}{},
		"rus": struct{}{},
		"sad": struct{}{},
		"sag": struct{}{},
		"sah": struct{}{},
		"sai": struct{}{},
		"sal": struct{}{},
		"sam": struct{}{},
		"san": struct{}{},
		"sas": struct{}{},
		"sat": struct{}{},
		"scn": struct{}{},
		"sco": struct{}{},
		"sel": struct{}{},
		"sem": struct{}{},
		"sga": struct{}{},
		"sgn": struct{}{},
		"shn": struct{}{},
		"sid": struct{}{},
		"sin": struct{}{},
		"sio": struct{}{},
		"sit": struct{}{},
		"sla": struct{}{},
		"slo": struct{}{},
		"slv": struct{}{},
		"sma": struct{}{},
		"sme": struct{}{},
		"smi": struct{}{},
		"smj": struct{}{},
		"smn": struct{}{},
		"smo": struct{}{},
		"sms": struct{}{},
		"sna": struct{}{},
		"snd": struct{}{},
		"snk": struct{}{},
		"sog": struct{}{},
		"som": struct{}{},
		"son": struct{}{},
		"sot": struct{}{},
		"spa": struct{}{},
		"srd": struct{}{},
		"srn": struct{}{},
		"srp": struct{}{},
		"srr": struct{}{},
		"ssa": struct{}{},
		"ssw": struct{}{},
		"suk": struct{}{},
		"sun": struct{}{},
		"sus": struct{}{},
		"sux": struct{}{},
		"swa": struct{}{},
		"swe": struct{}{},
		"syc": struct{}{},
		"syr": struct{}{},
		"tah": struct{}{},
		"tai": struct{}{},
		"tam": struct{}{},
		"tat": struct{}{},
		"tel": struct{}{},
		"tem": struct{}{},
		"ter": struct{}{},
		"tet": struct{}{},
		"tgk": struct{}{},
		"tgl": struct{}{},
		"tha": struct{}{},
		"tib": struct{}{},
		"tig": struct{}{},
		"tir": struct{}{},
		"tiv": struct{}{},
		"tkl": struct{}{},
		"tlh": struct{}{},
		"tli": struct{}{},
		"tmh": struct{}{},
		"tog": struct{}{},
		"ton": struct{}{},
		"tpi": struct{}{},
		"tsi": struct{}{},
		"tsn": struct{}{},
		"tso": struct{}{},
		"tuk": struct{}{},
		"tum": struct{}{},
		"tup": struct{}{},
		"tur": struct{}{},
		"tut": struct{}{},
		"tvl": struct{}{},
		"twi": struct{}{},
		"tyv": struct{}{},
		"udm": struct{}{},
		"uga": struct{}{},
		"uig": struct{}{},
		"ukr": struct{}{},
		"umb": struct{}{},
		"und": struct{}{},
		"urd": struct{}{},
		"uzb": struct{}{},
		"vai": struct{}{},
		"ven": struct{}{},
		"vie": struct{}{},
		"vol": struct{}{},
		"vot": struct{}{},
		"wak": struct{}{},
		"wal": struct{}{},
		"war": struct{}{},
		"was": struct{}{},
		"wel": struct{}{},
		"wen": struct{}{},
		"wln": struct{}{},
		"wol": struct{}{},
		"xal": struct{}{},
		"xho": struct{}{},
		"yao": struct{}{},
		"yap": struct{}{},
		"yid": struct{}{},
		"yor": struct{}{},
		"ypk": struct{}{},
		"zap": struct{}{},
		"zbl": struct{}{},
		"zen": struct{}{},
		"zgh": struct{}{},
		"zha": struct{}{},
		"znd": struct{}{},
		"zul": struct{}{},
		"zun": struct{}{},
		"zxx": struct{}{},
		"zza": struct{}{},
	}

	// https://github.com/wooorm/iso-639-3
	// https://github.com/abadojack/whatlanggo/pull/7/files
	// for ISO-639-3 macrolanguages => ISO-639-1 equivalents: https://iso639-3.sil.org/code_tables/639/data/all?title=&field_iso639_cd_st_mmbrshp_639_1_tid=All&name_3=&field_iso639_element_scope_tid=76&field_iso639_language_type_tid=51&items_per_page=200
	iso6933toiso6931 = map[string]string{
		"aar": "aa",
		"abk": "ab",
		"afr": "af",
		"aka": "ak",
		"fat": "ak", // Added 2018-05-29 https://iso639-3.sil.org/code/aka
		"amh": "am",
		"ara": "ar",
		"arb": "ar",
		"aao": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/aao
		"abh": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/abh
		"acm": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/acm
		"acq": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/acq
		"acw": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/acw
		"acx": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/acx
		"acy": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/acy
		"adf": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/adf
		"aeb": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/aeb
		"aec": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/aec
		"afb": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/afb
		"ajp": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/ajp
		"apc": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/apc
		"apd": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/apd
		"ars": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/ars
		"auz": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/auz
		"avl": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/avl
		"ayh": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/ayh
		"ayl": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/ayl
		"ayn": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/ayn
		"ayp": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/ayp
		"bbz": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/bbz
		"pga": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/pga
		"shu": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/shu
		"ssh": "ar", // Added 2018-05-29 https://iso639-3.sil.org/code/ssh
		"arg": "an",
		"asm": "as",
		"ava": "av",
		"ave": "ae",
		"aym": "ay",
		"ayc": "ay", // Added 2018-05-29 https://iso639-3.sil.org/code/ayc
		"ayr": "ay", // Added 2018-05-29 https://iso639-3.sil.org/code/ayr
		"aze": "az",
		"azb": "az", // Added 2018-05-29 https://iso639-3.sil.org/code/azb
		"azj": "az",
		"bak": "ba",
		"bam": "bm",
		"bel": "be",
		"ben": "bn",
		"bho": "bh",
		"bis": "bi",
		"bod": "bo",
		"bos": "bs",
		"bre": "br",
		"bul": "bg",
		"cat": "ca",
		"ces": "cs",
		"cha": "ch",
		"che": "ce",
		"chu": "cu",
		"chv": "cv",
		"cor": "kw",
		"cos": "co",
		"cre": "cr",
		"crj": "cr", // Added 2018-05-29 https://iso639-3.sil.org/code/crj
		"crk": "cr", // Added 2018-05-29 https://iso639-3.sil.org/code/crk
		"crl": "cr", // Added 2018-05-29 https://iso639-3.sil.org/code/crl
		"crm": "cr", // Added 2018-05-29 https://iso639-3.sil.org/code/crm
		"csw": "cr", // Added 2018-05-29 https://iso639-3.sil.org/code/csw
		"cwd": "cr", // Added 2018-05-29 https://iso639-3.sil.org/code/cwd
		"cym": "cy",
		"dan": "da",
		"deu": "de",
		"nds": "de", // Added 2018-05-25 https://iso639-3.sil.org/code/nds (Low German, Low Saxon)
		"div": "dv",
		"dzo": "dz",
		"ell": "el",
		"eng": "en",
		"epo": "eo",
		"est": "et",
		"ekk": "et", // Added 2018-05-29 https://iso639-3.sil.org/code/ekk
		"vro": "et", // Added 2018-05-29 https://iso639-3.sil.org/code/vro
		"eus": "eu",
		"ewe": "ee",
		"fao": "fo",
		"fas": "fa",
		"per": "fa", // Added 2018-05-29 https://iso639-3.sil.org/code/per
		"pes": "fa", // Added 2018-05-29 https://iso639-3.sil.org/code/pes
		"prs": "fa", // Added 2018-05-29 https://iso639-3.sil.org/code/prs
		"fij": "fj",
		"fin": "fi",
		"fra": "fr",
		"fry": "fy",
		"ful": "ff",
		"ffm": "ff", // Added 2018-05-29 https://iso639-3.sil.org/code/ffm
		"fub": "ff", // Added 2018-05-29 https://iso639-3.sil.org/code/ffm
		"fuc": "ff", // Added 2018-05-29 https://iso639-3.sil.org/code/ffm
		"fue": "ff", // Added 2018-05-29 https://iso639-3.sil.org/code/ffm
		"fuf": "ff", // Added 2018-05-29 https://iso639-3.sil.org/code/ffm
		"fuh": "ff", // Added 2018-05-29 https://iso639-3.sil.org/code/ffm
		"fui": "ff", // Added 2018-05-29 https://iso639-3.sil.org/code/ffm
		"fuq": "ff", // Added 2018-05-29 https://iso639-3.sil.org/code/ffm
		"fuv": "ff", // Added 2018-05-29 https://iso639-3.sil.org/code/ffm
		"gla": "gd",
		"gle": "ga",
		"glg": "gl",
		"glv": "gv",
		"grn": "gn",
		"gnw": "gn", // Added 2018-05-29 https://iso639-3.sil.org/code/gnw
		"gug": "gn", // Added 2018-05-29 https://iso639-3.sil.org/code/gug
		"gui": "gn", // Added 2018-05-29 https://iso639-3.sil.org/code/gui
		"gun": "gn", // Added 2018-05-29 https://iso639-3.sil.org/code/gun
		"hnd": "gn", // Added 2018-05-29 https://iso639-3.sil.org/code/hnd
		"guj": "gu",
		"hat": "ht",
		"hau": "ha",
		"hbs": "sh", // Serbian
		"cnr": "sh", // Added 2018-05-29 https://iso639-3.sil.org/code/cnr
		"heb": "he",
		"her": "hz",
		"hin": "hi",
		"hmo": "ho",
		"hrv": "hr",
		"hun": "hu",
		"hye": "hy",
		"ibo": "ig",
		"ido": "io",
		"iii": "ii",
		"iku": "iu",
		"ike": "iu", // Added 2018-05-29 https://iso639-3.sil.org/code/ike
		"ikt": "iu", // Added 2018-05-29 https://iso639-3.sil.org/code/ikt
		"ile": "ie",
		"ina": "ia",
		"ind": "id",
		"ipk": "ik",
		"esi": "ik", // Added 2018-05-29 https://iso639-3.sil.org/code/esi
		"esk": "ik", // Added 2018-05-29 https://iso639-3.sil.org/code/esk
		"isl": "is",
		"ita": "it",
		"jav": "jv",
		"jpn": "ja",
		"kal": "kl",
		"kan": "kn",
		"kas": "ks",
		"kat": "ka",
		"kau": "kr",
		"kby": "kr", // Added 2018-05-29 https://iso639-3.sil.org/code/kby
		"knc": "kr", // Added 2018-05-29 https://iso639-3.sil.org/code/knc
		"krt": "kr", // Added 2018-05-29 https://iso639-3.sil.org/code/krt
		"kaz": "kk",
		"khm": "km",
		"kik": "ki",
		"kin": "rw",
		"kir": "ky",
		"kom": "kv",
		"koi": "kv", // Added 2018-05-29 https://iso639-3.sil.org/code/koi
		"kpv": "kv", // Added 2018-05-29 https://iso639-3.sil.org/code/kpv
		"kon": "kg",
		"kng": "kg", // Added 2018-05-29 https://iso639-3.sil.org/code/kng
		"kwy": "kg", // Added 2018-05-29 https://iso639-3.sil.org/code/kwy
		"ldi": "kg", // Added 2018-05-29 https://iso639-3.sil.org/code/ldi
		"kor": "ko",
		"kua": "kj",
		"kur": "ku",
		"ckb": "ku", // Added 2018-05-29 https://iso639-3.sil.org/code/ckb
		"kmr": "ku", // Added 2018-05-29 https://iso639-3.sil.org/code/kmr
		"sdh": "ku", // Added 2018-05-29 https://iso639-3.sil.org/code/sdh
		"lao": "lo",
		"lat": "la",
		"lav": "lv",
		"ltg": "lv", // Added 2018-05-29 https://iso639-3.sil.org/code/ltg
		"lvs": "lv", // Added 2018-05-29 https://iso639-3.sil.org/code/lvs
		"lim": "li",
		"lin": "ln",
		"lit": "lt",
		"ltz": "lb",
		"lub": "lu",
		"lug": "lg",
		"mah": "mh",
		"mal": "ml",
		"mar": "mr",
		"mkd": "mk",
		"mlg": "mg", // Malagasy https://iso639-3.sil.org/code/mlg
		"bhr": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/bhr
		"bjq": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/bjq
		"bmm": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/bmm
		"bzc": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/bzc
		"msh": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/msh
		"plt": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/plt
		"skg": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/skg
		"tdx": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/tdx
		"tkg": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/tkg
		"txy": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/txy
		"xmv": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/xmv
		"xmw": "mg", // Added 2018-05-29 https://iso639-3.sil.org/code/xmw
		"mlt": "mt",
		"mon": "mn", // Mongolian
		"khk": "mn", // Added 2018-05-29 https://iso639-3.sil.org/code/khk
		"mvf": "mn", // Added 2018-05-29 https://iso639-3.sil.org/code/mvf
		"mri": "mi",
		"msa": "ms", // Malay https://iso639-3.sil.org/code/msa
		"bjn": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/bjn
		"btj": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/btj
		"bve": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/bve
		"bvu": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/bvu
		"coa": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/coa
		"jax": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/jax
		"max": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/max
		"may": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/may
		"meo": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/meo
		"mfa": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/mfa
		"mfb": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/mfb
		"mly": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/mly
		"mqg": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/mqg
		"msi": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/msi
		"pse": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/pse
		"vkt": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/vkt
		"xmm": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/xmm
		"zlm": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/zlm
		"zmi": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/zmi
		"zsm": "ms", // Added 2018-05-29 https://iso639-3.sil.org/code/zsm
		"mya": "my",
		"nau": "na",
		"nav": "nv",
		"nbl": "nr",
		"nde": "nd",
		"ndo": "ng",
		"nep": "ne", // Nepali https://iso639-3.sil.org/code/nep
		"dty": "ne", // Added 2018-05-25 https://iso639-3.sil.org/code/dty
		"npi": "ne", // Added 2018-05-25 https://iso639-3.sil.org/code/npi
		"nld": "nl",
		"nno": "nn",
		"nob": "nb",
		"nor": "no",
		"nya": "ny",
		"oci": "oc",
		"oji": "oj",
		"ciw": "oj", // Added 2018-05-25 https://iso639-3.sil.org/code/ciw
		"ojb": "oj", // Added 2018-05-25 https://iso639-3.sil.org/code/ojb
		"ojc": "oj", // Added 2018-05-25 https://iso639-3.sil.org/code/ojc
		"ojg": "oj", // Added 2018-05-25 https://iso639-3.sil.org/code/ojg
		"ojs": "oj", // Added 2018-05-25 https://iso639-3.sil.org/code/ojs
		"ojw": "oj", // Added 2018-05-25 https://iso639-3.sil.org/code/ojw
		"otw": "oj", // Added 2018-05-25 https://iso639-3.sil.org/code/otw
		"ori": "or",
		"ory": "or", // Added 2018-05-25 https://iso639-3.sil.org/code/ory
		"spv": "or", // Added 2018-05-25 https://iso639-3.sil.org/code/spv
		"orm": "om",
		"gax": "om", // Added 2018-05-25 https://iso639-3.sil.org/code/gax
		"gaz": "om", // Added 2018-05-25 https://iso639-3.sil.org/code/gaz
		"hae": "om", // Added 2018-05-25 https://iso639-3.sil.org/code/hae
		"orc": "om", // Added 2018-05-25 https://iso639-3.sil.org/code/orc
		"oss": "os",
		"pan": "pa",
		"pli": "pi",
		"pol": "pl",
		"por": "pt",
		"pus": "ps", // Pushto https://iso639-3.sil.org/code/pus
		"pbt": "ps", // Pushto https://iso639-3.sil.org/code/pbt
		"pbu": "ps", // Pushto https://iso639-3.sil.org/code/pbu
		"pst": "ps", // Pushto https://iso639-3.sil.org/code/pst
		"que": "qu",
		"cqu": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qub": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qud": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"quf": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qug": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"quh": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"quk": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qul": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qup": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qur": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qus": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"quw": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qux": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"quy": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"quz": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qva": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvc": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qve": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvh": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvi": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvj": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvl": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvm": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvn": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvo": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvp": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvs": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvw": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qvz": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qwa": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qwc": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qwh": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qws": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxa": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxc": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxh": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxl": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxn": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxo": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxp": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxr": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxt": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxu": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"qxw": "qu", // Added 2018-05-29 https://iso639-3.sil.org/code/que
		"roh": "rm",
		"ron": "ro",
		"rmc": "ro", // Added 2018-05-29 https://iso639-3.sil.org/code/rmn
		"rml": "ro", // Added 2018-05-29 https://iso639-3.sil.org/code/rmn
		"rmn": "ro", // Added 2018-05-29 https://iso639-3.sil.org/code/rmn
		"rmo": "ro", // Added 2018-05-29 https://iso639-3.sil.org/code/rmn
		"rmw": "ro", // Added 2018-05-29 https://iso639-3.sil.org/code/rmn
		"rmy": "ro", // Added 2018-05-29 https://iso639-3.sil.org/code/rmn
		"run": "rn",
		"rus": "ru",
		"sag": "sg",
		"san": "sa",
		"sco": "en", // Added 2018-05-22 https://iso639-3.sil.org/code/sco
		"sin": "si",
		"slk": "sk",
		"slv": "sl",
		"sme": "se",
		"smo": "sm",
		"sna": "sn",
		"snd": "sd",
		"som": "so",
		"sot": "st",
		"spa": "es",
		"sqi": "sq",
		"aae": "sq", // Added 2018-05-29 https://iso639-3.sil.org/code/aae
		"aat": "sq", // Added 2018-05-29 https://iso639-3.sil.org/code/aae
		"aln": "sq", // Added 2018-05-29 https://iso639-3.sil.org/code/aae
		"als": "sq", // Added 2018-05-29 https://iso639-3.sil.org/code/aae
		"srd": "sc", // Sardinian
		"sdc": "sc", // Added 2018-05-29 https://iso639-3.sil.org/code/sdc
		"sdn": "sc", // Added 2018-05-29 https://iso639-3.sil.org/code/sdn
		"src": "sc", // Added 2018-05-29 https://iso639-3.sil.org/code/src
		"sro": "sc", // Added 2018-05-29 https://iso639-3.sil.org/code/sro
		"srp": "sr",
		"ssw": "ss",
		"sun": "su",
		"swa": "sw",
		"swc": "sw", // Added 2018-05-29 https://iso639-3.sil.org/code/swc
		"swh": "sw", // Added 2018-05-29 https://iso639-3.sil.org/code/swh
		"swe": "sv",
		"tah": "ty",
		"tam": "ta",
		"tat": "tt",
		"tel": "te",
		"tgk": "tg",
		"tgl": "tl",
		"tha": "th",
		"tir": "ti",
		"ton": "to",
		"tsn": "tn",
		"tso": "ts",
		"tuk": "tk",
		"tur": "tr",
		"twi": "tw",
		"uig": "ug",
		"ukr": "uk",
		"urd": "ur",
		"uzb": "uz",
		"uzn": "uz", // Added 2018-05-29 https://iso639-3.sil.org/code/uzn
		"uzs": "uz", // Added 2018-05-29 https://iso639-3.sil.org/code/uzs
		"ven": "ve",
		"vie": "vi",
		"vol": "vo",
		"wln": "wa",
		"wol": "wo",
		"xho": "xh",
		"yid": "yi",
		"ydd": "yi", // Added 2018-05-29 https://iso639-3.sil.org/code/ydd
		"yih": "yi", // Added 2018-05-29 https://iso639-3.sil.org/code/yih
		"yor": "yo",
		"zha": "za",
		"ccx": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"ccy": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zch": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zeh": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zgb": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zgm": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zgn": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zhd": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zhn": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zlj": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zln": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zlq": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zqe": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zyb": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zyg": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zyj": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zyn": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zzj": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/zha
		"zho": "zh",
		"chi": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/chi
		"cjy": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/cjy
		"cmn": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/cmn No iso 639-1, but http://www.loc.gov/standards/iso639-2/faq.html#24
		"cpx": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/cpx
		"czh": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/czh
		"czo": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/czo
		"gan": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/gan
		"hak": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hak
		"hsn": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hsn
		"lzh": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/lzh
		"mnp": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/mnp
		"nan": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/nan
		"wuu": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/wuu
		"yue": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/yue
		"hmn": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"blu": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"cqd": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hea": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hma": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmc": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmd": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hme": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmg": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmh": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmi": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmj": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hml": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmm": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmp": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmq": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hms": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmw": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmy": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hmz": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hnj": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"hrm": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"huj": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"mmr": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"muq": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"mww": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"sfm": "zh", // Added 2018-05-29 https://iso639-3.sil.org/code/hmn
		"zul": "zu",
	}
}

func isValidLanguageCode(lang string) bool {
	if _, ok := iso6931languages[lang]; ok {
		return true
	}
	if _, ok := iso6932languages[lang]; ok {
		return true
	}
	return false
}

func iso6933toIso6931(iso3 string) string {
	if iso1, ok := iso6933toiso6931[iso3]; ok {
		return iso1
	}
	return iso3
}
